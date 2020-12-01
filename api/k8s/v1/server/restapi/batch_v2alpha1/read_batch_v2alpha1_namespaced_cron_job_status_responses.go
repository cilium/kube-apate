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

// ReadBatchV2alpha1NamespacedCronJobStatusOKCode is the HTTP code returned for type ReadBatchV2alpha1NamespacedCronJobStatusOK
const ReadBatchV2alpha1NamespacedCronJobStatusOKCode int = 200

/*ReadBatchV2alpha1NamespacedCronJobStatusOK OK

swagger:response readBatchV2alpha1NamespacedCronJobStatusOK
*/
type ReadBatchV2alpha1NamespacedCronJobStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV2alpha1CronJob `json:"body,omitempty"`
}

// NewReadBatchV2alpha1NamespacedCronJobStatusOK creates ReadBatchV2alpha1NamespacedCronJobStatusOK with default headers values
func NewReadBatchV2alpha1NamespacedCronJobStatusOK() *ReadBatchV2alpha1NamespacedCronJobStatusOK {

	return &ReadBatchV2alpha1NamespacedCronJobStatusOK{}
}

// WithPayload adds the payload to the read batch v2alpha1 namespaced cron job status o k response
func (o *ReadBatchV2alpha1NamespacedCronJobStatusOK) WithPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) *ReadBatchV2alpha1NamespacedCronJobStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read batch v2alpha1 namespaced cron job status o k response
func (o *ReadBatchV2alpha1NamespacedCronJobStatusOK) SetPayload(payload *models.IoK8sAPIBatchV2alpha1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadBatchV2alpha1NamespacedCronJobStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadBatchV2alpha1NamespacedCronJobStatusUnauthorizedCode is the HTTP code returned for type ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized
const ReadBatchV2alpha1NamespacedCronJobStatusUnauthorizedCode int = 401

/*ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized Unauthorized

swagger:response readBatchV2alpha1NamespacedCronJobStatusUnauthorized
*/
type ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized struct {
}

// NewReadBatchV2alpha1NamespacedCronJobStatusUnauthorized creates ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized with default headers values
func NewReadBatchV2alpha1NamespacedCronJobStatusUnauthorized() *ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized {

	return &ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReadBatchV2alpha1NamespacedCronJobStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
