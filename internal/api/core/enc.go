// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"net/http"

	"github.com/cilium/kube-apate/internal/encoders"
	logging "github.com/cilium/kube-apate/internal/log"
	"github.com/cilium/kube-apate/internal/log/logfields"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	k8sCoreV1 "k8s.io/api/core/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "v1/core")

var (
	coreEncoder k8sRuntime.Encoder
)

func init() {
	var (
		scheme = k8sRuntime.NewScheme()
		codecs serializer.CodecFactory
	)

	err := k8sCoreV1.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	codecs = serializer.NewCodecFactory(scheme)
	pb := protobuf.NewSerializer(scheme, scheme)
	coreEncoder = codecs.EncoderForVersion(pb, k8sCoreV1.SchemeGroupVersion)
}

func getResponder(obj k8sRuntime.Object) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
		w.Header().Set("Content-Type", encoders.ProtobufHeader)
		err := coreEncoder.Encode(obj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.WithError(err).Error("Unable to process request")
		}
	})
}
