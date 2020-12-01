// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateNetworkingV1beta1IngressClassOKCode is the HTTP code returned for type CreateNetworkingV1beta1IngressClassOK
const CreateNetworkingV1beta1IngressClassOKCode int = 200

/*CreateNetworkingV1beta1IngressClassOK OK

swagger:response createNetworkingV1beta1IngressClassOK
*/
type CreateNetworkingV1beta1IngressClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1IngressClass `json:"body,omitempty"`
}

// NewCreateNetworkingV1beta1IngressClassOK creates CreateNetworkingV1beta1IngressClassOK with default headers values
func NewCreateNetworkingV1beta1IngressClassOK() *CreateNetworkingV1beta1IngressClassOK {

	return &CreateNetworkingV1beta1IngressClassOK{}
}

// WithPayload adds the payload to the create networking v1beta1 ingress class o k response
func (o *CreateNetworkingV1beta1IngressClassOK) WithPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) *CreateNetworkingV1beta1IngressClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create networking v1beta1 ingress class o k response
func (o *CreateNetworkingV1beta1IngressClassOK) SetPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNetworkingV1beta1IngressClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNetworkingV1beta1IngressClassCreatedCode is the HTTP code returned for type CreateNetworkingV1beta1IngressClassCreated
const CreateNetworkingV1beta1IngressClassCreatedCode int = 201

/*CreateNetworkingV1beta1IngressClassCreated Created

swagger:response createNetworkingV1beta1IngressClassCreated
*/
type CreateNetworkingV1beta1IngressClassCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1IngressClass `json:"body,omitempty"`
}

// NewCreateNetworkingV1beta1IngressClassCreated creates CreateNetworkingV1beta1IngressClassCreated with default headers values
func NewCreateNetworkingV1beta1IngressClassCreated() *CreateNetworkingV1beta1IngressClassCreated {

	return &CreateNetworkingV1beta1IngressClassCreated{}
}

// WithPayload adds the payload to the create networking v1beta1 ingress class created response
func (o *CreateNetworkingV1beta1IngressClassCreated) WithPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) *CreateNetworkingV1beta1IngressClassCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create networking v1beta1 ingress class created response
func (o *CreateNetworkingV1beta1IngressClassCreated) SetPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNetworkingV1beta1IngressClassCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNetworkingV1beta1IngressClassAcceptedCode is the HTTP code returned for type CreateNetworkingV1beta1IngressClassAccepted
const CreateNetworkingV1beta1IngressClassAcceptedCode int = 202

/*CreateNetworkingV1beta1IngressClassAccepted Accepted

swagger:response createNetworkingV1beta1IngressClassAccepted
*/
type CreateNetworkingV1beta1IngressClassAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPINetworkingV1beta1IngressClass `json:"body,omitempty"`
}

// NewCreateNetworkingV1beta1IngressClassAccepted creates CreateNetworkingV1beta1IngressClassAccepted with default headers values
func NewCreateNetworkingV1beta1IngressClassAccepted() *CreateNetworkingV1beta1IngressClassAccepted {

	return &CreateNetworkingV1beta1IngressClassAccepted{}
}

// WithPayload adds the payload to the create networking v1beta1 ingress class accepted response
func (o *CreateNetworkingV1beta1IngressClassAccepted) WithPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) *CreateNetworkingV1beta1IngressClassAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create networking v1beta1 ingress class accepted response
func (o *CreateNetworkingV1beta1IngressClassAccepted) SetPayload(payload *models.IoK8sAPINetworkingV1beta1IngressClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateNetworkingV1beta1IngressClassAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateNetworkingV1beta1IngressClassUnauthorizedCode is the HTTP code returned for type CreateNetworkingV1beta1IngressClassUnauthorized
const CreateNetworkingV1beta1IngressClassUnauthorizedCode int = 401

/*CreateNetworkingV1beta1IngressClassUnauthorized Unauthorized

swagger:response createNetworkingV1beta1IngressClassUnauthorized
*/
type CreateNetworkingV1beta1IngressClassUnauthorized struct {
}

// NewCreateNetworkingV1beta1IngressClassUnauthorized creates CreateNetworkingV1beta1IngressClassUnauthorized with default headers values
func NewCreateNetworkingV1beta1IngressClassUnauthorized() *CreateNetworkingV1beta1IngressClassUnauthorized {

	return &CreateNetworkingV1beta1IngressClassUnauthorized{}
}

// WriteResponse to the client
func (o *CreateNetworkingV1beta1IngressClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
