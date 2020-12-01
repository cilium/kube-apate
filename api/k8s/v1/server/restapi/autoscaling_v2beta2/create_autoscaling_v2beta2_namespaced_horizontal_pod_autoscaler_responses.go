// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
const CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOKCode int = 200

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK OK

swagger:response createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK creates CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {

	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the create autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create autoscaling v2beta2 namespaced horizontal pod autoscaler o k response
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreatedCode is the HTTP code returned for type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated
const CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreatedCode int = 201

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated Created

swagger:response createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated creates CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated {

	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated{}
}

// WithPayload adds the payload to the create autoscaling v2beta2 namespaced horizontal pod autoscaler created response
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create autoscaling v2beta2 namespaced horizontal pod autoscaler created response
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAcceptedCode is the HTTP code returned for type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted
const CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAcceptedCode int = 202

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted Accepted

swagger:response createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted creates CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted {

	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted{}
}

// WithPayload adds the payload to the create autoscaling v2beta2 namespaced horizontal pod autoscaler accepted response
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create autoscaling v2beta2 namespaced horizontal pod autoscaler accepted response
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
const CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized creates CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized {

	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
