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

// ConnectCoreV1PatchNodeProxyOKCode is the HTTP code returned for type ConnectCoreV1PatchNodeProxyOK
const ConnectCoreV1PatchNodeProxyOKCode int = 200

/*ConnectCoreV1PatchNodeProxyOK OK

swagger:response connectCoreV1PatchNodeProxyOK
*/
type ConnectCoreV1PatchNodeProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PatchNodeProxyOK creates ConnectCoreV1PatchNodeProxyOK with default headers values
func NewConnectCoreV1PatchNodeProxyOK() *ConnectCoreV1PatchNodeProxyOK {

	return &ConnectCoreV1PatchNodeProxyOK{}
}

// WithPayload adds the payload to the connect core v1 patch node proxy o k response
func (o *ConnectCoreV1PatchNodeProxyOK) WithPayload(payload string) *ConnectCoreV1PatchNodeProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 patch node proxy o k response
func (o *ConnectCoreV1PatchNodeProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PatchNodeProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PatchNodeProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PatchNodeProxyUnauthorized
const ConnectCoreV1PatchNodeProxyUnauthorizedCode int = 401

/*ConnectCoreV1PatchNodeProxyUnauthorized Unauthorized

swagger:response connectCoreV1PatchNodeProxyUnauthorized
*/
type ConnectCoreV1PatchNodeProxyUnauthorized struct {
}

// NewConnectCoreV1PatchNodeProxyUnauthorized creates ConnectCoreV1PatchNodeProxyUnauthorized with default headers values
func NewConnectCoreV1PatchNodeProxyUnauthorized() *ConnectCoreV1PatchNodeProxyUnauthorized {

	return &ConnectCoreV1PatchNodeProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PatchNodeProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}