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

// PostApisCiliumIoV2CiliumEndpointOKCode is the HTTP code returned for type PostApisCiliumIoV2CiliumEndpointOK
const PostApisCiliumIoV2CiliumEndpointOKCode int = 200

/*PostApisCiliumIoV2CiliumEndpointOK OK

swagger:response postApisCiliumIoV2CiliumEndpointOK
*/
type PostApisCiliumIoV2CiliumEndpointOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostApisCiliumIoV2CiliumEndpointOK creates PostApisCiliumIoV2CiliumEndpointOK with default headers values
func NewPostApisCiliumIoV2CiliumEndpointOK() *PostApisCiliumIoV2CiliumEndpointOK {

	return &PostApisCiliumIoV2CiliumEndpointOK{}
}

// WithPayload adds the payload to the post apis cilium io v2 cilium endpoint o k response
func (o *PostApisCiliumIoV2CiliumEndpointOK) WithPayload(payload interface{}) *PostApisCiliumIoV2CiliumEndpointOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post apis cilium io v2 cilium endpoint o k response
func (o *PostApisCiliumIoV2CiliumEndpointOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostApisCiliumIoV2CiliumEndpointOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostApisCiliumIoV2CiliumEndpointCreatedCode is the HTTP code returned for type PostApisCiliumIoV2CiliumEndpointCreated
const PostApisCiliumIoV2CiliumEndpointCreatedCode int = 201

/*PostApisCiliumIoV2CiliumEndpointCreated Created

swagger:response postApisCiliumIoV2CiliumEndpointCreated
*/
type PostApisCiliumIoV2CiliumEndpointCreated struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostApisCiliumIoV2CiliumEndpointCreated creates PostApisCiliumIoV2CiliumEndpointCreated with default headers values
func NewPostApisCiliumIoV2CiliumEndpointCreated() *PostApisCiliumIoV2CiliumEndpointCreated {

	return &PostApisCiliumIoV2CiliumEndpointCreated{}
}

// WithPayload adds the payload to the post apis cilium io v2 cilium endpoint created response
func (o *PostApisCiliumIoV2CiliumEndpointCreated) WithPayload(payload interface{}) *PostApisCiliumIoV2CiliumEndpointCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post apis cilium io v2 cilium endpoint created response
func (o *PostApisCiliumIoV2CiliumEndpointCreated) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostApisCiliumIoV2CiliumEndpointCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostApisCiliumIoV2CiliumEndpointAcceptedCode is the HTTP code returned for type PostApisCiliumIoV2CiliumEndpointAccepted
const PostApisCiliumIoV2CiliumEndpointAcceptedCode int = 202

/*PostApisCiliumIoV2CiliumEndpointAccepted Accepted

swagger:response postApisCiliumIoV2CiliumEndpointAccepted
*/
type PostApisCiliumIoV2CiliumEndpointAccepted struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPostApisCiliumIoV2CiliumEndpointAccepted creates PostApisCiliumIoV2CiliumEndpointAccepted with default headers values
func NewPostApisCiliumIoV2CiliumEndpointAccepted() *PostApisCiliumIoV2CiliumEndpointAccepted {

	return &PostApisCiliumIoV2CiliumEndpointAccepted{}
}

// WithPayload adds the payload to the post apis cilium io v2 cilium endpoint accepted response
func (o *PostApisCiliumIoV2CiliumEndpointAccepted) WithPayload(payload interface{}) *PostApisCiliumIoV2CiliumEndpointAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post apis cilium io v2 cilium endpoint accepted response
func (o *PostApisCiliumIoV2CiliumEndpointAccepted) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostApisCiliumIoV2CiliumEndpointAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostApisCiliumIoV2CiliumEndpointUnauthorizedCode is the HTTP code returned for type PostApisCiliumIoV2CiliumEndpointUnauthorized
const PostApisCiliumIoV2CiliumEndpointUnauthorizedCode int = 401

/*PostApisCiliumIoV2CiliumEndpointUnauthorized Unauthorized

swagger:response postApisCiliumIoV2CiliumEndpointUnauthorized
*/
type PostApisCiliumIoV2CiliumEndpointUnauthorized struct {
}

// NewPostApisCiliumIoV2CiliumEndpointUnauthorized creates PostApisCiliumIoV2CiliumEndpointUnauthorized with default headers values
func NewPostApisCiliumIoV2CiliumEndpointUnauthorized() *PostApisCiliumIoV2CiliumEndpointUnauthorized {

	return &PostApisCiliumIoV2CiliumEndpointUnauthorized{}
}

// WriteResponse to the client
func (o *PostApisCiliumIoV2CiliumEndpointUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
