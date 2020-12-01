// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package encoders

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sMetaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sWatch "k8s.io/apimachinery/pkg/watch"
)

type lister interface {
	List(start, limit int64) (k8sRuntime.Object, error)
}

type streamLister interface {
	lister
	Stream() chan k8sRuntime.Object
}

type WatchOrListResponderCfg struct {
	Watch          *bool
	TimeoutSeconds *int64
	Limit          *int64
	Cont           *string
}

type getResponder func(obj k8sRuntime.Object) middleware.Responder

func WatchOrListResponder(lister lister, params *WatchOrListResponderCfg, getResponder getResponder) middleware.Responder {
	if params.Watch != nil && *params.Watch {
		var timeout int64
		if params.TimeoutSeconds != nil {
			timeout = *params.TimeoutSeconds
		}
		return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
			time.Sleep(time.Duration(timeout) * time.Second)
		})
	}

	var (
		obj   k8sRuntime.Object
		err   error
		start int64
		limit int64
	)
	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Cont != nil {
		start, err = strconv.ParseInt(*params.Cont, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	obj, err = lister.List(start, limit)
	if err != nil {
		panic(err)
	}

	return getResponder(obj)
}

func WatchOrListResponderWithEvents(lister streamLister, params *WatchOrListResponderCfg, getResponder getResponder, enc k8sRuntime.Encoder) middleware.Responder {
	if params.Watch != nil && *params.Watch {
		var timeout int64
		if params.TimeoutSeconds != nil {
			timeout = *params.TimeoutSeconds
		}
		return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
			c := time.NewTimer(time.Duration(timeout) * time.Second)
			st := lister.Stream()
			for {
				select {
				case <-c.C:
					return
				case obj, ok := <-st:
					if !ok {
						st = lister.Stream()
						continue
					}
					err := enc.Encode(obj, w)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
					}
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
				}
				c.Reset(time.Duration(timeout) * time.Second)
			}
		})
	}

	var (
		obj   k8sRuntime.Object
		err   error
		start int64
		limit int64
	)
	if params.Limit != nil {
		limit = *params.Limit
	}
	if params.Cont != nil {
		start, err = strconv.ParseInt(*params.Cont, 10, 64)
		if err != nil {
			panic(err)
		}
	}
	obj, err = lister.List(start, limit)
	if err != nil {
		panic(err)
	}

	return getResponder(obj)
}

type ObjHandler interface {
	AddElements(elem int64) int64
	GenObjs(start, maxElemts int64) <-chan k8sRuntime.Object
	Stream() chan k8sRuntime.Object
	Log() *logrus.Logger
}

func GenerateK8sEvents(o ObjHandler, addElemts int64) {
	if addElemts == 0 {
		return
	}
	prevTotal := o.AddElements(addElemts)
	total := prevTotal + addElemts

	var (
		eventType   string
		totalEvents int64
		objs        <-chan k8sRuntime.Object
	)
	switch {
	case addElemts < 0:
		objs = o.GenObjs(total, prevTotal)
		eventType = string(k8sWatch.Deleted)
		totalEvents = -addElemts
	case addElemts > 0:
		objs = o.GenObjs(prevTotal, total)
		eventType = string(k8sWatch.Added)
		totalEvents = addElemts
	}
	streamer := o.Stream()
	for ce := range objs {
		we := &k8sMetaV1.WatchEvent{
			Type: eventType,
			Object: k8sRuntime.RawExtension{
				Object: ce,
			},
		}
		streamer <- we
	}
	o.Log().WithFields(logrus.Fields{
		"total-events": totalEvents,
		"event-type":   eventType,
	}).Info("Streamed Objects")
}
