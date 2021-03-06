// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListBatchV1JobForAllNamespacesOKCode is the HTTP code returned for type ListBatchV1JobForAllNamespacesOK
const ListBatchV1JobForAllNamespacesOKCode int = 200

/*ListBatchV1JobForAllNamespacesOK OK

swagger:response listBatchV1JobForAllNamespacesOK
*/
type ListBatchV1JobForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1JobList `json:"body,omitempty"`
}

// NewListBatchV1JobForAllNamespacesOK creates ListBatchV1JobForAllNamespacesOK with default headers values
func NewListBatchV1JobForAllNamespacesOK() *ListBatchV1JobForAllNamespacesOK {

	return &ListBatchV1JobForAllNamespacesOK{}
}

// WithPayload adds the payload to the list batch v1 job for all namespaces o k response
func (o *ListBatchV1JobForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIBatchV1JobList) *ListBatchV1JobForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list batch v1 job for all namespaces o k response
func (o *ListBatchV1JobForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIBatchV1JobList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBatchV1JobForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListBatchV1JobForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListBatchV1JobForAllNamespacesUnauthorized
const ListBatchV1JobForAllNamespacesUnauthorizedCode int = 401

/*ListBatchV1JobForAllNamespacesUnauthorized Unauthorized

swagger:response listBatchV1JobForAllNamespacesUnauthorized
*/
type ListBatchV1JobForAllNamespacesUnauthorized struct {
}

// NewListBatchV1JobForAllNamespacesUnauthorized creates ListBatchV1JobForAllNamespacesUnauthorized with default headers values
func NewListBatchV1JobForAllNamespacesUnauthorized() *ListBatchV1JobForAllNamespacesUnauthorized {

	return &ListBatchV1JobForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListBatchV1JobForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
