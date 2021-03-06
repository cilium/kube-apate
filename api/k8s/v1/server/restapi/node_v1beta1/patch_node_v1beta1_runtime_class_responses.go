// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchNodeV1beta1RuntimeClassOKCode is the HTTP code returned for type PatchNodeV1beta1RuntimeClassOK
const PatchNodeV1beta1RuntimeClassOKCode int = 200

/*PatchNodeV1beta1RuntimeClassOK OK

swagger:response patchNodeV1beta1RuntimeClassOK
*/
type PatchNodeV1beta1RuntimeClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINodeV1beta1RuntimeClass `json:"body,omitempty"`
}

// NewPatchNodeV1beta1RuntimeClassOK creates PatchNodeV1beta1RuntimeClassOK with default headers values
func NewPatchNodeV1beta1RuntimeClassOK() *PatchNodeV1beta1RuntimeClassOK {

	return &PatchNodeV1beta1RuntimeClassOK{}
}

// WithPayload adds the payload to the patch node v1beta1 runtime class o k response
func (o *PatchNodeV1beta1RuntimeClassOK) WithPayload(payload *models.IoK8sAPINodeV1beta1RuntimeClass) *PatchNodeV1beta1RuntimeClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch node v1beta1 runtime class o k response
func (o *PatchNodeV1beta1RuntimeClassOK) SetPayload(payload *models.IoK8sAPINodeV1beta1RuntimeClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchNodeV1beta1RuntimeClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchNodeV1beta1RuntimeClassUnauthorizedCode is the HTTP code returned for type PatchNodeV1beta1RuntimeClassUnauthorized
const PatchNodeV1beta1RuntimeClassUnauthorizedCode int = 401

/*PatchNodeV1beta1RuntimeClassUnauthorized Unauthorized

swagger:response patchNodeV1beta1RuntimeClassUnauthorized
*/
type PatchNodeV1beta1RuntimeClassUnauthorized struct {
}

// NewPatchNodeV1beta1RuntimeClassUnauthorized creates PatchNodeV1beta1RuntimeClassUnauthorized with default headers values
func NewPatchNodeV1beta1RuntimeClassUnauthorized() *PatchNodeV1beta1RuntimeClassUnauthorized {

	return &PatchNodeV1beta1RuntimeClassUnauthorized{}
}

// WriteResponse to the client
func (o *PatchNodeV1beta1RuntimeClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
