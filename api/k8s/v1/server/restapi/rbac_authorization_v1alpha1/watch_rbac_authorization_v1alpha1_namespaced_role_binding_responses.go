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

// WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOKCode is the HTTP code returned for type WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK
const WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOKCode int = 200

/*WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK OK

swagger:response watchRbacAuthorizationV1alpha1NamespacedRoleBindingOK
*/
type WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK creates WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK with default headers values
func NewWatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK() *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK {

	return &WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1alpha1 namespaced role binding o k response
func (o *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1alpha1 namespaced role binding o k response
func (o *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized
const WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized
*/
type WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized struct {
}

// NewWatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized creates WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized with default headers values
func NewWatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized() *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized {

	return &WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
