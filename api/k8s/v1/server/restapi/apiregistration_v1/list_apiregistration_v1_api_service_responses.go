// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListApiregistrationV1APIServiceOKCode is the HTTP code returned for type ListApiregistrationV1APIServiceOK
const ListApiregistrationV1APIServiceOKCode int = 200

/*ListApiregistrationV1APIServiceOK OK

swagger:response listApiregistrationV1ApiServiceOK
*/
type ListApiregistrationV1APIServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList `json:"body,omitempty"`
}

// NewListApiregistrationV1APIServiceOK creates ListApiregistrationV1APIServiceOK with default headers values
func NewListApiregistrationV1APIServiceOK() *ListApiregistrationV1APIServiceOK {

	return &ListApiregistrationV1APIServiceOK{}
}

// WithPayload adds the payload to the list apiregistration v1 Api service o k response
func (o *ListApiregistrationV1APIServiceOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) *ListApiregistrationV1APIServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list apiregistration v1 Api service o k response
func (o *ListApiregistrationV1APIServiceOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListApiregistrationV1APIServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListApiregistrationV1APIServiceUnauthorizedCode is the HTTP code returned for type ListApiregistrationV1APIServiceUnauthorized
const ListApiregistrationV1APIServiceUnauthorizedCode int = 401

/*ListApiregistrationV1APIServiceUnauthorized Unauthorized

swagger:response listApiregistrationV1ApiServiceUnauthorized
*/
type ListApiregistrationV1APIServiceUnauthorized struct {
}

// NewListApiregistrationV1APIServiceUnauthorized creates ListApiregistrationV1APIServiceUnauthorized with default headers values
func NewListApiregistrationV1APIServiceUnauthorized() *ListApiregistrationV1APIServiceUnauthorized {

	return &ListApiregistrationV1APIServiceUnauthorized{}
}

// WriteResponse to the client
func (o *ListApiregistrationV1APIServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}