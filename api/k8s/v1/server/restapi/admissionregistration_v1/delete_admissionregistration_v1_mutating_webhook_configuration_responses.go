// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteAdmissionregistrationV1MutatingWebhookConfigurationOKCode is the HTTP code returned for type DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK
const DeleteAdmissionregistrationV1MutatingWebhookConfigurationOKCode int = 200

/*DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK OK

swagger:response deleteAdmissionregistrationV1MutatingWebhookConfigurationOK
*/
type DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationOK creates DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK with default headers values
func NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationOK() *DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK {

	return &DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the delete admissionregistration v1 mutating webhook configuration o k response
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete admissionregistration v1 mutating webhook configuration o k response
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAdmissionregistrationV1MutatingWebhookConfigurationAcceptedCode is the HTTP code returned for type DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted
const DeleteAdmissionregistrationV1MutatingWebhookConfigurationAcceptedCode int = 202

/*DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted Accepted

swagger:response deleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted
*/
type DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted creates DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted with default headers values
func NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted() *DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted {

	return &DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted{}
}

// WithPayload adds the payload to the delete admissionregistration v1 mutating webhook configuration accepted response
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete admissionregistration v1 mutating webhook configuration accepted response
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized
const DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorizedCode int = 401

/*DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized Unauthorized

swagger:response deleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized
*/
type DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized struct {
}

// NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized creates DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized with default headers values
func NewDeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized() *DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized {

	return &DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteAdmissionregistrationV1MutatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}