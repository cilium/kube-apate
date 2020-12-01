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

// WatchCoreV1ConfigMapListForAllNamespacesOKCode is the HTTP code returned for type WatchCoreV1ConfigMapListForAllNamespacesOK
const WatchCoreV1ConfigMapListForAllNamespacesOKCode int = 200

/*WatchCoreV1ConfigMapListForAllNamespacesOK OK

swagger:response watchCoreV1ConfigMapListForAllNamespacesOK
*/
type WatchCoreV1ConfigMapListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1ConfigMapListForAllNamespacesOK creates WatchCoreV1ConfigMapListForAllNamespacesOK with default headers values
func NewWatchCoreV1ConfigMapListForAllNamespacesOK() *WatchCoreV1ConfigMapListForAllNamespacesOK {

	return &WatchCoreV1ConfigMapListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch core v1 config map list for all namespaces o k response
func (o *WatchCoreV1ConfigMapListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1ConfigMapListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 config map list for all namespaces o k response
func (o *WatchCoreV1ConfigMapListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1ConfigMapListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1ConfigMapListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchCoreV1ConfigMapListForAllNamespacesUnauthorized
const WatchCoreV1ConfigMapListForAllNamespacesUnauthorizedCode int = 401

/*WatchCoreV1ConfigMapListForAllNamespacesUnauthorized Unauthorized

swagger:response watchCoreV1ConfigMapListForAllNamespacesUnauthorized
*/
type WatchCoreV1ConfigMapListForAllNamespacesUnauthorized struct {
}

// NewWatchCoreV1ConfigMapListForAllNamespacesUnauthorized creates WatchCoreV1ConfigMapListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1ConfigMapListForAllNamespacesUnauthorized() *WatchCoreV1ConfigMapListForAllNamespacesUnauthorized {

	return &WatchCoreV1ConfigMapListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1ConfigMapListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
