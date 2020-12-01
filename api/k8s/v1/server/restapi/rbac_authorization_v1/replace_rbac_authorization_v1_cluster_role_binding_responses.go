// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceRbacAuthorizationV1ClusterRoleBindingOKCode is the HTTP code returned for type ReplaceRbacAuthorizationV1ClusterRoleBindingOK
const ReplaceRbacAuthorizationV1ClusterRoleBindingOKCode int = 200

/*ReplaceRbacAuthorizationV1ClusterRoleBindingOK OK

swagger:response replaceRbacAuthorizationV1ClusterRoleBindingOK
*/
type ReplaceRbacAuthorizationV1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRoleBinding `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1ClusterRoleBindingOK creates ReplaceRbacAuthorizationV1ClusterRoleBindingOK with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleBindingOK() *ReplaceRbacAuthorizationV1ClusterRoleBindingOK {

	return &ReplaceRbacAuthorizationV1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the replace rbac authorization v1 cluster role binding o k response
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBinding) *ReplaceRbacAuthorizationV1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1 cluster role binding o k response
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1ClusterRoleBindingCreatedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1ClusterRoleBindingCreated
const ReplaceRbacAuthorizationV1ClusterRoleBindingCreatedCode int = 201

/*ReplaceRbacAuthorizationV1ClusterRoleBindingCreated Created

swagger:response replaceRbacAuthorizationV1ClusterRoleBindingCreated
*/
type ReplaceRbacAuthorizationV1ClusterRoleBindingCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRoleBinding `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1ClusterRoleBindingCreated creates ReplaceRbacAuthorizationV1ClusterRoleBindingCreated with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleBindingCreated() *ReplaceRbacAuthorizationV1ClusterRoleBindingCreated {

	return &ReplaceRbacAuthorizationV1ClusterRoleBindingCreated{}
}

// WithPayload adds the payload to the replace rbac authorization v1 cluster role binding created response
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingCreated) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBinding) *ReplaceRbacAuthorizationV1ClusterRoleBindingCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1 cluster role binding created response
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingCreated) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized
const ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorizedCode int = 401

/*ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized Unauthorized

swagger:response replaceRbacAuthorizationV1ClusterRoleBindingUnauthorized
*/
type ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized struct {
}

// NewReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized creates ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized() *ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized {

	return &ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
