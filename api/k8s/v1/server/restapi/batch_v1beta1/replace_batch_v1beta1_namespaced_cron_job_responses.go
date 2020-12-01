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

// ReplaceBatchV1beta1NamespacedCronJobOKCode is the HTTP code returned for type ReplaceBatchV1beta1NamespacedCronJobOK
const ReplaceBatchV1beta1NamespacedCronJobOKCode int = 200

/*ReplaceBatchV1beta1NamespacedCronJobOK OK

swagger:response replaceBatchV1beta1NamespacedCronJobOK
*/
type ReplaceBatchV1beta1NamespacedCronJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJob `json:"body,omitempty"`
}

// NewReplaceBatchV1beta1NamespacedCronJobOK creates ReplaceBatchV1beta1NamespacedCronJobOK with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobOK() *ReplaceBatchV1beta1NamespacedCronJobOK {

	return &ReplaceBatchV1beta1NamespacedCronJobOK{}
}

// WithPayload adds the payload to the replace batch v1beta1 namespaced cron job o k response
func (o *ReplaceBatchV1beta1NamespacedCronJobOK) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) *ReplaceBatchV1beta1NamespacedCronJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1beta1 namespaced cron job o k response
func (o *ReplaceBatchV1beta1NamespacedCronJobOK) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1beta1NamespacedCronJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1beta1NamespacedCronJobCreatedCode is the HTTP code returned for type ReplaceBatchV1beta1NamespacedCronJobCreated
const ReplaceBatchV1beta1NamespacedCronJobCreatedCode int = 201

/*ReplaceBatchV1beta1NamespacedCronJobCreated Created

swagger:response replaceBatchV1beta1NamespacedCronJobCreated
*/
type ReplaceBatchV1beta1NamespacedCronJobCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV1beta1CronJob `json:"body,omitempty"`
}

// NewReplaceBatchV1beta1NamespacedCronJobCreated creates ReplaceBatchV1beta1NamespacedCronJobCreated with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobCreated() *ReplaceBatchV1beta1NamespacedCronJobCreated {

	return &ReplaceBatchV1beta1NamespacedCronJobCreated{}
}

// WithPayload adds the payload to the replace batch v1beta1 namespaced cron job created response
func (o *ReplaceBatchV1beta1NamespacedCronJobCreated) WithPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) *ReplaceBatchV1beta1NamespacedCronJobCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace batch v1beta1 namespaced cron job created response
func (o *ReplaceBatchV1beta1NamespacedCronJobCreated) SetPayload(payload *models.IoK8sAPIBatchV1beta1CronJob) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBatchV1beta1NamespacedCronJobCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBatchV1beta1NamespacedCronJobUnauthorizedCode is the HTTP code returned for type ReplaceBatchV1beta1NamespacedCronJobUnauthorized
const ReplaceBatchV1beta1NamespacedCronJobUnauthorizedCode int = 401

/*ReplaceBatchV1beta1NamespacedCronJobUnauthorized Unauthorized

swagger:response replaceBatchV1beta1NamespacedCronJobUnauthorized
*/
type ReplaceBatchV1beta1NamespacedCronJobUnauthorized struct {
}

// NewReplaceBatchV1beta1NamespacedCronJobUnauthorized creates ReplaceBatchV1beta1NamespacedCronJobUnauthorized with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobUnauthorized() *ReplaceBatchV1beta1NamespacedCronJobUnauthorized {

	return &ReplaceBatchV1beta1NamespacedCronJobUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceBatchV1beta1NamespacedCronJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
