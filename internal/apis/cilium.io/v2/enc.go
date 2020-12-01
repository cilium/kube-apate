// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/cilium/kube-apate/internal/encoders"
	logging "github.com/cilium/kube-apate/internal/log"
	"github.com/cilium/kube-apate/internal/log/logfields"

	ciliumV2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	k8sJson "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "v2/cilium.io")

type logger struct {
}

func (*logger) Log() *logrus.Logger {
	return log.Logger
}

var (
	scheme        = k8sRuntime.NewScheme()
	ciliumEncoder k8sRuntime.Encoder
	codecs        serializer.CodecFactory
)

func init() {
	err := ciliumV2.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	codecs = serializer.NewCodecFactory(scheme)

	jsonSerializer := k8sJson.NewSerializerWithOptions(
		nil, scheme, scheme,
		k8sJson.SerializerOptions{Yaml: false, Pretty: false, Strict: true},
	)

	ciliumEncoder = codecs.EncoderForVersion(jsonSerializer, ciliumV2.SchemeGroupVersion)
}

func postResponder(obj k8sRuntime.Object) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", encoders.JSONHeader)
		err := ciliumEncoder.Encode(obj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.WithError(err).Error("Unable to process request")
		}
	})
}

func getResponder(obj k8sRuntime.Object) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
		w.Header().Set("Content-Type", encoders.JSONHeader)
		err := ciliumEncoder.Encode(obj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.WithError(err).Error("Unable to process request")
		}
	})
}
