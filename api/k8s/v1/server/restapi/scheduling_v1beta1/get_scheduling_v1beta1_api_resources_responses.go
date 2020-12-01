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

// GetSchedulingV1beta1APIResourcesOKCode is the HTTP code returned for type GetSchedulingV1beta1APIResourcesOK
const GetSchedulingV1beta1APIResourcesOKCode int = 200

/*GetSchedulingV1beta1APIResourcesOK OK

swagger:response getSchedulingV1beta1ApiResourcesOK
*/
type GetSchedulingV1beta1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetSchedulingV1beta1APIResourcesOK creates GetSchedulingV1beta1APIResourcesOK with default headers values
func NewGetSchedulingV1beta1APIResourcesOK() *GetSchedulingV1beta1APIResourcesOK {

	return &GetSchedulingV1beta1APIResourcesOK{}
}

// WithPayload adds the payload to the get scheduling v1beta1 Api resources o k response
func (o *GetSchedulingV1beta1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetSchedulingV1beta1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get scheduling v1beta1 Api resources o k response
func (o *GetSchedulingV1beta1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSchedulingV1beta1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSchedulingV1beta1APIResourcesUnauthorizedCode is the HTTP code returned for type GetSchedulingV1beta1APIResourcesUnauthorized
const GetSchedulingV1beta1APIResourcesUnauthorizedCode int = 401

/*GetSchedulingV1beta1APIResourcesUnauthorized Unauthorized

swagger:response getSchedulingV1beta1ApiResourcesUnauthorized
*/
type GetSchedulingV1beta1APIResourcesUnauthorized struct {
}

// NewGetSchedulingV1beta1APIResourcesUnauthorized creates GetSchedulingV1beta1APIResourcesUnauthorized with default headers values
func NewGetSchedulingV1beta1APIResourcesUnauthorized() *GetSchedulingV1beta1APIResourcesUnauthorized {

	return &GetSchedulingV1beta1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetSchedulingV1beta1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
