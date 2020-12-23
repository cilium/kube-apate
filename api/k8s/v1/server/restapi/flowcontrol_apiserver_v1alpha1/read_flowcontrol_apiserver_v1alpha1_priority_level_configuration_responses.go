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

// ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOKCode is the HTTP code returned for type ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK
const ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOKCode int = 200

/*ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK OK

swagger:response readFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK
*/
type ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration `json:"body,omitempty"`
}

// NewReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK creates ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK with default headers values
func NewReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK() *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {

	return &ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK{}
}

// WithPayload adds the payload to the read flowcontrol apiserver v1alpha1 priority level configuration o k response
func (o *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) WithPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read flowcontrol apiserver v1alpha1 priority level configuration o k response
func (o *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) SetPayload(payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorizedCode is the HTTP code returned for type ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized
const ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorizedCode int = 401

/*ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized Unauthorized

swagger:response readFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized
*/
type ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized struct {
}

// NewReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized creates ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized with default headers values
func NewReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized() *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized {

	return &ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized{}
}

// WriteResponse to the client
func (o *ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}