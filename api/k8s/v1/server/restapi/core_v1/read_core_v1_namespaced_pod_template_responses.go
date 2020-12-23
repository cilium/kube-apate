// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadCoreV1NamespacedPodTemplateOKCode is the HTTP code returned for type ReadCoreV1NamespacedPodTemplateOK
const ReadCoreV1NamespacedPodTemplateOKCode int = 200

/*ReadCoreV1NamespacedPodTemplateOK OK

swagger:response readCoreV1NamespacedPodTemplateOK
*/
type ReadCoreV1NamespacedPodTemplateOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICoreV1PodTemplate `json:"body,omitempty"`
}

// NewReadCoreV1NamespacedPodTemplateOK creates ReadCoreV1NamespacedPodTemplateOK with default headers values
func NewReadCoreV1NamespacedPodTemplateOK() *ReadCoreV1NamespacedPodTemplateOK {

	return &ReadCoreV1NamespacedPodTemplateOK{}
}

// WithPayload adds the payload to the read core v1 namespaced pod template o k response
func (o *ReadCoreV1NamespacedPodTemplateOK) WithPayload(payload *models.IoK8sAPICoreV1PodTemplate) *ReadCoreV1NamespacedPodTemplateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read core v1 namespaced pod template o k response
func (o *ReadCoreV1NamespacedPodTemplateOK) SetPayload(payload *models.IoK8sAPICoreV1PodTemplate) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedPodTemplateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCoreV1NamespacedPodTemplateUnauthorizedCode is the HTTP code returned for type ReadCoreV1NamespacedPodTemplateUnauthorized
const ReadCoreV1NamespacedPodTemplateUnauthorizedCode int = 401

/*ReadCoreV1NamespacedPodTemplateUnauthorized Unauthorized

swagger:response readCoreV1NamespacedPodTemplateUnauthorized
*/
type ReadCoreV1NamespacedPodTemplateUnauthorized struct {
}

// NewReadCoreV1NamespacedPodTemplateUnauthorized creates ReadCoreV1NamespacedPodTemplateUnauthorized with default headers values
func NewReadCoreV1NamespacedPodTemplateUnauthorized() *ReadCoreV1NamespacedPodTemplateUnauthorized {

	return &ReadCoreV1NamespacedPodTemplateUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCoreV1NamespacedPodTemplateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}