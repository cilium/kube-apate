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

// DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOKCode is the HTTP code returned for type DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK
const DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOKCode int = 200

/*DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK OK

swagger:response deleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK
*/
type DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK creates DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK with default headers values
func NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK() *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK {

	return &DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK{}
}

// WithPayload adds the payload to the delete rbac authorization v1beta1 collection namespaced role o k response
func (o *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete rbac authorization v1beta1 collection namespaced role o k response
func (o *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorizedCode is the HTTP code returned for type DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized
const DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorizedCode int = 401

/*DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized Unauthorized

swagger:response deleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized
*/
type DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized struct {
}

// NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized creates DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized() *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized {

	return &DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}