// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchNetworkingV1beta1NamespacedIngressOKCode is the HTTP code returned for type WatchNetworkingV1beta1NamespacedIngressOK
const WatchNetworkingV1beta1NamespacedIngressOKCode int = 200

/*WatchNetworkingV1beta1NamespacedIngressOK OK

swagger:response watchNetworkingV1beta1NamespacedIngressOK
*/
type WatchNetworkingV1beta1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchNetworkingV1beta1NamespacedIngressOK creates WatchNetworkingV1beta1NamespacedIngressOK with default headers values
func NewWatchNetworkingV1beta1NamespacedIngressOK() *WatchNetworkingV1beta1NamespacedIngressOK {

	return &WatchNetworkingV1beta1NamespacedIngressOK{}
}

// WithPayload adds the payload to the watch networking v1beta1 namespaced ingress o k response
func (o *WatchNetworkingV1beta1NamespacedIngressOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchNetworkingV1beta1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch networking v1beta1 namespaced ingress o k response
func (o *WatchNetworkingV1beta1NamespacedIngressOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchNetworkingV1beta1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchNetworkingV1beta1NamespacedIngressUnauthorizedCode is the HTTP code returned for type WatchNetworkingV1beta1NamespacedIngressUnauthorized
const WatchNetworkingV1beta1NamespacedIngressUnauthorizedCode int = 401

/*WatchNetworkingV1beta1NamespacedIngressUnauthorized Unauthorized

swagger:response watchNetworkingV1beta1NamespacedIngressUnauthorized
*/
type WatchNetworkingV1beta1NamespacedIngressUnauthorized struct {
}

// NewWatchNetworkingV1beta1NamespacedIngressUnauthorized creates WatchNetworkingV1beta1NamespacedIngressUnauthorized with default headers values
func NewWatchNetworkingV1beta1NamespacedIngressUnauthorized() *WatchNetworkingV1beta1NamespacedIngressUnauthorized {

	return &WatchNetworkingV1beta1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *WatchNetworkingV1beta1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}