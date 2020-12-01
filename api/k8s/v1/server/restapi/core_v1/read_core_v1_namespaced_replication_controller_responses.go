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

// ReadCoreV1NamespacedReplicationControllerOKCode is the HTTP code returned for type ReadCoreV1NamespacedReplicationControllerOK
const ReadCoreV1NamespacedReplicationControllerOKCode int = 200

/*ReadCoreV1NamespacedReplicationControllerOK OK

swagger:response readCoreV1NamespacedReplicationControllerOK
*/
type ReadCoreV1NamespacedReplicationControllerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1ReplicationController `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedReplicationControllerOK creates ReadCoreV1NamespacedReplicationControllerOK with default headers values
func NewReadCoreV1NamespacedReplicationControllerOK() *ReadCoreV1NamespacedReplicationControllerOK {

	return &ReadCoreV1NamespacedReplicationControllerOK{}
}

// WithPayload adds the payload to the read core v1 namespaced replication controller o k response
func (o *ReadCoreV1NamespacedReplicationControllerOK) WithPayload(payload *models.IoK8sAPICoreV1ReplicationController) *ReadCoreV1NamespacedReplicationControllerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced replication controller o k response
func (o *ReadCoreV1NamespacedReplicationControllerOK) SetPayload(payload *models.IoK8sAPICoreV1ReplicationController) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedReplicationControllerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedReplicationControllerUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedReplicationControllerUnauthorized
const ReadCoreV1NamespacedReplicationControllerUnauthorizedCode int = 401

/*ReadCoreV1NamespacedReplicationControllerUnauthorized Unauthorized

swagger:response readCoreV1NamespacedReplicationControllerUnauthorized
*/
type ReadCoreV1NamespacedReplicationControllerUnauthorized struct {
}

// NewReadCoreV1NamespacedReplicationControllerUnauthorized creates ReadCoreV1NamespacedReplicationControllerUnauthorized with default headers values
func NewReadCoreV1NamespacedReplicationControllerUnauthorized() *ReadCoreV1NamespacedReplicationControllerUnauthorized {

	return &ReadCoreV1NamespacedReplicationControllerUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedReplicationControllerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
