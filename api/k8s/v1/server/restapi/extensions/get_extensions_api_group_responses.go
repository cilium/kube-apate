// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetExtensionsAPIGroupOKCode is the HTTP code returned for type GetExtensionsAPIGroupOK
const GetExtensionsAPIGroupOKCode int = 200

/*GetExtensionsAPIGroupOK OK

swagger:response getExtensionsApiGroupOK
*/
type GetExtensionsAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetExtensionsAPIGroupOK creates GetExtensionsAPIGroupOK with default headers values
func NewGetExtensionsAPIGroupOK() *GetExtensionsAPIGroupOK {

	return &GetExtensionsAPIGroupOK{}
}

// WithPayload adds the payload to the get extensions Api group o k response
func (o *GetExtensionsAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetExtensionsAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get extensions Api group o k response
func (o *GetExtensionsAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetExtensionsAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetExtensionsAPIGroupUnauthorizedCode is the HTTP code returned for type GetExtensionsAPIGroupUnauthorized
const GetExtensionsAPIGroupUnauthorizedCode int = 401

/*GetExtensionsAPIGroupUnauthorized Unauthorized

swagger:response getExtensionsApiGroupUnauthorized
*/
type GetExtensionsAPIGroupUnauthorized struct {
}

// NewGetExtensionsAPIGroupUnauthorized creates GetExtensionsAPIGroupUnauthorized with default headers values
func NewGetExtensionsAPIGroupUnauthorized() *GetExtensionsAPIGroupUnauthorized {

	return &GetExtensionsAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetExtensionsAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
