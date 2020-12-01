// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchCoreV1NamespacedEventOKCode is the HTTP code returned for type PatchCoreV1NamespacedEventOK
const PatchCoreV1NamespacedEventOKCode int = 200

/*PatchCoreV1NamespacedEventOK OK

swagger:response patchCoreV1NamespacedEventOK
*/
type PatchCoreV1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Event `json:"body,omitempty"`
}

// NewPatchCoreV1NamespacedEventOK creates PatchCoreV1NamespacedEventOK with default headers values
func NewPatchCoreV1NamespacedEventOK() *PatchCoreV1NamespacedEventOK {

	return &PatchCoreV1NamespacedEventOK{}
}

// WithPayload adds the payload to the patch core v1 namespaced event o k response
func (o *PatchCoreV1NamespacedEventOK) WithPayload(payload *models.IoK8sAPICoreV1Event) *PatchCoreV1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch core v1 namespaced event o k response
func (o *PatchCoreV1NamespacedEventOK) SetPayload(payload *models.IoK8sAPICoreV1Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCoreV1NamespacedEventUnauthorizedCode is the HTTP code returned for type PatchCoreV1NamespacedEventUnauthorized
const PatchCoreV1NamespacedEventUnauthorizedCode int = 401

/*PatchCoreV1NamespacedEventUnauthorized Unauthorized

swagger:response patchCoreV1NamespacedEventUnauthorized
*/
type PatchCoreV1NamespacedEventUnauthorized struct {
}

// NewPatchCoreV1NamespacedEventUnauthorized creates PatchCoreV1NamespacedEventUnauthorized with default headers values
func NewPatchCoreV1NamespacedEventUnauthorized() *PatchCoreV1NamespacedEventUnauthorized {

	return &PatchCoreV1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCoreV1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
