// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetCoordinationV1APIResourcesOKCode is the HTTP code returned for type GetCoordinationV1APIResourcesOK
const GetCoordinationV1APIResourcesOKCode int = 200

/*GetCoordinationV1APIResourcesOK OK

swagger:response getCoordinationV1ApiResourcesOK
*/
type GetCoordinationV1APIResourcesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList `json:"body,omitempty"`
}

// NewGetCoordinationV1APIResourcesOK creates GetCoordinationV1APIResourcesOK with default headers values
func NewGetCoordinationV1APIResourcesOK() *GetCoordinationV1APIResourcesOK {

	return &GetCoordinationV1APIResourcesOK{}
}

// WithPayload adds the payload to the get coordination v1 Api resources o k response
func (o *GetCoordinationV1APIResourcesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) *GetCoordinationV1APIResourcesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get coordination v1 Api resources o k response
func (o *GetCoordinationV1APIResourcesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCoordinationV1APIResourcesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCoordinationV1APIResourcesUnauthorizedCode is the HTTP code returned for type GetCoordinationV1APIResourcesUnauthorized
const GetCoordinationV1APIResourcesUnauthorizedCode int = 401

/*GetCoordinationV1APIResourcesUnauthorized Unauthorized

swagger:response getCoordinationV1ApiResourcesUnauthorized
*/
type GetCoordinationV1APIResourcesUnauthorized struct {
}

// NewGetCoordinationV1APIResourcesUnauthorized creates GetCoordinationV1APIResourcesUnauthorized with default headers values
func NewGetCoordinationV1APIResourcesUnauthorized() *GetCoordinationV1APIResourcesUnauthorized {

	return &GetCoordinationV1APIResourcesUnauthorized{}
}

// WriteResponse to the client
func (o *GetCoordinationV1APIResourcesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
