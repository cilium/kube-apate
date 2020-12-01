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

// WatchCoreV1PersistentVolumeClaimListForAllNamespacesOKCode is the HTTP code returned for type WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK
const WatchCoreV1PersistentVolumeClaimListForAllNamespacesOKCode int = 200

/*WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK OK

swagger:response watchCoreV1PersistentVolumeClaimListForAllNamespacesOK
*/
type WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesOK creates WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK with default headers values
func NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesOK() *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK {

	return &WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch core v1 persistent volume claim list for all namespaces o k response
func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 persistent volume claim list for all namespaces o k response
func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized
const WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorizedCode int = 401

/*WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized Unauthorized

swagger:response watchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized
*/
type WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized struct {
}

// NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized creates WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized() *WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized {

	return &WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
