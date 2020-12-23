// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchStorageV1CSIDriverOKCode is the HTTP code returned for type PatchStorageV1CSIDriverOK
const PatchStorageV1CSIDriverOKCode int = 200

/*PatchStorageV1CSIDriverOK OK

swagger:response patchStorageV1CSIDriverOK
*/
type PatchStorageV1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSIDriver `json:"body,omitempty"`
}

// NewPatchStorageV1CSIDriverOK creates PatchStorageV1CSIDriverOK with default headers values
func NewPatchStorageV1CSIDriverOK() *PatchStorageV1CSIDriverOK {

	return &PatchStorageV1CSIDriverOK{}
}

// WithPayload adds the payload to the patch storage v1 c s i driver o k response
func (o *PatchStorageV1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1CSIDriver) *PatchStorageV1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch storage v1 c s i driver o k response
func (o *PatchStorageV1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchStorageV1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchStorageV1CSIDriverUnauthorizedCode is the HTTP code returned for type PatchStorageV1CSIDriverUnauthorized
const PatchStorageV1CSIDriverUnauthorizedCode int = 401

/*PatchStorageV1CSIDriverUnauthorized Unauthorized

swagger:response patchStorageV1CSIDriverUnauthorized
*/
type PatchStorageV1CSIDriverUnauthorized struct {
}

// NewPatchStorageV1CSIDriverUnauthorized creates PatchStorageV1CSIDriverUnauthorized with default headers values
func NewPatchStorageV1CSIDriverUnauthorized() *PatchStorageV1CSIDriverUnauthorized {

	return &PatchStorageV1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *PatchStorageV1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}