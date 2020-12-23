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

// ReplaceAppsV1NamespacedDaemonSetStatusOKCode is the HTTP code returned for type ReplaceAppsV1NamespacedDaemonSetStatusOK
const ReplaceAppsV1NamespacedDaemonSetStatusOKCode int = 200

/*ReplaceAppsV1NamespacedDaemonSetStatusOK OK

swagger:response replaceAppsV1NamespacedDaemonSetStatusOK
*/
type ReplaceAppsV1NamespacedDaemonSetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedDaemonSetStatusOK creates ReplaceAppsV1NamespacedDaemonSetStatusOK with default headers values
func NewReplaceAppsV1NamespacedDaemonSetStatusOK() *ReplaceAppsV1NamespacedDaemonSetStatusOK {

	return &ReplaceAppsV1NamespacedDaemonSetStatusOK{}
}

// WithPayload adds the payload to the replace apps v1 namespaced daemon set status o k response
func (o *ReplaceAppsV1NamespacedDaemonSetStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *ReplaceAppsV1NamespacedDaemonSetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced daemon set status o k response
func (o *ReplaceAppsV1NamespacedDaemonSetStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedDaemonSetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedDaemonSetStatusCreatedCode is the HTTP code returned for type ReplaceAppsV1NamespacedDaemonSetStatusCreated
const ReplaceAppsV1NamespacedDaemonSetStatusCreatedCode int = 201

/*ReplaceAppsV1NamespacedDaemonSetStatusCreated Created

swagger:response replaceAppsV1NamespacedDaemonSetStatusCreated
*/
type ReplaceAppsV1NamespacedDaemonSetStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1DaemonSet `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedDaemonSetStatusCreated creates ReplaceAppsV1NamespacedDaemonSetStatusCreated with default headers values
func NewReplaceAppsV1NamespacedDaemonSetStatusCreated() *ReplaceAppsV1NamespacedDaemonSetStatusCreated {

	return &ReplaceAppsV1NamespacedDaemonSetStatusCreated{}
}

// WithPayload adds the payload to the replace apps v1 namespaced daemon set status created response
func (o *ReplaceAppsV1NamespacedDaemonSetStatusCreated) WithPayload(payload *models.IoK8sAPIAppsV1DaemonSet) *ReplaceAppsV1NamespacedDaemonSetStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced daemon set status created response
func (o *ReplaceAppsV1NamespacedDaemonSetStatusCreated) SetPayload(payload *models.IoK8sAPIAppsV1DaemonSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedDaemonSetStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedDaemonSetStatusUnauthorizedCode is the HTTP code returned for type ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized
const ReplaceAppsV1NamespacedDaemonSetStatusUnauthorizedCode int = 401

/*ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized Unauthorized

swagger:response replaceAppsV1NamespacedDaemonSetStatusUnauthorized
*/
type ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized struct {
}

// NewReplaceAppsV1NamespacedDaemonSetStatusUnauthorized creates ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized with default headers values
func NewReplaceAppsV1NamespacedDaemonSetStatusUnauthorized() *ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized {

	return &ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}