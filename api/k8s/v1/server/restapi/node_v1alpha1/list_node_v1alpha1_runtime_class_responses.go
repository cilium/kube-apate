// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListNodeV1alpha1RuntimeClassOKCode is the HTTP code returned for type ListNodeV1alpha1RuntimeClassOK
const ListNodeV1alpha1RuntimeClassOKCode int = 200

/*ListNodeV1alpha1RuntimeClassOK OK

swagger:response listNodeV1alpha1RuntimeClassOK
*/
type ListNodeV1alpha1RuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClassList `json:"body,omitempty"`
}

// NewListNodeV1alpha1RuntimeClassOK creates ListNodeV1alpha1RuntimeClassOK with default headers values
func NewListNodeV1alpha1RuntimeClassOK() *ListNodeV1alpha1RuntimeClassOK {

	return &ListNodeV1alpha1RuntimeClassOK{}
}

// WithPayload adds the payload to the list node v1alpha1 runtime class o k response
func (o *ListNodeV1alpha1RuntimeClassOK) WithPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClassList) *ListNodeV1alpha1RuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list node v1alpha1 runtime class o k response
func (o *ListNodeV1alpha1RuntimeClassOK) SetPayload(payload *models.IoK8sAPINodeV1alpha1RuntimeClassList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListNodeV1alpha1RuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListNodeV1alpha1RuntimeClassUnauthorizedCode is the HTTP code returned for type ListNodeV1alpha1RuntimeClassUnauthorized
const ListNodeV1alpha1RuntimeClassUnauthorizedCode int = 401

/*ListNodeV1alpha1RuntimeClassUnauthorized Unauthorized

swagger:response listNodeV1alpha1RuntimeClassUnauthorized
*/
type ListNodeV1alpha1RuntimeClassUnauthorized struct {
}

// NewListNodeV1alpha1RuntimeClassUnauthorized creates ListNodeV1alpha1RuntimeClassUnauthorized with default headers values
func NewListNodeV1alpha1RuntimeClassUnauthorized() *ListNodeV1alpha1RuntimeClassUnauthorized {

	return &ListNodeV1alpha1RuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *ListNodeV1alpha1RuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
