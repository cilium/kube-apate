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

// ListCertificatesV1beta1CertificateSigningRequestOKCode is the HTTP code returned for type ListCertificatesV1beta1CertificateSigningRequestOK
const ListCertificatesV1beta1CertificateSigningRequestOKCode int = 200

/*ListCertificatesV1beta1CertificateSigningRequestOK OK

swagger:response listCertificatesV1beta1CertificateSigningRequestOK
*/
type ListCertificatesV1beta1CertificateSigningRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequestList `json:"body,omitempty"`
}

// NewListCertificatesV1beta1CertificateSigningRequestOK creates ListCertificatesV1beta1CertificateSigningRequestOK with default headers values
func NewListCertificatesV1beta1CertificateSigningRequestOK() *ListCertificatesV1beta1CertificateSigningRequestOK {

	return &ListCertificatesV1beta1CertificateSigningRequestOK{}
}

// WithPayload adds the payload to the list certificates v1beta1 certificate signing request o k response
func (o *ListCertificatesV1beta1CertificateSigningRequestOK) WithPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequestList) *ListCertificatesV1beta1CertificateSigningRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list certificates v1beta1 certificate signing request o k response
func (o *ListCertificatesV1beta1CertificateSigningRequestOK) SetPayload(payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequestList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListCertificatesV1beta1CertificateSigningRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListCertificatesV1beta1CertificateSigningRequestUnauthorizedCode is the HTTP code returned for type ListCertificatesV1beta1CertificateSigningRequestUnauthorized
const ListCertificatesV1beta1CertificateSigningRequestUnauthorizedCode int = 401

/*ListCertificatesV1beta1CertificateSigningRequestUnauthorized Unauthorized

swagger:response listCertificatesV1beta1CertificateSigningRequestUnauthorized
*/
type ListCertificatesV1beta1CertificateSigningRequestUnauthorized struct {
}

// NewListCertificatesV1beta1CertificateSigningRequestUnauthorized creates ListCertificatesV1beta1CertificateSigningRequestUnauthorized with default headers values
func NewListCertificatesV1beta1CertificateSigningRequestUnauthorized() *ListCertificatesV1beta1CertificateSigningRequestUnauthorized {

	return &ListCertificatesV1beta1CertificateSigningRequestUnauthorized{}
}

// WriteResponse to the client
func (o *ListCertificatesV1beta1CertificateSigningRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}
