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

// ListEventsV1beta1EventForAllNamespacesOKCode is the HTTP code returned for type ListEventsV1beta1EventForAllNamespacesOK
const ListEventsV1beta1EventForAllNamespacesOKCode int = 200

/*ListEventsV1beta1EventForAllNamespacesOK OK

swagger:response listEventsV1beta1EventForAllNamespacesOK
*/
type ListEventsV1beta1EventForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIEventsV1beta1EventList `json:"body,omitempty"`
}

// NewListEventsV1beta1EventForAllNamespacesOK creates ListEventsV1beta1EventForAllNamespacesOK with default headers values
func NewListEventsV1beta1EventForAllNamespacesOK() *ListEventsV1beta1EventForAllNamespacesOK {

	return &ListEventsV1beta1EventForAllNamespacesOK{}
}

// WithPayload adds the payload to the list events v1beta1 event for all namespaces o k response
func (o *ListEventsV1beta1EventForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIEventsV1beta1EventList) *ListEventsV1beta1EventForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list events v1beta1 event for all namespaces o k response
func (o *ListEventsV1beta1EventForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIEventsV1beta1EventList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListEventsV1beta1EventForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListEventsV1beta1EventForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListEventsV1beta1EventForAllNamespacesUnauthorized
const ListEventsV1beta1EventForAllNamespacesUnauthorizedCode int = 401

/*ListEventsV1beta1EventForAllNamespacesUnauthorized Unauthorized

swagger:response listEventsV1beta1EventForAllNamespacesUnauthorized
*/
type ListEventsV1beta1EventForAllNamespacesUnauthorized struct {
}

// NewListEventsV1beta1EventForAllNamespacesUnauthorized creates ListEventsV1beta1EventForAllNamespacesUnauthorized with default headers values
func NewListEventsV1beta1EventForAllNamespacesUnauthorized() *ListEventsV1beta1EventForAllNamespacesUnauthorized {

	return &ListEventsV1beta1EventForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListEventsV1beta1EventForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}