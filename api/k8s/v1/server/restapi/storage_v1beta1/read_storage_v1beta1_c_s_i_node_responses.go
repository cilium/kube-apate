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

// ReadStorageV1beta1CSINodeOKCode is the HTTP code returned for type ReadStorageV1beta1CSINodeOK
const ReadStorageV1beta1CSINodeOKCode int = 200

/*ReadStorageV1beta1CSINodeOK OK

swagger:response readStorageV1beta1CSINodeOK
*/
type ReadStorageV1beta1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSINode `json:"body,omitempty"`
}

// NewReadStorageV1beta1CSINodeOK creates ReadStorageV1beta1CSINodeOK with default headers values
func NewReadStorageV1beta1CSINodeOK() *ReadStorageV1beta1CSINodeOK {

	return &ReadStorageV1beta1CSINodeOK{}
}

// WithPayload adds the payload to the read storage v1beta1 c s i node o k response
func (o *ReadStorageV1beta1CSINodeOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) *ReadStorageV1beta1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read storage v1beta1 c s i node o k response
func (o *ReadStorageV1beta1CSINodeOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadStorageV1beta1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadStorageV1beta1CSINodeUnauthorizedCode is the HTTP code returned for type ReadStorageV1beta1CSINodeUnauthorized
const ReadStorageV1beta1CSINodeUnauthorizedCode int = 401

/*ReadStorageV1beta1CSINodeUnauthorized Unauthorized

swagger:response readStorageV1beta1CSINodeUnauthorized
*/
type ReadStorageV1beta1CSINodeUnauthorized struct {
}

// NewReadStorageV1beta1CSINodeUnauthorized creates ReadStorageV1beta1CSINodeUnauthorized with default headers values
func NewReadStorageV1beta1CSINodeUnauthorized() *ReadStorageV1beta1CSINodeUnauthorized {

	return &ReadStorageV1beta1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *ReadStorageV1beta1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
