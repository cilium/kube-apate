// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAppsV1NamespacedDaemonSetOKCode is the HTTP code returned for type CreateAppsV1NamespacedDaemonSetOK
const CreateAppsV1NamespacedDaemonSetOKCode int = 200

/*CreateAppsV1NamespacedDaemonSetOK OK

swagger:response createAppsV1NamespacedDaemonSetOK
*/
type CreateAppsV1NamespacedDaemonSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedDaemonSetOK creates CreateAppsV1NamespacedDaemonSetOK with default headers values
func NewCreateAppsV1NamespacedDaemonSetOK() *CreateAppsV1NamespacedDaemonSetOK {

	return &CreateAppsV1NamespacedDaemonSetOK{}
}

// WithPayload adds the payload to the create apps v1 namespaced daemon set o k response
func (o *CreateAppsV1NamespacedDaemonSetOK) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *CreateAppsV1NamespacedDaemonSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced daemon set o k response
func (o *CreateAppsV1NamespacedDaemonSetOK) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDaemonSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedDaemonSetCreatedCode is the HTTP code returned for type CreateAppsV1NamespacedDaemonSetCreated
const CreateAppsV1NamespacedDaemonSetCreatedCode int = 201

/*CreateAppsV1NamespacedDaemonSetCreated Created

swagger:response createAppsV1NamespacedDaemonSetCreated
*/
type CreateAppsV1NamespacedDaemonSetCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedDaemonSetCreated creates CreateAppsV1NamespacedDaemonSetCreated with default headers values
func NewCreateAppsV1NamespacedDaemonSetCreated() *CreateAppsV1NamespacedDaemonSetCreated {

	return &CreateAppsV1NamespacedDaemonSetCreated{}
}

// WithPayload adds the payload to the create apps v1 namespaced daemon set created response
func (o *CreateAppsV1NamespacedDaemonSetCreated) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *CreateAppsV1NamespacedDaemonSetCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced daemon set created response
func (o *CreateAppsV1NamespacedDaemonSetCreated) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDaemonSetCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedDaemonSetAcceptedCode is the HTTP code returned for type CreateAppsV1NamespacedDaemonSetAccepted
const CreateAppsV1NamespacedDaemonSetAcceptedCode int = 202

/*CreateAppsV1NamespacedDaemonSetAccepted Accepted

swagger:response createAppsV1NamespacedDaemonSetAccepted
*/
type CreateAppsV1NamespacedDaemonSetAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewCreateAppsV1NamespacedDaemonSetAccepted creates CreateAppsV1NamespacedDaemonSetAccepted with default headers values
func NewCreateAppsV1NamespacedDaemonSetAccepted() *CreateAppsV1NamespacedDaemonSetAccepted {

	return &CreateAppsV1NamespacedDaemonSetAccepted{}
}

// WithPayload adds the payload to the create apps v1 namespaced daemon set accepted response
func (o *CreateAppsV1NamespacedDaemonSetAccepted) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *CreateAppsV1NamespacedDaemonSetAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apps v1 namespaced daemon set accepted response
func (o *CreateAppsV1NamespacedDaemonSetAccepted) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDaemonSetAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAppsV1NamespacedDaemonSetUnauthorizedCode is the HTTP code returned for type CreateAppsV1NamespacedDaemonSetUnauthorized
const CreateAppsV1NamespacedDaemonSetUnauthorizedCode int = 401

/*CreateAppsV1NamespacedDaemonSetUnauthorized Unauthorized

swagger:response createAppsV1NamespacedDaemonSetUnauthorized
*/
type CreateAppsV1NamespacedDaemonSetUnauthorized struct {
}

// NewCreateAppsV1NamespacedDaemonSetUnauthorized creates CreateAppsV1NamespacedDaemonSetUnauthorized with default headers values
func NewCreateAppsV1NamespacedDaemonSetUnauthorized() *CreateAppsV1NamespacedDaemonSetUnauthorized {

	return &CreateAppsV1NamespacedDaemonSetUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAppsV1NamespacedDaemonSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
