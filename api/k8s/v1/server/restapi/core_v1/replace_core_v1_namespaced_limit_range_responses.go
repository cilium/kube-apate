// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceCoreV1NamespacedLimitRangeOKCode is the HTTP code returned for type ReplaceCoreV1NamespacedLimitRangeOK
const ReplaceCoreV1NamespacedLimitRangeOKCode int = 200

/*ReplaceCoreV1NamespacedLimitRangeOK OK

swagger:response replaceCoreV1NamespacedLimitRangeOK
*/
type ReplaceCoreV1NamespacedLimitRangeOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1LimitRange `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedLimitRangeOK creates ReplaceCoreV1NamespacedLimitRangeOK with default headers values
func NewReplaceCoreV1NamespacedLimitRangeOK() *ReplaceCoreV1NamespacedLimitRangeOK {

	return &ReplaceCoreV1NamespacedLimitRangeOK{}
}

// WithPayload adds the payload to the replace core v1 namespaced limit range o k response
func (o *ReplaceCoreV1NamespacedLimitRangeOK) WithPayload(payload *models.IoK8sAPICoreV1LimitRange) *ReplaceCoreV1NamespacedLimitRangeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced limit range o k response
func (o *ReplaceCoreV1NamespacedLimitRangeOK) SetPayload(payload *models.IoK8sAPICoreV1LimitRange) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedLimitRangeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedLimitRangeCreatedCode is the HTTP code returned for type ReplaceCoreV1NamespacedLimitRangeCreated
const ReplaceCoreV1NamespacedLimitRangeCreatedCode int = 201

/*ReplaceCoreV1NamespacedLimitRangeCreated Created

swagger:response replaceCoreV1NamespacedLimitRangeCreated
*/
type ReplaceCoreV1NamespacedLimitRangeCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1LimitRange `json:"body,omitempty"`
}

// NewReplaceCoreV1NamespacedLimitRangeCreated creates ReplaceCoreV1NamespacedLimitRangeCreated with default headers values
func NewReplaceCoreV1NamespacedLimitRangeCreated() *ReplaceCoreV1NamespacedLimitRangeCreated {

	return &ReplaceCoreV1NamespacedLimitRangeCreated{}
}

// WithPayload adds the payload to the replace core v1 namespaced limit range created response
func (o *ReplaceCoreV1NamespacedLimitRangeCreated) WithPayload(payload *models.IoK8sAPICoreV1LimitRange) *ReplaceCoreV1NamespacedLimitRangeCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace core v1 namespaced limit range created response
func (o *ReplaceCoreV1NamespacedLimitRangeCreated) SetPayload(payload *models.IoK8sAPICoreV1LimitRange) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedLimitRangeCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceCoreV1NamespacedLimitRangeUnauthorizedCode is the HTTP code returned for type ReplaceCoreV1NamespacedLimitRangeUnauthorized
const ReplaceCoreV1NamespacedLimitRangeUnauthorizedCode int = 401

/*ReplaceCoreV1NamespacedLimitRangeUnauthorized Unauthorized

swagger:response replaceCoreV1NamespacedLimitRangeUnauthorized
*/
type ReplaceCoreV1NamespacedLimitRangeUnauthorized struct {
}

// NewReplaceCoreV1NamespacedLimitRangeUnauthorized creates ReplaceCoreV1NamespacedLimitRangeUnauthorized with default headers values
func NewReplaceCoreV1NamespacedLimitRangeUnauthorized() *ReplaceCoreV1NamespacedLimitRangeUnauthorized {

	return &ReplaceCoreV1NamespacedLimitRangeUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceCoreV1NamespacedLimitRangeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
