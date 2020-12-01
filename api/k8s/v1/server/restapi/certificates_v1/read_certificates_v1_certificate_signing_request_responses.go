// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadCertificatesV1CertificateSigningRequestOKCode is the HTTP code returned for type ReadCertificatesV1CertificateSigningRequestOK
const ReadCertificatesV1CertificateSigningRequestOKCode int = 200

/*ReadCertificatesV1CertificateSigningRequestOK OK

swagger:response readCertificatesV1CertificateSigningRequestOK
*/
type ReadCertificatesV1CertificateSigningRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1CertificateSigningRequest `json:"body,omitempty"`
}

// NewReadCertificatesV1CertificateSigningRequestOK creates ReadCertificatesV1CertificateSigningRequestOK with default headers values
func NewReadCertificatesV1CertificateSigningRequestOK() *ReadCertificatesV1CertificateSigningRequestOK {

	return &ReadCertificatesV1CertificateSigningRequestOK{}
}

// WithPayload adds the payload to the read certificates v1 certificate signing request o k response
func (o *ReadCertificatesV1CertificateSigningRequestOK) WithPayload(payload *models.IoK8sAPICertificatesV1CertificateSigningRequest) *ReadCertificatesV1CertificateSigningRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the read certificates v1 certificate signing request o k response
func (o *ReadCertificatesV1CertificateSigningRequestOK) SetPayload(payload *models.IoK8sAPICertificatesV1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReadCertificatesV1CertificateSigningRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReadCertificatesV1CertificateSigningRequestUnauthorizedCode is the HTTP code returned for type ReadCertificatesV1CertificateSigningRequestUnauthorized
const ReadCertificatesV1CertificateSigningRequestUnauthorizedCode int = 401

/*ReadCertificatesV1CertificateSigningRequestUnauthorized Unauthorized

swagger:response readCertificatesV1CertificateSigningRequestUnauthorized
*/
type ReadCertificatesV1CertificateSigningRequestUnauthorized struct {
}

// NewReadCertificatesV1CertificateSigningRequestUnauthorized creates ReadCertificatesV1CertificateSigningRequestUnauthorized with default headers values
func NewReadCertificatesV1CertificateSigningRequestUnauthorized() *ReadCertificatesV1CertificateSigningRequestUnauthorized {

	return &ReadCertificatesV1CertificateSigningRequestUnauthorized{}
}

// WriteResponse to the client
func (o *ReadCertificatesV1CertificateSigningRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
