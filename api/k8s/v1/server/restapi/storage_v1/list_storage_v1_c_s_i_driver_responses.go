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

// ListStorageV1CSIDriverOKCode is the HTTP code returned for type ListStorageV1CSIDriverOK
const ListStorageV1CSIDriverOKCode int = 200

/*ListStorageV1CSIDriverOK OK

swagger:response listStorageV1CSIDriverOK
*/
type ListStorageV1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSIDriverList `json:"body,omitempty"`
}

// NewListStorageV1CSIDriverOK creates ListStorageV1CSIDriverOK with default headers values
func NewListStorageV1CSIDriverOK() *ListStorageV1CSIDriverOK {

	return &ListStorageV1CSIDriverOK{}
}

// WithPayload adds the payload to the list storage v1 c s i driver o k response
func (o *ListStorageV1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1CSIDriverList) *ListStorageV1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list storage v1 c s i driver o k response
func (o *ListStorageV1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1CSIDriverList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStorageV1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStorageV1CSIDriverUnauthorizedCode is the HTTP code returned for type ListStorageV1CSIDriverUnauthorized
const ListStorageV1CSIDriverUnauthorizedCode int = 401

/*ListStorageV1CSIDriverUnauthorized Unauthorized

swagger:response listStorageV1CSIDriverUnauthorized
*/
type ListStorageV1CSIDriverUnauthorized struct {
}

// NewListStorageV1CSIDriverUnauthorized creates ListStorageV1CSIDriverUnauthorized with default headers values
func NewListStorageV1CSIDriverUnauthorized() *ListStorageV1CSIDriverUnauthorized {

	return &ListStorageV1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *ListStorageV1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}