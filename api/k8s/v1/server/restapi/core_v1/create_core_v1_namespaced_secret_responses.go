// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateCoreV1NamespacedSecretOKCode is the HTTP code returned for type CreateCoreV1NamespacedSecretOK
const CreateCoreV1NamespacedSecretOKCode int = 200

/*CreateCoreV1NamespacedSecretOK OK

swagger:response createCoreV1NamespacedSecretOK
*/
type CreateCoreV1NamespacedSecretOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Secret `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedSecretOK creates CreateCoreV1NamespacedSecretOK with default headers values
func NewCreateCoreV1NamespacedSecretOK() *CreateCoreV1NamespacedSecretOK {

	return &CreateCoreV1NamespacedSecretOK{}
}

// WithPayload adds the payload to the create core v1 namespaced secret o k response
func (o *CreateCoreV1NamespacedSecretOK) WithPayload(payload *models.IoK8sAPICoreV1Secret) *CreateCoreV1NamespacedSecretOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced secret o k response
func (o *CreateCoreV1NamespacedSecretOK) SetPayload(payload *models.IoK8sAPICoreV1Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedSecretOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedSecretCreatedCode is the HTTP code returned for type CreateCoreV1NamespacedSecretCreated
const CreateCoreV1NamespacedSecretCreatedCode int = 201

/*CreateCoreV1NamespacedSecretCreated Created

swagger:response createCoreV1NamespacedSecretCreated
*/
type CreateCoreV1NamespacedSecretCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Secret `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedSecretCreated creates CreateCoreV1NamespacedSecretCreated with default headers values
func NewCreateCoreV1NamespacedSecretCreated() *CreateCoreV1NamespacedSecretCreated {

	return &CreateCoreV1NamespacedSecretCreated{}
}

// WithPayload adds the payload to the create core v1 namespaced secret created response
func (o *CreateCoreV1NamespacedSecretCreated) WithPayload(payload *models.IoK8sAPICoreV1Secret) *CreateCoreV1NamespacedSecretCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced secret created response
func (o *CreateCoreV1NamespacedSecretCreated) SetPayload(payload *models.IoK8sAPICoreV1Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedSecretCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedSecretAcceptedCode is the HTTP code returned for type CreateCoreV1NamespacedSecretAccepted
const CreateCoreV1NamespacedSecretAcceptedCode int = 202

/*CreateCoreV1NamespacedSecretAccepted Accepted

swagger:response createCoreV1NamespacedSecretAccepted
*/
type CreateCoreV1NamespacedSecretAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Secret `json:"body,omitempty"`
}

// NewCreateCoreV1NamespacedSecretAccepted creates CreateCoreV1NamespacedSecretAccepted with default headers values
func NewCreateCoreV1NamespacedSecretAccepted() *CreateCoreV1NamespacedSecretAccepted {

	return &CreateCoreV1NamespacedSecretAccepted{}
}

// WithPayload adds the payload to the create core v1 namespaced secret accepted response
func (o *CreateCoreV1NamespacedSecretAccepted) WithPayload(payload *models.IoK8sAPICoreV1Secret) *CreateCoreV1NamespacedSecretAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create core v1 namespaced secret accepted response
func (o *CreateCoreV1NamespacedSecretAccepted) SetPayload(payload *models.IoK8sAPICoreV1Secret) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedSecretAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateCoreV1NamespacedSecretUnauthorizedCode is the HTTP code returned for type CreateCoreV1NamespacedSecretUnauthorized
const CreateCoreV1NamespacedSecretUnauthorizedCode int = 401

/*CreateCoreV1NamespacedSecretUnauthorized Unauthorized

swagger:response createCoreV1NamespacedSecretUnauthorized
*/
type CreateCoreV1NamespacedSecretUnauthorized struct {
}

// NewCreateCoreV1NamespacedSecretUnauthorized creates CreateCoreV1NamespacedSecretUnauthorized with default headers values
func NewCreateCoreV1NamespacedSecretUnauthorized() *CreateCoreV1NamespacedSecretUnauthorized {

	return &CreateCoreV1NamespacedSecretUnauthorized{}
}

// WriteResponse to the client
func (o *CreateCoreV1NamespacedSecretUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}