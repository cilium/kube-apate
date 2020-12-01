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

// ListBatchV2alpha1CronJobForAllNamespacesOKCode is the HTTP code returned for type ListBatchV2alpha1CronJobForAllNamespacesOK
const ListBatchV2alpha1CronJobForAllNamespacesOKCode int = 200

/*ListBatchV2alpha1CronJobForAllNamespacesOK OK

swagger:response listBatchV2alpha1CronJobForAllNamespacesOK
*/
type ListBatchV2alpha1CronJobForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIBatchV2alpha1CronJobList `json:"body,omitempty"`
}

// NewListBatchV2alpha1CronJobForAllNamespacesOK creates ListBatchV2alpha1CronJobForAllNamespacesOK with default headers values
func NewListBatchV2alpha1CronJobForAllNamespacesOK() *ListBatchV2alpha1CronJobForAllNamespacesOK {

	return &ListBatchV2alpha1CronJobForAllNamespacesOK{}
}

// WithPayload adds the payload to the list batch v2alpha1 cron job for all namespaces o k response
func (o *ListBatchV2alpha1CronJobForAllNamespacesOK) WithPayload(payload *models.IoK8sAPIBatchV2alpha1CronJobList) *ListBatchV2alpha1CronJobForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list batch v2alpha1 cron job for all namespaces o k response
func (o *ListBatchV2alpha1CronJobForAllNamespacesOK) SetPayload(payload *models.IoK8sAPIBatchV2alpha1CronJobList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListBatchV2alpha1CronJobForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListBatchV2alpha1CronJobForAllNamespacesUnauthorizedCode is the HTTP code returned for type ListBatchV2alpha1CronJobForAllNamespacesUnauthorized
const ListBatchV2alpha1CronJobForAllNamespacesUnauthorizedCode int = 401

/*ListBatchV2alpha1CronJobForAllNamespacesUnauthorized Unauthorized

swagger:response listBatchV2alpha1CronJobForAllNamespacesUnauthorized
*/
type ListBatchV2alpha1CronJobForAllNamespacesUnauthorized struct {
}

// NewListBatchV2alpha1CronJobForAllNamespacesUnauthorized creates ListBatchV2alpha1CronJobForAllNamespacesUnauthorized with default headers values
func NewListBatchV2alpha1CronJobForAllNamespacesUnauthorized() *ListBatchV2alpha1CronJobForAllNamespacesUnauthorized {

	return &ListBatchV2alpha1CronJobForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *ListBatchV2alpha1CronJobForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
