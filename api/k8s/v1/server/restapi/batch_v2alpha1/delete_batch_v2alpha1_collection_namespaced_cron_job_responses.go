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

// DeleteBatchV2alpha1CollectionNamespacedCronJobOKCode is the HTTP code returned for type DeleteBatchV2alpha1CollectionNamespacedCronJobOK
const DeleteBatchV2alpha1CollectionNamespacedCronJobOKCode int = 200

/*DeleteBatchV2alpha1CollectionNamespacedCronJobOK OK

swagger:response deleteBatchV2alpha1CollectionNamespacedCronJobOK
*/
type DeleteBatchV2alpha1CollectionNamespacedCronJobOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteBatchV2alpha1CollectionNamespacedCronJobOK creates DeleteBatchV2alpha1CollectionNamespacedCronJobOK with default headers values
func NewDeleteBatchV2alpha1CollectionNamespacedCronJobOK() *DeleteBatchV2alpha1CollectionNamespacedCronJobOK {

	return &DeleteBatchV2alpha1CollectionNamespacedCronJobOK{}
}

// WithPayload adds the payload to the delete batch v2alpha1 collection namespaced cron job o k response
func (o *DeleteBatchV2alpha1CollectionNamespacedCronJobOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteBatchV2alpha1CollectionNamespacedCronJobOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete batch v2alpha1 collection namespaced cron job o k response
func (o *DeleteBatchV2alpha1CollectionNamespacedCronJobOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteBatchV2alpha1CollectionNamespacedCronJobOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorizedCode is the HTTP code returned for type DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized
const DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorizedCode int = 401

/*DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized Unauthorized

swagger:response deleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized
*/
type DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized struct {
}

// NewDeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized creates DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized with default headers values
func NewDeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized() *DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized {

	return &DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteBatchV2alpha1CollectionNamespacedCronJobUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
