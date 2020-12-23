// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateStorageV1beta1CSIDriverOKCode is the HTTP code returned for type CreateStorageV1beta1CSIDriverOK
const CreateStorageV1beta1CSIDriverOKCode int = 200

/*CreateStorageV1beta1CSIDriverOK OK

swagger:response createStorageV1beta1CSIDriverOK
*/
type CreateStorageV1beta1CSIDriverOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewCreateStorageV1beta1CSIDriverOK creates CreateStorageV1beta1CSIDriverOK with default headers values
func NewCreateStorageV1beta1CSIDriverOK() *CreateStorageV1beta1CSIDriverOK {

	return &CreateStorageV1beta1CSIDriverOK{}
}

// WithPayload adds the payload to the create storage v1beta1 c s i driver o k response
func (o *CreateStorageV1beta1CSIDriverOK) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *CreateStorageV1beta1CSIDriverOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1beta1 c s i driver o k response
func (o *CreateStorageV1beta1CSIDriverOK) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1beta1CSIDriverOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1beta1CSIDriverCreatedCode is the HTTP code returned for type CreateStorageV1beta1CSIDriverCreated
const CreateStorageV1beta1CSIDriverCreatedCode int = 201

/*CreateStorageV1beta1CSIDriverCreated Created

swagger:response createStorageV1beta1CSIDriverCreated
*/
type CreateStorageV1beta1CSIDriverCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewCreateStorageV1beta1CSIDriverCreated creates CreateStorageV1beta1CSIDriverCreated with default headers values
func NewCreateStorageV1beta1CSIDriverCreated() *CreateStorageV1beta1CSIDriverCreated {

	return &CreateStorageV1beta1CSIDriverCreated{}
}

// WithPayload adds the payload to the create storage v1beta1 c s i driver created response
func (o *CreateStorageV1beta1CSIDriverCreated) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *CreateStorageV1beta1CSIDriverCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1beta1 c s i driver created response
func (o *CreateStorageV1beta1CSIDriverCreated) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1beta1CSIDriverCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1beta1CSIDriverAcceptedCode is the HTTP code returned for type CreateStorageV1beta1CSIDriverAccepted
const CreateStorageV1beta1CSIDriverAcceptedCode int = 202

/*CreateStorageV1beta1CSIDriverAccepted Accepted

swagger:response createStorageV1beta1CSIDriverAccepted
*/
type CreateStorageV1beta1CSIDriverAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver `json:"body,omitempty"`
}

// NewCreateStorageV1beta1CSIDriverAccepted creates CreateStorageV1beta1CSIDriverAccepted with default headers values
func NewCreateStorageV1beta1CSIDriverAccepted() *CreateStorageV1beta1CSIDriverAccepted {

	return &CreateStorageV1beta1CSIDriverAccepted{}
}

// WithPayload adds the payload to the create storage v1beta1 c s i driver accepted response
func (o *CreateStorageV1beta1CSIDriverAccepted) WithPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) *CreateStorageV1beta1CSIDriverAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create storage v1beta1 c s i driver accepted response
func (o *CreateStorageV1beta1CSIDriverAccepted) SetPayload(payload *models.IoK8sAPIStorageV1beta1CSIDriver) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateStorageV1beta1CSIDriverAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateStorageV1beta1CSIDriverUnauthorizedCode is the HTTP code returned for type CreateStorageV1beta1CSIDriverUnauthorized
const CreateStorageV1beta1CSIDriverUnauthorizedCode int = 401

/*CreateStorageV1beta1CSIDriverUnauthorized Unauthorized

swagger:response createStorageV1beta1CSIDriverUnauthorized
*/
type CreateStorageV1beta1CSIDriverUnauthorized struct {
}

// NewCreateStorageV1beta1CSIDriverUnauthorized creates CreateStorageV1beta1CSIDriverUnauthorized with default headers values
func NewCreateStorageV1beta1CSIDriverUnauthorized() *CreateStorageV1beta1CSIDriverUnauthorized {

	return &CreateStorageV1beta1CSIDriverUnauthorized{}
}

// WriteResponse to the client
func (o *CreateStorageV1beta1CSIDriverUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}