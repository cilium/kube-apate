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

// ConnectCoreV1PostNodeProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1PostNodeProxyWithPathOK
const ConnectCoreV1PostNodeProxyWithPathOKCode int = 200

/*ConnectCoreV1PostNodeProxyWithPathOK OK

swagger:response connectCoreV1PostNodeProxyWithPathOK
*/
type ConnectCoreV1PostNodeProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PostNodeProxyWithPathOK creates ConnectCoreV1PostNodeProxyWithPathOK with default headers values
func NewConnectCoreV1PostNodeProxyWithPathOK() *ConnectCoreV1PostNodeProxyWithPathOK {

	return &ConnectCoreV1PostNodeProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 post node proxy with path o k response
func (o *ConnectCoreV1PostNodeProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1PostNodeProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 post node proxy with path o k response
func (o *ConnectCoreV1PostNodeProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PostNodeProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PostNodeProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PostNodeProxyWithPathUnauthorized
const ConnectCoreV1PostNodeProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1PostNodeProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1PostNodeProxyWithPathUnauthorized
*/
type ConnectCoreV1PostNodeProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1PostNodeProxyWithPathUnauthorized creates ConnectCoreV1PostNodeProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PostNodeProxyWithPathUnauthorized() *ConnectCoreV1PostNodeProxyWithPathUnauthorized {

	return &ConnectCoreV1PostNodeProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PostNodeProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
