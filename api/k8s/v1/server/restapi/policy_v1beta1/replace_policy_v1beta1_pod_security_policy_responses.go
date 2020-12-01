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

// ReplacePolicyV1beta1PodSecurityPolicyOKCode is the HTTP code returned for type ReplacePolicyV1beta1PodSecurityPolicyOK
const ReplacePolicyV1beta1PodSecurityPolicyOKCode int = 200

/*ReplacePolicyV1beta1PodSecurityPolicyOK OK

swagger:response replacePolicyV1beta1PodSecurityPolicyOK
*/
type ReplacePolicyV1beta1PodSecurityPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewReplacePolicyV1beta1PodSecurityPolicyOK creates ReplacePolicyV1beta1PodSecurityPolicyOK with default headers values
func NewReplacePolicyV1beta1PodSecurityPolicyOK() *ReplacePolicyV1beta1PodSecurityPolicyOK {

	return &ReplacePolicyV1beta1PodSecurityPolicyOK{}
}

// WithPayload adds the payload to the replace policy v1beta1 pod security policy o k response
func (o *ReplacePolicyV1beta1PodSecurityPolicyOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *ReplacePolicyV1beta1PodSecurityPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace policy v1beta1 pod security policy o k response
func (o *ReplacePolicyV1beta1PodSecurityPolicyOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePolicyV1beta1PodSecurityPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplacePolicyV1beta1PodSecurityPolicyCreatedCode is the HTTP code returned for type ReplacePolicyV1beta1PodSecurityPolicyCreated
const ReplacePolicyV1beta1PodSecurityPolicyCreatedCode int = 201

/*ReplacePolicyV1beta1PodSecurityPolicyCreated Created

swagger:response replacePolicyV1beta1PodSecurityPolicyCreated
*/
type ReplacePolicyV1beta1PodSecurityPolicyCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewReplacePolicyV1beta1PodSecurityPolicyCreated creates ReplacePolicyV1beta1PodSecurityPolicyCreated with default headers values
func NewReplacePolicyV1beta1PodSecurityPolicyCreated() *ReplacePolicyV1beta1PodSecurityPolicyCreated {

	return &ReplacePolicyV1beta1PodSecurityPolicyCreated{}
}

// WithPayload adds the payload to the replace policy v1beta1 pod security policy created response
func (o *ReplacePolicyV1beta1PodSecurityPolicyCreated) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *ReplacePolicyV1beta1PodSecurityPolicyCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace policy v1beta1 pod security policy created response
func (o *ReplacePolicyV1beta1PodSecurityPolicyCreated) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplacePolicyV1beta1PodSecurityPolicyCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplacePolicyV1beta1PodSecurityPolicyUnauthorizedCode is the HTTP code returned for type ReplacePolicyV1beta1PodSecurityPolicyUnauthorized
const ReplacePolicyV1beta1PodSecurityPolicyUnauthorizedCode int = 401

/*ReplacePolicyV1beta1PodSecurityPolicyUnauthorized Unauthorized

swagger:response replacePolicyV1beta1PodSecurityPolicyUnauthorized
*/
type ReplacePolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

// NewReplacePolicyV1beta1PodSecurityPolicyUnauthorized creates ReplacePolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewReplacePolicyV1beta1PodSecurityPolicyUnauthorized() *ReplacePolicyV1beta1PodSecurityPolicyUnauthorized {

	return &ReplacePolicyV1beta1PodSecurityPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *ReplacePolicyV1beta1PodSecurityPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
