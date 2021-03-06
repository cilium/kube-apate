// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteAppsV1NamespacedStatefulSetOKCode is the HTTP code returned for type DeleteAppsV1NamespacedStatefulSetOK
const DeleteAppsV1NamespacedStatefulSetOKCode int = 200

/*DeleteAppsV1NamespacedStatefulSetOK OK

swagger:response deleteAppsV1NamespacedStatefulSetOK
*/
type DeleteAppsV1NamespacedStatefulSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAppsV1NamespacedStatefulSetOK creates DeleteAppsV1NamespacedStatefulSetOK with default headers values
func NewDeleteAppsV1NamespacedStatefulSetOK() *DeleteAppsV1NamespacedStatefulSetOK {

	return &DeleteAppsV1NamespacedStatefulSetOK{}
}

// WithPayload adds the payload to the delete apps v1 namespaced stateful set o k response
func (o *DeleteAppsV1NamespacedStatefulSetOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAppsV1NamespacedStatefulSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apps v1 namespaced stateful set o k response
func (o *DeleteAppsV1NamespacedStatefulSetOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppsV1NamespacedStatefulSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAppsV1NamespacedStatefulSetAcceptedCode is the HTTP code returned for type DeleteAppsV1NamespacedStatefulSetAccepted
const DeleteAppsV1NamespacedStatefulSetAcceptedCode int = 202

/*DeleteAppsV1NamespacedStatefulSetAccepted Accepted

swagger:response deleteAppsV1NamespacedStatefulSetAccepted
*/
type DeleteAppsV1NamespacedStatefulSetAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAppsV1NamespacedStatefulSetAccepted creates DeleteAppsV1NamespacedStatefulSetAccepted with default headers values
func NewDeleteAppsV1NamespacedStatefulSetAccepted() *DeleteAppsV1NamespacedStatefulSetAccepted {

	return &DeleteAppsV1NamespacedStatefulSetAccepted{}
}

// WithPayload adds the payload to the delete apps v1 namespaced stateful set accepted response
func (o *DeleteAppsV1NamespacedStatefulSetAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAppsV1NamespacedStatefulSetAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apps v1 namespaced stateful set accepted response
func (o *DeleteAppsV1NamespacedStatefulSetAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppsV1NamespacedStatefulSetAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAppsV1NamespacedStatefulSetUnauthorizedCode is the HTTP code returned for type DeleteAppsV1NamespacedStatefulSetUnauthorized
const DeleteAppsV1NamespacedStatefulSetUnauthorizedCode int = 401

/*DeleteAppsV1NamespacedStatefulSetUnauthorized Unauthorized

swagger:response deleteAppsV1NamespacedStatefulSetUnauthorized
*/
type DeleteAppsV1NamespacedStatefulSetUnauthorized struct {
}

// NewDeleteAppsV1NamespacedStatefulSetUnauthorized creates DeleteAppsV1NamespacedStatefulSetUnauthorized with default headers values
func NewDeleteAppsV1NamespacedStatefulSetUnauthorized() *DeleteAppsV1NamespacedStatefulSetUnauthorized {

	return &DeleteAppsV1NamespacedStatefulSetUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAppsV1NamespacedStatefulSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
