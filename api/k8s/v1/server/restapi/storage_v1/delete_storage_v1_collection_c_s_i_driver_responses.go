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

// DeleteStorageV1CollectionCSIDriverOKCode is the HTTP code returned for type DeleteStorageV1CollectionCSIDriverOK
const DeleteStorageV1CollectionCSIDriverOKCode int = 200

/*DeleteStorageV1CollectionCSIDriverOK OK

swagger:response deleteStorageV1CollectionCSIDriverOK
*/
type DeleteStorageV1CollectionCSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteStorageV1CollectionCSIDriverOK creates DeleteStorageV1CollectionCSIDriverOK with default headers values
func NewDeleteStorageV1CollectionCSIDriverOK() *DeleteStorageV1CollectionCSIDriverOK {

	return &DeleteStorageV1CollectionCSIDriverOK{}
}

// WithPayload adds the payload to the delete storage v1 collection c s i driver o k response
func (o *DeleteStorageV1CollectionCSIDriverOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteStorageV1CollectionCSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete storage v1 collection c s i driver o k response
func (o *DeleteStorageV1CollectionCSIDriverOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStorageV1CollectionCSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteStorageV1CollectionCSIDriverUnauthorizedCode is the HTTP code returned for type DeleteStorageV1CollectionCSIDriverUnauthorized
const DeleteStorageV1CollectionCSIDriverUnauthorizedCode int = 401

/*DeleteStorageV1CollectionCSIDriverUnauthorized Unauthorized

swagger:response deleteStorageV1CollectionCSIDriverUnauthorized
*/
type DeleteStorageV1CollectionCSIDriverUnauthorized struct {
}

// NewDeleteStorageV1CollectionCSIDriverUnauthorized creates DeleteStorageV1CollectionCSIDriverUnauthorized with default headers values
func NewDeleteStorageV1CollectionCSIDriverUnauthorized() *DeleteStorageV1CollectionCSIDriverUnauthorized {

	return &DeleteStorageV1CollectionCSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteStorageV1CollectionCSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
