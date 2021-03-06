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

// WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOKCode is the HTTP code returned for type WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK
const WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOKCode int = 200

/*WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK OK

swagger:response watchRbacAuthorizationV1RoleBindingListForAllNamespacesOK
*/
type WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK creates WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK with default headers values
func NewWatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK() *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK {

	return &WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK{}
}

// WithPayload adds the payload to the watch rbac authorization v1 role binding list for all namespaces o k response
func (o *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch rbac authorization v1 role binding list for all namespaces o k response
func (o *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorizedCode is the HTTP code returned for type WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized
const WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorizedCode int = 401

/*WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized Unauthorized

swagger:response watchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized
*/
type WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized struct {
}

// NewWatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized creates WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized with default headers values
func NewWatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized() *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized {

	return &WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized{}
}

// WriteResponse to the client
func (o *WatchRbacAuthorizationV1RoleBindingListForAllNamespacesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
