// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetBatchAPIGroupOKCode is the HTTP code returned for type GetBatchAPIGroupOK
const GetBatchAPIGroupOKCode int = 200

/*GetBatchAPIGroupOK OK

swagger:response getBatchApiGroupOK
*/
type GetBatchAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetBatchAPIGroupOK creates GetBatchAPIGroupOK with default headers values
func NewGetBatchAPIGroupOK() *GetBatchAPIGroupOK {

	return &GetBatchAPIGroupOK{}
}

// WithPayload adds the payload to the get batch Api group o k response
func (o *GetBatchAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetBatchAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get batch Api group o k response
func (o *GetBatchAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetBatchAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetBatchAPIGroupUnauthorizedCode is the HTTP code returned for type GetBatchAPIGroupUnauthorized
const GetBatchAPIGroupUnauthorizedCode int = 401

/*GetBatchAPIGroupUnauthorized Unauthorized

swagger:response getBatchApiGroupUnauthorized
*/
type GetBatchAPIGroupUnauthorized struct {
}

// NewGetBatchAPIGroupUnauthorized creates GetBatchAPIGroupUnauthorized with default headers values
func NewGetBatchAPIGroupUnauthorized() *GetBatchAPIGroupUnauthorized {

	return &GetBatchAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetBatchAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
