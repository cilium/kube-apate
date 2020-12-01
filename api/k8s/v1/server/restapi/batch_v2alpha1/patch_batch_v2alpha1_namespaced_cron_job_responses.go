// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchBatchV2alpha1NamespacedCronJobOKCode is the HTTP code returned for type PatchBatchV2alpha1NamespacedCronJobOK
const PatchBatchV2alpha1NamespacedCronJobOKCode int = 200

/*PatchBatchV2alpha1NamespacedCronJobOK OK

swagger:response patchBatchV2alpha1NamespacedCronJobOK
*/
type PatchBatchV2alpha1NamespacedCronJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV2alpha1CronJob `json:"body,omitempty"`
}

// NewPatchBatchV2alpha1NamespacedCronJobOK creates PatchBatchV2alpha1NamespacedCronJobOK with default headers values
func NewPatchBatchV2alpha1NamespacedCronJobOK() *PatchBatchV2alpha1NamespacedCronJobOK {

	return &PatchBatchV2alpha1NamespacedCronJobOK{}
}

// WithPayload adds the payload to the patch batch v2alpha1 namespaced cron job o k response
func (o *PatchBatchV2alpha1NamespacedCronJobOK) WithPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) *PatchBatchV2alpha1NamespacedCronJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch batch v2alpha1 namespaced cron job o k response
func (o *PatchBatchV2alpha1NamespacedCronJobOK) SetPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBatchV2alpha1NamespacedCronJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchBatchV2alpha1NamespacedCronJobUnauthorizedCode is the HTTP code returned for type PatchBatchV2alpha1NamespacedCronJobUnauthorized
const PatchBatchV2alpha1NamespacedCronJobUnauthorizedCode int = 401

/*PatchBatchV2alpha1NamespacedCronJobUnauthorized Unauthorized

swagger:response patchBatchV2alpha1NamespacedCronJobUnauthorized
*/
type PatchBatchV2alpha1NamespacedCronJobUnauthorized struct {
}

// NewPatchBatchV2alpha1NamespacedCronJobUnauthorized creates PatchBatchV2alpha1NamespacedCronJobUnauthorized with default headers values
func NewPatchBatchV2alpha1NamespacedCronJobUnauthorized() *PatchBatchV2alpha1NamespacedCronJobUnauthorized {

	return &PatchBatchV2alpha1NamespacedCronJobUnauthorized{}
}

// WriteResponse to the client
func (o *PatchBatchV2alpha1NamespacedCronJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
