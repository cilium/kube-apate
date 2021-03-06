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

// ReplaceCoreV1NamespaceStatusOKCode is the HTTP code returned for type ReplaceCoreV1NamespaceStatusOK
const ReplaceCoreV1NamespaceStatusOKCode int = 200

/*ReplaceCoreV1NamespaceStatusOK OK

swagger:response replaceCoreV1NamespaceStatusOK
*/
type ReplaceCoreV1NamespaceStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespaceStatusOK creates ReplaceCoreV1NamespaceStatusOK with default headers values
func NewReplaceCoreV1NamespaceStatusOK() *ReplaceCoreV1NamespaceStatusOK {

	return &ReplaceCoreV1NamespaceStatusOK{}
}

// WithPayload adds the payload to the replace core v1 namespace status o k response
func (o *ReplaceCoreV1NamespaceStatusOK) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *ReplaceCoreV1NamespaceStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespace status o k response
func (o *ReplaceCoreV1NamespaceStatusOK) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespaceStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespaceStatusCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespaceStatusCreated
const ReplaceCoreV1NamespaceStatusCreatedCode int = 201

/*ReplaceCoreV1NamespaceStatusCreated Created

swagger:response replaceCoreV1NamespaceStatusCreated
*/
type ReplaceCoreV1NamespaceStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1Namespace `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespaceStatusCreated creates ReplaceCoreV1NamespaceStatusCreated with default headers values
func NewReplaceCoreV1NamespaceStatusCreated() *ReplaceCoreV1NamespaceStatusCreated {

	return &ReplaceCoreV1NamespaceStatusCreated{}
}

// WithPayload adds the payload to the replace core v1 namespace status created response
func (o *ReplaceCoreV1NamespaceStatusCreated) WithPayload(payload *models.IoK8sAPICoreV1Namespace) *ReplaceCoreV1NamespaceStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespace status created response
func (o *ReplaceCoreV1NamespaceStatusCreated) SetPayload(payload *models.IoK8sAPICoreV1Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespaceStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespaceStatusUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespaceStatusUnauthorized
const ReplaceCoreV1NamespaceStatusUnauthorizedCode int = 401

/*ReplaceCoreV1NamespaceStatusUnauthorized Unauthorized

swagger:response replaceCoreV1NamespaceStatusUnauthorized
*/
type ReplaceCoreV1NamespaceStatusUnauthorized struct {
}

// NewReplaceCoreV1NamespaceStatusUnauthorized creates ReplaceCoreV1NamespaceStatusUnauthorized with default headers values
func NewReplaceCoreV1NamespaceStatusUnauthorized() *ReplaceCoreV1NamespaceStatusUnauthorized {

	return &ReplaceCoreV1NamespaceStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespaceStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
