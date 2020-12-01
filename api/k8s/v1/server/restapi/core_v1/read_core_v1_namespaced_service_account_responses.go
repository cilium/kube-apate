// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadCoreV1NamespacedServiceAccountOKCode is the HTTP code returned for type ReadCoreV1NamespacedServiceAccountOK
const ReadCoreV1NamespacedServiceAccountOKCode int = 200

/*ReadCoreV1NamespacedServiceAccountOK OK

swagger:response readCoreV1NamespacedServiceAccountOK
*/
type ReadCoreV1NamespacedServiceAccountOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ServiceAccount `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedServiceAccountOK creates ReadCoreV1NamespacedServiceAccountOK with default headers values
func NewReadCoreV1NamespacedServiceAccountOK() *ReadCoreV1NamespacedServiceAccountOK {

	return &ReadCoreV1NamespacedServiceAccountOK{}
}

// WithPayload adds the payload to the read core v1 namespaced service account o k response
func (o *ReadCoreV1NamespacedServiceAccountOK) WithPayload(payload *models.IoK8sAPICoreV1ServiceAccount) *ReadCoreV1NamespacedServiceAccountOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced service account o k response
func (o *ReadCoreV1NamespacedServiceAccountOK) SetPayload(payload *models.IoK8sAPICoreV1ServiceAccount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedServiceAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedServiceAccountUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedServiceAccountUnauthorized
const ReadCoreV1NamespacedServiceAccountUnauthorizedCode int = 401

/*ReadCoreV1NamespacedServiceAccountUnauthorized Unauthorized

swagger:response readCoreV1NamespacedServiceAccountUnauthorized
*/
type ReadCoreV1NamespacedServiceAccountUnauthorized struct {
}

// NewReadCoreV1NamespacedServiceAccountUnauthorized creates ReadCoreV1NamespacedServiceAccountUnauthorized with default headers values
func NewReadCoreV1NamespacedServiceAccountUnauthorized() *ReadCoreV1NamespacedServiceAccountUnauthorized {

	return &ReadCoreV1NamespacedServiceAccountUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedServiceAccountUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
