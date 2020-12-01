// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchApiextensionsV1CustomResourceDefinitionListOKCode is the HTTP code returned for type WatchApiextensionsV1CustomResourceDefinitionListOK
const WatchApiextensionsV1CustomResourceDefinitionListOKCode int = 200

/*WatchApiextensionsV1CustomResourceDefinitionListOK OK

swagger:response watchApiextensionsV1CustomResourceDefinitionListOK
*/
type WatchApiextensionsV1CustomResourceDefinitionListOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchApiextensionsV1CustomResourceDefinitionListOK creates WatchApiextensionsV1CustomResourceDefinitionListOK with default headers values
func NewWatchApiextensionsV1CustomResourceDefinitionListOK() *WatchApiextensionsV1CustomResourceDefinitionListOK {

	return &WatchApiextensionsV1CustomResourceDefinitionListOK{}
}

// WithPayload adds the payload to the watch apiextensions v1 custom resource definition list o k response
func (o *WatchApiextensionsV1CustomResourceDefinitionListOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchApiextensionsV1CustomResourceDefinitionListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch apiextensions v1 custom resource definition list o k response
func (o *WatchApiextensionsV1CustomResourceDefinitionListOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchApiextensionsV1CustomResourceDefinitionListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchApiextensionsV1CustomResourceDefinitionListUnauthorizedCode is the HTTP code returned for type WatchApiextensionsV1CustomResourceDefinitionListUnauthorized
const WatchApiextensionsV1CustomResourceDefinitionListUnauthorizedCode int = 401

/*WatchApiextensionsV1CustomResourceDefinitionListUnauthorized Unauthorized

swagger:response watchApiextensionsV1CustomResourceDefinitionListUnauthorized
*/
type WatchApiextensionsV1CustomResourceDefinitionListUnauthorized struct {
}

// NewWatchApiextensionsV1CustomResourceDefinitionListUnauthorized creates WatchApiextensionsV1CustomResourceDefinitionListUnauthorized with default headers values
func NewWatchApiextensionsV1CustomResourceDefinitionListUnauthorized() *WatchApiextensionsV1CustomResourceDefinitionListUnauthorized {

	return &WatchApiextensionsV1CustomResourceDefinitionListUnauthorized{}
}

// WriteResponse to the client
func (o *WatchApiextensionsV1CustomResourceDefinitionListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
