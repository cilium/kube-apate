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

// ConnectCoreV1OptionsNodeProxyOKCode is the HTTP code returned for type ConnectCoreV1OptionsNodeProxyOK
const ConnectCoreV1OptionsNodeProxyOKCode int = 200

/*ConnectCoreV1OptionsNodeProxyOK OK

swagger:response connectCoreV1OptionsNodeProxyOK
*/
type ConnectCoreV1OptionsNodeProxyOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1OptionsNodeProxyOK creates ConnectCoreV1OptionsNodeProxyOK with default headers values
func NewConnectCoreV1OptionsNodeProxyOK() *ConnectCoreV1OptionsNodeProxyOK {

	return &ConnectCoreV1OptionsNodeProxyOK{}
}

// WithPayload adds the payload to the connect core v1 options node proxy o k response
func (o *ConnectCoreV1OptionsNodeProxyOK) WithPayload(payload string) *ConnectCoreV1OptionsNodeProxyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 options node proxy o k response
func (o *ConnectCoreV1OptionsNodeProxyOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1OptionsNodeProxyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1OptionsNodeProxyUnauthorizedCode is the HTTP code returned for type ConnectCoreV1OptionsNodeProxyUnauthorized
const ConnectCoreV1OptionsNodeProxyUnauthorizedCode int = 401

/*ConnectCoreV1OptionsNodeProxyUnauthorized Unauthorized

swagger:response connectCoreV1OptionsNodeProxyUnauthorized
*/
type ConnectCoreV1OptionsNodeProxyUnauthorized struct {
}

// NewConnectCoreV1OptionsNodeProxyUnauthorized creates ConnectCoreV1OptionsNodeProxyUnauthorized with default headers values
func NewConnectCoreV1OptionsNodeProxyUnauthorized() *ConnectCoreV1OptionsNodeProxyUnauthorized {

	return &ConnectCoreV1OptionsNodeProxyUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1OptionsNodeProxyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}