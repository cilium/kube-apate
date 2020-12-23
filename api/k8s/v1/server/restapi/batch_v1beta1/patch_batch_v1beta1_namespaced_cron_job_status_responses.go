// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchBatchV1beta1NamespacedCronJobStatusOKCode is the HTTP code returned for type PatchBatchV1beta1NamespacedCronJobStatusOK
const PatchBatchV1beta1NamespacedCronJobStatusOKCode int = 200

/*PatchBatchV1beta1NamespacedCronJobStatusOK OK

swagger:response patchBatchV1beta1NamespacedCronJobStatusOK
*/
type PatchBatchV1beta1NamespacedCronJobStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJob `json:"body,omitempty"`
}

// NewPatchBatchV1beta1NamespacedCronJobStatusOK creates PatchBatchV1beta1NamespacedCronJobStatusOK with default headers values
func NewPatchBatchV1beta1NamespacedCronJobStatusOK() *PatchBatchV1beta1NamespacedCronJobStatusOK {

	return &PatchBatchV1beta1NamespacedCronJobStatusOK{}
}

// WithPayload adds the payload to the patch batch v1beta1 namespaced cron job status o k response
func (o *PatchBatchV1beta1NamespacedCronJobStatusOK) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) *PatchBatchV1beta1NamespacedCronJobStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch batch v1beta1 namespaced cron job status o k response
func (o *PatchBatchV1beta1NamespacedCronJobStatusOK) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchBatchV1beta1NamespacedCronJobStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchBatchV1beta1NamespacedCronJobStatusUnauthorizedCode is the HTTP code returned for type PatchBatchV1beta1NamespacedCronJobStatusUnauthorized
const PatchBatchV1beta1NamespacedCronJobStatusUnauthorizedCode int = 401

/*PatchBatchV1beta1NamespacedCronJobStatusUnauthorized Unauthorized

swagger:response patchBatchV1beta1NamespacedCronJobStatusUnauthorized
*/
type PatchBatchV1beta1NamespacedCronJobStatusUnauthorized struct {
}

// NewPatchBatchV1beta1NamespacedCronJobStatusUnauthorized creates PatchBatchV1beta1NamespacedCronJobStatusUnauthorized with default headers values
func NewPatchBatchV1beta1NamespacedCronJobStatusUnauthorized() *PatchBatchV1beta1NamespacedCronJobStatusUnauthorized {

	return &PatchBatchV1beta1NamespacedCronJobStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchBatchV1beta1NamespacedCronJobStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}