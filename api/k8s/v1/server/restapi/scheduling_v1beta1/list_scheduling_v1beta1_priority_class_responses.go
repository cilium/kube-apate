// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListSchedulingV1beta1PriorityClassOKCode is the HTTP code returned for type ListSchedulingV1beta1PriorityClassOK
const ListSchedulingV1beta1PriorityClassOKCode int = 200

/*ListSchedulingV1beta1PriorityClassOK OK

swagger:response listSchedulingV1beta1PriorityClassOK
*/
type ListSchedulingV1beta1PriorityClassOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPISchedulingV1beta1PriorityClassList `json:"body,omitempty"`
}

// NewListSchedulingV1beta1PriorityClassOK creates ListSchedulingV1beta1PriorityClassOK with default headers values
func NewListSchedulingV1beta1PriorityClassOK() *ListSchedulingV1beta1PriorityClassOK {

	return &ListSchedulingV1beta1PriorityClassOK{}
}

// WithPayload adds the payload to the list scheduling v1beta1 priority class o k response
func (o *ListSchedulingV1beta1PriorityClassOK) WithPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClassList) *ListSchedulingV1beta1PriorityClassOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list scheduling v1beta1 priority class o k response
func (o *ListSchedulingV1beta1PriorityClassOK) SetPayload(payload *models.IoK8sAPISchedulingV1beta1PriorityClassList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListSchedulingV1beta1PriorityClassOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListSchedulingV1beta1PriorityClassUnauthorizedCode is the HTTP code returned for type ListSchedulingV1beta1PriorityClassUnauthorized
const ListSchedulingV1beta1PriorityClassUnauthorizedCode int = 401

/*ListSchedulingV1beta1PriorityClassUnauthorized Unauthorized

swagger:response listSchedulingV1beta1PriorityClassUnauthorized
*/
type ListSchedulingV1beta1PriorityClassUnauthorized struct {
}

// NewListSchedulingV1beta1PriorityClassUnauthorized creates ListSchedulingV1beta1PriorityClassUnauthorized with default headers values
func NewListSchedulingV1beta1PriorityClassUnauthorized() *ListSchedulingV1beta1PriorityClassUnauthorized {

	return &ListSchedulingV1beta1PriorityClassUnauthorized{}
}

// WriteResponse to the client
func (o *ListSchedulingV1beta1PriorityClassUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
