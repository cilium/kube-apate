// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListStorageV1CSINodeOKCode is the HTTP code returned for type ListStorageV1CSINodeOK
const ListStorageV1CSINodeOKCode int = 200

/*ListStorageV1CSINodeOK OK

swagger:response listStorageV1CSINodeOK
*/
type ListStorageV1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSINodeList `json:"body,omitempty"`
}

// NewListStorageV1CSINodeOK creates ListStorageV1CSINodeOK with default headers values
func NewListStorageV1CSINodeOK() *ListStorageV1CSINodeOK {

	return &ListStorageV1CSINodeOK{}
}

// WithPayload adds the payload to the list storage v1 c s i node o k response
func (o *ListStorageV1CSINodeOK) WithPayload(payload *models.IoK8sAPIStorageV1CSINodeList) *ListStorageV1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list storage v1 c s i node o k response
func (o *ListStorageV1CSINodeOK) SetPayload(payload *models.IoK8sAPIStorageV1CSINodeList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListStorageV1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListStorageV1CSINodeUnauthorizedCode is the HTTP code returned for type ListStorageV1CSINodeUnauthorized
const ListStorageV1CSINodeUnauthorizedCode int = 401

/*ListStorageV1CSINodeUnauthorized Unauthorized

swagger:response listStorageV1CSINodeUnauthorized
*/
type ListStorageV1CSINodeUnauthorized struct {
}

// NewListStorageV1CSINodeUnauthorized creates ListStorageV1CSINodeUnauthorized with default headers values
func NewListStorageV1CSINodeUnauthorized() *ListStorageV1CSINodeUnauthorized {

	return &ListStorageV1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *ListStorageV1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
