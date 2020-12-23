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

// PatchApiextensionsV1beta1CustomResourceDefinitionOKCode is the HTTP code returned for type PatchApiextensionsV1beta1CustomResourceDefinitionOK
const PatchApiextensionsV1beta1CustomResourceDefinitionOKCode int = 200

/*PatchApiextensionsV1beta1CustomResourceDefinitionOK OK

swagger:response patchApiextensionsV1beta1CustomResourceDefinitionOK
*/
type PatchApiextensionsV1beta1CustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition `json:"body,omitempty"`
}

// NewPatchApiextensionsV1beta1CustomResourceDefinitionOK creates PatchApiextensionsV1beta1CustomResourceDefinitionOK with default headers values
func NewPatchApiextensionsV1beta1CustomResourceDefinitionOK() *PatchApiextensionsV1beta1CustomResourceDefinitionOK {

	return &PatchApiextensionsV1beta1CustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the patch apiextensions v1beta1 custom resource definition o k response
func (o *PatchApiextensionsV1beta1CustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) *PatchApiextensionsV1beta1CustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apiextensions v1beta1 custom resource definition o k response
func (o *PatchApiextensionsV1beta1CustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApiextensionsV1beta1CustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized
const PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorizedCode int = 401

/*PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized Unauthorized

swagger:response patchApiextensionsV1beta1CustomResourceDefinitionUnauthorized
*/
type PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized struct {
}

// NewPatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized creates PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized with default headers values
func NewPatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized() *PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized {

	return &PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *PatchApiextensionsV1beta1CustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}