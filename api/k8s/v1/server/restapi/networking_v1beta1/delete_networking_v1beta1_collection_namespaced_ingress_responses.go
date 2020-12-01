// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteNetworkingV1beta1CollectionNamespacedIngressOKCode is the HTTP code returned for type DeleteNetworkingV1beta1CollectionNamespacedIngressOK
const DeleteNetworkingV1beta1CollectionNamespacedIngressOKCode int = 200

/*DeleteNetworkingV1beta1CollectionNamespacedIngressOK OK

swagger:response deleteNetworkingV1beta1CollectionNamespacedIngressOK
*/
type DeleteNetworkingV1beta1CollectionNamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteNetworkingV1beta1CollectionNamespacedIngressOK creates DeleteNetworkingV1beta1CollectionNamespacedIngressOK with default headers values
func NewDeleteNetworkingV1beta1CollectionNamespacedIngressOK() *DeleteNetworkingV1beta1CollectionNamespacedIngressOK {

	return &DeleteNetworkingV1beta1CollectionNamespacedIngressOK{}
}

// WithPayload adds the payload to the delete networking v1beta1 collection namespaced ingress o k response
func (o *DeleteNetworkingV1beta1CollectionNamespacedIngressOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteNetworkingV1beta1CollectionNamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete networking v1beta1 collection namespaced ingress o k response
func (o *DeleteNetworkingV1beta1CollectionNamespacedIngressOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNetworkingV1beta1CollectionNamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorizedCode is the HTTP code returned for type DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized
const DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorizedCode int = 401

/*DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized Unauthorized

swagger:response deleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized
*/
type DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized struct {
}

// NewDeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized creates DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized with default headers values
func NewDeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized() *DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized {

	return &DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteNetworkingV1beta1CollectionNamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
