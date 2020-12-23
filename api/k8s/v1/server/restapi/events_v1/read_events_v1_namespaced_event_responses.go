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

// ReadEventsV1NamespacedEventOKCode is the HTTP code returned for type ReadEventsV1NamespacedEventOK
const ReadEventsV1NamespacedEventOKCode int = 200

/*ReadEventsV1NamespacedEventOK OK

swagger:response readEventsV1NamespacedEventOK
*/
type ReadEventsV1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1Event `json:"body,omitempty"`
}

// NewReadEventsV1NamespacedEventOK creates ReadEventsV1NamespacedEventOK with default headers values
func NewReadEventsV1NamespacedEventOK() *ReadEventsV1NamespacedEventOK {

	return &ReadEventsV1NamespacedEventOK{}
}

// WithPayload adds the payload to the read events v1 namespaced event o k response
func (o *ReadEventsV1NamespacedEventOK) WithPayload(payload *models.IoK8sAPIEventsV1Event) *ReadEventsV1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read events v1 namespaced event o k response
func (o *ReadEventsV1NamespacedEventOK) SetPayload(payload *models.IoK8sAPIEventsV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadEventsV1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadEventsV1NamespacedEventUnauthorizedCode is the HTTP code returned for type ReadEventsV1NamespacedEventUnauthorized
const ReadEventsV1NamespacedEventUnauthorizedCode int = 401

/*ReadEventsV1NamespacedEventUnauthorized Unauthorized

swagger:response readEventsV1NamespacedEventUnauthorized
*/
type ReadEventsV1NamespacedEventUnauthorized struct {
}

// NewReadEventsV1NamespacedEventUnauthorized creates ReadEventsV1NamespacedEventUnauthorized with default headers values
func NewReadEventsV1NamespacedEventUnauthorized() *ReadEventsV1NamespacedEventUnauthorized {

	return &ReadEventsV1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *ReadEventsV1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}