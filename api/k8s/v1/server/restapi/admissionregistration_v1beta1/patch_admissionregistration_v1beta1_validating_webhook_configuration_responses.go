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

// PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOKCode is the HTTP code returned for type PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK
const PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOKCode int = 200

/*PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK OK

swagger:response patchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK
*/
type PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration `json:"body,omitempty"`
}

// NewPatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK creates PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK with default headers values
func NewPatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK() *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK {

	return &PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK{}
}

// WithPayload adds the payload to the patch admissionregistration v1beta1 validating webhook configuration o k response
func (o *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) WithPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch admissionregistration v1beta1 validating webhook configuration o k response
func (o *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) SetPayload(payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorizedCode is the HTTP code returned for type PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized
const PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorizedCode int = 401

/*PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized Unauthorized

swagger:response patchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized
*/
type PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized struct {
}

// NewPatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized creates PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewPatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized() *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized {

	return &PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
