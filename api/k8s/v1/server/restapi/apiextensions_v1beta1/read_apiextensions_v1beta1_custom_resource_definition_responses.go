// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadApiextensionsV1beta1CustomResourceDefinitionOKCode is the HTTP code returned for type ReadApiextensionsV1beta1CustomResourceDefinitionOK
const ReadApiextensionsV1beta1CustomResourceDefinitionOKCode int = 200

/*ReadApiextensionsV1beta1CustomResourceDefinitionOK OK

swagger:response readApiextensionsV1beta1CustomResourceDefinitionOK
*/
type ReadApiextensionsV1beta1CustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition `json:"body,omitempty"`
}

// NewReadApiextensionsV1beta1CustomResourceDefinitionOK creates ReadApiextensionsV1beta1CustomResourceDefinitionOK with default headers values
func NewReadApiextensionsV1beta1CustomResourceDefinitionOK() *ReadApiextensionsV1beta1CustomResourceDefinitionOK {

	return &ReadApiextensionsV1beta1CustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the read apiextensions v1beta1 custom resource definition o k response
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) *ReadApiextensionsV1beta1CustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apiextensions v1beta1 custom resource definition o k response
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized
const ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorizedCode int = 401

/*ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized Unauthorized

swagger:response readApiextensionsV1beta1CustomResourceDefinitionUnauthorized
*/
type ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized struct {
}

// NewReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized creates ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized with default headers values
func NewReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized() *ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized {

	return &ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}