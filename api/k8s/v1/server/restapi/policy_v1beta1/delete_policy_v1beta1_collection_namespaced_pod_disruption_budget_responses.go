// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOKCode is the HTTP code returned for type DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK
const DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOKCode int = 200

/*DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK OK

swagger:response deletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK
*/
type DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK creates DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK with default headers values
func NewDeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK() *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK {

	return &DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK{}
}

// WithPayload adds the payload to the delete policy v1beta1 collection namespaced pod disruption budget o k response
func (o *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy v1beta1 collection namespaced pod disruption budget o k response
func (o *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorizedCode is the HTTP code returned for type DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized
const DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorizedCode int = 401

/*DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized Unauthorized

swagger:response deletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized
*/
type DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized struct {
}

// NewDeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized creates DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized with default headers values
func NewDeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized() *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized {

	return &DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized{}
}

// WriteResponse to the client
func (o *DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
