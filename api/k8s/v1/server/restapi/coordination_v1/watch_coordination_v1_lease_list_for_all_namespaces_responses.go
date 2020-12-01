// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchCoordinationV1LeaseListForAllNamespacesOKCode is the HTTP code returned for type WatchCoordinationV1LeaseListForAllNamespacesOK
const WatchCoordinationV1LeaseListForAllNamespacesOKCode int = 200

/*WatchCoordinationV1LeaseListForAllNamespacesOK OK

swagger:response watchCoordinationV1LeaseListForAllNamespacesOK
*/
type WatchCoordinationV1LeaseListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoordinationV1LeaseListForAllNamespacesOK creates WatchCoordinationV1LeaseListForAllNamespacesOK with default headers values
func NewWatchCoordinationV1LeaseListForAllNamespacesOK() *WatchCoordinationV1LeaseListForAllNamespacesOK {

	return &WatchCoordinationV1LeaseListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch coordination v1 lease list for all namespaces o k response
func (o *WatchCoordinationV1LeaseListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoordinationV1LeaseListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch coordination v1 lease list for all namespaces o k response
func (o *WatchCoordinationV1LeaseListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoordinationV1LeaseListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoordinationV1LeaseListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchCoordinationV1LeaseListForAllNamespacesUnauthorized
const WatchCoordinationV1LeaseListForAllNamespacesUnauthorizedCode int = 401

/*WatchCoordinationV1LeaseListForAllNamespacesUnauthorized Unauthorized

swagger:response watchCoordinationV1LeaseListForAllNamespacesUnauthorized
*/
type WatchCoordinationV1LeaseListForAllNamespacesUnauthorized struct {
}

// NewWatchCoordinationV1LeaseListForAllNamespacesUnauthorized creates WatchCoordinationV1LeaseListForAllNamespacesUnauthorized with default headers values
func NewWatchCoordinationV1LeaseListForAllNamespacesUnauthorized() *WatchCoordinationV1LeaseListForAllNamespacesUnauthorized {

	return &WatchCoordinationV1LeaseListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoordinationV1LeaseListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
