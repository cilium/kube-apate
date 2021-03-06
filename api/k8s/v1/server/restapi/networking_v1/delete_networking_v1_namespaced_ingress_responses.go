// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteNetworkingV1NamespacedIngressOKCode is the HTTP code returned for type DeleteNetworkingV1NamespacedIngressOK
const DeleteNetworkingV1NamespacedIngressOKCode int = 200

/*DeleteNetworkingV1NamespacedIngressOK OK

swagger:response deleteNetworkingV1NamespacedIngressOK
*/
type DeleteNetworkingV1NamespacedIngressOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteNetworkingV1NamespacedIngressOK creates DeleteNetworkingV1NamespacedIngressOK with default headers values
func NewDeleteNetworkingV1NamespacedIngressOK() *DeleteNetworkingV1NamespacedIngressOK {

	return &DeleteNetworkingV1NamespacedIngressOK{}
}

// WithPayload adds the payload to the delete networking v1 namespaced ingress o k response
func (o *DeleteNetworkingV1NamespacedIngressOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteNetworkingV1NamespacedIngressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete networking v1 namespaced ingress o k response
func (o *DeleteNetworkingV1NamespacedIngressOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNetworkingV1NamespacedIngressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNetworkingV1NamespacedIngressAcceptedCode is the HTTP code returned for type DeleteNetworkingV1NamespacedIngressAccepted
const DeleteNetworkingV1NamespacedIngressAcceptedCode int = 202

/*DeleteNetworkingV1NamespacedIngressAccepted Accepted

swagger:response deleteNetworkingV1NamespacedIngressAccepted
*/
type DeleteNetworkingV1NamespacedIngressAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteNetworkingV1NamespacedIngressAccepted creates DeleteNetworkingV1NamespacedIngressAccepted with default headers values
func NewDeleteNetworkingV1NamespacedIngressAccepted() *DeleteNetworkingV1NamespacedIngressAccepted {

	return &DeleteNetworkingV1NamespacedIngressAccepted{}
}

// WithPayload adds the payload to the delete networking v1 namespaced ingress accepted response
func (o *DeleteNetworkingV1NamespacedIngressAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteNetworkingV1NamespacedIngressAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete networking v1 namespaced ingress accepted response
func (o *DeleteNetworkingV1NamespacedIngressAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteNetworkingV1NamespacedIngressAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteNetworkingV1NamespacedIngressUnauthorizedCode is the HTTP code returned for type DeleteNetworkingV1NamespacedIngressUnauthorized
const DeleteNetworkingV1NamespacedIngressUnauthorizedCode int = 401

/*DeleteNetworkingV1NamespacedIngressUnauthorized Unauthorized

swagger:response deleteNetworkingV1NamespacedIngressUnauthorized
*/
type DeleteNetworkingV1NamespacedIngressUnauthorized struct {
}

// NewDeleteNetworkingV1NamespacedIngressUnauthorized creates DeleteNetworkingV1NamespacedIngressUnauthorized with default headers values
func NewDeleteNetworkingV1NamespacedIngressUnauthorized() *DeleteNetworkingV1NamespacedIngressUnauthorized {

	return &DeleteNetworkingV1NamespacedIngressUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteNetworkingV1NamespacedIngressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
