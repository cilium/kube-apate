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

// WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOKCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK
const WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOKCode int = 200

/*WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK OK

swagger:response watchRbacAuthorizationV1beta1NamespacedRoleBindingListOK
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK creates WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK() *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK {

	return &WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1beta1 namespaced role binding list o k response
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1beta1 namespaced role binding list o k response
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized
const WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized struct {
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized creates WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized() *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized {

	return &WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleBindingListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
