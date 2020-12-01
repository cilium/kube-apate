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

// DeleteCoreV1PersistentVolumeOKCode is the HTTP code returned for type DeleteCoreV1PersistentVolumeOK
const DeleteCoreV1PersistentVolumeOKCode int = 200

/*DeleteCoreV1PersistentVolumeOK OK

swagger:response deleteCoreV1PersistentVolumeOK
*/
type DeleteCoreV1PersistentVolumeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolume `json:"body,omitempty"`
}

// NewDeleteCoreV1PersistentVolumeOK creates DeleteCoreV1PersistentVolumeOK with default headers values
func NewDeleteCoreV1PersistentVolumeOK() *DeleteCoreV1PersistentVolumeOK {

	return &DeleteCoreV1PersistentVolumeOK{}
}

// WithPayload adds the payload to the delete core v1 persistent volume o k response
func (o *DeleteCoreV1PersistentVolumeOK) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolume) *DeleteCoreV1PersistentVolumeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 persistent volume o k response
func (o *DeleteCoreV1PersistentVolumeOK) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1PersistentVolumeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1PersistentVolumeAcceptedCode is the HTTP code returned for type DeleteCoreV1PersistentVolumeAccepted
const DeleteCoreV1PersistentVolumeAcceptedCode int = 202

/*DeleteCoreV1PersistentVolumeAccepted Accepted

swagger:response deleteCoreV1PersistentVolumeAccepted
*/
type DeleteCoreV1PersistentVolumeAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PersistentVolume `json:"body,omitempty"`
}

// NewDeleteCoreV1PersistentVolumeAccepted creates DeleteCoreV1PersistentVolumeAccepted with default headers values
func NewDeleteCoreV1PersistentVolumeAccepted() *DeleteCoreV1PersistentVolumeAccepted {

	return &DeleteCoreV1PersistentVolumeAccepted{}
}

// WithPayload adds the payload to the delete core v1 persistent volume accepted response
func (o *DeleteCoreV1PersistentVolumeAccepted) WithPayload(payload *models.IoK8sAPICoreV1PersistentVolume) *DeleteCoreV1PersistentVolumeAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete core v1 persistent volume accepted response
func (o *DeleteCoreV1PersistentVolumeAccepted) SetPayload(payload *models.IoK8sAPICoreV1PersistentVolume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCoreV1PersistentVolumeAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCoreV1PersistentVolumeUnauthorizedCode is the HTTP code returned for type DeleteCoreV1PersistentVolumeUnauthorized
const DeleteCoreV1PersistentVolumeUnauthorizedCode int = 401

/*DeleteCoreV1PersistentVolumeUnauthorized Unauthorized

swagger:response deleteCoreV1PersistentVolumeUnauthorized
*/
type DeleteCoreV1PersistentVolumeUnauthorized struct {
}

// NewDeleteCoreV1PersistentVolumeUnauthorized creates DeleteCoreV1PersistentVolumeUnauthorized with default headers values
func NewDeleteCoreV1PersistentVolumeUnauthorized() *DeleteCoreV1PersistentVolumeUnauthorized {

	return &DeleteCoreV1PersistentVolumeUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCoreV1PersistentVolumeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
