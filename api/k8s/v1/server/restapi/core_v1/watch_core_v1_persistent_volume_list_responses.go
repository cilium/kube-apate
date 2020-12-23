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

// WatchCoreV1PersistentVolumeListOKCode is the HTTP code returned for type WatchCoreV1PersistentVolumeListOK
const WatchCoreV1PersistentVolumeListOKCode int = 200

/*WatchCoreV1PersistentVolumeListOK OK

swagger:response watchCoreV1PersistentVolumeListOK
*/
type WatchCoreV1PersistentVolumeListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCoreV1PersistentVolumeListOK creates WatchCoreV1PersistentVolumeListOK with default headers values
func NewWatchCoreV1PersistentVolumeListOK() *WatchCoreV1PersistentVolumeListOK {

	return &WatchCoreV1PersistentVolumeListOK{}
}

// WithPayload adds the payload to the watch core v1 persistent volume list o k response
func (o *WatchCoreV1PersistentVolumeListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCoreV1PersistentVolumeListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch core v1 persistent volume list o k response
func (o *WatchCoreV1PersistentVolumeListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCoreV1PersistentVolumeListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCoreV1PersistentVolumeListUnauthorizedCode is the HTTP code returned for type WatchCoreV1PersistentVolumeListUnauthorized
const WatchCoreV1PersistentVolumeListUnauthorizedCode int = 401

/*WatchCoreV1PersistentVolumeListUnauthorized Unauthorized

swagger:response watchCoreV1PersistentVolumeListUnauthorized
*/
type WatchCoreV1PersistentVolumeListUnauthorized struct {
}

// NewWatchCoreV1PersistentVolumeListUnauthorized creates WatchCoreV1PersistentVolumeListUnauthorized with default headers values
func NewWatchCoreV1PersistentVolumeListUnauthorized() *WatchCoreV1PersistentVolumeListUnauthorized {

	return &WatchCoreV1PersistentVolumeListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCoreV1PersistentVolumeListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}