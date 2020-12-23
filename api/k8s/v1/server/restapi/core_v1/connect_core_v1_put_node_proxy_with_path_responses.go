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

// ConnectCoreV1PutNodeProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1PutNodeProxyWithPathOK
const ConnectCoreV1PutNodeProxyWithPathOKCode int = 200

/*ConnectCoreV1PutNodeProxyWithPathOK OK

swagger:response connectCoreV1PutNodeProxyWithPathOK
*/
type ConnectCoreV1PutNodeProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PutNodeProxyWithPathOK creates ConnectCoreV1PutNodeProxyWithPathOK with default headers values
func NewConnectCoreV1PutNodeProxyWithPathOK() *ConnectCoreV1PutNodeProxyWithPathOK {

	return &ConnectCoreV1PutNodeProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 put node proxy with path o k response
func (o *ConnectCoreV1PutNodeProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1PutNodeProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 put node proxy with path o k response
func (o *ConnectCoreV1PutNodeProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PutNodeProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PutNodeProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PutNodeProxyWithPathUnauthorized
const ConnectCoreV1PutNodeProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1PutNodeProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1PutNodeProxyWithPathUnauthorized
*/
type ConnectCoreV1PutNodeProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1PutNodeProxyWithPathUnauthorized creates ConnectCoreV1PutNodeProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PutNodeProxyWithPathUnauthorized() *ConnectCoreV1PutNodeProxyWithPathUnauthorized {

	return &ConnectCoreV1PutNodeProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PutNodeProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}