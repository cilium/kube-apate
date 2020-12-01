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

// ConnectCoreV1DeleteNodeProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1DeleteNodeProxyWithPathOK
const ConnectCoreV1DeleteNodeProxyWithPathOKCode int = 200

/*ConnectCoreV1DeleteNodeProxyWithPathOK OK

swagger:response connectCoreV1DeleteNodeProxyWithPathOK
*/
type ConnectCoreV1DeleteNodeProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1DeleteNodeProxyWithPathOK creates ConnectCoreV1DeleteNodeProxyWithPathOK with default headers values
func NewConnectCoreV1DeleteNodeProxyWithPathOK() *ConnectCoreV1DeleteNodeProxyWithPathOK {

	return &ConnectCoreV1DeleteNodeProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 delete node proxy with path o k response
func (o *ConnectCoreV1DeleteNodeProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1DeleteNodeProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 delete node proxy with path o k response
func (o *ConnectCoreV1DeleteNodeProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1DeleteNodeProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1DeleteNodeProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1DeleteNodeProxyWithPathUnauthorized
const ConnectCoreV1DeleteNodeProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1DeleteNodeProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1DeleteNodeProxyWithPathUnauthorized
*/
type ConnectCoreV1DeleteNodeProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1DeleteNodeProxyWithPathUnauthorized creates ConnectCoreV1DeleteNodeProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1DeleteNodeProxyWithPathUnauthorized() *ConnectCoreV1DeleteNodeProxyWithPathUnauthorized {

	return &ConnectCoreV1DeleteNodeProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1DeleteNodeProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
