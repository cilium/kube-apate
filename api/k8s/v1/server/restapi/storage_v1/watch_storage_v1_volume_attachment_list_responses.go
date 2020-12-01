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

// WatchStorageV1VolumeAttachmentListOKCode is the HTTP code returned for type WatchStorageV1VolumeAttachmentListOK
const WatchStorageV1VolumeAttachmentListOKCode int = 200

/*WatchStorageV1VolumeAttachmentListOK OK

swagger:response watchStorageV1VolumeAttachmentListOK
*/
type WatchStorageV1VolumeAttachmentListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1VolumeAttachmentListOK creates WatchStorageV1VolumeAttachmentListOK with default headers values
func NewWatchStorageV1VolumeAttachmentListOK() *WatchStorageV1VolumeAttachmentListOK {

	return &WatchStorageV1VolumeAttachmentListOK{}
}

// WithPayload adds the payload to the watch storage v1 volume attachment list o k response
func (o *WatchStorageV1VolumeAttachmentListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1VolumeAttachmentListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1 volume attachment list o k response
func (o *WatchStorageV1VolumeAttachmentListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1VolumeAttachmentListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1VolumeAttachmentListUnauthorizedCode is the HTTP code returned for type WatchStorageV1VolumeAttachmentListUnauthorized
const WatchStorageV1VolumeAttachmentListUnauthorizedCode int = 401

/*WatchStorageV1VolumeAttachmentListUnauthorized Unauthorized

swagger:response watchStorageV1VolumeAttachmentListUnauthorized
*/
type WatchStorageV1VolumeAttachmentListUnauthorized struct {
}

// NewWatchStorageV1VolumeAttachmentListUnauthorized creates WatchStorageV1VolumeAttachmentListUnauthorized with default headers values
func NewWatchStorageV1VolumeAttachmentListUnauthorized() *WatchStorageV1VolumeAttachmentListUnauthorized {

	return &WatchStorageV1VolumeAttachmentListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1VolumeAttachmentListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
