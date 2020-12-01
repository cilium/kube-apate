// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode is the HTTP code returned for type ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
const ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOKCode int = 200

/*ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK OK

swagger:response replaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK
*/
type ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK creates ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK with default headers values
func NewReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK() *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {

	return &ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the replace admissionregistration v1beta1 mutating webhook configuration o k response
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace admissionregistration v1beta1 mutating webhook configuration o k response
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreatedCode is the HTTP code returned for type ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated
const ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreatedCode int = 201

/*ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated Created

swagger:response replaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated
*/
type ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration `json:"body,omitempty"`
}

// NewReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated creates ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated with default headers values
func NewReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated() *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated {

	return &ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated{}
}

// WithPayload adds the payload to the replace admissionregistration v1beta1 mutating webhook configuration created response
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace admissionregistration v1beta1 mutating webhook configuration created response
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1MutatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
const ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response replaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized
*/
type ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized struct {
}

// NewReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized creates ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized with default headers values
func NewReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized() *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized {

	return &ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
