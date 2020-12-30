// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package encoders

import (
	"net/http"
	"strconv"
	"sync"
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

type ObjHandlerNoOp struct{}

func (*ObjHandlerNoOp) AddElements(int64) int64 {
	return 0
}

func (*ObjHandlerNoOp) GenObjs(_, _ int64) <-chan k8sRuntime.Object {
	ch := make(chan k8sRuntime.Object)
	close(ch)
	return ch
}

func (*ObjHandlerNoOp) Stream() chan k8sRuntime.Object {
	return nil
}

func (*ObjHandlerNoOp) Log() *logrus.Logger {
	return logrus.New()
}
func GenerateK8sEventsWithDependent(owner, dependent ObjHandler, addElemts int64) {
	generateK8sEventsWithDependent(owner, dependent, true, addElemts)
}

func GenerateK8sEvents(owner ObjHandler, addElemts int64) {
	generateK8sEventsWithDependent(owner, &ObjHandlerNoOp{}, false, addElemts)
}

func generateK8sEventsWithDependent(owner, dependent ObjHandler, withDependent bool, addElemts int64) {
	if addElemts == 0 {
		return
	}
	prevOwnerTotal := owner.AddElements(addElemts)
	ownerTotal := prevOwnerTotal + addElemts
	dependent.AddElements(addElemts)

	var (
		eventType     string
		totalEvents   int64
		objs, depObjs <-chan k8sRuntime.Object
	)
	switch {
	case addElemts < 0:
		objs = owner.GenObjs(ownerTotal, prevOwnerTotal)
		depObjs = dependent.GenObjs(ownerTotal, prevOwnerTotal)
		eventType = string(k8sWatch.Deleted)
		totalEvents = -addElemts
	case addElemts > 0:
		objs = owner.GenObjs(prevOwnerTotal, ownerTotal)
		depObjs = dependent.GenObjs(prevOwnerTotal, ownerTotal)
		eventType = string(k8sWatch.Added)
		totalEvents = addElemts
	}
	var (
		wg           sync.WaitGroup
		newDependent chan struct{}
	)
	if withDependent {
		newDependent = make(chan struct{})
	}

	wg.Add(1)
	go func() {
		streamObjs(owner, objs, eventType, true, newDependent, totalEvents)
		wg.Done()
	}()

	if withDependent {
		wg.Add(1)
		streamObjs(dependent, depObjs, eventType, false, newDependent, totalEvents)
		wg.Done()
	}

	wg.Wait()
}

func streamObjs(
	objHandler ObjHandler,
	objs <-chan k8sRuntime.Object,
	eventType string,
	isOwner bool,
	newDependent chan struct{},
	totalEvents int64) {

	streamer := objHandler.Stream()
	for ce := range objs {
		we := &k8sMetaV1.WatchEvent{
			Type: eventType,
			Object: k8sRuntime.RawExtension{
				Object: ce,
			},
		}
		if newDependent != nil {
			if !isOwner {
				// wait until the owner of this dependent is created first
				<-newDependent
				streamer <- we
			} else {
				streamer <- we
				// wait until the dependent of this owner is created first
				newDependent <- struct{}{}
			}
		} else {
			streamer <- we
		}
	}
	objHandler.Log().WithFields(logrus.Fields{
		"total-events": totalEvents,
		"event-type":   eventType,
	}).Info("Streamed Objects")
}
