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

// ListCoreV1NamespacedServiceOKCode is the HTTP code returned for type ListCoreV1NamespacedServiceOK
const ListCoreV1NamespacedServiceOKCode int = 200

/*ListCoreV1NamespacedServiceOK OK

swagger:response listCoreV1NamespacedServiceOK
*/
type ListCoreV1NamespacedServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ServiceList `json:"body,omitempty"`
}

// NewListCoreV1NamespacedServiceOK creates ListCoreV1NamespacedServiceOK with default headers values
func NewListCoreV1NamespacedServiceOK() *ListCoreV1NamespacedServiceOK {

	return &ListCoreV1NamespacedServiceOK{}
}

// WithPayload adds the payload to the list core v1 namespaced service o k response
func (o *ListCoreV1NamespacedServiceOK) WithPayload(payload *models.IoK8sAPICoreV1ServiceList) *ListCoreV1NamespacedServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 namespaced service o k response
func (o *ListCoreV1NamespacedServiceOK) SetPayload(payload *models.IoK8sAPICoreV1ServiceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1NamespacedServiceUnauthorizedCode is the HTTP code returned for type ListCoreV1NamespacedServiceUnauthorized
const ListCoreV1NamespacedServiceUnauthorizedCode int = 401

/*ListCoreV1NamespacedServiceUnauthorized Unauthorized

swagger:response listCoreV1NamespacedServiceUnauthorized
*/
type ListCoreV1NamespacedServiceUnauthorized struct {
}

// NewListCoreV1NamespacedServiceUnauthorized creates ListCoreV1NamespacedServiceUnauthorized with default headers values
func NewListCoreV1NamespacedServiceUnauthorized() *ListCoreV1NamespacedServiceUnauthorized {

	return &ListCoreV1NamespacedServiceUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1NamespacedServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
