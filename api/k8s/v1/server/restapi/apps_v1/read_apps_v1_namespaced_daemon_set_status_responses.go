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

// ReadAppsV1NamespacedDaemonSetStatusOKCode is the HTTP code returned for type ReadAppsV1NamespacedDaemonSetStatusOK
const ReadAppsV1NamespacedDaemonSetStatusOKCode int = 200

/*ReadAppsV1NamespacedDaemonSetStatusOK OK

swagger:response readAppsV1NamespacedDaemonSetStatusOK
*/
type ReadAppsV1NamespacedDaemonSetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewReadAppsV1NamespacedDaemonSetStatusOK creates ReadAppsV1NamespacedDaemonSetStatusOK with default headers values
func NewReadAppsV1NamespacedDaemonSetStatusOK() *ReadAppsV1NamespacedDaemonSetStatusOK {

	return &ReadAppsV1NamespacedDaemonSetStatusOK{}
}

// WithPayload adds the payload to the read apps v1 namespaced daemon set status o k response
func (o *ReadAppsV1NamespacedDaemonSetStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *ReadAppsV1NamespacedDaemonSetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apps v1 namespaced daemon set status o k response
func (o *ReadAppsV1NamespacedDaemonSetStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedDaemonSetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadAppsV1NamespacedDaemonSetStatusUnauthorizedCode is the HTTP code returned for type ReadAppsV1NamespacedDaemonSetStatusUnauthorized
const ReadAppsV1NamespacedDaemonSetStatusUnauthorizedCode int = 401

/*ReadAppsV1NamespacedDaemonSetStatusUnauthorized Unauthorized

swagger:response readAppsV1NamespacedDaemonSetStatusUnauthorized
*/
type ReadAppsV1NamespacedDaemonSetStatusUnauthorized struct {
}

// NewReadAppsV1NamespacedDaemonSetStatusUnauthorized creates ReadAppsV1NamespacedDaemonSetStatusUnauthorized with default headers values
func NewReadAppsV1NamespacedDaemonSetStatusUnauthorized() *ReadAppsV1NamespacedDaemonSetStatusUnauthorized {

	return &ReadAppsV1NamespacedDaemonSetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedDaemonSetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
