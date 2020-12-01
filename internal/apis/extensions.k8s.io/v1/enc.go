// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"net/http"

	"github.com/cilium/kube-apate/internal/encoders"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	k8sAPIExtensionsV1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	k8sMetaInternalVersion "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	k8sJson "k8s.io/apimachinery/pkg/runtime/serializer/json"
)

var (
	crdEncoder k8sRuntime.Encoder
)

func init() {
	var (
		scheme = k8sRuntime.NewScheme()
		codecs serializer.CodecFactory
	)

	err := k8sAPIExtensionsV1.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	err = k8sMetaInternalVersion.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	codecs = serializer.NewCodecFactory(scheme)

	jsonSerializer := k8sJson.NewSerializerWithOptions(
		nil, scheme, scheme,
		k8sJson.SerializerOptions{Yaml: false, Pretty: false, Strict: true},
	)
	crdEncoder = codecs.EncoderForVersion(jsonSerializer, k8sMetaInternalVersion.SchemeGroupVersion)
}

func getResponder(obj k8sRuntime.Object) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
		w.Header().Set("Content-Type", encoders.JSONHeader)
		err := crdEncoder.Encode(obj, w)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.WithError(err).Error("Unable to process request")
		}
	})
}
