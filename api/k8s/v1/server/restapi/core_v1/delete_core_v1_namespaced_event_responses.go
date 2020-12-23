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

// DeleteCoreV1NamespacedEventOKCode is the HTTP code returned for type DeleteCoreV1NamespacedEventOK
const DeleteCoreV1NamespacedEventOKCode int = 200

/*DeleteCoreV1NamespacedEventOK OK

swagger:response deleteCoreV1NamespacedEventOK
*/
type DeleteCoreV1NamespacedEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedEventOK creates DeleteCoreV1NamespacedEventOK with default headers values
func NewDeleteCoreV1NamespacedEventOK() *DeleteCoreV1NamespacedEventOK {

	return &DeleteCoreV1NamespacedEventOK{}
}

// WithPayload adds the payload to the delete core v1 namespaced event o k response
func (o *DeleteCoreV1NamespacedEventOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced event o k response
func (o *DeleteCoreV1NamespacedEventOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedEventAcceptedCode is the HTTP code returned for type DeleteCoreV1NamespacedEventAccepted
const DeleteCoreV1NamespacedEventAcceptedCode int = 202

/*DeleteCoreV1NamespacedEventAccepted Accepted

swagger:response deleteCoreV1NamespacedEventAccepted
*/
type DeleteCoreV1NamespacedEventAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedEventAccepted creates DeleteCoreV1NamespacedEventAccepted with default headers values
func NewDeleteCoreV1NamespacedEventAccepted() *DeleteCoreV1NamespacedEventAccepted {

	return &DeleteCoreV1NamespacedEventAccepted{}
}

// WithPayload adds the payload to the delete core v1 namespaced event accepted response
func (o *DeleteCoreV1NamespacedEventAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedEventAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced event accepted response
func (o *DeleteCoreV1NamespacedEventAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedEventAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedEventUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NamespacedEventUnauthorized
const DeleteCoreV1NamespacedEventUnauthorizedCode int = 401

/*DeleteCoreV1NamespacedEventUnauthorized Unauthorized

swagger:response deleteCoreV1NamespacedEventUnauthorized
*/
type DeleteCoreV1NamespacedEventUnauthorized struct {
}

// NewDeleteCoreV1NamespacedEventUnauthorized creates DeleteCoreV1NamespacedEventUnauthorized with default headers values
func NewDeleteCoreV1NamespacedEventUnauthorized() *DeleteCoreV1NamespacedEventUnauthorized {

	return &DeleteCoreV1NamespacedEventUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedEventUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}