// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchStorageV1CSINodeOKCode is the HTTP code returned for type WatchStorageV1CSINodeOK
const WatchStorageV1CSINodeOKCode int = 200

/*WatchStorageV1CSINodeOK OK

swagger:response watchStorageV1CSINodeOK
*/
type WatchStorageV1CSINodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1CSINodeOK creates WatchStorageV1CSINodeOK with default headers values
func NewWatchStorageV1CSINodeOK() *WatchStorageV1CSINodeOK {

	return &WatchStorageV1CSINodeOK{}
}

// WithPayload adds the payload to the watch storage v1 c s i node o k response
func (o *WatchStorageV1CSINodeOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1CSINodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1 c s i node o k response
func (o *WatchStorageV1CSINodeOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1CSINodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1CSINodeUnauthorizedCode is the HTTP code returned for type WatchStorageV1CSINodeUnauthorized
const WatchStorageV1CSINodeUnauthorizedCode int = 401

/*WatchStorageV1CSINodeUnauthorized Unauthorized

swagger:response watchStorageV1CSINodeUnauthorized
*/
type WatchStorageV1CSINodeUnauthorized struct {
}

// NewWatchStorageV1CSINodeUnauthorized creates WatchStorageV1CSINodeUnauthorized with default headers values
func NewWatchStorageV1CSINodeUnauthorized() *WatchStorageV1CSINodeUnauthorized {

	return &WatchStorageV1CSINodeUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1CSINodeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
