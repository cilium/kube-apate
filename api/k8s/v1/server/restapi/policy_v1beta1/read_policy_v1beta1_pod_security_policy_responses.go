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

// ReadPolicyV1beta1PodSecurityPolicyOKCode is the HTTP code returned for type ReadPolicyV1beta1PodSecurityPolicyOK
const ReadPolicyV1beta1PodSecurityPolicyOKCode int = 200

/*ReadPolicyV1beta1PodSecurityPolicyOK OK

swagger:response readPolicyV1beta1PodSecurityPolicyOK
*/
type ReadPolicyV1beta1PodSecurityPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy `json:"body,omitempty"`
}

// NewReadPolicyV1beta1PodSecurityPolicyOK creates ReadPolicyV1beta1PodSecurityPolicyOK with default headers values
func NewReadPolicyV1beta1PodSecurityPolicyOK() *ReadPolicyV1beta1PodSecurityPolicyOK {

	return &ReadPolicyV1beta1PodSecurityPolicyOK{}
}

// WithPayload adds the payload to the read policy v1beta1 pod security policy o k response
func (o *ReadPolicyV1beta1PodSecurityPolicyOK) WithPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) *ReadPolicyV1beta1PodSecurityPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read policy v1beta1 pod security policy o k response
func (o *ReadPolicyV1beta1PodSecurityPolicyOK) SetPayload(payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadPolicyV1beta1PodSecurityPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadPolicyV1beta1PodSecurityPolicyUnauthorizedCode is the HTTP code returned for type ReadPolicyV1beta1PodSecurityPolicyUnauthorized
const ReadPolicyV1beta1PodSecurityPolicyUnauthorizedCode int = 401

/*ReadPolicyV1beta1PodSecurityPolicyUnauthorized Unauthorized

swagger:response readPolicyV1beta1PodSecurityPolicyUnauthorized
*/
type ReadPolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

// NewReadPolicyV1beta1PodSecurityPolicyUnauthorized creates ReadPolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewReadPolicyV1beta1PodSecurityPolicyUnauthorized() *ReadPolicyV1beta1PodSecurityPolicyUnauthorized {

	return &ReadPolicyV1beta1PodSecurityPolicyUnauthorized{}
}

// WriteResponse to the client
func (o *ReadPolicyV1beta1PodSecurityPolicyUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}