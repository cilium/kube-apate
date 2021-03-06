// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateRbacAuthorizationV1beta1ClusterRoleBindingOKCode is the HTTP code returned for type CreateRbacAuthorizationV1beta1ClusterRoleBindingOK
const CreateRbacAuthorizationV1beta1ClusterRoleBindingOKCode int = 200

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingOK OK

swagger:response createRbacAuthorizationV1beta1ClusterRoleBindingOK
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding `json:"body,omitempty"`
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingOK creates CreateRbacAuthorizationV1beta1ClusterRoleBindingOK with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingOK() *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK {

	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the create rbac authorization v1beta1 cluster role binding o k response
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create rbac authorization v1beta1 cluster role binding o k response
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRbacAuthorizationV1beta1ClusterRoleBindingCreatedCode is the HTTP code returned for type CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated
const CreateRbacAuthorizationV1beta1ClusterRoleBindingCreatedCode int = 201

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated Created

swagger:response createRbacAuthorizationV1beta1ClusterRoleBindingCreated
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding `json:"body,omitempty"`
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingCreated creates CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingCreated() *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated {

	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated{}
}

// WithPayload adds the payload to the create rbac authorization v1beta1 cluster role binding created response
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated) WithPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create rbac authorization v1beta1 cluster role binding created response
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated) SetPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRbacAuthorizationV1beta1ClusterRoleBindingAcceptedCode is the HTTP code returned for type CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted
const CreateRbacAuthorizationV1beta1ClusterRoleBindingAcceptedCode int = 202

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted Accepted

swagger:response createRbacAuthorizationV1beta1ClusterRoleBindingAccepted
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding `json:"body,omitempty"`
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted creates CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted() *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted {

	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted{}
}

// WithPayload adds the payload to the create rbac authorization v1beta1 cluster role binding accepted response
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted) WithPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create rbac authorization v1beta1 cluster role binding accepted response
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted) SetPayload(payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized
const CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorizedCode int = 401

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized Unauthorized

swagger:response createRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized struct {
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized creates CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized() *CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized {

	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
