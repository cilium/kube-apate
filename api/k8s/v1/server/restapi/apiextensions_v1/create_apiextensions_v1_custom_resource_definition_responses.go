// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateApiextensionsV1CustomResourceDefinitionOKCode is the HTTP code returned for type CreateApiextensionsV1CustomResourceDefinitionOK
const CreateApiextensionsV1CustomResourceDefinitionOKCode int = 200

/*CreateApiextensionsV1CustomResourceDefinitionOK OK

swagger:response createApiextensionsV1CustomResourceDefinitionOK
*/
type CreateApiextensionsV1CustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewCreateApiextensionsV1CustomResourceDefinitionOK creates CreateApiextensionsV1CustomResourceDefinitionOK with default headers values
func NewCreateApiextensionsV1CustomResourceDefinitionOK() *CreateApiextensionsV1CustomResourceDefinitionOK {

	return &CreateApiextensionsV1CustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the create apiextensions v1 custom resource definition o k response
func (o *CreateApiextensionsV1CustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *CreateApiextensionsV1CustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apiextensions v1 custom resource definition o k response
func (o *CreateApiextensionsV1CustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApiextensionsV1CustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApiextensionsV1CustomResourceDefinitionCreatedCode is the HTTP code returned for type CreateApiextensionsV1CustomResourceDefinitionCreated
const CreateApiextensionsV1CustomResourceDefinitionCreatedCode int = 201

/*CreateApiextensionsV1CustomResourceDefinitionCreated Created

swagger:response createApiextensionsV1CustomResourceDefinitionCreated
*/
type CreateApiextensionsV1CustomResourceDefinitionCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewCreateApiextensionsV1CustomResourceDefinitionCreated creates CreateApiextensionsV1CustomResourceDefinitionCreated with default headers values
func NewCreateApiextensionsV1CustomResourceDefinitionCreated() *CreateApiextensionsV1CustomResourceDefinitionCreated {

	return &CreateApiextensionsV1CustomResourceDefinitionCreated{}
}

// WithPayload adds the payload to the create apiextensions v1 custom resource definition created response
func (o *CreateApiextensionsV1CustomResourceDefinitionCreated) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *CreateApiextensionsV1CustomResourceDefinitionCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apiextensions v1 custom resource definition created response
func (o *CreateApiextensionsV1CustomResourceDefinitionCreated) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApiextensionsV1CustomResourceDefinitionCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApiextensionsV1CustomResourceDefinitionAcceptedCode is the HTTP code returned for type CreateApiextensionsV1CustomResourceDefinitionAccepted
const CreateApiextensionsV1CustomResourceDefinitionAcceptedCode int = 202

/*CreateApiextensionsV1CustomResourceDefinitionAccepted Accepted

swagger:response createApiextensionsV1CustomResourceDefinitionAccepted
*/
type CreateApiextensionsV1CustomResourceDefinitionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewCreateApiextensionsV1CustomResourceDefinitionAccepted creates CreateApiextensionsV1CustomResourceDefinitionAccepted with default headers values
func NewCreateApiextensionsV1CustomResourceDefinitionAccepted() *CreateApiextensionsV1CustomResourceDefinitionAccepted {

	return &CreateApiextensionsV1CustomResourceDefinitionAccepted{}
}

// WithPayload adds the payload to the create apiextensions v1 custom resource definition accepted response
func (o *CreateApiextensionsV1CustomResourceDefinitionAccepted) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *CreateApiextensionsV1CustomResourceDefinitionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create apiextensions v1 custom resource definition accepted response
func (o *CreateApiextensionsV1CustomResourceDefinitionAccepted) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateApiextensionsV1CustomResourceDefinitionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateApiextensionsV1CustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type CreateApiextensionsV1CustomResourceDefinitionUnauthorized
const CreateApiextensionsV1CustomResourceDefinitionUnauthorizedCode int = 401

/*CreateApiextensionsV1CustomResourceDefinitionUnauthorized Unauthorized

swagger:response createApiextensionsV1CustomResourceDefinitionUnauthorized
*/
type CreateApiextensionsV1CustomResourceDefinitionUnauthorized struct {
}

// NewCreateApiextensionsV1CustomResourceDefinitionUnauthorized creates CreateApiextensionsV1CustomResourceDefinitionUnauthorized with default headers values
func NewCreateApiextensionsV1CustomResourceDefinitionUnauthorized() *CreateApiextensionsV1CustomResourceDefinitionUnauthorized {

	return &CreateApiextensionsV1CustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *CreateApiextensionsV1CustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
