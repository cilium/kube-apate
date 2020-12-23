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

// ReplaceRbacAuthorizationV1beta1NamespacedRoleOKCode is the HTTP code returned for type ReplaceRbacAuthorizationV1beta1NamespacedRoleOK
const ReplaceRbacAuthorizationV1beta1NamespacedRoleOKCode int = 200

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleOK OK

swagger:response replaceRbacAuthorizationV1beta1NamespacedRoleOK
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1Role `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleOK creates ReplaceRbacAuthorizationV1beta1NamespacedRoleOK with default headers values
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleOK() *ReplaceRbacAuthorizationV1beta1NamespacedRoleOK {

	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleOK{}
}

// WithPayload adds the payload to the replace rbac authorization v1beta1 namespaced role o k response
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleOK) WithPayload(payload *models.IoK8sAPIRbacV1beta1Role) *ReplaceRbacAuthorizationV1beta1NamespacedRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1beta1 namespaced role o k response
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleOK) SetPayload(payload *models.IoK8sAPIRbacV1beta1Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1beta1NamespacedRoleCreatedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated
const ReplaceRbacAuthorizationV1beta1NamespacedRoleCreatedCode int = 201

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated Created

swagger:response replaceRbacAuthorizationV1beta1NamespacedRoleCreated
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1Role `json:"body,omitempty"`
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleCreated creates ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated with default headers values
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleCreated() *ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated {

	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated{}
}

// WithPayload adds the payload to the replace rbac authorization v1beta1 namespaced role created response
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated) WithPayload(payload *models.IoK8sAPIRbacV1beta1Role) *ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace rbac authorization v1beta1 namespaced role created response
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated) SetPayload(payload *models.IoK8sAPIRbacV1beta1Role) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorizedCode is the HTTP code returned for type ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized
const ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorizedCode int = 401

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized Unauthorized

swagger:response replaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized struct {
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized creates ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized() *ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized {

	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}