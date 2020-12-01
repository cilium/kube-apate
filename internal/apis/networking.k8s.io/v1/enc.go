// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"net/http"

	"github.com/cilium/kube-apate/internal/encoders"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	k8sNetworkingV1 "k8s.io/api/networking/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/protobuf"
)

var (
	networkingEncoder k8sRuntime.Encoder
)

func init() {
	var (
		scheme = k8sRuntime.NewScheme()
		codecs serializer.CodecFactory
	)

	err := k8sNetworkingV1.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	codecs = serializer.NewCodecFactory(scheme)
	pb := protobuf.NewSerializer(scheme, scheme)
	networkingEncoder = codecs.EncoderForVersion(pb, k8sNetworkingV1.SchemeGroupVersion)
}

func getResponder(obj k8sRuntime.Object) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
		w.Header().Set("Content-Type", encoders.ProtobufHeader)
		err := networkingEncoder.Encode(obj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.WithError(err).Error("Unable to process request")
		}
	})
}
