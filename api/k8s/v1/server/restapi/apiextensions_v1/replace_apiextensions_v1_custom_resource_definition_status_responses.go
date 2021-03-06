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

// ReplaceApiextensionsV1CustomResourceDefinitionStatusOKCode is the HTTP code returned for type ReplaceApiextensionsV1CustomResourceDefinitionStatusOK
const ReplaceApiextensionsV1CustomResourceDefinitionStatusOKCode int = 200

/*ReplaceApiextensionsV1CustomResourceDefinitionStatusOK OK

swagger:response replaceApiextensionsV1CustomResourceDefinitionStatusOK
*/
type ReplaceApiextensionsV1CustomResourceDefinitionStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewReplaceApiextensionsV1CustomResourceDefinitionStatusOK creates ReplaceApiextensionsV1CustomResourceDefinitionStatusOK with default headers values
func NewReplaceApiextensionsV1CustomResourceDefinitionStatusOK() *ReplaceApiextensionsV1CustomResourceDefinitionStatusOK {

	return &ReplaceApiextensionsV1CustomResourceDefinitionStatusOK{}
}

// WithPayload adds the payload to the replace apiextensions v1 custom resource definition status o k response
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *ReplaceApiextensionsV1CustomResourceDefinitionStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apiextensions v1 custom resource definition status o k response
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceApiextensionsV1CustomResourceDefinitionStatusCreatedCode is the HTTP code returned for type ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated
const ReplaceApiextensionsV1CustomResourceDefinitionStatusCreatedCode int = 201

/*ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated Created

swagger:response replaceApiextensionsV1CustomResourceDefinitionStatusCreated
*/
type ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewReplaceApiextensionsV1CustomResourceDefinitionStatusCreated creates ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated with default headers values
func NewReplaceApiextensionsV1CustomResourceDefinitionStatusCreated() *ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated {

	return &ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated{}
}

// WithPayload adds the payload to the replace apiextensions v1 custom resource definition status created response
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apiextensions v1 custom resource definition status created response
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorizedCode is the HTTP code returned for type ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized
const ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorizedCode int = 401

/*ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized Unauthorized

swagger:response replaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized
*/
type ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized struct {
}

// NewReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized creates ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized with default headers values
func NewReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized() *ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized {

	return &ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceApiextensionsV1CustomResourceDefinitionStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
