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

// WatchBatchV1beta1NamespacedCronJobListOKCode is the HTTP code returned for type WatchBatchV1beta1NamespacedCronJobListOK
const WatchBatchV1beta1NamespacedCronJobListOKCode int = 200

/*WatchBatchV1beta1NamespacedCronJobListOK OK

swagger:response watchBatchV1beta1NamespacedCronJobListOK
*/
type WatchBatchV1beta1NamespacedCronJobListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchBatchV1beta1NamespacedCronJobListOK creates WatchBatchV1beta1NamespacedCronJobListOK with default headers values
func NewWatchBatchV1beta1NamespacedCronJobListOK() *WatchBatchV1beta1NamespacedCronJobListOK {

	return &WatchBatchV1beta1NamespacedCronJobListOK{}
}

// WithPayload adds the payload to the watch batch v1beta1 namespaced cron job list o k response
func (o *WatchBatchV1beta1NamespacedCronJobListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchBatchV1beta1NamespacedCronJobListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch batch v1beta1 namespaced cron job list o k response
func (o *WatchBatchV1beta1NamespacedCronJobListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchBatchV1beta1NamespacedCronJobListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchBatchV1beta1NamespacedCronJobListUnauthorizedCode is the HTTP code returned for type WatchBatchV1beta1NamespacedCronJobListUnauthorized
const WatchBatchV1beta1NamespacedCronJobListUnauthorizedCode int = 401

/*WatchBatchV1beta1NamespacedCronJobListUnauthorized Unauthorized

swagger:response watchBatchV1beta1NamespacedCronJobListUnauthorized
*/
type WatchBatchV1beta1NamespacedCronJobListUnauthorized struct {
}

// NewWatchBatchV1beta1NamespacedCronJobListUnauthorized creates WatchBatchV1beta1NamespacedCronJobListUnauthorized with default headers values
func NewWatchBatchV1beta1NamespacedCronJobListUnauthorized() *WatchBatchV1beta1NamespacedCronJobListUnauthorized {

	return &WatchBatchV1beta1NamespacedCronJobListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchBatchV1beta1NamespacedCronJobListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
