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

// ConnectCoreV1PatchNamespacedPodProxyWithPathOKCode is the HTTP code returned for type ConnectCoreV1PatchNamespacedPodProxyWithPathOK
const ConnectCoreV1PatchNamespacedPodProxyWithPathOKCode int = 200

/*ConnectCoreV1PatchNamespacedPodProxyWithPathOK OK

swagger:response connectCoreV1PatchNamespacedPodProxyWithPathOK
*/
type ConnectCoreV1PatchNamespacedPodProxyWithPathOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewConnectCoreV1PatchNamespacedPodProxyWithPathOK creates ConnectCoreV1PatchNamespacedPodProxyWithPathOK with default headers values
func NewConnectCoreV1PatchNamespacedPodProxyWithPathOK() *ConnectCoreV1PatchNamespacedPodProxyWithPathOK {

	return &ConnectCoreV1PatchNamespacedPodProxyWithPathOK{}
}

// WithPayload adds the payload to the connect core v1 patch namespaced pod proxy with path o k response
func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathOK) WithPayload(payload string) *ConnectCoreV1PatchNamespacedPodProxyWithPathOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the connect core v1 patch namespaced pod proxy with path o k response
func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorizedCode is the HTTP code returned for type ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized
const ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorizedCode int = 401

/*ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized Unauthorized

swagger:response connectCoreV1PatchNamespacedPodProxyWithPathUnauthorized
*/
type ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized struct {
}

// NewConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized creates ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized() *ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized {

	return &ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized{}
}

// WriteResponse to the client
func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}