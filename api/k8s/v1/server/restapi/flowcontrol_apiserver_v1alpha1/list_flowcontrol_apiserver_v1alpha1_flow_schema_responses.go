// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListFlowcontrolApiserverV1alpha1FlowSchemaOKCode is the HTTP code returned for type ListFlowcontrolApiserverV1alpha1FlowSchemaOK
const ListFlowcontrolApiserverV1alpha1FlowSchemaOKCode int = 200

/*ListFlowcontrolApiserverV1alpha1FlowSchemaOK OK

swagger:response listFlowcontrolApiserverV1alpha1FlowSchemaOK
*/
type ListFlowcontrolApiserverV1alpha1FlowSchemaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchemaList `json:"body,omitempty"`
}

// NewListFlowcontrolApiserverV1alpha1FlowSchemaOK creates ListFlowcontrolApiserverV1alpha1FlowSchemaOK with default headers values
func NewListFlowcontrolApiserverV1alpha1FlowSchemaOK() *ListFlowcontrolApiserverV1alpha1FlowSchemaOK {

	return &ListFlowcontrolApiserverV1alpha1FlowSchemaOK{}
}

// WithPayload adds the payload to the list flowcontrol apiserver v1alpha1 flow schema o k response
func (o *ListFlowcontrolApiserverV1alpha1FlowSchemaOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchemaList) *ListFlowcontrolApiserverV1alpha1FlowSchemaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list flowcontrol apiserver v1alpha1 flow schema o k response
func (o *ListFlowcontrolApiserverV1alpha1FlowSchemaOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchemaList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListFlowcontrolApiserverV1alpha1FlowSchemaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode is the HTTP code returned for type ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
const ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode int = 401

/*ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized Unauthorized

swagger:response listFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
*/
type ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized struct {
}

// NewListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized creates ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized with default headers values
func NewListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized() *ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized {

	return &ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized{}
}

// WriteResponse to the client
func (o *ListFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
