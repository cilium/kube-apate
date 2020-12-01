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

// DeleteApiextensionsV1beta1CustomResourceDefinitionOKCode is the HTTP code returned for type DeleteApiextensionsV1beta1CustomResourceDefinitionOK
const DeleteApiextensionsV1beta1CustomResourceDefinitionOKCode int = 200

/*DeleteApiextensionsV1beta1CustomResourceDefinitionOK OK

swagger:response deleteApiextensionsV1beta1CustomResourceDefinitionOK
*/
type DeleteApiextensionsV1beta1CustomResourceDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteApiextensionsV1beta1CustomResourceDefinitionOK creates DeleteApiextensionsV1beta1CustomResourceDefinitionOK with default headers values
func NewDeleteApiextensionsV1beta1CustomResourceDefinitionOK() *DeleteApiextensionsV1beta1CustomResourceDefinitionOK {

	return &DeleteApiextensionsV1beta1CustomResourceDefinitionOK{}
}

// WithPayload adds the payload to the delete apiextensions v1beta1 custom resource definition o k response
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteApiextensionsV1beta1CustomResourceDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apiextensions v1beta1 custom resource definition o k response
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApiextensionsV1beta1CustomResourceDefinitionAcceptedCode is the HTTP code returned for type DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted
const DeleteApiextensionsV1beta1CustomResourceDefinitionAcceptedCode int = 202

/*DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted Accepted

swagger:response deleteApiextensionsV1beta1CustomResourceDefinitionAccepted
*/
type DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteApiextensionsV1beta1CustomResourceDefinitionAccepted creates DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted with default headers values
func NewDeleteApiextensionsV1beta1CustomResourceDefinitionAccepted() *DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted {

	return &DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted{}
}

// WithPayload adds the payload to the delete apiextensions v1beta1 custom resource definition accepted response
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apiextensions v1beta1 custom resource definition accepted response
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorizedCode is the HTTP code returned for type DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized
const DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorizedCode int = 401

/*DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized Unauthorized

swagger:response deleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized
*/
type DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized struct {
}

// NewDeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized creates DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized with default headers values
func NewDeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized() *DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized {

	return &DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteApiextensionsV1beta1CustomResourceDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
