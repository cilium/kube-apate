// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetStorageAPIGroupOKCode is the HTTP code returned for type GetStorageAPIGroupOK
const GetStorageAPIGroupOKCode int = 200

/*GetStorageAPIGroupOK OK

swagger:response getStorageApiGroupOK
*/
type GetStorageAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetStorageAPIGroupOK creates GetStorageAPIGroupOK with default headers values
func NewGetStorageAPIGroupOK() *GetStorageAPIGroupOK {

	return &GetStorageAPIGroupOK{}
}

// WithPayload adds the payload to the get storage Api group o k response
func (o *GetStorageAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetStorageAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get storage Api group o k response
func (o *GetStorageAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStorageAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetStorageAPIGroupUnauthorizedCode is the HTTP code returned for type GetStorageAPIGroupUnauthorized
const GetStorageAPIGroupUnauthorizedCode int = 401

/*GetStorageAPIGroupUnauthorized Unauthorized

swagger:response getStorageApiGroupUnauthorized
*/
type GetStorageAPIGroupUnauthorized struct {
}

// NewGetStorageAPIGroupUnauthorized creates GetStorageAPIGroupUnauthorized with default headers values
func NewGetStorageAPIGroupUnauthorized() *GetStorageAPIGroupUnauthorized {

	return &GetStorageAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetStorageAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}