// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceEventsV1beta1NamespacedEventOKCode is the HTTP code returned for type ReplaceEventsV1beta1NamespacedEventOK
const ReplaceEventsV1beta1NamespacedEventOKCode int = 200

/*ReplaceEventsV1beta1NamespacedEventOK OK

swagger:response replaceEventsV1beta1NamespacedEventOK
*/
type ReplaceEventsV1beta1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1Event `json:"body,omitempty"`
}

// NewReplaceEventsV1beta1NamespacedEventOK creates ReplaceEventsV1beta1NamespacedEventOK with default headers values
func NewReplaceEventsV1beta1NamespacedEventOK() *ReplaceEventsV1beta1NamespacedEventOK {

	return &ReplaceEventsV1beta1NamespacedEventOK{}
}

// WithPayload adds the payload to the replace events v1beta1 namespaced event o k response
func (o *ReplaceEventsV1beta1NamespacedEventOK) WithPayload(payload *models.IoK8sAPIEventsV1beta1Event) *ReplaceEventsV1beta1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace events v1beta1 namespaced event o k response
func (o *ReplaceEventsV1beta1NamespacedEventOK) SetPayload(payload *models.IoK8sAPIEventsV1beta1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceEventsV1beta1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceEventsV1beta1NamespacedEventCreatedCode is the HTTP code returned for type ReplaceEventsV1beta1NamespacedEventCreated
const ReplaceEventsV1beta1NamespacedEventCreatedCode int = 201

/*ReplaceEventsV1beta1NamespacedEventCreated Created

swagger:response replaceEventsV1beta1NamespacedEventCreated
*/
type ReplaceEventsV1beta1NamespacedEventCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1Event `json:"body,omitempty"`
}

// NewReplaceEventsV1beta1NamespacedEventCreated creates ReplaceEventsV1beta1NamespacedEventCreated with default headers values
func NewReplaceEventsV1beta1NamespacedEventCreated() *ReplaceEventsV1beta1NamespacedEventCreated {

	return &ReplaceEventsV1beta1NamespacedEventCreated{}
}

// WithPayload adds the payload to the replace events v1beta1 namespaced event created response
func (o *ReplaceEventsV1beta1NamespacedEventCreated) WithPayload(payload *models.IoK8sAPIEventsV1beta1Event) *ReplaceEventsV1beta1NamespacedEventCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace events v1beta1 namespaced event created response
func (o *ReplaceEventsV1beta1NamespacedEventCreated) SetPayload(payload *models.IoK8sAPIEventsV1beta1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceEventsV1beta1NamespacedEventCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceEventsV1beta1NamespacedEventUnauthorizedCode is the HTTP code returned for type ReplaceEventsV1beta1NamespacedEventUnauthorized
const ReplaceEventsV1beta1NamespacedEventUnauthorizedCode int = 401

/*ReplaceEventsV1beta1NamespacedEventUnauthorized Unauthorized

swagger:response replaceEventsV1beta1NamespacedEventUnauthorized
*/
type ReplaceEventsV1beta1NamespacedEventUnauthorized struct {
}

// NewReplaceEventsV1beta1NamespacedEventUnauthorized creates ReplaceEventsV1beta1NamespacedEventUnauthorized with default headers values
func NewReplaceEventsV1beta1NamespacedEventUnauthorized() *ReplaceEventsV1beta1NamespacedEventUnauthorized {

	return &ReplaceEventsV1beta1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceEventsV1beta1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
