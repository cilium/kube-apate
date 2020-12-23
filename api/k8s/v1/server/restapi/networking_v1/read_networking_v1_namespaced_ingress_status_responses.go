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

// ReadNetworkingV1NamespacedIngressStatusOKCode is the HTTP code returned for type ReadNetworkingV1NamespacedIngressStatusOK
const ReadNetworkingV1NamespacedIngressStatusOKCode int = 200

/*ReadNetworkingV1NamespacedIngressStatusOK OK

swagger:response readNetworkingV1NamespacedIngressStatusOK
*/
type ReadNetworkingV1NamespacedIngressStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1Ingress `json:"body,omitempty"`
}

// NewReadNetworkingV1NamespacedIngressStatusOK creates ReadNetworkingV1NamespacedIngressStatusOK with default headers values
func NewReadNetworkingV1NamespacedIngressStatusOK() *ReadNetworkingV1NamespacedIngressStatusOK {

	return &ReadNetworkingV1NamespacedIngressStatusOK{}
}

// WithPayload adds the payload to the read networking v1 namespaced ingress status o k response
func (o *ReadNetworkingV1NamespacedIngressStatusOK) WithPayload(payload *models.IoK8sAPINetworkingV1Ingress) *ReadNetworkingV1NamespacedIngressStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read networking v1 namespaced ingress status o k response
func (o *ReadNetworkingV1NamespacedIngressStatusOK) SetPayload(payload *models.IoK8sAPINetworkingV1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadNetworkingV1NamespacedIngressStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadNetworkingV1NamespacedIngressStatusUnauthorizedCode is the HTTP code returned for type ReadNetworkingV1NamespacedIngressStatusUnauthorized
const ReadNetworkingV1NamespacedIngressStatusUnauthorizedCode int = 401

/*ReadNetworkingV1NamespacedIngressStatusUnauthorized Unauthorized

swagger:response readNetworkingV1NamespacedIngressStatusUnauthorized
*/
type ReadNetworkingV1NamespacedIngressStatusUnauthorized struct {
}

// NewReadNetworkingV1NamespacedIngressStatusUnauthorized creates ReadNetworkingV1NamespacedIngressStatusUnauthorized with default headers values
func NewReadNetworkingV1NamespacedIngressStatusUnauthorized() *ReadNetworkingV1NamespacedIngressStatusUnauthorized {

	return &ReadNetworkingV1NamespacedIngressStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadNetworkingV1NamespacedIngressStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}