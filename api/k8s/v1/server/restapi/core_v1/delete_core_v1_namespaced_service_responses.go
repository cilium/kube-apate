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

// DeleteCoreV1NamespacedServiceOKCode is the HTTP code returned for type DeleteCoreV1NamespacedServiceOK
const DeleteCoreV1NamespacedServiceOKCode int = 200

/*DeleteCoreV1NamespacedServiceOK OK

swagger:response deleteCoreV1NamespacedServiceOK
*/
type DeleteCoreV1NamespacedServiceOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedServiceOK creates DeleteCoreV1NamespacedServiceOK with default headers values
func NewDeleteCoreV1NamespacedServiceOK() *DeleteCoreV1NamespacedServiceOK {

	return &DeleteCoreV1NamespacedServiceOK{}
}

// WithPayload adds the payload to the delete core v1 namespaced service o k response
func (o *DeleteCoreV1NamespacedServiceOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedServiceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced service o k response
func (o *DeleteCoreV1NamespacedServiceOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedServiceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedServiceAcceptedCode is the HTTP code returned for type DeleteCoreV1NamespacedServiceAccepted
const DeleteCoreV1NamespacedServiceAcceptedCode int = 202

/*DeleteCoreV1NamespacedServiceAccepted Accepted

swagger:response deleteCoreV1NamespacedServiceAccepted
*/
type DeleteCoreV1NamespacedServiceAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoreV1NamespacedServiceAccepted creates DeleteCoreV1NamespacedServiceAccepted with default headers values
func NewDeleteCoreV1NamespacedServiceAccepted() *DeleteCoreV1NamespacedServiceAccepted {

	return &DeleteCoreV1NamespacedServiceAccepted{}
}

// WithPayload adds the payload to the delete core v1 namespaced service accepted response
func (o *DeleteCoreV1NamespacedServiceAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoreV1NamespacedServiceAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 namespaced service accepted response
func (o *DeleteCoreV1NamespacedServiceAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedServiceAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1NamespacedServiceUnauthorizedCode is the HTTP code returned for type DeleteCoreV1NamespacedServiceUnauthorized
const DeleteCoreV1NamespacedServiceUnauthorizedCode int = 401

/*DeleteCoreV1NamespacedServiceUnauthorized Unauthorized

swagger:response deleteCoreV1NamespacedServiceUnauthorized
*/
type DeleteCoreV1NamespacedServiceUnauthorized struct {
}

// NewDeleteCoreV1NamespacedServiceUnauthorized creates DeleteCoreV1NamespacedServiceUnauthorized with default headers values
func NewDeleteCoreV1NamespacedServiceUnauthorized() *DeleteCoreV1NamespacedServiceUnauthorized {

	return &DeleteCoreV1NamespacedServiceUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1NamespacedServiceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
