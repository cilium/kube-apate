// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateExtensionsV1beta1NamespacedIngressOKCode is the HTTP code returned for type CreateExtensionsV1beta1NamespacedIngressOK
const CreateExtensionsV1beta1NamespacedIngressOKCode int = 200

/*CreateExtensionsV1beta1NamespacedIngressOK OK

swagger:response createExtensionsV1beta1NamespacedIngressOK
*/
type CreateExtensionsV1beta1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewCreateExtensionsV1beta1NamespacedIngressOK creates CreateExtensionsV1beta1NamespacedIngressOK with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressOK() *CreateExtensionsV1beta1NamespacedIngressOK {

	return &CreateExtensionsV1beta1NamespacedIngressOK{}
}

// WithPayload adds the payload to the create extensions v1beta1 namespaced ingress o k response
func (o *CreateExtensionsV1beta1NamespacedIngressOK) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *CreateExtensionsV1beta1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create extensions v1beta1 namespaced ingress o k response
func (o *CreateExtensionsV1beta1NamespacedIngressOK) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExtensionsV1beta1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExtensionsV1beta1NamespacedIngressCreatedCode is the HTTP code returned for type CreateExtensionsV1beta1NamespacedIngressCreated
const CreateExtensionsV1beta1NamespacedIngressCreatedCode int = 201

/*CreateExtensionsV1beta1NamespacedIngressCreated Created

swagger:response createExtensionsV1beta1NamespacedIngressCreated
*/
type CreateExtensionsV1beta1NamespacedIngressCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewCreateExtensionsV1beta1NamespacedIngressCreated creates CreateExtensionsV1beta1NamespacedIngressCreated with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressCreated() *CreateExtensionsV1beta1NamespacedIngressCreated {

	return &CreateExtensionsV1beta1NamespacedIngressCreated{}
}

// WithPayload adds the payload to the create extensions v1beta1 namespaced ingress created response
func (o *CreateExtensionsV1beta1NamespacedIngressCreated) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *CreateExtensionsV1beta1NamespacedIngressCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create extensions v1beta1 namespaced ingress created response
func (o *CreateExtensionsV1beta1NamespacedIngressCreated) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExtensionsV1beta1NamespacedIngressCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExtensionsV1beta1NamespacedIngressAcceptedCode is the HTTP code returned for type CreateExtensionsV1beta1NamespacedIngressAccepted
const CreateExtensionsV1beta1NamespacedIngressAcceptedCode int = 202

/*CreateExtensionsV1beta1NamespacedIngressAccepted Accepted

swagger:response createExtensionsV1beta1NamespacedIngressAccepted
*/
type CreateExtensionsV1beta1NamespacedIngressAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewCreateExtensionsV1beta1NamespacedIngressAccepted creates CreateExtensionsV1beta1NamespacedIngressAccepted with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressAccepted() *CreateExtensionsV1beta1NamespacedIngressAccepted {

	return &CreateExtensionsV1beta1NamespacedIngressAccepted{}
}

// WithPayload adds the payload to the create extensions v1beta1 namespaced ingress accepted response
func (o *CreateExtensionsV1beta1NamespacedIngressAccepted) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *CreateExtensionsV1beta1NamespacedIngressAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create extensions v1beta1 namespaced ingress accepted response
func (o *CreateExtensionsV1beta1NamespacedIngressAccepted) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateExtensionsV1beta1NamespacedIngressAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateExtensionsV1beta1NamespacedIngressUnauthorizedCode is the HTTP code returned for type CreateExtensionsV1beta1NamespacedIngressUnauthorized
const CreateExtensionsV1beta1NamespacedIngressUnauthorizedCode int = 401

/*CreateExtensionsV1beta1NamespacedIngressUnauthorized Unauthorized

swagger:response createExtensionsV1beta1NamespacedIngressUnauthorized
*/
type CreateExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

// NewCreateExtensionsV1beta1NamespacedIngressUnauthorized creates CreateExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressUnauthorized() *CreateExtensionsV1beta1NamespacedIngressUnauthorized {

	return &CreateExtensionsV1beta1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *CreateExtensionsV1beta1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
