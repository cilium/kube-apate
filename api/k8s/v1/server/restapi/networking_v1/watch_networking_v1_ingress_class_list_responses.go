// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchNetworkingV1IngressClassListOKCode is the HTTP code returned for type WatchNetworkingV1IngressClassListOK
const WatchNetworkingV1IngressClassListOKCode int = 200

/*WatchNetworkingV1IngressClassListOK OK

swagger:response watchNetworkingV1IngressClassListOK
*/
type WatchNetworkingV1IngressClassListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchNetworkingV1IngressClassListOK creates WatchNetworkingV1IngressClassListOK with default headers values
func NewWatchNetworkingV1IngressClassListOK() *WatchNetworkingV1IngressClassListOK {

	return &WatchNetworkingV1IngressClassListOK{}
}

// WithPayload adds the payload to the watch networking v1 ingress class list o k response
func (o *WatchNetworkingV1IngressClassListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchNetworkingV1IngressClassListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch networking v1 ingress class list o k response
func (o *WatchNetworkingV1IngressClassListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchNetworkingV1IngressClassListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchNetworkingV1IngressClassListUnauthorizedCode is the HTTP code returned for type WatchNetworkingV1IngressClassListUnauthorized
const WatchNetworkingV1IngressClassListUnauthorizedCode int = 401

/*WatchNetworkingV1IngressClassListUnauthorized Unauthorized

swagger:response watchNetworkingV1IngressClassListUnauthorized
*/
type WatchNetworkingV1IngressClassListUnauthorized struct {
}

// NewWatchNetworkingV1IngressClassListUnauthorized creates WatchNetworkingV1IngressClassListUnauthorized with default headers values
func NewWatchNetworkingV1IngressClassListUnauthorized() *WatchNetworkingV1IngressClassListUnauthorized {

	return &WatchNetworkingV1IngressClassListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchNetworkingV1IngressClassListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
