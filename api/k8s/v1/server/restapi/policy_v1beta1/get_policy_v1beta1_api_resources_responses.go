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

// GetPolicyV1beta1APIResourcesOKCode is the HTTP code returned for type GetPolicyV1beta1APIResourcesOK
const GetPolicyV1beta1APIResourcesOKCode int = 200

/*GetPolicyV1beta1APIResourcesOK OK

swagger:response getPolicyV1beta1ApiResourcesOK
*/
type GetPolicyV1beta1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetPolicyV1beta1APIResourcesOK creates GetPolicyV1beta1APIResourcesOK with default headers values
func NewGetPolicyV1beta1APIResourcesOK() *GetPolicyV1beta1APIResourcesOK {

	return &GetPolicyV1beta1APIResourcesOK{}
}

// WithPayload adds the payload to the get policy v1beta1 Api resources o k response
func (o *GetPolicyV1beta1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetPolicyV1beta1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy v1beta1 Api resources o k response
func (o *GetPolicyV1beta1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyV1beta1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPolicyV1beta1APIResourcesUnauthorizedCode is the HTTP code returned for type GetPolicyV1beta1APIResourcesUnauthorized
const GetPolicyV1beta1APIResourcesUnauthorizedCode int = 401

/*GetPolicyV1beta1APIResourcesUnauthorized Unauthorized

swagger:response getPolicyV1beta1ApiResourcesUnauthorized
*/
type GetPolicyV1beta1APIResourcesUnauthorized struct {
}

// NewGetPolicyV1beta1APIResourcesUnauthorized creates GetPolicyV1beta1APIResourcesUnauthorized with default headers values
func NewGetPolicyV1beta1APIResourcesUnauthorized() *GetPolicyV1beta1APIResourcesUnauthorized {

	return &GetPolicyV1beta1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetPolicyV1beta1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}