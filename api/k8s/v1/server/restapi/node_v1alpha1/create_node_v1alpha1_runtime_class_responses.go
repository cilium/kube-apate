// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateNodeV1alpha1RuntimeClassOKCode is the HTTP code returned for type CreateNodeV1alpha1RuntimeClassOK
const CreateNodeV1alpha1RuntimeClassOKCode int = 200

/*CreateNodeV1alpha1RuntimeClassOK OK

swagger:response createNodeV1alpha1RuntimeClassOK
*/
type CreateNodeV1alpha1RuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClass `json:"body,omitempty"`
}

// NewCreateNodeV1alpha1RuntimeClassOK creates CreateNodeV1alpha1RuntimeClassOK with default headers values
func NewCreateNodeV1alpha1RuntimeClassOK() *CreateNodeV1alpha1RuntimeClassOK {

	return &CreateNodeV1alpha1RuntimeClassOK{}
}

// WithPayload adds the payload to the create node v1alpha1 runtime class o k response
func (o *CreateNodeV1alpha1RuntimeClassOK) WithPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) *CreateNodeV1alpha1RuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create node v1alpha1 runtime class o k response
func (o *CreateNodeV1alpha1RuntimeClassOK) SetPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNodeV1alpha1RuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNodeV1alpha1RuntimeClassCreatedCode is the HTTP code returned for type CreateNodeV1alpha1RuntimeClassCreated
const CreateNodeV1alpha1RuntimeClassCreatedCode int = 201

/*CreateNodeV1alpha1RuntimeClassCreated Created

swagger:response createNodeV1alpha1RuntimeClassCreated
*/
type CreateNodeV1alpha1RuntimeClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClass `json:"body,omitempty"`
}

// NewCreateNodeV1alpha1RuntimeClassCreated creates CreateNodeV1alpha1RuntimeClassCreated with default headers values
func NewCreateNodeV1alpha1RuntimeClassCreated() *CreateNodeV1alpha1RuntimeClassCreated {

	return &CreateNodeV1alpha1RuntimeClassCreated{}
}

// WithPayload adds the payload to the create node v1alpha1 runtime class created response
func (o *CreateNodeV1alpha1RuntimeClassCreated) WithPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) *CreateNodeV1alpha1RuntimeClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create node v1alpha1 runtime class created response
func (o *CreateNodeV1alpha1RuntimeClassCreated) SetPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNodeV1alpha1RuntimeClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNodeV1alpha1RuntimeClassAcceptedCode is the HTTP code returned for type CreateNodeV1alpha1RuntimeClassAccepted
const CreateNodeV1alpha1RuntimeClassAcceptedCode int = 202

/*CreateNodeV1alpha1RuntimeClassAccepted Accepted

swagger:response createNodeV1alpha1RuntimeClassAccepted
*/
type CreateNodeV1alpha1RuntimeClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClass `json:"body,omitempty"`
}

// NewCreateNodeV1alpha1RuntimeClassAccepted creates CreateNodeV1alpha1RuntimeClassAccepted with default headers values
func NewCreateNodeV1alpha1RuntimeClassAccepted() *CreateNodeV1alpha1RuntimeClassAccepted {

	return &CreateNodeV1alpha1RuntimeClassAccepted{}
}

// WithPayload adds the payload to the create node v1alpha1 runtime class accepted response
func (o *CreateNodeV1alpha1RuntimeClassAccepted) WithPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) *CreateNodeV1alpha1RuntimeClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create node v1alpha1 runtime class accepted response
func (o *CreateNodeV1alpha1RuntimeClassAccepted) SetPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNodeV1alpha1RuntimeClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNodeV1alpha1RuntimeClassUnauthorizedCode is the HTTP code returned for type CreateNodeV1alpha1RuntimeClassUnauthorized
const CreateNodeV1alpha1RuntimeClassUnauthorizedCode int = 401

/*CreateNodeV1alpha1RuntimeClassUnauthorized Unauthorized

swagger:response createNodeV1alpha1RuntimeClassUnauthorized
*/
type CreateNodeV1alpha1RuntimeClassUnauthorized struct {
}

// NewCreateNodeV1alpha1RuntimeClassUnauthorized creates CreateNodeV1alpha1RuntimeClassUnauthorized with default headers values
func NewCreateNodeV1alpha1RuntimeClassUnauthorized() *CreateNodeV1alpha1RuntimeClassUnauthorized {

	return &CreateNodeV1alpha1RuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *CreateNodeV1alpha1RuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}