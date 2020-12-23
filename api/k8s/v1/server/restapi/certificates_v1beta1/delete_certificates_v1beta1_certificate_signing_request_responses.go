// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteCertificatesV1beta1CertificateSigningRequestOKCode is the HTTP code returned for type DeleteCertificatesV1beta1CertificateSigningRequestOK
const DeleteCertificatesV1beta1CertificateSigningRequestOKCode int = 200

/*DeleteCertificatesV1beta1CertificateSigningRequestOK OK

swagger:response deleteCertificatesV1beta1CertificateSigningRequestOK
*/
type DeleteCertificatesV1beta1CertificateSigningRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCertificatesV1beta1CertificateSigningRequestOK creates DeleteCertificatesV1beta1CertificateSigningRequestOK with default headers values
func NewDeleteCertificatesV1beta1CertificateSigningRequestOK() *DeleteCertificatesV1beta1CertificateSigningRequestOK {

	return &DeleteCertificatesV1beta1CertificateSigningRequestOK{}
}

// WithPayload adds the payload to the delete certificates v1beta1 certificate signing request o k response
func (o *DeleteCertificatesV1beta1CertificateSigningRequestOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCertificatesV1beta1CertificateSigningRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete certificates v1beta1 certificate signing request o k response
func (o *DeleteCertificatesV1beta1CertificateSigningRequestOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCertificatesV1beta1CertificateSigningRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCertificatesV1beta1CertificateSigningRequestAcceptedCode is the HTTP code returned for type DeleteCertificatesV1beta1CertificateSigningRequestAccepted
const DeleteCertificatesV1beta1CertificateSigningRequestAcceptedCode int = 202

/*DeleteCertificatesV1beta1CertificateSigningRequestAccepted Accepted

swagger:response deleteCertificatesV1beta1CertificateSigningRequestAccepted
*/
type DeleteCertificatesV1beta1CertificateSigningRequestAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status `json:"body,omitempty"`
}

// NewDeleteCertificatesV1beta1CertificateSigningRequestAccepted creates DeleteCertificatesV1beta1CertificateSigningRequestAccepted with default headers values
func NewDeleteCertificatesV1beta1CertificateSigningRequestAccepted() *DeleteCertificatesV1beta1CertificateSigningRequestAccepted {

	return &DeleteCertificatesV1beta1CertificateSigningRequestAccepted{}
}

// WithPayload adds the payload to the delete certificates v1beta1 certificate signing request accepted response
func (o *DeleteCertificatesV1beta1CertificateSigningRequestAccepted) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) *DeleteCertificatesV1beta1CertificateSigningRequestAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete certificates v1beta1 certificate signing request accepted response
func (o *DeleteCertificatesV1beta1CertificateSigningRequestAccepted) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1Status) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteCertificatesV1beta1CertificateSigningRequestAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteCertificatesV1beta1CertificateSigningRequestUnauthorizedCode is the HTTP code returned for type DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized
const DeleteCertificatesV1beta1CertificateSigningRequestUnauthorizedCode int = 401

/*DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized Unauthorized

swagger:response deleteCertificatesV1beta1CertificateSigningRequestUnauthorized
*/
type DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized struct {
}

// NewDeleteCertificatesV1beta1CertificateSigningRequestUnauthorized creates DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized with default headers values
func NewDeleteCertificatesV1beta1CertificateSigningRequestUnauthorized() *DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized {

	return &DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteCertificatesV1beta1CertificateSigningRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}