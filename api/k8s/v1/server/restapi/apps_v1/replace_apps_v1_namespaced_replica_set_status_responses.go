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

// ReplaceAppsV1NamespacedReplicaSetStatusOKCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetStatusOK
const ReplaceAppsV1NamespacedReplicaSetStatusOKCode int = 200

/*ReplaceAppsV1NamespacedReplicaSetStatusOK OK

swagger:response replaceAppsV1NamespacedReplicaSetStatusOK
*/
type ReplaceAppsV1NamespacedReplicaSetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedReplicaSetStatusOK creates ReplaceAppsV1NamespacedReplicaSetStatusOK with default headers values
func NewReplaceAppsV1NamespacedReplicaSetStatusOK() *ReplaceAppsV1NamespacedReplicaSetStatusOK {

	return &ReplaceAppsV1NamespacedReplicaSetStatusOK{}
}

// WithPayload adds the payload to the replace apps v1 namespaced replica set status o k response
func (o *ReplaceAppsV1NamespacedReplicaSetStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *ReplaceAppsV1NamespacedReplicaSetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced replica set status o k response
func (o *ReplaceAppsV1NamespacedReplicaSetStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedReplicaSetStatusCreatedCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetStatusCreated
const ReplaceAppsV1NamespacedReplicaSetStatusCreatedCode int = 201

/*ReplaceAppsV1NamespacedReplicaSetStatusCreated Created

swagger:response replaceAppsV1NamespacedReplicaSetStatusCreated
*/
type ReplaceAppsV1NamespacedReplicaSetStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1ReplicaSet `json:"body,omitempty"`
}

// NewReplaceAppsV1NamespacedReplicaSetStatusCreated creates ReplaceAppsV1NamespacedReplicaSetStatusCreated with default headers values
func NewReplaceAppsV1NamespacedReplicaSetStatusCreated() *ReplaceAppsV1NamespacedReplicaSetStatusCreated {

	return &ReplaceAppsV1NamespacedReplicaSetStatusCreated{}
}

// WithPayload adds the payload to the replace apps v1 namespaced replica set status created response
func (o *ReplaceAppsV1NamespacedReplicaSetStatusCreated) WithPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) *ReplaceAppsV1NamespacedReplicaSetStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace apps v1 namespaced replica set status created response
func (o *ReplaceAppsV1NamespacedReplicaSetStatusCreated) SetPayload(payload *models.IoK8sAPIAppsV1ReplicaSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAppsV1NamespacedReplicaSetStatusUnauthorizedCode is the HTTP code returned for type ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized
const ReplaceAppsV1NamespacedReplicaSetStatusUnauthorizedCode int = 401

/*ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized Unauthorized

swagger:response replaceAppsV1NamespacedReplicaSetStatusUnauthorized
*/
type ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized struct {
}

// NewReplaceAppsV1NamespacedReplicaSetStatusUnauthorized creates ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized with default headers values
func NewReplaceAppsV1NamespacedReplicaSetStatusUnauthorized() *ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized {

	return &ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAppsV1NamespacedReplicaSetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
