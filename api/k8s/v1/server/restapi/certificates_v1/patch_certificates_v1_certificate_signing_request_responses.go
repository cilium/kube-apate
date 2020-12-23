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

// PatchCertificatesV1CertificateSigningRequestOKCode is the HTTP code returned for type PatchCertificatesV1CertificateSigningRequestOK
const PatchCertificatesV1CertificateSigningRequestOKCode int = 200

/*PatchCertificatesV1CertificateSigningRequestOK OK

swagger:response patchCertificatesV1CertificateSigningRequestOK
*/
type PatchCertificatesV1CertificateSigningRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1CertificateSigningRequest `json:"body,omitempty"`
}

// NewPatchCertificatesV1CertificateSigningRequestOK creates PatchCertificatesV1CertificateSigningRequestOK with default headers values
func NewPatchCertificatesV1CertificateSigningRequestOK() *PatchCertificatesV1CertificateSigningRequestOK {

	return &PatchCertificatesV1CertificateSigningRequestOK{}
}

// WithPayload adds the payload to the patch certificates v1 certificate signing request o k response
func (o *PatchCertificatesV1CertificateSigningRequestOK) WithPayload(payload *models.IoK8sAPICertificatesV1CertificateSigningRequest) *PatchCertificatesV1CertificateSigningRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch certificates v1 certificate signing request o k response
func (o *PatchCertificatesV1CertificateSigningRequestOK) SetPayload(payload *models.IoK8sAPICertificatesV1CertificateSigningRequest) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchCertificatesV1CertificateSigningRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchCertificatesV1CertificateSigningRequestUnauthorizedCode is the HTTP code returned for type PatchCertificatesV1CertificateSigningRequestUnauthorized
const PatchCertificatesV1CertificateSigningRequestUnauthorizedCode int = 401

/*PatchCertificatesV1CertificateSigningRequestUnauthorized Unauthorized

swagger:response patchCertificatesV1CertificateSigningRequestUnauthorized
*/
type PatchCertificatesV1CertificateSigningRequestUnauthorized struct {
}

// NewPatchCertificatesV1CertificateSigningRequestUnauthorized creates PatchCertificatesV1CertificateSigningRequestUnauthorized with default headers values
func NewPatchCertificatesV1CertificateSigningRequestUnauthorized() *PatchCertificatesV1CertificateSigningRequestUnauthorized {

	return &PatchCertificatesV1CertificateSigningRequestUnauthorized{}
}

// WriteResponse to the client
func (o *PatchCertificatesV1CertificateSigningRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}