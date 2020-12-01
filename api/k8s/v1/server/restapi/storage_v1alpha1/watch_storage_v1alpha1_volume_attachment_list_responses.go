// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchStorageV1alpha1VolumeAttachmentListOKCode is the HTTP code returned for type WatchStorageV1alpha1VolumeAttachmentListOK
const WatchStorageV1alpha1VolumeAttachmentListOKCode int = 200

/*WatchStorageV1alpha1VolumeAttachmentListOK OK

swagger:response watchStorageV1alpha1VolumeAttachmentListOK
*/
type WatchStorageV1alpha1VolumeAttachmentListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchStorageV1alpha1VolumeAttachmentListOK creates WatchStorageV1alpha1VolumeAttachmentListOK with default headers values
func NewWatchStorageV1alpha1VolumeAttachmentListOK() *WatchStorageV1alpha1VolumeAttachmentListOK {

	return &WatchStorageV1alpha1VolumeAttachmentListOK{}
}

// WithPayload adds the payload to the watch storage v1alpha1 volume attachment list o k response
func (o *WatchStorageV1alpha1VolumeAttachmentListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchStorageV1alpha1VolumeAttachmentListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch storage v1alpha1 volume attachment list o k response
func (o *WatchStorageV1alpha1VolumeAttachmentListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchStorageV1alpha1VolumeAttachmentListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchStorageV1alpha1VolumeAttachmentListUnauthorizedCode is the HTTP code returned for type WatchStorageV1alpha1VolumeAttachmentListUnauthorized
const WatchStorageV1alpha1VolumeAttachmentListUnauthorizedCode int = 401

/*WatchStorageV1alpha1VolumeAttachmentListUnauthorized Unauthorized

swagger:response watchStorageV1alpha1VolumeAttachmentListUnauthorized
*/
type WatchStorageV1alpha1VolumeAttachmentListUnauthorized struct {
}

// NewWatchStorageV1alpha1VolumeAttachmentListUnauthorized creates WatchStorageV1alpha1VolumeAttachmentListUnauthorized with default headers values
func NewWatchStorageV1alpha1VolumeAttachmentListUnauthorized() *WatchStorageV1alpha1VolumeAttachmentListUnauthorized {

	return &WatchStorageV1alpha1VolumeAttachmentListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchStorageV1alpha1VolumeAttachmentListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
