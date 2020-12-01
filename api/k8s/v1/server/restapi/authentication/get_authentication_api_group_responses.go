// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetAuthenticationAPIGroupOKCode is the HTTP code returned for type GetAuthenticationAPIGroupOK
const GetAuthenticationAPIGroupOKCode int = 200

/*GetAuthenticationAPIGroupOK OK

swagger:response getAuthenticationApiGroupOK
*/
type GetAuthenticationAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetAuthenticationAPIGroupOK creates GetAuthenticationAPIGroupOK with default headers values
func NewGetAuthenticationAPIGroupOK() *GetAuthenticationAPIGroupOK {

	return &GetAuthenticationAPIGroupOK{}
}

// WithPayload adds the payload to the get authentication Api group o k response
func (o *GetAuthenticationAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetAuthenticationAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get authentication Api group o k response
func (o *GetAuthenticationAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAuthenticationAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAuthenticationAPIGroupUnauthorizedCode is the HTTP code returned for type GetAuthenticationAPIGroupUnauthorized
const GetAuthenticationAPIGroupUnauthorizedCode int = 401

/*GetAuthenticationAPIGroupUnauthorized Unauthorized

swagger:response getAuthenticationApiGroupUnauthorized
*/
type GetAuthenticationAPIGroupUnauthorized struct {
}

// NewGetAuthenticationAPIGroupUnauthorized creates GetAuthenticationAPIGroupUnauthorized with default headers values
func NewGetAuthenticationAPIGroupUnauthorized() *GetAuthenticationAPIGroupUnauthorized {

	return &GetAuthenticationAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetAuthenticationAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
