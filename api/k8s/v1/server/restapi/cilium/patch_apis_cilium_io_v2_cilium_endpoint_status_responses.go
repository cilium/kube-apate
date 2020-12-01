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

// PatchApisCiliumIoV2CiliumEndpointStatusOKCode is the HTTP code returned for type PatchApisCiliumIoV2CiliumEndpointStatusOK
const PatchApisCiliumIoV2CiliumEndpointStatusOKCode int = 200

/*PatchApisCiliumIoV2CiliumEndpointStatusOK OK

swagger:response patchApisCiliumIoV2CiliumEndpointStatusOK
*/
type PatchApisCiliumIoV2CiliumEndpointStatusOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchApisCiliumIoV2CiliumEndpointStatusOK creates PatchApisCiliumIoV2CiliumEndpointStatusOK with default headers values
func NewPatchApisCiliumIoV2CiliumEndpointStatusOK() *PatchApisCiliumIoV2CiliumEndpointStatusOK {

	return &PatchApisCiliumIoV2CiliumEndpointStatusOK{}
}

// WithPayload adds the payload to the patch apis cilium io v2 cilium endpoint status o k response
func (o *PatchApisCiliumIoV2CiliumEndpointStatusOK) WithPayload(payload interface{}) *PatchApisCiliumIoV2CiliumEndpointStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apis cilium io v2 cilium endpoint status o k response
func (o *PatchApisCiliumIoV2CiliumEndpointStatusOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApisCiliumIoV2CiliumEndpointStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchApisCiliumIoV2CiliumEndpointStatusCreatedCode is the HTTP code returned for type PatchApisCiliumIoV2CiliumEndpointStatusCreated
const PatchApisCiliumIoV2CiliumEndpointStatusCreatedCode int = 201

/*PatchApisCiliumIoV2CiliumEndpointStatusCreated Created

swagger:response patchApisCiliumIoV2CiliumEndpointStatusCreated
*/
type PatchApisCiliumIoV2CiliumEndpointStatusCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchApisCiliumIoV2CiliumEndpointStatusCreated creates PatchApisCiliumIoV2CiliumEndpointStatusCreated with default headers values
func NewPatchApisCiliumIoV2CiliumEndpointStatusCreated() *PatchApisCiliumIoV2CiliumEndpointStatusCreated {

	return &PatchApisCiliumIoV2CiliumEndpointStatusCreated{}
}

// WithPayload adds the payload to the patch apis cilium io v2 cilium endpoint status created response
func (o *PatchApisCiliumIoV2CiliumEndpointStatusCreated) WithPayload(payload interface{}) *PatchApisCiliumIoV2CiliumEndpointStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apis cilium io v2 cilium endpoint status created response
func (o *PatchApisCiliumIoV2CiliumEndpointStatusCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApisCiliumIoV2CiliumEndpointStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchApisCiliumIoV2CiliumEndpointStatusAcceptedCode is the HTTP code returned for type PatchApisCiliumIoV2CiliumEndpointStatusAccepted
const PatchApisCiliumIoV2CiliumEndpointStatusAcceptedCode int = 202

/*PatchApisCiliumIoV2CiliumEndpointStatusAccepted Accepted

swagger:response patchApisCiliumIoV2CiliumEndpointStatusAccepted
*/
type PatchApisCiliumIoV2CiliumEndpointStatusAccepted struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPatchApisCiliumIoV2CiliumEndpointStatusAccepted creates PatchApisCiliumIoV2CiliumEndpointStatusAccepted with default headers values
func NewPatchApisCiliumIoV2CiliumEndpointStatusAccepted() *PatchApisCiliumIoV2CiliumEndpointStatusAccepted {

	return &PatchApisCiliumIoV2CiliumEndpointStatusAccepted{}
}

// WithPayload adds the payload to the patch apis cilium io v2 cilium endpoint status accepted response
func (o *PatchApisCiliumIoV2CiliumEndpointStatusAccepted) WithPayload(payload interface{}) *PatchApisCiliumIoV2CiliumEndpointStatusAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apis cilium io v2 cilium endpoint status accepted response
func (o *PatchApisCiliumIoV2CiliumEndpointStatusAccepted) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApisCiliumIoV2CiliumEndpointStatusAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchApisCiliumIoV2CiliumEndpointStatusUnauthorizedCode is the HTTP code returned for type PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized
const PatchApisCiliumIoV2CiliumEndpointStatusUnauthorizedCode int = 401

/*PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized Unauthorized

swagger:response patchApisCiliumIoV2CiliumEndpointStatusUnauthorized
*/
type PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized struct {
}

// NewPatchApisCiliumIoV2CiliumEndpointStatusUnauthorized creates PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized with default headers values
func NewPatchApisCiliumIoV2CiliumEndpointStatusUnauthorized() *PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized {

	return &PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchApisCiliumIoV2CiliumEndpointStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
