// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateEventsV1NamespacedEventOKCode is the HTTP code returned for type CreateEventsV1NamespacedEventOK
const CreateEventsV1NamespacedEventOKCode int = 200

/*CreateEventsV1NamespacedEventOK OK

swagger:response createEventsV1NamespacedEventOK
*/
type CreateEventsV1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1Event `json:"body,omitempty"`
}

// NewCreateEventsV1NamespacedEventOK creates CreateEventsV1NamespacedEventOK with default headers values
func NewCreateEventsV1NamespacedEventOK() *CreateEventsV1NamespacedEventOK {

	return &CreateEventsV1NamespacedEventOK{}
}

// WithPayload adds the payload to the create events v1 namespaced event o k response
func (o *CreateEventsV1NamespacedEventOK) WithPayload(payload *models.IoK8sAPIEventsV1Event) *CreateEventsV1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create events v1 namespaced event o k response
func (o *CreateEventsV1NamespacedEventOK) SetPayload(payload *models.IoK8sAPIEventsV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventsV1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventsV1NamespacedEventCreatedCode is the HTTP code returned for type CreateEventsV1NamespacedEventCreated
const CreateEventsV1NamespacedEventCreatedCode int = 201

/*CreateEventsV1NamespacedEventCreated Created

swagger:response createEventsV1NamespacedEventCreated
*/
type CreateEventsV1NamespacedEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1Event `json:"body,omitempty"`
}

// NewCreateEventsV1NamespacedEventCreated creates CreateEventsV1NamespacedEventCreated with default headers values
func NewCreateEventsV1NamespacedEventCreated() *CreateEventsV1NamespacedEventCreated {

	return &CreateEventsV1NamespacedEventCreated{}
}

// WithPayload adds the payload to the create events v1 namespaced event created response
func (o *CreateEventsV1NamespacedEventCreated) WithPayload(payload *models.IoK8sAPIEventsV1Event) *CreateEventsV1NamespacedEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create events v1 namespaced event created response
func (o *CreateEventsV1NamespacedEventCreated) SetPayload(payload *models.IoK8sAPIEventsV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventsV1NamespacedEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventsV1NamespacedEventAcceptedCode is the HTTP code returned for type CreateEventsV1NamespacedEventAccepted
const CreateEventsV1NamespacedEventAcceptedCode int = 202

/*CreateEventsV1NamespacedEventAccepted Accepted

swagger:response createEventsV1NamespacedEventAccepted
*/
type CreateEventsV1NamespacedEventAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1Event `json:"body,omitempty"`
}

// NewCreateEventsV1NamespacedEventAccepted creates CreateEventsV1NamespacedEventAccepted with default headers values
func NewCreateEventsV1NamespacedEventAccepted() *CreateEventsV1NamespacedEventAccepted {

	return &CreateEventsV1NamespacedEventAccepted{}
}

// WithPayload adds the payload to the create events v1 namespaced event accepted response
func (o *CreateEventsV1NamespacedEventAccepted) WithPayload(payload *models.IoK8sAPIEventsV1Event) *CreateEventsV1NamespacedEventAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create events v1 namespaced event accepted response
func (o *CreateEventsV1NamespacedEventAccepted) SetPayload(payload *models.IoK8sAPIEventsV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateEventsV1NamespacedEventAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateEventsV1NamespacedEventUnauthorizedCode is the HTTP code returned for type CreateEventsV1NamespacedEventUnauthorized
const CreateEventsV1NamespacedEventUnauthorizedCode int = 401

/*CreateEventsV1NamespacedEventUnauthorized Unauthorized

swagger:response createEventsV1NamespacedEventUnauthorized
*/
type CreateEventsV1NamespacedEventUnauthorized struct {
}

// NewCreateEventsV1NamespacedEventUnauthorized creates CreateEventsV1NamespacedEventUnauthorized with default headers values
func NewCreateEventsV1NamespacedEventUnauthorized() *CreateEventsV1NamespacedEventUnauthorized {

	return &CreateEventsV1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *CreateEventsV1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}