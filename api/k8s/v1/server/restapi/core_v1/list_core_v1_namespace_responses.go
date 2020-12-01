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

// ListCoreV1NamespaceOKCode is the HTTP code returned for type ListCoreV1NamespaceOK
const ListCoreV1NamespaceOKCode int = 200

/*ListCoreV1NamespaceOK OK

swagger:response listCoreV1NamespaceOK
*/
type ListCoreV1NamespaceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1NamespaceList `json:"body,omitempty"`
}

// NewListCoreV1NamespaceOK creates ListCoreV1NamespaceOK with default headers values
func NewListCoreV1NamespaceOK() *ListCoreV1NamespaceOK {

	return &ListCoreV1NamespaceOK{}
}

// WithPayload adds the payload to the list core v1 namespace o k response
func (o *ListCoreV1NamespaceOK) WithPayload(payload *models.IoK8sAPICoreV1NamespaceList) *ListCoreV1NamespaceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list core v1 namespace o k response
func (o *ListCoreV1NamespaceOK) SetPayload(payload *models.IoK8sAPICoreV1NamespaceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoreV1NamespaceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoreV1NamespaceUnauthorizedCode is the HTTP code returned for type ListCoreV1NamespaceUnauthorized
const ListCoreV1NamespaceUnauthorizedCode int = 401

/*ListCoreV1NamespaceUnauthorized Unauthorized

swagger:response listCoreV1NamespaceUnauthorized
*/
type ListCoreV1NamespaceUnauthorized struct {
}

// NewListCoreV1NamespaceUnauthorized creates ListCoreV1NamespaceUnauthorized with default headers values
func NewListCoreV1NamespaceUnauthorized() *ListCoreV1NamespaceUnauthorized {

	return &ListCoreV1NamespaceUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoreV1NamespaceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
