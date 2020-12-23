// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteSettingsV1alpha1CollectionNamespacedPodPresetOKCode is the HTTP code returned for type DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK
const DeleteSettingsV1alpha1CollectionNamespacedPodPresetOKCode int = 200

/*DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK OK

swagger:response deleteSettingsV1alpha1CollectionNamespacedPodPresetOK
*/
type DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteSettingsV1alpha1CollectionNamespacedPodPresetOK creates DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK with default headers values
func NewDeleteSettingsV1alpha1CollectionNamespacedPodPresetOK() *DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK {

	return &DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK{}
}

// WithPayload adds the payload to the delete settings v1alpha1 collection namespaced pod preset o k response
func (o *DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete settings v1alpha1 collection namespaced pod preset o k response
func (o *DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSettingsV1alpha1CollectionNamespacedPodPresetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorizedCode is the HTTP code returned for type DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized
const DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorizedCode int = 401

/*DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized Unauthorized

swagger:response deleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized
*/
type DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized struct {
}

// NewDeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized creates DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized with default headers values
func NewDeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized() *DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized {

	return &DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteSettingsV1alpha1CollectionNamespacedPodPresetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}