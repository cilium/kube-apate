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

// ReplaceRbacAuthorizationV1NamespacedRoleOKCode is the HTTP code returned for type ReplaceRbacAuthorizationV1NamespacedRoleOK
const ReplaceRbacAuthorizationV1NamespacedRoleOKCode int = 200

/*ReplaceRbacAuthorizationV1NamespacedRoleOK OK

swagger:response replaceRbacAuthorizationV1NamespacedRoleOK
*/
type ReplaceRbacAuthorizationV1NamespacedRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1Role `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1NamespacedRoleOK creates ReplaceRbacAuthorizationV1NamespacedRoleOK with default headers values
func NewReplaceRbacAuthorizationV1NamespacedRoleOK() *ReplaceRbacAuthorizationV1NamespacedRoleOK {

	return &ReplaceRbacAuthorizationV1NamespacedRoleOK{}
}

// WithPayload adds the payload to the replace rbac authorization v1 namespaced role o k response
func (o *ReplaceRbacAuthorizationV1NamespacedRoleOK) WithPayload(payload *models.IoK8sAPIRbacV1Role) *ReplaceRbacAuthorizationV1NamespacedRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1 namespaced role o k response
func (o *ReplaceRbacAuthorizationV1NamespacedRoleOK) SetPayload(payload *models.IoK8sAPIRbacV1Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1NamespacedRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1NamespacedRoleCreatedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1NamespacedRoleCreated
const ReplaceRbacAuthorizationV1NamespacedRoleCreatedCode int = 201

/*ReplaceRbacAuthorizationV1NamespacedRoleCreated Created

swagger:response replaceRbacAuthorizationV1NamespacedRoleCreated
*/
type ReplaceRbacAuthorizationV1NamespacedRoleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1Role `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1NamespacedRoleCreated creates ReplaceRbacAuthorizationV1NamespacedRoleCreated with default headers values
func NewReplaceRbacAuthorizationV1NamespacedRoleCreated() *ReplaceRbacAuthorizationV1NamespacedRoleCreated {

	return &ReplaceRbacAuthorizationV1NamespacedRoleCreated{}
}

// WithPayload adds the payload to the replace rbac authorization v1 namespaced role created response
func (o *ReplaceRbacAuthorizationV1NamespacedRoleCreated) WithPayload(payload *models.IoK8sAPIRbacV1Role) *ReplaceRbacAuthorizationV1NamespacedRoleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1 namespaced role created response
func (o *ReplaceRbacAuthorizationV1NamespacedRoleCreated) SetPayload(payload *models.IoK8sAPIRbacV1Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1NamespacedRoleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1NamespacedRoleUnauthorizedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized
const ReplaceRbacAuthorizationV1NamespacedRoleUnauthorizedCode int = 401

/*ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized Unauthorized

swagger:response replaceRbacAuthorizationV1NamespacedRoleUnauthorized
*/
type ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized struct {
}

// NewReplaceRbacAuthorizationV1NamespacedRoleUnauthorized creates ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1NamespacedRoleUnauthorized() *ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized {

	return &ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1NamespacedRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}