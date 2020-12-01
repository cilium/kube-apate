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

// DeleteStorageV1CSINodeOKCode is the HTTP code returned for type DeleteStorageV1CSINodeOK
const DeleteStorageV1CSINodeOKCode int = 200

/*DeleteStorageV1CSINodeOK OK

swagger:response deleteStorageV1CSINodeOK
*/
type DeleteStorageV1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSINode `json:"body,omitempty"`
}

// NewDeleteStorageV1CSINodeOK creates DeleteStorageV1CSINodeOK with default headers values
func NewDeleteStorageV1CSINodeOK() *DeleteStorageV1CSINodeOK {

	return &DeleteStorageV1CSINodeOK{}
}

// WithPayload adds the payload to the delete storage v1 c s i node o k response
func (o *DeleteStorageV1CSINodeOK) WithPayload(payload *models.IoK8sAPIStorageV1CSINode) *DeleteStorageV1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1 c s i node o k response
func (o *DeleteStorageV1CSINodeOK) SetPayload(payload *models.IoK8sAPIStorageV1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1CSINodeAcceptedCode is the HTTP code returned for type DeleteStorageV1CSINodeAccepted
const DeleteStorageV1CSINodeAcceptedCode int = 202

/*DeleteStorageV1CSINodeAccepted Accepted

swagger:response deleteStorageV1CSINodeAccepted
*/
type DeleteStorageV1CSINodeAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1CSINode `json:"body,omitempty"`
}

// NewDeleteStorageV1CSINodeAccepted creates DeleteStorageV1CSINodeAccepted with default headers values
func NewDeleteStorageV1CSINodeAccepted() *DeleteStorageV1CSINodeAccepted {

	return &DeleteStorageV1CSINodeAccepted{}
}

// WithPayload adds the payload to the delete storage v1 c s i node accepted response
func (o *DeleteStorageV1CSINodeAccepted) WithPayload(payload *models.IoK8sAPIStorageV1CSINode) *DeleteStorageV1CSINodeAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1 c s i node accepted response
func (o *DeleteStorageV1CSINodeAccepted) SetPayload(payload *models.IoK8sAPIStorageV1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1CSINodeAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1CSINodeUnauthorizedCode is the HTTP code returned for type DeleteStorageV1CSINodeUnauthorized
const DeleteStorageV1CSINodeUnauthorizedCode int = 401

/*DeleteStorageV1CSINodeUnauthorized Unauthorized

swagger:response deleteStorageV1CSINodeUnauthorized
*/
type DeleteStorageV1CSINodeUnauthorized struct {
}

// NewDeleteStorageV1CSINodeUnauthorized creates DeleteStorageV1CSINodeUnauthorized with default headers values
func NewDeleteStorageV1CSINodeUnauthorized() *DeleteStorageV1CSINodeUnauthorized {

	return &DeleteStorageV1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteStorageV1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
