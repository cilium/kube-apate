// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadApiregistrationV1beta1APIServiceStatusOKCode is the HTTP code returned for type ReadApiregistrationV1beta1APIServiceStatusOK
const ReadApiregistrationV1beta1APIServiceStatusOKCode int = 200

/*ReadApiregistrationV1beta1APIServiceStatusOK OK

swagger:response readApiregistrationV1beta1ApiServiceStatusOK
*/
type ReadApiregistrationV1beta1APIServiceStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService `json:"body,omitempty"`
}

// NewReadApiregistrationV1beta1APIServiceStatusOK creates ReadApiregistrationV1beta1APIServiceStatusOK with default headers values
func NewReadApiregistrationV1beta1APIServiceStatusOK() *ReadApiregistrationV1beta1APIServiceStatusOK {

	return &ReadApiregistrationV1beta1APIServiceStatusOK{}
}

// WithPayload adds the payload to the read apiregistration v1beta1 Api service status o k response
func (o *ReadApiregistrationV1beta1APIServiceStatusOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) *ReadApiregistrationV1beta1APIServiceStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apiregistration v1beta1 Api service status o k response
func (o *ReadApiregistrationV1beta1APIServiceStatusOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadApiregistrationV1beta1APIServiceStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadApiregistrationV1beta1APIServiceStatusUnauthorizedCode is the HTTP code returned for type ReadApiregistrationV1beta1APIServiceStatusUnauthorized
const ReadApiregistrationV1beta1APIServiceStatusUnauthorizedCode int = 401

/*ReadApiregistrationV1beta1APIServiceStatusUnauthorized Unauthorized

swagger:response readApiregistrationV1beta1ApiServiceStatusUnauthorized
*/
type ReadApiregistrationV1beta1APIServiceStatusUnauthorized struct {
}

// NewReadApiregistrationV1beta1APIServiceStatusUnauthorized creates ReadApiregistrationV1beta1APIServiceStatusUnauthorized with default headers values
func NewReadApiregistrationV1beta1APIServiceStatusUnauthorized() *ReadApiregistrationV1beta1APIServiceStatusUnauthorized {

	return &ReadApiregistrationV1beta1APIServiceStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadApiregistrationV1beta1APIServiceStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}