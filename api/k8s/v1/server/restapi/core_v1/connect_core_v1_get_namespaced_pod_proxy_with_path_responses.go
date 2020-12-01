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

// ConnectCoreV1GetNamespacedPodProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1GetNamespacedPodProxyWithPathOK
const ConnectCoreV1GetNamespacedPodProxyWithPathOKCode int = 200

/*ConnectCoreV1GetNamespacedPodProxyWithPathOK OK

swagger:response connectCoreV1GetNamespacedPodProxyWithPathOK
*/
type ConnectCoreV1GetNamespacedPodProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1GetNamespacedPodProxyWithPathOK creates ConnectCoreV1GetNamespacedPodProxyWithPathOK with default headers values
func NewConnectCoreV1GetNamespacedPodProxyWithPathOK() *ConnectCoreV1GetNamespacedPodProxyWithPathOK {

	return &ConnectCoreV1GetNamespacedPodProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 get namespaced pod proxy with path o k response
func (o *ConnectCoreV1GetNamespacedPodProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1GetNamespacedPodProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 get namespaced pod proxy with path o k response
func (o *ConnectCoreV1GetNamespacedPodProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1GetNamespacedPodProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized
const ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1GetNamespacedPodProxyWithPathUnauthorized
*/
type ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized creates ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized() *ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized {

	return &ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
