// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteStorageV1beta1CSINodeOKCode is the HTTP code returned for type DeleteStorageV1beta1CSINodeOK
const DeleteStorageV1beta1CSINodeOKCode int = 200

/*DeleteStorageV1beta1CSINodeOK OK

swagger:response deleteStorageV1beta1CSINodeOK
*/
type DeleteStorageV1beta1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSINode `json:"body,omitempty"`
}

// NewDeleteStorageV1beta1CSINodeOK creates DeleteStorageV1beta1CSINodeOK with default headers values
func NewDeleteStorageV1beta1CSINodeOK() *DeleteStorageV1beta1CSINodeOK {

	return &DeleteStorageV1beta1CSINodeOK{}
}

// WithPayload adds the payload to the delete storage v1beta1 c s i node o k response
func (o *DeleteStorageV1beta1CSINodeOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) *DeleteStorageV1beta1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1beta1 c s i node o k response
func (o *DeleteStorageV1beta1CSINodeOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1beta1CSINodeAcceptedCode is the HTTP code returned for type DeleteStorageV1beta1CSINodeAccepted
const DeleteStorageV1beta1CSINodeAcceptedCode int = 202

/*DeleteStorageV1beta1CSINodeAccepted Accepted

swagger:response deleteStorageV1beta1CSINodeAccepted
*/
type DeleteStorageV1beta1CSINodeAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSINode `json:"body,omitempty"`
}

// NewDeleteStorageV1beta1CSINodeAccepted creates DeleteStorageV1beta1CSINodeAccepted with default headers values
func NewDeleteStorageV1beta1CSINodeAccepted() *DeleteStorageV1beta1CSINodeAccepted {

	return &DeleteStorageV1beta1CSINodeAccepted{}
}

// WithPayload adds the payload to the delete storage v1beta1 c s i node accepted response
func (o *DeleteStorageV1beta1CSINodeAccepted) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) *DeleteStorageV1beta1CSINodeAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1beta1 c s i node accepted response
func (o *DeleteStorageV1beta1CSINodeAccepted) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSINode) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1CSINodeAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1beta1CSINodeUnauthorizedCode is the HTTP code returned for type DeleteStorageV1beta1CSINodeUnauthorized
const DeleteStorageV1beta1CSINodeUnauthorizedCode int = 401

/*DeleteStorageV1beta1CSINodeUnauthorized Unauthorized

swagger:response deleteStorageV1beta1CSINodeUnauthorized
*/
type DeleteStorageV1beta1CSINodeUnauthorized struct {
}

// NewDeleteStorageV1beta1CSINodeUnauthorized creates DeleteStorageV1beta1CSINodeUnauthorized with default headers values
func NewDeleteStorageV1beta1CSINodeUnauthorized() *DeleteStorageV1beta1CSINodeUnauthorized {

	return &DeleteStorageV1beta1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteStorageV1beta1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
