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

// ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOKCode is the HTTP code returned for type ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK
const ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOKCode int = 200

/*ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK OK

swagger:response replacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK
*/
type ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget `json:"body,omitempty"`
}

// NewReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK creates ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK with default headers values
func NewReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK() *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK {

	return &ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK{}
}

// WithPayload adds the payload to the replace policy v1beta1 namespaced pod disruption budget status o k response
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget) *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace policy v1beta1 namespaced pod disruption budget status o k response
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreatedCode is the HTTP code returned for type ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated
const ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreatedCode int = 201

/*ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated Created

swagger:response replacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated
*/
type ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget `json:"body,omitempty"`
}

// NewReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated creates ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated with default headers values
func NewReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated() *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated {

	return &ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated{}
}

// WithPayload adds the payload to the replace policy v1beta1 namespaced pod disruption budget status created response
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget) *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace policy v1beta1 namespaced pod disruption budget status created response
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorizedCode is the HTTP code returned for type ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized
const ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorizedCode int = 401

/*ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized Unauthorized

swagger:response replacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized
*/
type ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized struct {
}

// NewReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized creates ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized with default headers values
func NewReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized() *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized {

	return &ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}