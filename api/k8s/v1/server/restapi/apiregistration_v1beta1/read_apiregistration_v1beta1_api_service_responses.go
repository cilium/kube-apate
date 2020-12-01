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

// ReadApiregistrationV1beta1APIServiceOKCode is the HTTP code returned for type ReadApiregistrationV1beta1APIServiceOK
const ReadApiregistrationV1beta1APIServiceOKCode int = 200

/*ReadApiregistrationV1beta1APIServiceOK OK

swagger:response readApiregistrationV1beta1ApiServiceOK
*/
type ReadApiregistrationV1beta1APIServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService `json:"body,omitempty"`
}

// NewReadApiregistrationV1beta1APIServiceOK creates ReadApiregistrationV1beta1APIServiceOK with default headers values
func NewReadApiregistrationV1beta1APIServiceOK() *ReadApiregistrationV1beta1APIServiceOK {

	return &ReadApiregistrationV1beta1APIServiceOK{}
}

// WithPayload adds the payload to the read apiregistration v1beta1 Api service o k response
func (o *ReadApiregistrationV1beta1APIServiceOK) WithPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) *ReadApiregistrationV1beta1APIServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apiregistration v1beta1 Api service o k response
func (o *ReadApiregistrationV1beta1APIServiceOK) SetPayload(payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadApiregistrationV1beta1APIServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadApiregistrationV1beta1APIServiceUnauthorizedCode is the HTTP code returned for type ReadApiregistrationV1beta1APIServiceUnauthorized
const ReadApiregistrationV1beta1APIServiceUnauthorizedCode int = 401

/*ReadApiregistrationV1beta1APIServiceUnauthorized Unauthorized

swagger:response readApiregistrationV1beta1ApiServiceUnauthorized
*/
type ReadApiregistrationV1beta1APIServiceUnauthorized struct {
}

// NewReadApiregistrationV1beta1APIServiceUnauthorized creates ReadApiregistrationV1beta1APIServiceUnauthorized with default headers values
func NewReadApiregistrationV1beta1APIServiceUnauthorized() *ReadApiregistrationV1beta1APIServiceUnauthorized {

	return &ReadApiregistrationV1beta1APIServiceUnauthorized{}
}

// WriteResponse to the client
func (o *ReadApiregistrationV1beta1APIServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
