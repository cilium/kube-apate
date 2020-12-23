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

// ReplaceRbacAuthorizationV1ClusterRoleOKCode is the HTTP code returned for type ReplaceRbacAuthorizationV1ClusterRoleOK
const ReplaceRbacAuthorizationV1ClusterRoleOKCode int = 200

/*ReplaceRbacAuthorizationV1ClusterRoleOK OK

swagger:response replaceRbacAuthorizationV1ClusterRoleOK
*/
type ReplaceRbacAuthorizationV1ClusterRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRole `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1ClusterRoleOK creates ReplaceRbacAuthorizationV1ClusterRoleOK with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleOK() *ReplaceRbacAuthorizationV1ClusterRoleOK {

	return &ReplaceRbacAuthorizationV1ClusterRoleOK{}
}

// WithPayload adds the payload to the replace rbac authorization v1 cluster role o k response
func (o *ReplaceRbacAuthorizationV1ClusterRoleOK) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRole) *ReplaceRbacAuthorizationV1ClusterRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1 cluster role o k response
func (o *ReplaceRbacAuthorizationV1ClusterRoleOK) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1ClusterRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1ClusterRoleCreatedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1ClusterRoleCreated
const ReplaceRbacAuthorizationV1ClusterRoleCreatedCode int = 201

/*ReplaceRbacAuthorizationV1ClusterRoleCreated Created

swagger:response replaceRbacAuthorizationV1ClusterRoleCreated
*/
type ReplaceRbacAuthorizationV1ClusterRoleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRole `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1ClusterRoleCreated creates ReplaceRbacAuthorizationV1ClusterRoleCreated with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleCreated() *ReplaceRbacAuthorizationV1ClusterRoleCreated {

	return &ReplaceRbacAuthorizationV1ClusterRoleCreated{}
}

// WithPayload adds the payload to the replace rbac authorization v1 cluster role created response
func (o *ReplaceRbacAuthorizationV1ClusterRoleCreated) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRole) *ReplaceRbacAuthorizationV1ClusterRoleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1 cluster role created response
func (o *ReplaceRbacAuthorizationV1ClusterRoleCreated) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRole) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1ClusterRoleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1ClusterRoleUnauthorizedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1ClusterRoleUnauthorized
const ReplaceRbacAuthorizationV1ClusterRoleUnauthorizedCode int = 401

/*ReplaceRbacAuthorizationV1ClusterRoleUnauthorized Unauthorized

swagger:response replaceRbacAuthorizationV1ClusterRoleUnauthorized
*/
type ReplaceRbacAuthorizationV1ClusterRoleUnauthorized struct {
}

// NewReplaceRbacAuthorizationV1ClusterRoleUnauthorized creates ReplaceRbacAuthorizationV1ClusterRoleUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleUnauthorized() *ReplaceRbacAuthorizationV1ClusterRoleUnauthorized {

	return &ReplaceRbacAuthorizationV1ClusterRoleUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1ClusterRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}