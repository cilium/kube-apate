// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceStorageV1beta1CSINodeOKCode is the HTTP code returned for type ReplaceStorageV1beta1CSINodeOK
const ReplaceStorageV1beta1CSINodeOKCode int = 200

/*ReplaceStorageV1beta1CSINodeOK OK

swagger:response replaceStorageV1beta1CSINodeOK
*/
type ReplaceStorageV1beta1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSINode `json:"body,omitempty"`
}

// NewReplaceStorageV1beta1CSINodeOK creates ReplaceStorageV1beta1CSINodeOK with default headers values
func NewReplaceStorageV1beta1CSINodeOK() *ReplaceStorageV1beta1CSINodeOK {

	return &ReplaceStorageV1beta1CSINodeOK{}
}

// WithPayload adds the payload to the replace storage v1beta1 c s i node o k response
func (o *ReplaceStorageV1beta1CSINodeOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) *ReplaceStorageV1beta1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1beta1 c s i node o k response
func (o *ReplaceStorageV1beta1CSINodeOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1beta1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1beta1CSINodeCreatedCode is the HTTP code returned for type ReplaceStorageV1beta1CSINodeCreated
const ReplaceStorageV1beta1CSINodeCreatedCode int = 201

/*ReplaceStorageV1beta1CSINodeCreated Created

swagger:response replaceStorageV1beta1CSINodeCreated
*/
type ReplaceStorageV1beta1CSINodeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSINode `json:"body,omitempty"`
}

// NewReplaceStorageV1beta1CSINodeCreated creates ReplaceStorageV1beta1CSINodeCreated with default headers values
func NewReplaceStorageV1beta1CSINodeCreated() *ReplaceStorageV1beta1CSINodeCreated {

	return &ReplaceStorageV1beta1CSINodeCreated{}
}

// WithPayload adds the payload to the replace storage v1beta1 c s i node created response
func (o *ReplaceStorageV1beta1CSINodeCreated) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) *ReplaceStorageV1beta1CSINodeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace storage v1beta1 c s i node created response
func (o *ReplaceStorageV1beta1CSINodeCreated) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceStorageV1beta1CSINodeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceStorageV1beta1CSINodeUnauthorizedCode is the HTTP code returned for type ReplaceStorageV1beta1CSINodeUnauthorized
const ReplaceStorageV1beta1CSINodeUnauthorizedCode int = 401

/*ReplaceStorageV1beta1CSINodeUnauthorized Unauthorized

swagger:response replaceStorageV1beta1CSINodeUnauthorized
*/
type ReplaceStorageV1beta1CSINodeUnauthorized struct {
}

// NewReplaceStorageV1beta1CSINodeUnauthorized creates ReplaceStorageV1beta1CSINodeUnauthorized with default headers values
func NewReplaceStorageV1beta1CSINodeUnauthorized() *ReplaceStorageV1beta1CSINodeUnauthorized {

	return &ReplaceStorageV1beta1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceStorageV1beta1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
