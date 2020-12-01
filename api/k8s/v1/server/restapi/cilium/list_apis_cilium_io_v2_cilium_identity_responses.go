// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListApisCiliumIoV2CiliumIdentityOKCode is the HTTP code returned for type ListApisCiliumIoV2CiliumIdentityOK
const ListApisCiliumIoV2CiliumIdentityOKCode int = 200

/*ListApisCiliumIoV2CiliumIdentityOK OK

swagger:response listApisCiliumIoV2CiliumIdentityOK
*/
type ListApisCiliumIoV2CiliumIdentityOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewListApisCiliumIoV2CiliumIdentityOK creates ListApisCiliumIoV2CiliumIdentityOK with default headers values
func NewListApisCiliumIoV2CiliumIdentityOK() *ListApisCiliumIoV2CiliumIdentityOK {

	return &ListApisCiliumIoV2CiliumIdentityOK{}
}

// WithPayload adds the payload to the list apis cilium io v2 cilium identity o k response
func (o *ListApisCiliumIoV2CiliumIdentityOK) WithPayload(payload interface{}) *ListApisCiliumIoV2CiliumIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list apis cilium io v2 cilium identity o k response
func (o *ListApisCiliumIoV2CiliumIdentityOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListApisCiliumIoV2CiliumIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListApisCiliumIoV2CiliumIdentityUnauthorizedCode is the HTTP code returned for type ListApisCiliumIoV2CiliumIdentityUnauthorized
const ListApisCiliumIoV2CiliumIdentityUnauthorizedCode int = 401

/*ListApisCiliumIoV2CiliumIdentityUnauthorized Unauthorized

swagger:response listApisCiliumIoV2CiliumIdentityUnauthorized
*/
type ListApisCiliumIoV2CiliumIdentityUnauthorized struct {
}

// NewListApisCiliumIoV2CiliumIdentityUnauthorized creates ListApisCiliumIoV2CiliumIdentityUnauthorized with default headers values
func NewListApisCiliumIoV2CiliumIdentityUnauthorized() *ListApisCiliumIoV2CiliumIdentityUnauthorized {

	return &ListApisCiliumIoV2CiliumIdentityUnauthorized{}
}

// WriteResponse to the client
func (o *ListApisCiliumIoV2CiliumIdentityUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
