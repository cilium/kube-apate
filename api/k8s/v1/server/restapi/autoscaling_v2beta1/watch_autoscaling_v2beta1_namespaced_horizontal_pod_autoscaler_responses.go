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

// WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
const WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK OK

swagger:response watchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
*/
type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK creates WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK() *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {

	return &WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the watch autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
const WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response watchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized creates WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewWatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized() *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}