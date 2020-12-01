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

// ReplaceBatchV1beta1NamespacedCronJobStatusOKCode is the HTTP code returned for type ReplaceBatchV1beta1NamespacedCronJobStatusOK
const ReplaceBatchV1beta1NamespacedCronJobStatusOKCode int = 200

/*ReplaceBatchV1beta1NamespacedCronJobStatusOK OK

swagger:response replaceBatchV1beta1NamespacedCronJobStatusOK
*/
type ReplaceBatchV1beta1NamespacedCronJobStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJob `json:"body,omitempty"`
}

// NewReplaceBatchV1beta1NamespacedCronJobStatusOK creates ReplaceBatchV1beta1NamespacedCronJobStatusOK with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobStatusOK() *ReplaceBatchV1beta1NamespacedCronJobStatusOK {

	return &ReplaceBatchV1beta1NamespacedCronJobStatusOK{}
}

// WithPayload adds the payload to the replace batch v1beta1 namespaced cron job status o k response
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusOK) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) *ReplaceBatchV1beta1NamespacedCronJobStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1beta1 namespaced cron job status o k response
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusOK) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1beta1NamespacedCronJobStatusCreatedCode is the HTTP code returned for type ReplaceBatchV1beta1NamespacedCronJobStatusCreated
const ReplaceBatchV1beta1NamespacedCronJobStatusCreatedCode int = 201

/*ReplaceBatchV1beta1NamespacedCronJobStatusCreated Created

swagger:response replaceBatchV1beta1NamespacedCronJobStatusCreated
*/
type ReplaceBatchV1beta1NamespacedCronJobStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJob `json:"body,omitempty"`
}

// NewReplaceBatchV1beta1NamespacedCronJobStatusCreated creates ReplaceBatchV1beta1NamespacedCronJobStatusCreated with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobStatusCreated() *ReplaceBatchV1beta1NamespacedCronJobStatusCreated {

	return &ReplaceBatchV1beta1NamespacedCronJobStatusCreated{}
}

// WithPayload adds the payload to the replace batch v1beta1 namespaced cron job status created response
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusCreated) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) *ReplaceBatchV1beta1NamespacedCronJobStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1beta1 namespaced cron job status created response
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusCreated) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorizedCode is the HTTP code returned for type ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized
const ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorizedCode int = 401

/*ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized Unauthorized

swagger:response replaceBatchV1beta1NamespacedCronJobStatusUnauthorized
*/
type ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized struct {
}

// NewReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized creates ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized() *ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized {

	return &ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceBatchV1beta1NamespacedCronJobStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
