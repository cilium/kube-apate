// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode is the HTTP code returned for type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
const CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOKCode int = 200

/*CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK OK

swagger:response createAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK
*/
type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK creates CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK() *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {

	return &CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK{}
}

// WithPayload adds the payload to the create autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create autoscaling v2beta1 namespaced horizontal pod autoscaler o k response
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreatedCode is the HTTP code returned for type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated
const CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreatedCode int = 201

/*CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated Created

swagger:response createAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated
*/
type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated creates CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated() *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated {

	return &CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated{}
}

// WithPayload adds the payload to the create autoscaling v2beta1 namespaced horizontal pod autoscaler created response
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create autoscaling v2beta1 namespaced horizontal pod autoscaler created response
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAcceptedCode is the HTTP code returned for type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted
const CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAcceptedCode int = 202

/*CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted Accepted

swagger:response createAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted
*/
type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler `json:"body,omitempty"`
}

// NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted creates CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted with default headers values
func NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted() *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted {

	return &CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted{}
}

// WithPayload adds the payload to the create autoscaling v2beta1 namespaced horizontal pod autoscaler accepted response
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted) WithPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create autoscaling v2beta1 namespaced horizontal pod autoscaler accepted response
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted) SetPayload(payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode is the HTTP code returned for type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
const CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorizedCode int = 401

/*CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized Unauthorized

swagger:response createAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized
*/
type CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

// NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized creates CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewCreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized() *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized {

	return &CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
