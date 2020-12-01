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

// PatchRbacAuthorizationV1beta1NamespacedRoleBindingOKCode is the HTTP code returned for type PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK
const PatchRbacAuthorizationV1beta1NamespacedRoleBindingOKCode int = 200

/*PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK OK

swagger:response patchRbacAuthorizationV1beta1NamespacedRoleBindingOK
*/
type PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1beta1RoleBinding `json:"body,omitempty"`
}

// NewPatchRbacAuthorizationV1beta1NamespacedRoleBindingOK creates PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK with default headers values
func NewPatchRbacAuthorizationV1beta1NamespacedRoleBindingOK() *PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK {

	return &PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK{}
}

// WithPayload adds the payload to the patch rbac authorization v1beta1 namespaced role binding o k response
func (o *PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1beta1RoleBinding) *PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch rbac authorization v1beta1 namespaced role binding o k response
func (o *PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1beta1RoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1beta1NamespacedRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorizedCode is the HTTP code returned for type PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized
const PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorizedCode int = 401

/*PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized Unauthorized

swagger:response patchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized
*/
type PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized struct {
}

// NewPatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized creates PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized with default headers values
func NewPatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized() *PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized {

	return &PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *PatchRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
