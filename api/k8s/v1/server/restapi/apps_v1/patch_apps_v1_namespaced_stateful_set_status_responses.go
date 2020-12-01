// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchAppsV1NamespacedStatefulSetStatusOKCode is the HTTP code returned for type PatchAppsV1NamespacedStatefulSetStatusOK
const PatchAppsV1NamespacedStatefulSetStatusOKCode int = 200

/*PatchAppsV1NamespacedStatefulSetStatusOK OK

swagger:response patchAppsV1NamespacedStatefulSetStatusOK
*/
type PatchAppsV1NamespacedStatefulSetStatusOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIAppsV1StatefulSet `json:"body,omitempty"`
}

// NewPatchAppsV1NamespacedStatefulSetStatusOK creates PatchAppsV1NamespacedStatefulSetStatusOK with default headers values
func NewPatchAppsV1NamespacedStatefulSetStatusOK() *PatchAppsV1NamespacedStatefulSetStatusOK {

	return &PatchAppsV1NamespacedStatefulSetStatusOK{}
}

// WithPayload adds the payload to the patch apps v1 namespaced stateful set status o k response
func (o *PatchAppsV1NamespacedStatefulSetStatusOK) WithPayload(payload *models.IoK8sAPIAppsV1StatefulSet) *PatchAppsV1NamespacedStatefulSetStatusOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch apps v1 namespaced stateful set status o k response
func (o *PatchAppsV1NamespacedStatefulSetStatusOK) SetPayload(payload *models.IoK8sAPIAppsV1StatefulSet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedStatefulSetStatusOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchAppsV1NamespacedStatefulSetStatusUnauthorizedCode is the HTTP code returned for type PatchAppsV1NamespacedStatefulSetStatusUnauthorized
const PatchAppsV1NamespacedStatefulSetStatusUnauthorizedCode int = 401

/*PatchAppsV1NamespacedStatefulSetStatusUnauthorized Unauthorized

swagger:response patchAppsV1NamespacedStatefulSetStatusUnauthorized
*/
type PatchAppsV1NamespacedStatefulSetStatusUnauthorized struct {
}

// NewPatchAppsV1NamespacedStatefulSetStatusUnauthorized creates PatchAppsV1NamespacedStatefulSetStatusUnauthorized with default headers values
func NewPatchAppsV1NamespacedStatefulSetStatusUnauthorized() *PatchAppsV1NamespacedStatefulSetStatusUnauthorized {

	return &PatchAppsV1NamespacedStatefulSetStatusUnauthorized{}
}

// WriteResponse to the client
func (o *PatchAppsV1NamespacedStatefulSetStatusUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
