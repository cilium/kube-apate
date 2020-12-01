// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListCoordinationV1beta1LeaseForAllNamespacesOKCode is the HTTP code returned for type ListCoordinationV1beta1LeaseForAllNamespacesOK
const ListCoordinationV1beta1LeaseForAllNamespacesOKCode int = 200

/*ListCoordinationV1beta1LeaseForAllNamespacesOK OK

swagger:response listCoordinationV1beta1LeaseForAllNamespacesOK
*/
type ListCoordinationV1beta1LeaseForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoordinationV1beta1LeaseList `json:"body,omitempty"`
}

// NewListCoordinationV1beta1LeaseForAllNamespacesOK creates ListCoordinationV1beta1LeaseForAllNamespacesOK with default headers values
func NewListCoordinationV1beta1LeaseForAllNamespacesOK() *ListCoordinationV1beta1LeaseForAllNamespacesOK {

	return &ListCoordinationV1beta1LeaseForAllNamespacesOK{}
}

// WithPayload adds the payload to the list coordination v1beta1 lease for all namespaces o k response
func (o *ListCoordinationV1beta1LeaseForAllNamespacesOK) WithPayload(payload *models.IoK8sAPICoordinationV1beta1LeaseList) *ListCoordinationV1beta1LeaseForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list coordination v1beta1 lease for all namespaces o k response
func (o *ListCoordinationV1beta1LeaseForAllNamespacesOK) SetPayload(payload *models.IoK8sAPICoordinationV1beta1LeaseList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCoordinationV1beta1LeaseForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCoordinationV1beta1LeaseForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized
const ListCoordinationV1beta1LeaseForAllNamespacesUnauthorizedCode int = 401

/*ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized Unauthorized

swagger:response listCoordinationV1beta1LeaseForAllNamespacesUnauthorized
*/
type ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized struct {
}

// NewListCoordinationV1beta1LeaseForAllNamespacesUnauthorized creates ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized with default headers values
func NewListCoordinationV1beta1LeaseForAllNamespacesUnauthorized() *ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized {

	return &ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
