// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchNetworkingV1IngressClassOKCode is the HTTP code returned for type PatchNetworkingV1IngressClassOK
const PatchNetworkingV1IngressClassOKCode int = 200

/*PatchNetworkingV1IngressClassOK OK

swagger:response patchNetworkingV1IngressClassOK
*/
type PatchNetworkingV1IngressClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1IngressClass `json:"body,omitempty"`
}

// NewPatchNetworkingV1IngressClassOK creates PatchNetworkingV1IngressClassOK with default headers values
func NewPatchNetworkingV1IngressClassOK() *PatchNetworkingV1IngressClassOK {

	return &PatchNetworkingV1IngressClassOK{}
}

// WithPayload adds the payload to the patch networking v1 ingress class o k response
func (o *PatchNetworkingV1IngressClassOK) WithPayload(payload *models.IoK8sAPINetworkingV1IngressClass) *PatchNetworkingV1IngressClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch networking v1 ingress class o k response
func (o *PatchNetworkingV1IngressClassOK) SetPayload(payload *models.IoK8sAPINetworkingV1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchNetworkingV1IngressClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchNetworkingV1IngressClassUnauthorizedCode is the HTTP code returned for type PatchNetworkingV1IngressClassUnauthorized
const PatchNetworkingV1IngressClassUnauthorizedCode int = 401

/*PatchNetworkingV1IngressClassUnauthorized Unauthorized

swagger:response patchNetworkingV1IngressClassUnauthorized
*/
type PatchNetworkingV1IngressClassUnauthorized struct {
}

// NewPatchNetworkingV1IngressClassUnauthorized creates PatchNetworkingV1IngressClassUnauthorized with default headers values
func NewPatchNetworkingV1IngressClassUnauthorized() *PatchNetworkingV1IngressClassUnauthorized {

	return &PatchNetworkingV1IngressClassUnauthorized{}
}

// WriteResponse to the client
func (o *PatchNetworkingV1IngressClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
