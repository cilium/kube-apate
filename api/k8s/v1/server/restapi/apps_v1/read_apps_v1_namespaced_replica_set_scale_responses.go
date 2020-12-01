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

// ReadAppsV1NamespacedReplicaSetScaleOKCode is the HTTP code returned for type ReadAppsV1NamespacedReplicaSetScaleOK
const ReadAppsV1NamespacedReplicaSetScaleOKCode int = 200

/*ReadAppsV1NamespacedReplicaSetScaleOK OK

swagger:response readAppsV1NamespacedReplicaSetScaleOK
*/
type ReadAppsV1NamespacedReplicaSetScaleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV1Scale `json:"body,omitempty"`
}

// NewReadAppsV1NamespacedReplicaSetScaleOK creates ReadAppsV1NamespacedReplicaSetScaleOK with default headers values
func NewReadAppsV1NamespacedReplicaSetScaleOK() *ReadAppsV1NamespacedReplicaSetScaleOK {

	return &ReadAppsV1NamespacedReplicaSetScaleOK{}
}

// WithPayload adds the payload to the read apps v1 namespaced replica set scale o k response
func (o *ReadAppsV1NamespacedReplicaSetScaleOK) WithPayload(payload *models.IoK8sAPIAutoscalingV1Scale) *ReadAppsV1NamespacedReplicaSetScaleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read apps v1 namespaced replica set scale o k response
func (o *ReadAppsV1NamespacedReplicaSetScaleOK) SetPayload(payload *models.IoK8sAPIAutoscalingV1Scale) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedReplicaSetScaleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadAppsV1NamespacedReplicaSetScaleUnauthorizedCode is the HTTP code returned for type ReadAppsV1NamespacedReplicaSetScaleUnauthorized
const ReadAppsV1NamespacedReplicaSetScaleUnauthorizedCode int = 401

/*ReadAppsV1NamespacedReplicaSetScaleUnauthorized Unauthorized

swagger:response readAppsV1NamespacedReplicaSetScaleUnauthorized
*/
type ReadAppsV1NamespacedReplicaSetScaleUnauthorized struct {
}

// NewReadAppsV1NamespacedReplicaSetScaleUnauthorized creates ReadAppsV1NamespacedReplicaSetScaleUnauthorized with default headers values
func NewReadAppsV1NamespacedReplicaSetScaleUnauthorized() *ReadAppsV1NamespacedReplicaSetScaleUnauthorized {

	return &ReadAppsV1NamespacedReplicaSetScaleUnauthorized{}
}

// WriteResponse to the client
func (o *ReadAppsV1NamespacedReplicaSetScaleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
