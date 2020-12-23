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

// ReadRbacAuthorizationV1ClusterRoleBindingOKCode is the HTTP code returned for type ReadRbacAuthorizationV1ClusterRoleBindingOK
const ReadRbacAuthorizationV1ClusterRoleBindingOKCode int = 200

/*ReadRbacAuthorizationV1ClusterRoleBindingOK OK

swagger:response readRbacAuthorizationV1ClusterRoleBindingOK
*/
type ReadRbacAuthorizationV1ClusterRoleBindingOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIRbacV1ClusterRoleBinding `json:"body,omitempty"`
}

// NewReadRbacAuthorizationV1ClusterRoleBindingOK creates ReadRbacAuthorizationV1ClusterRoleBindingOK with default headers values
func NewReadRbacAuthorizationV1ClusterRoleBindingOK() *ReadRbacAuthorizationV1ClusterRoleBindingOK {

	return &ReadRbacAuthorizationV1ClusterRoleBindingOK{}
}

// WithPayload adds the payload to the read rbac authorization v1 cluster role binding o k response
func (o *ReadRbacAuthorizationV1ClusterRoleBindingOK) WithPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBinding) *ReadRbacAuthorizationV1ClusterRoleBindingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read rbac authorization v1 cluster role binding o k response
func (o *ReadRbacAuthorizationV1ClusterRoleBindingOK) SetPayload(payload *models.IoK8sAPIRbacV1ClusterRoleBinding) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadRbacAuthorizationV1ClusterRoleBindingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadRbacAuthorizationV1ClusterRoleBindingUnauthorizedCode is the HTTP code returned for type ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized
const ReadRbacAuthorizationV1ClusterRoleBindingUnauthorizedCode int = 401

/*ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized Unauthorized

swagger:response readRbacAuthorizationV1ClusterRoleBindingUnauthorized
*/
type ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized struct {
}

// NewReadRbacAuthorizationV1ClusterRoleBindingUnauthorized creates ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized with default headers values
func NewReadRbacAuthorizationV1ClusterRoleBindingUnauthorized() *ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized {

	return &ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized{}
}

// WriteResponse to the client
func (o *ReadRbacAuthorizationV1ClusterRoleBindingUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}