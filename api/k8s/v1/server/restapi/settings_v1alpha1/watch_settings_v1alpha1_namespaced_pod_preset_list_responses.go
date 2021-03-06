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

// WatchSettingsV1alpha1NamespacedPodPresetListOKCode is the HTTP code returned for type WatchSettingsV1alpha1NamespacedPodPresetListOK
const WatchSettingsV1alpha1NamespacedPodPresetListOKCode int = 200

/*WatchSettingsV1alpha1NamespacedPodPresetListOK OK

swagger:response watchSettingsV1alpha1NamespacedPodPresetListOK
*/
type WatchSettingsV1alpha1NamespacedPodPresetListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchSettingsV1alpha1NamespacedPodPresetListOK creates WatchSettingsV1alpha1NamespacedPodPresetListOK with default headers values
func NewWatchSettingsV1alpha1NamespacedPodPresetListOK() *WatchSettingsV1alpha1NamespacedPodPresetListOK {

	return &WatchSettingsV1alpha1NamespacedPodPresetListOK{}
}

// WithPayload adds the payload to the watch settings v1alpha1 namespaced pod preset list o k response
func (o *WatchSettingsV1alpha1NamespacedPodPresetListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchSettingsV1alpha1NamespacedPodPresetListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch settings v1alpha1 namespaced pod preset list o k response
func (o *WatchSettingsV1alpha1NamespacedPodPresetListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchSettingsV1alpha1NamespacedPodPresetListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchSettingsV1alpha1NamespacedPodPresetListUnauthorizedCode is the HTTP code returned for type WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized
const WatchSettingsV1alpha1NamespacedPodPresetListUnauthorizedCode int = 401

/*WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized Unauthorized

swagger:response watchSettingsV1alpha1NamespacedPodPresetListUnauthorized
*/
type WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized struct {
}

// NewWatchSettingsV1alpha1NamespacedPodPresetListUnauthorized creates WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized with default headers values
func NewWatchSettingsV1alpha1NamespacedPodPresetListUnauthorized() *WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized {

	return &WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchSettingsV1alpha1NamespacedPodPresetListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
