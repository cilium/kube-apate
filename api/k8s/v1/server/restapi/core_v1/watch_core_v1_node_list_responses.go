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

// WatchCoreV1NodeListOKCode is the HTTP code returned for type WatchCoreV1NodeListOK
const WatchCoreV1NodeListOKCode int = 200

/*WatchCoreV1NodeListOK OK

swagger:response watchCoreV1NodeListOK
*/
type WatchCoreV1NodeListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1NodeListOK creates WatchCoreV1NodeListOK with default headers values
func NewWatchCoreV1NodeListOK() *WatchCoreV1NodeListOK {

	return &WatchCoreV1NodeListOK{}
}

// WithPayload adds the payload to the watch core v1 node list o k response
func (o *WatchCoreV1NodeListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1NodeListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 node list o k response
func (o *WatchCoreV1NodeListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1NodeListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1NodeListUnauthorizedCode is the HTTP code returned for type WatchCoreV1NodeListUnauthorized
const WatchCoreV1NodeListUnauthorizedCode int = 401

/*WatchCoreV1NodeListUnauthorized Unauthorized

swagger:response watchCoreV1NodeListUnauthorized
*/
type WatchCoreV1NodeListUnauthorized struct {
}

// NewWatchCoreV1NodeListUnauthorized creates WatchCoreV1NodeListUnauthorized with default headers values
func NewWatchCoreV1NodeListUnauthorized() *WatchCoreV1NodeListUnauthorized {

	return &WatchCoreV1NodeListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1NodeListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
