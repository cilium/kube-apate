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

// DeleteCoordinationV1beta1CollectionNamespacedLeaseOKCode is the HTTP code returned for type DeleteCoordinationV1beta1CollectionNamespacedLeaseOK
const DeleteCoordinationV1beta1CollectionNamespacedLeaseOKCode int = 200

/*DeleteCoordinationV1beta1CollectionNamespacedLeaseOK OK

swagger:response deleteCoordinationV1beta1CollectionNamespacedLeaseOK
*/
type DeleteCoordinationV1beta1CollectionNamespacedLeaseOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCoordinationV1beta1CollectionNamespacedLeaseOK creates DeleteCoordinationV1beta1CollectionNamespacedLeaseOK with default headers values
func NewDeleteCoordinationV1beta1CollectionNamespacedLeaseOK() *DeleteCoordinationV1beta1CollectionNamespacedLeaseOK {

	return &DeleteCoordinationV1beta1CollectionNamespacedLeaseOK{}
}

// WithPayload adds the payload to the delete coordination v1beta1 collection namespaced lease o k response
func (o *DeleteCoordinationV1beta1CollectionNamespacedLeaseOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCoordinationV1beta1CollectionNamespacedLeaseOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete coordination v1beta1 collection namespaced lease o k response
func (o *DeleteCoordinationV1beta1CollectionNamespacedLeaseOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoordinationV1beta1CollectionNamespacedLeaseOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorizedCode is the HTTP code returned for type DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized
const DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorizedCode int = 401

/*DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized Unauthorized

swagger:response deleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized
*/
type DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized struct {
}

// NewDeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized creates DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized with default headers values
func NewDeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized() *DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized {

	return &DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoordinationV1beta1CollectionNamespacedLeaseUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}