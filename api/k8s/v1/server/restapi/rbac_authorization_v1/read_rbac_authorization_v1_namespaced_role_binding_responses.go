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

// ReadRbacAuthorizationV1NamespacedRoleBindingOKCode is the HTTP code returned for type ReadRbacAuthorizationV1NamespacedRoleBindingOK
const ReadRbacAuthorizationV1NamespacedRoleBindingOKCode int = 200

/*ReadRbacAuthorizationV1NamespacedRoleBindingOK OK

swagger:response readRbacAuthorizationV1NamespacedRoleBindingOK
*/
type ReadRbacAuthorizationV1NamespacedRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1RoleBinding `json:"body,omitempty"`
}

// NewReadRbacAuthorizationV1NamespacedRoleBindingOK creates ReadRbacAuthorizationV1NamespacedRoleBindingOK with default headers values
func NewReadRbacAuthorizationV1NamespacedRoleBindingOK() *ReadRbacAuthorizationV1NamespacedRoleBindingOK {

	return &ReadRbacAuthorizationV1NamespacedRoleBindingOK{}
}

// WithPayload adds the payload to the read rbac authorization v1 namespaced role binding o k response
func (o *ReadRbacAuthorizationV1NamespacedRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1RoleBinding) *ReadRbacAuthorizationV1NamespacedRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read rbac authorization v1 namespaced role binding o k response
func (o *ReadRbacAuthorizationV1NamespacedRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1RoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadRbacAuthorizationV1NamespacedRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorizedCode is the HTTP code returned for type ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized
const ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorizedCode int = 401

/*ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized Unauthorized

swagger:response readRbacAuthorizationV1NamespacedRoleBindingUnauthorized
*/
type ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized struct {
}

// NewReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized creates ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized with default headers values
func NewReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized() *ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized {

	return &ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *ReadRbacAuthorizationV1NamespacedRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
