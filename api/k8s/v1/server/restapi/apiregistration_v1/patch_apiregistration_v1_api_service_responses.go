// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchApiregistrationV1APIServiceOKCode is the HTTP code returned for type PatchApiregistrationV1APIServiceOK
const PatchApiregistrationV1APIServiceOKCode int = 200

/*PatchApiregistrationV1APIServiceOK OK

swagger:response patchApiregistrationV1ApiServiceOK
*/
type PatchApiregistrationV1APIServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService `json:"body,omitempty"`
}

// NewPatchApiregistrationV1APIServiceOK creates PatchApiregistrationV1APIServiceOK with default headers values
func NewPatchApiregistrationV1APIServiceOK() *PatchApiregistrationV1APIServiceOK {

	return &PatchApiregistrationV1APIServiceOK{}
}

// WithPayload adds the payload to the patch apiregistration v1 Api service o k response
func (o *PatchApiregistrationV1APIServiceOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) *PatchApiregistrationV1APIServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apiregistration v1 Api service o k response
func (o *PatchApiregistrationV1APIServiceOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchApiregistrationV1APIServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchApiregistrationV1APIServiceUnauthorizedCode is the HTTP code returned for type PatchApiregistrationV1APIServiceUnauthorized
const PatchApiregistrationV1APIServiceUnauthorizedCode int = 401

/*PatchApiregistrationV1APIServiceUnauthorized Unauthorized

swagger:response patchApiregistrationV1ApiServiceUnauthorized
*/
type PatchApiregistrationV1APIServiceUnauthorized struct {
}

// NewPatchApiregistrationV1APIServiceUnauthorized creates PatchApiregistrationV1APIServiceUnauthorized with default headers values
func NewPatchApiregistrationV1APIServiceUnauthorized() *PatchApiregistrationV1APIServiceUnauthorized {

	return &PatchApiregistrationV1APIServiceUnauthorized{}
}

// WriteResponse to the client
func (o *PatchApiregistrationV1APIServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}