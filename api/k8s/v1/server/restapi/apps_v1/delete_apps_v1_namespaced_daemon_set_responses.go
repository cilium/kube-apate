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

// DeleteAppsV1NamespacedDaemonSetOKCode is the HTTP code returned for type DeleteAppsV1NamespacedDaemonSetOK
const DeleteAppsV1NamespacedDaemonSetOKCode int = 200

/*DeleteAppsV1NamespacedDaemonSetOK OK

swagger:response deleteAppsV1NamespacedDaemonSetOK
*/
type DeleteAppsV1NamespacedDaemonSetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAppsV1NamespacedDaemonSetOK creates DeleteAppsV1NamespacedDaemonSetOK with default headers values
func NewDeleteAppsV1NamespacedDaemonSetOK() *DeleteAppsV1NamespacedDaemonSetOK {

	return &DeleteAppsV1NamespacedDaemonSetOK{}
}

// WithPayload adds the payload to the delete apps v1 namespaced daemon set o k response
func (o *DeleteAppsV1NamespacedDaemonSetOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAppsV1NamespacedDaemonSetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apps v1 namespaced daemon set o k response
func (o *DeleteAppsV1NamespacedDaemonSetOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppsV1NamespacedDaemonSetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAppsV1NamespacedDaemonSetAcceptedCode is the HTTP code returned for type DeleteAppsV1NamespacedDaemonSetAccepted
const DeleteAppsV1NamespacedDaemonSetAcceptedCode int = 202

/*DeleteAppsV1NamespacedDaemonSetAccepted Accepted

swagger:response deleteAppsV1NamespacedDaemonSetAccepted
*/
type DeleteAppsV1NamespacedDaemonSetAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAppsV1NamespacedDaemonSetAccepted creates DeleteAppsV1NamespacedDaemonSetAccepted with default headers values
func NewDeleteAppsV1NamespacedDaemonSetAccepted() *DeleteAppsV1NamespacedDaemonSetAccepted {

	return &DeleteAppsV1NamespacedDaemonSetAccepted{}
}

// WithPayload adds the payload to the delete apps v1 namespaced daemon set accepted response
func (o *DeleteAppsV1NamespacedDaemonSetAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAppsV1NamespacedDaemonSetAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete apps v1 namespaced daemon set accepted response
func (o *DeleteAppsV1NamespacedDaemonSetAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAppsV1NamespacedDaemonSetAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAppsV1NamespacedDaemonSetUnauthorizedCode is the HTTP code returned for type DeleteAppsV1NamespacedDaemonSetUnauthorized
const DeleteAppsV1NamespacedDaemonSetUnauthorizedCode int = 401

/*DeleteAppsV1NamespacedDaemonSetUnauthorized Unauthorized

swagger:response deleteAppsV1NamespacedDaemonSetUnauthorized
*/
type DeleteAppsV1NamespacedDaemonSetUnauthorized struct {
}

// NewDeleteAppsV1NamespacedDaemonSetUnauthorized creates DeleteAppsV1NamespacedDaemonSetUnauthorized with default headers values
func NewDeleteAppsV1NamespacedDaemonSetUnauthorized() *DeleteAppsV1NamespacedDaemonSetUnauthorized {

	return &DeleteAppsV1NamespacedDaemonSetUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAppsV1NamespacedDaemonSetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
