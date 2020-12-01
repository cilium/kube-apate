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

// ReadApiextensionsV1CustomResourceDefinitionStatusOKCode is the HTTP code returned for type ReadApiextensionsV1CustomResourceDefinitionStatusOK
const ReadApiextensionsV1CustomResourceDefinitionStatusOKCode int = 200

/*ReadApiextensionsV1CustomResourceDefinitionStatusOK OK

swagger:response readApiextensionsV1CustomResourceDefinitionStatusOK
*/
type ReadApiextensionsV1CustomResourceDefinitionStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition `json:"body,omitempty"`
}

// NewReadApiextensionsV1CustomResourceDefinitionStatusOK creates ReadApiextensionsV1CustomResourceDefinitionStatusOK with default headers values
func NewReadApiextensionsV1CustomResourceDefinitionStatusOK() *ReadApiextensionsV1CustomResourceDefinitionStatusOK {

	return &ReadApiextensionsV1CustomResourceDefinitionStatusOK{}
}

// WithPayload adds the payload to the read apiextensions v1 custom resource definition status o k response
func (o *ReadApiextensionsV1CustomResourceDefinitionStatusOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) *ReadApiextensionsV1CustomResourceDefinitionStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apiextensions v1 custom resource definition status o k response
func (o *ReadApiextensionsV1CustomResourceDefinitionStatusOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadApiextensionsV1CustomResourceDefinitionStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorizedCode is the HTTP code returned for type ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized
const ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorizedCode int = 401

/*ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized Unauthorized

swagger:response readApiextensionsV1CustomResourceDefinitionStatusUnauthorized
*/
type ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized struct {
}

// NewReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized creates ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized with default headers values
func NewReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized() *ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized {

	return &ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadApiextensionsV1CustomResourceDefinitionStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
