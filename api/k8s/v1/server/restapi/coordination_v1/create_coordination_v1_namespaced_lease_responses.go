// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateCoordinationV1NamespacedLeaseOKCode is the HTTP code returned for type CreateCoordinationV1NamespacedLeaseOK
const CreateCoordinationV1NamespacedLeaseOKCode int = 200

/*CreateCoordinationV1NamespacedLeaseOK OK

swagger:response createCoordinationV1NamespacedLeaseOK
*/
type CreateCoordinationV1NamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewCreateCoordinationV1NamespacedLeaseOK creates CreateCoordinationV1NamespacedLeaseOK with default headers values
func NewCreateCoordinationV1NamespacedLeaseOK() *CreateCoordinationV1NamespacedLeaseOK {

	return &CreateCoordinationV1NamespacedLeaseOK{}
}

// WithPayload adds the payload to the create coordination v1 namespaced lease o k response
func (o *CreateCoordinationV1NamespacedLeaseOK) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *CreateCoordinationV1NamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create coordination v1 namespaced lease o k response
func (o *CreateCoordinationV1NamespacedLeaseOK) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoordinationV1NamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoordinationV1NamespacedLeaseCreatedCode is the HTTP code returned for type CreateCoordinationV1NamespacedLeaseCreated
const CreateCoordinationV1NamespacedLeaseCreatedCode int = 201

/*CreateCoordinationV1NamespacedLeaseCreated Created

swagger:response createCoordinationV1NamespacedLeaseCreated
*/
type CreateCoordinationV1NamespacedLeaseCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewCreateCoordinationV1NamespacedLeaseCreated creates CreateCoordinationV1NamespacedLeaseCreated with default headers values
func NewCreateCoordinationV1NamespacedLeaseCreated() *CreateCoordinationV1NamespacedLeaseCreated {

	return &CreateCoordinationV1NamespacedLeaseCreated{}
}

// WithPayload adds the payload to the create coordination v1 namespaced lease created response
func (o *CreateCoordinationV1NamespacedLeaseCreated) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *CreateCoordinationV1NamespacedLeaseCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create coordination v1 namespaced lease created response
func (o *CreateCoordinationV1NamespacedLeaseCreated) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoordinationV1NamespacedLeaseCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoordinationV1NamespacedLeaseAcceptedCode is the HTTP code returned for type CreateCoordinationV1NamespacedLeaseAccepted
const CreateCoordinationV1NamespacedLeaseAcceptedCode int = 202

/*CreateCoordinationV1NamespacedLeaseAccepted Accepted

swagger:response createCoordinationV1NamespacedLeaseAccepted
*/
type CreateCoordinationV1NamespacedLeaseAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1Lease `json:"body,omitempty"`
}

// NewCreateCoordinationV1NamespacedLeaseAccepted creates CreateCoordinationV1NamespacedLeaseAccepted with default headers values
func NewCreateCoordinationV1NamespacedLeaseAccepted() *CreateCoordinationV1NamespacedLeaseAccepted {

	return &CreateCoordinationV1NamespacedLeaseAccepted{}
}

// WithPayload adds the payload to the create coordination v1 namespaced lease accepted response
func (o *CreateCoordinationV1NamespacedLeaseAccepted) WithPayload(payload *models.IoK8sAPICoordinationV1Lease) *CreateCoordinationV1NamespacedLeaseAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create coordination v1 namespaced lease accepted response
func (o *CreateCoordinationV1NamespacedLeaseAccepted) SetPayload(payload *models.IoK8sAPICoordinationV1Lease) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoordinationV1NamespacedLeaseAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoordinationV1NamespacedLeaseUnauthorizedCode is the HTTP code returned for type CreateCoordinationV1NamespacedLeaseUnauthorized
const CreateCoordinationV1NamespacedLeaseUnauthorizedCode int = 401

/*CreateCoordinationV1NamespacedLeaseUnauthorized Unauthorized

swagger:response createCoordinationV1NamespacedLeaseUnauthorized
*/
type CreateCoordinationV1NamespacedLeaseUnauthorized struct {
}

// NewCreateCoordinationV1NamespacedLeaseUnauthorized creates CreateCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewCreateCoordinationV1NamespacedLeaseUnauthorized() *CreateCoordinationV1NamespacedLeaseUnauthorized {

	return &CreateCoordinationV1NamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoordinationV1NamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
