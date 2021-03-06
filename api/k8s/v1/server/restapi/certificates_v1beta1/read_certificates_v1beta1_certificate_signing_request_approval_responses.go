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

// ReadCertificatesV1beta1CertificateSigningRequestApprovalOKCode is the HTTP code returned for type ReadCertificatesV1beta1CertificateSigningRequestApprovalOK
const ReadCertificatesV1beta1CertificateSigningRequestApprovalOKCode int = 200

/*ReadCertificatesV1beta1CertificateSigningRequestApprovalOK OK

swagger:response readCertificatesV1beta1CertificateSigningRequestApprovalOK
*/
type ReadCertificatesV1beta1CertificateSigningRequestApprovalOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest `json:"body,omitempty"`
}

// NewReadCertificatesV1beta1CertificateSigningRequestApprovalOK creates ReadCertificatesV1beta1CertificateSigningRequestApprovalOK with default headers values
func NewReadCertificatesV1beta1CertificateSigningRequestApprovalOK() *ReadCertificatesV1beta1CertificateSigningRequestApprovalOK {

	return &ReadCertificatesV1beta1CertificateSigningRequestApprovalOK{}
}

// WithPayload adds the payload to the read certificates v1beta1 certificate signing request approval o k response
func (o *ReadCertificatesV1beta1CertificateSigningRequestApprovalOK) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) *ReadCertificatesV1beta1CertificateSigningRequestApprovalOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read certificates v1beta1 certificate signing request approval o k response
func (o *ReadCertificatesV1beta1CertificateSigningRequestApprovalOK) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCertificatesV1beta1CertificateSigningRequestApprovalOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorizedCode is the HTTP code returned for type ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized
const ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorizedCode int = 401

/*ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized Unauthorized

swagger:response readCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized
*/
type ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized struct {
}

// NewReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized creates ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized with default headers values
func NewReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized() *ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized {

	return &ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
