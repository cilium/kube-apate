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

// WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOKCode is the HTTP code returned for type WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK
const WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOKCode int = 200

/*WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK OK

swagger:response watchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK
*/
type WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK creates WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK with default headers values
func NewWatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK() *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK {

	return &WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1alpha1 role binding list for all namespaces o k response
func (o *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1alpha1 role binding list for all namespaces o k response
func (o *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized
const WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized
*/
type WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized struct {
}

// NewWatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized creates WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized with default headers values
func NewWatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized() *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized {

	return &WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
