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

// ConnectCoreV1PutNodeProxyOKCode is the HTTP code returned for type ConnectCoreV1PutNodeProxyOK
const ConnectCoreV1PutNodeProxyOKCode int = 200

/*ConnectCoreV1PutNodeProxyOK OK

swagger:response connectCoreV1PutNodeProxyOK
*/
type ConnectCoreV1PutNodeProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PutNodeProxyOK creates ConnectCoreV1PutNodeProxyOK with default headers values
func NewConnectCoreV1PutNodeProxyOK() *ConnectCoreV1PutNodeProxyOK {

	return &ConnectCoreV1PutNodeProxyOK{}
}

// WithPayload adds the payload to the connect core v1 put node proxy o k response
func (o *ConnectCoreV1PutNodeProxyOK) WithPayload(payload string) *ConnectCoreV1PutNodeProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 put node proxy o k response
func (o *ConnectCoreV1PutNodeProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PutNodeProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PutNodeProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PutNodeProxyUnauthorized
const ConnectCoreV1PutNodeProxyUnauthorizedCode int = 401

/*ConnectCoreV1PutNodeProxyUnauthorized Unauthorized

swagger:response connectCoreV1PutNodeProxyUnauthorized
*/
type ConnectCoreV1PutNodeProxyUnauthorized struct {
}

// NewConnectCoreV1PutNodeProxyUnauthorized creates ConnectCoreV1PutNodeProxyUnauthorized with default headers values
func NewConnectCoreV1PutNodeProxyUnauthorized() *ConnectCoreV1PutNodeProxyUnauthorized {

	return &ConnectCoreV1PutNodeProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PutNodeProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
