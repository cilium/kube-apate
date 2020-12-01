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

// DeleteFlowcontrolApiserverV1alpha1FlowSchemaOKCode is the HTTP code returned for type DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK
const DeleteFlowcontrolApiserverV1alpha1FlowSchemaOKCode int = 200

/*DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK OK

swagger:response deleteFlowcontrolApiserverV1alpha1FlowSchemaOK
*/
type DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaOK creates DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK with default headers values
func NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaOK() *DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK {

	return &DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK{}
}

// WithPayload adds the payload to the delete flowcontrol apiserver v1alpha1 flow schema o k response
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete flowcontrol apiserver v1alpha1 flow schema o k response
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteFlowcontrolApiserverV1alpha1FlowSchemaAcceptedCode is the HTTP code returned for type DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted
const DeleteFlowcontrolApiserverV1alpha1FlowSchemaAcceptedCode int = 202

/*DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted Accepted

swagger:response deleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted
*/
type DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted creates DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted with default headers values
func NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted() *DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted {

	return &DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted{}
}

// WithPayload adds the payload to the delete flowcontrol apiserver v1alpha1 flow schema accepted response
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete flowcontrol apiserver v1alpha1 flow schema accepted response
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode is the HTTP code returned for type DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
const DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorizedCode int = 401

/*DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized Unauthorized

swagger:response deleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized
*/
type DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized struct {
}

// NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized creates DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized with default headers values
func NewDeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized() *DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized {

	return &DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteFlowcontrolApiserverV1alpha1FlowSchemaUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
