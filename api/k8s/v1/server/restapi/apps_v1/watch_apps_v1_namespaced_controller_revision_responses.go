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

// WatchAppsV1NamespacedControllerRevisionOKCode is the HTTP code returned for type WatchAppsV1NamespacedControllerRevisionOK
const WatchAppsV1NamespacedControllerRevisionOKCode int = 200

/*WatchAppsV1NamespacedControllerRevisionOK OK

swagger:response watchAppsV1NamespacedControllerRevisionOK
*/
type WatchAppsV1NamespacedControllerRevisionOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchAppsV1NamespacedControllerRevisionOK creates WatchAppsV1NamespacedControllerRevisionOK with default headers values
func NewWatchAppsV1NamespacedControllerRevisionOK() *WatchAppsV1NamespacedControllerRevisionOK {

	return &WatchAppsV1NamespacedControllerRevisionOK{}
}

// WithPayload adds the payload to the watch apps v1 namespaced controller revision o k response
func (o *WatchAppsV1NamespacedControllerRevisionOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchAppsV1NamespacedControllerRevisionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch apps v1 namespaced controller revision o k response
func (o *WatchAppsV1NamespacedControllerRevisionOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchAppsV1NamespacedControllerRevisionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchAppsV1NamespacedControllerRevisionUnauthorizedCode is the HTTP code returned for type WatchAppsV1NamespacedControllerRevisionUnauthorized
const WatchAppsV1NamespacedControllerRevisionUnauthorizedCode int = 401

/*WatchAppsV1NamespacedControllerRevisionUnauthorized Unauthorized

swagger:response watchAppsV1NamespacedControllerRevisionUnauthorized
*/
type WatchAppsV1NamespacedControllerRevisionUnauthorized struct {
}

// NewWatchAppsV1NamespacedControllerRevisionUnauthorized creates WatchAppsV1NamespacedControllerRevisionUnauthorized with default headers values
func NewWatchAppsV1NamespacedControllerRevisionUnauthorized() *WatchAppsV1NamespacedControllerRevisionUnauthorized {

	return &WatchAppsV1NamespacedControllerRevisionUnauthorized{}
}

// WriteResponse to the client
func (o *WatchAppsV1NamespacedControllerRevisionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}