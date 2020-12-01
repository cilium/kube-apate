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

// CreateStorageV1VolumeAttachmentOKCode is the HTTP code returned for type CreateStorageV1VolumeAttachmentOK
const CreateStorageV1VolumeAttachmentOKCode int = 200

/*CreateStorageV1VolumeAttachmentOK OK

swagger:response createStorageV1VolumeAttachmentOK
*/
type CreateStorageV1VolumeAttachmentOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachment `json:"body,omitempty"`
}

// NewCreateStorageV1VolumeAttachmentOK creates CreateStorageV1VolumeAttachmentOK with default headers values
func NewCreateStorageV1VolumeAttachmentOK() *CreateStorageV1VolumeAttachmentOK {

	return &CreateStorageV1VolumeAttachmentOK{}
}

// WithPayload adds the payload to the create storage v1 volume attachment o k response
func (o *CreateStorageV1VolumeAttachmentOK) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) *CreateStorageV1VolumeAttachmentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 volume attachment o k response
func (o *CreateStorageV1VolumeAttachmentOK) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1VolumeAttachmentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1VolumeAttachmentCreatedCode is the HTTP code returned for type CreateStorageV1VolumeAttachmentCreated
const CreateStorageV1VolumeAttachmentCreatedCode int = 201

/*CreateStorageV1VolumeAttachmentCreated Created

swagger:response createStorageV1VolumeAttachmentCreated
*/
type CreateStorageV1VolumeAttachmentCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachment `json:"body,omitempty"`
}

// NewCreateStorageV1VolumeAttachmentCreated creates CreateStorageV1VolumeAttachmentCreated with default headers values
func NewCreateStorageV1VolumeAttachmentCreated() *CreateStorageV1VolumeAttachmentCreated {

	return &CreateStorageV1VolumeAttachmentCreated{}
}

// WithPayload adds the payload to the create storage v1 volume attachment created response
func (o *CreateStorageV1VolumeAttachmentCreated) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) *CreateStorageV1VolumeAttachmentCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 volume attachment created response
func (o *CreateStorageV1VolumeAttachmentCreated) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1VolumeAttachmentCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1VolumeAttachmentAcceptedCode is the HTTP code returned for type CreateStorageV1VolumeAttachmentAccepted
const CreateStorageV1VolumeAttachmentAcceptedCode int = 202

/*CreateStorageV1VolumeAttachmentAccepted Accepted

swagger:response createStorageV1VolumeAttachmentAccepted
*/
type CreateStorageV1VolumeAttachmentAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1VolumeAttachment `json:"body,omitempty"`
}

// NewCreateStorageV1VolumeAttachmentAccepted creates CreateStorageV1VolumeAttachmentAccepted with default headers values
func NewCreateStorageV1VolumeAttachmentAccepted() *CreateStorageV1VolumeAttachmentAccepted {

	return &CreateStorageV1VolumeAttachmentAccepted{}
}

// WithPayload adds the payload to the create storage v1 volume attachment accepted response
func (o *CreateStorageV1VolumeAttachmentAccepted) WithPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) *CreateStorageV1VolumeAttachmentAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1 volume attachment accepted response
func (o *CreateStorageV1VolumeAttachmentAccepted) SetPayload(payload *models.IoK8sAPIStorageV1VolumeAttachment) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1VolumeAttachmentAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1VolumeAttachmentUnauthorizedCode is the HTTP code returned for type CreateStorageV1VolumeAttachmentUnauthorized
const CreateStorageV1VolumeAttachmentUnauthorizedCode int = 401

/*CreateStorageV1VolumeAttachmentUnauthorized Unauthorized

swagger:response createStorageV1VolumeAttachmentUnauthorized
*/
type CreateStorageV1VolumeAttachmentUnauthorized struct {
}

// NewCreateStorageV1VolumeAttachmentUnauthorized creates CreateStorageV1VolumeAttachmentUnauthorized with default headers values
func NewCreateStorageV1VolumeAttachmentUnauthorized() *CreateStorageV1VolumeAttachmentUnauthorized {

	return &CreateStorageV1VolumeAttachmentUnauthorized{}
}

// WriteResponse to the client
func (o *CreateStorageV1VolumeAttachmentUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
