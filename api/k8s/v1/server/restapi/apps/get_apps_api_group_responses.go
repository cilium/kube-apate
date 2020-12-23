// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetAppsAPIGroupOKCode is the HTTP code returned for type GetAppsAPIGroupOK
const GetAppsAPIGroupOKCode int = 200

/*GetAppsAPIGroupOK OK

swagger:response getAppsApiGroupOK
*/
type GetAppsAPIGroupOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup `json:"body,omitempty"`
}

// NewGetAppsAPIGroupOK creates GetAppsAPIGroupOK with default headers values
func NewGetAppsAPIGroupOK() *GetAppsAPIGroupOK {

	return &GetAppsAPIGroupOK{}
}

// WithPayload adds the payload to the get apps Api group o k response
func (o *GetAppsAPIGroupOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) *GetAppsAPIGroupOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps Api group o k response
func (o *GetAppsAPIGroupOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsAPIGroupOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppsAPIGroupUnauthorizedCode is the HTTP code returned for type GetAppsAPIGroupUnauthorized
const GetAppsAPIGroupUnauthorizedCode int = 401

/*GetAppsAPIGroupUnauthorized Unauthorized

swagger:response getAppsApiGroupUnauthorized
*/
type GetAppsAPIGroupUnauthorized struct {
}

// NewGetAppsAPIGroupUnauthorized creates GetAppsAPIGroupUnauthorized with default headers values
func NewGetAppsAPIGroupUnauthorized() *GetAppsAPIGroupUnauthorized {

	return &GetAppsAPIGroupUnauthorized{}
}

// WriteResponse to the client
func (o *GetAppsAPIGroupUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}