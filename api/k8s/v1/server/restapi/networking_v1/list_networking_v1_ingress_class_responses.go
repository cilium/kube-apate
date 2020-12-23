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

// ListNetworkingV1IngressClassOKCode is the HTTP code returned for type ListNetworkingV1IngressClassOK
const ListNetworkingV1IngressClassOKCode int = 200

/*ListNetworkingV1IngressClassOK OK

swagger:response listNetworkingV1IngressClassOK
*/
type ListNetworkingV1IngressClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1IngressClassList `json:"body,omitempty"`
}

// NewListNetworkingV1IngressClassOK creates ListNetworkingV1IngressClassOK with default headers values
func NewListNetworkingV1IngressClassOK() *ListNetworkingV1IngressClassOK {

	return &ListNetworkingV1IngressClassOK{}
}

// WithPayload adds the payload to the list networking v1 ingress class o k response
func (o *ListNetworkingV1IngressClassOK) WithPayload(payload *models.IoK8sAPINetworkingV1IngressClassList) *ListNetworkingV1IngressClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list networking v1 ingress class o k response
func (o *ListNetworkingV1IngressClassOK) SetPayload(payload *models.IoK8sAPINetworkingV1IngressClassList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNetworkingV1IngressClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListNetworkingV1IngressClassUnauthorizedCode is the HTTP code returned for type ListNetworkingV1IngressClassUnauthorized
const ListNetworkingV1IngressClassUnauthorizedCode int = 401

/*ListNetworkingV1IngressClassUnauthorized Unauthorized

swagger:response listNetworkingV1IngressClassUnauthorized
*/
type ListNetworkingV1IngressClassUnauthorized struct {
}

// NewListNetworkingV1IngressClassUnauthorized creates ListNetworkingV1IngressClassUnauthorized with default headers values
func NewListNetworkingV1IngressClassUnauthorized() *ListNetworkingV1IngressClassUnauthorized {

	return &ListNetworkingV1IngressClassUnauthorized{}
}

// WriteResponse to the client
func (o *ListNetworkingV1IngressClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}