// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ConnectCoreV1PostNamespacedServiceProxyOKCode is the HTTP code returned for type ConnectCoreV1PostNamespacedServiceProxyOK
const ConnectCoreV1PostNamespacedServiceProxyOKCode int = 200

/*ConnectCoreV1PostNamespacedServiceProxyOK OK

swagger:response connectCoreV1PostNamespacedServiceProxyOK
*/
type ConnectCoreV1PostNamespacedServiceProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PostNamespacedServiceProxyOK creates ConnectCoreV1PostNamespacedServiceProxyOK with default headers values
func NewConnectCoreV1PostNamespacedServiceProxyOK() *ConnectCoreV1PostNamespacedServiceProxyOK {

	return &ConnectCoreV1PostNamespacedServiceProxyOK{}
}

// WithPayload adds the payload to the connect core v1 post namespaced service proxy o k response
func (o *ConnectCoreV1PostNamespacedServiceProxyOK) WithPayload(payload string) *ConnectCoreV1PostNamespacedServiceProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 post namespaced service proxy o k response
func (o *ConnectCoreV1PostNamespacedServiceProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PostNamespacedServiceProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PostNamespacedServiceProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PostNamespacedServiceProxyUnauthorized
const ConnectCoreV1PostNamespacedServiceProxyUnauthorizedCode int = 401

/*ConnectCoreV1PostNamespacedServiceProxyUnauthorized Unauthorized

swagger:response connectCoreV1PostNamespacedServiceProxyUnauthorized
*/
type ConnectCoreV1PostNamespacedServiceProxyUnauthorized struct {
}

// NewConnectCoreV1PostNamespacedServiceProxyUnauthorized creates ConnectCoreV1PostNamespacedServiceProxyUnauthorized with default headers values
func NewConnectCoreV1PostNamespacedServiceProxyUnauthorized() *ConnectCoreV1PostNamespacedServiceProxyUnauthorized {

	return &ConnectCoreV1PostNamespacedServiceProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PostNamespacedServiceProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
