// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceExtensionsV1beta1NamespacedIngressOKCode is the HTTP code returned for type ReplaceExtensionsV1beta1NamespacedIngressOK
const ReplaceExtensionsV1beta1NamespacedIngressOKCode int = 200

/*ReplaceExtensionsV1beta1NamespacedIngressOK OK

swagger:response replaceExtensionsV1beta1NamespacedIngressOK
*/
type ReplaceExtensionsV1beta1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewReplaceExtensionsV1beta1NamespacedIngressOK creates ReplaceExtensionsV1beta1NamespacedIngressOK with default headers values
func NewReplaceExtensionsV1beta1NamespacedIngressOK() *ReplaceExtensionsV1beta1NamespacedIngressOK {

	return &ReplaceExtensionsV1beta1NamespacedIngressOK{}
}

// WithPayload adds the payload to the replace extensions v1beta1 namespaced ingress o k response
func (o *ReplaceExtensionsV1beta1NamespacedIngressOK) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *ReplaceExtensionsV1beta1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace extensions v1beta1 namespaced ingress o k response
func (o *ReplaceExtensionsV1beta1NamespacedIngressOK) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceExtensionsV1beta1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceExtensionsV1beta1NamespacedIngressCreatedCode is the HTTP code returned for type ReplaceExtensionsV1beta1NamespacedIngressCreated
const ReplaceExtensionsV1beta1NamespacedIngressCreatedCode int = 201

/*ReplaceExtensionsV1beta1NamespacedIngressCreated Created

swagger:response replaceExtensionsV1beta1NamespacedIngressCreated
*/
type ReplaceExtensionsV1beta1NamespacedIngressCreated struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress `json:"body,omitempty"`
}

// NewReplaceExtensionsV1beta1NamespacedIngressCreated creates ReplaceExtensionsV1beta1NamespacedIngressCreated with default headers values
func NewReplaceExtensionsV1beta1NamespacedIngressCreated() *ReplaceExtensionsV1beta1NamespacedIngressCreated {

	return &ReplaceExtensionsV1beta1NamespacedIngressCreated{}
}

// WithPayload adds the payload to the replace extensions v1beta1 namespaced ingress created response
func (o *ReplaceExtensionsV1beta1NamespacedIngressCreated) WithPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) *ReplaceExtensionsV1beta1NamespacedIngressCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace extensions v1beta1 namespaced ingress created response
func (o *ReplaceExtensionsV1beta1NamespacedIngressCreated) SetPayload(payload *models.IoK8sAPIExtensionsV1beta1Ingress) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceExtensionsV1beta1NamespacedIngressCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceExtensionsV1beta1NamespacedIngressUnauthorizedCode is the HTTP code returned for type ReplaceExtensionsV1beta1NamespacedIngressUnauthorized
const ReplaceExtensionsV1beta1NamespacedIngressUnauthorizedCode int = 401

/*ReplaceExtensionsV1beta1NamespacedIngressUnauthorized Unauthorized

swagger:response replaceExtensionsV1beta1NamespacedIngressUnauthorized
*/
type ReplaceExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

// NewReplaceExtensionsV1beta1NamespacedIngressUnauthorized creates ReplaceExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewReplaceExtensionsV1beta1NamespacedIngressUnauthorized() *ReplaceExtensionsV1beta1NamespacedIngressUnauthorized {

	return &ReplaceExtensionsV1beta1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *ReplaceExtensionsV1beta1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
