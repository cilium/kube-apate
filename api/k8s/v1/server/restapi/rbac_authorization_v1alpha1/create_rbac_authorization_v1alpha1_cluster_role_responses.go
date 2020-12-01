// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateRbacAuthorizationV1alpha1ClusterRoleOKCode is the HTTP code returned for type CreateRbacAuthorizationV1alpha1ClusterRoleOK
const CreateRbacAuthorizationV1alpha1ClusterRoleOKCode int = 200

/*CreateRbacAuthorizationV1alpha1ClusterRoleOK OK

swagger:response createRbacAuthorizationV1alpha1ClusterRoleOK
*/
type CreateRbacAuthorizationV1alpha1ClusterRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRole `json:"body,omitempty"`
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleOK creates CreateRbacAuthorizationV1alpha1ClusterRoleOK with default headers values
func NewCreateRbacAuthorizationV1alpha1ClusterRoleOK() *CreateRbacAuthorizationV1alpha1ClusterRoleOK {

	return &CreateRbacAuthorizationV1alpha1ClusterRoleOK{}
}

// WithPayload adds the payload to the create rbac authorization v1alpha1 cluster role o k response
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleOK) WithPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRole) *CreateRbacAuthorizationV1alpha1ClusterRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create rbac authorization v1alpha1 cluster role o k response
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleOK) SetPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRbacAuthorizationV1alpha1ClusterRoleCreatedCode is the HTTP code returned for type CreateRbacAuthorizationV1alpha1ClusterRoleCreated
const CreateRbacAuthorizationV1alpha1ClusterRoleCreatedCode int = 201

/*CreateRbacAuthorizationV1alpha1ClusterRoleCreated Created

swagger:response createRbacAuthorizationV1alpha1ClusterRoleCreated
*/
type CreateRbacAuthorizationV1alpha1ClusterRoleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRole `json:"body,omitempty"`
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleCreated creates CreateRbacAuthorizationV1alpha1ClusterRoleCreated with default headers values
func NewCreateRbacAuthorizationV1alpha1ClusterRoleCreated() *CreateRbacAuthorizationV1alpha1ClusterRoleCreated {

	return &CreateRbacAuthorizationV1alpha1ClusterRoleCreated{}
}

// WithPayload adds the payload to the create rbac authorization v1alpha1 cluster role created response
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleCreated) WithPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRole) *CreateRbacAuthorizationV1alpha1ClusterRoleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create rbac authorization v1alpha1 cluster role created response
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleCreated) SetPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRbacAuthorizationV1alpha1ClusterRoleAcceptedCode is the HTTP code returned for type CreateRbacAuthorizationV1alpha1ClusterRoleAccepted
const CreateRbacAuthorizationV1alpha1ClusterRoleAcceptedCode int = 202

/*CreateRbacAuthorizationV1alpha1ClusterRoleAccepted Accepted

swagger:response createRbacAuthorizationV1alpha1ClusterRoleAccepted
*/
type CreateRbacAuthorizationV1alpha1ClusterRoleAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRole `json:"body,omitempty"`
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleAccepted creates CreateRbacAuthorizationV1alpha1ClusterRoleAccepted with default headers values
func NewCreateRbacAuthorizationV1alpha1ClusterRoleAccepted() *CreateRbacAuthorizationV1alpha1ClusterRoleAccepted {

	return &CreateRbacAuthorizationV1alpha1ClusterRoleAccepted{}
}

// WithPayload adds the payload to the create rbac authorization v1alpha1 cluster role accepted response
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleAccepted) WithPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRole) *CreateRbacAuthorizationV1alpha1ClusterRoleAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create rbac authorization v1alpha1 cluster role accepted response
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleAccepted) SetPayload(payload *models.IoK8sAPIRbacV1alpha1ClusterRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorizedCode is the HTTP code returned for type CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized
const CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorizedCode int = 401

/*CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized Unauthorized

swagger:response createRbacAuthorizationV1alpha1ClusterRoleUnauthorized
*/
type CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized struct {
}

// NewCreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized creates CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized with default headers values
func NewCreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized() *CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized {

	return &CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized{}
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1alpha1ClusterRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
