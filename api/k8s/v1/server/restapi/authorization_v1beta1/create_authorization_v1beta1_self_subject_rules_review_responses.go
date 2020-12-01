// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAuthorizationV1beta1SelfSubjectRulesReviewOKCode is the HTTP code returned for type CreateAuthorizationV1beta1SelfSubjectRulesReviewOK
const CreateAuthorizationV1beta1SelfSubjectRulesReviewOKCode int = 200

/*CreateAuthorizationV1beta1SelfSubjectRulesReviewOK OK

swagger:response createAuthorizationV1beta1SelfSubjectRulesReviewOK
*/
type CreateAuthorizationV1beta1SelfSubjectRulesReviewOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview `json:"body,omitempty"`
}

// NewCreateAuthorizationV1beta1SelfSubjectRulesReviewOK creates CreateAuthorizationV1beta1SelfSubjectRulesReviewOK with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectRulesReviewOK() *CreateAuthorizationV1beta1SelfSubjectRulesReviewOK {

	return &CreateAuthorizationV1beta1SelfSubjectRulesReviewOK{}
}

// WithPayload adds the payload to the create authorization v1beta1 self subject rules review o k response
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewOK) WithPayload(payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview) *CreateAuthorizationV1beta1SelfSubjectRulesReviewOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authorization v1beta1 self subject rules review o k response
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewOK) SetPayload(payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthorizationV1beta1SelfSubjectRulesReviewCreatedCode is the HTTP code returned for type CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated
const CreateAuthorizationV1beta1SelfSubjectRulesReviewCreatedCode int = 201

/*CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated Created

swagger:response createAuthorizationV1beta1SelfSubjectRulesReviewCreated
*/
type CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview `json:"body,omitempty"`
}

// NewCreateAuthorizationV1beta1SelfSubjectRulesReviewCreated creates CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectRulesReviewCreated() *CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated {

	return &CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated{}
}

// WithPayload adds the payload to the create authorization v1beta1 self subject rules review created response
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated) WithPayload(payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview) *CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authorization v1beta1 self subject rules review created response
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated) SetPayload(payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthorizationV1beta1SelfSubjectRulesReviewAcceptedCode is the HTTP code returned for type CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted
const CreateAuthorizationV1beta1SelfSubjectRulesReviewAcceptedCode int = 202

/*CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted Accepted

swagger:response createAuthorizationV1beta1SelfSubjectRulesReviewAccepted
*/
type CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview `json:"body,omitempty"`
}

// NewCreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted creates CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted() *CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted {

	return &CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted{}
}

// WithPayload adds the payload to the create authorization v1beta1 self subject rules review accepted response
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted) WithPayload(payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview) *CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create authorization v1beta1 self subject rules review accepted response
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted) SetPayload(payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorizedCode is the HTTP code returned for type CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized
const CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorizedCode int = 401

/*CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized Unauthorized

swagger:response createAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized
*/
type CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized struct {
}

// NewCreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized creates CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized() *CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized {

	return &CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized{}
}

// WriteResponse to the client
func (o *CreateAuthorizationV1beta1SelfSubjectRulesReviewUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
