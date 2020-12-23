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

// ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOKCode is the HTTP code returned for type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK
const ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOKCode int = 200

/*ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK OK

swagger:response replaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK
*/
type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema `json:"body,omitempty"`
}

// NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK creates ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK with default headers values
func NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK() *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK {

	return &ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK{}
}

// WithPayload adds the payload to the replace flowcontrol apiserver v1alpha1 flow schema status o k response
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema) *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace flowcontrol apiserver v1alpha1 flow schema status o k response
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreatedCode is the HTTP code returned for type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated
const ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreatedCode int = 201

/*ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated Created

swagger:response replaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated
*/
type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema `json:"body,omitempty"`
}

// NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated creates ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated with default headers values
func NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated() *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated {

	return &ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated{}
}

// WithPayload adds the payload to the replace flowcontrol apiserver v1alpha1 flow schema status created response
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema) *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace flowcontrol apiserver v1alpha1 flow schema status created response
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1FlowSchema) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorizedCode is the HTTP code returned for type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized
const ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorizedCode int = 401

/*ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized Unauthorized

swagger:response replaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized
*/
type ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized struct {
}

// NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized creates ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized with default headers values
func NewReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized() *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized {

	return &ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}