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

// ReadStorageV1beta1StorageClassOKCode is the HTTP code returned for type ReadStorageV1beta1StorageClassOK
const ReadStorageV1beta1StorageClassOKCode int = 200

/*ReadStorageV1beta1StorageClassOK OK

swagger:response readStorageV1beta1StorageClassOK
*/
type ReadStorageV1beta1StorageClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1StorageClass `json:"body,omitempty"`
}

// NewReadStorageV1beta1StorageClassOK creates ReadStorageV1beta1StorageClassOK with default headers values
func NewReadStorageV1beta1StorageClassOK() *ReadStorageV1beta1StorageClassOK {

	return &ReadStorageV1beta1StorageClassOK{}
}

// WithPayload adds the payload to the read storage v1beta1 storage class o k response
func (o *ReadStorageV1beta1StorageClassOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1StorageClass) *ReadStorageV1beta1StorageClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read storage v1beta1 storage class o k response
func (o *ReadStorageV1beta1StorageClassOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1StorageClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadStorageV1beta1StorageClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadStorageV1beta1StorageClassUnauthorizedCode is the HTTP code returned for type ReadStorageV1beta1StorageClassUnauthorized
const ReadStorageV1beta1StorageClassUnauthorizedCode int = 401

/*ReadStorageV1beta1StorageClassUnauthorized Unauthorized

swagger:response readStorageV1beta1StorageClassUnauthorized
*/
type ReadStorageV1beta1StorageClassUnauthorized struct {
}

// NewReadStorageV1beta1StorageClassUnauthorized creates ReadStorageV1beta1StorageClassUnauthorized with default headers values
func NewReadStorageV1beta1StorageClassUnauthorized() *ReadStorageV1beta1StorageClassUnauthorized {

	return &ReadStorageV1beta1StorageClassUnauthorized{}
}

// WriteResponse to the client
func (o *ReadStorageV1beta1StorageClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}