// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOKCode is the HTTP code returned for type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK
const WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOKCode int = 200

/*WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK OK

swagger:response watchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK
*/
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK creates WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK with default headers values
func NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK() *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK {

	return &WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch autoscaling v2beta1 horizontal pod autoscaler list for all namespaces o k response
func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch autoscaling v2beta1 horizontal pod autoscaler list for all namespaces o k response
func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized
const WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorizedCode int = 401

/*WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized Unauthorized

swagger:response watchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized
*/
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized struct {
}

// NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized creates WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized with default headers values
func NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized() *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized {

	return &WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
