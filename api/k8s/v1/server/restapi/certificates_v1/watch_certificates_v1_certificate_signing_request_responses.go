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

// WatchCertificatesV1CertificateSigningRequestOKCode is the HTTP code returned for type WatchCertificatesV1CertificateSigningRequestOK
const WatchCertificatesV1CertificateSigningRequestOKCode int = 200

/*WatchCertificatesV1CertificateSigningRequestOK OK

swagger:response watchCertificatesV1CertificateSigningRequestOK
*/
type WatchCertificatesV1CertificateSigningRequestOK struct {

	/*
	  In: Body
	*/
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent `json:"body,omitempty"`
}

// NewWatchCertificatesV1CertificateSigningRequestOK creates WatchCertificatesV1CertificateSigningRequestOK with default headers values
func NewWatchCertificatesV1CertificateSigningRequestOK() *WatchCertificatesV1CertificateSigningRequestOK {

	return &WatchCertificatesV1CertificateSigningRequestOK{}
}

// WithPayload adds the payload to the watch certificates v1 certificate signing request o k response
func (o *WatchCertificatesV1CertificateSigningRequestOK) WithPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) *WatchCertificatesV1CertificateSigningRequestOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the watch certificates v1 certificate signing request o k response
func (o *WatchCertificatesV1CertificateSigningRequestOK) SetPayload(payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WatchCertificatesV1CertificateSigningRequestOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WatchCertificatesV1CertificateSigningRequestUnauthorizedCode is the HTTP code returned for type WatchCertificatesV1CertificateSigningRequestUnauthorized
const WatchCertificatesV1CertificateSigningRequestUnauthorizedCode int = 401

/*WatchCertificatesV1CertificateSigningRequestUnauthorized Unauthorized

swagger:response watchCertificatesV1CertificateSigningRequestUnauthorized
*/
type WatchCertificatesV1CertificateSigningRequestUnauthorized struct {
}

// NewWatchCertificatesV1CertificateSigningRequestUnauthorized creates WatchCertificatesV1CertificateSigningRequestUnauthorized with default headers values
func NewWatchCertificatesV1CertificateSigningRequestUnauthorized() *WatchCertificatesV1CertificateSigningRequestUnauthorized {

	return &WatchCertificatesV1CertificateSigningRequestUnauthorized{}
}

// WriteResponse to the client
func (o *WatchCertificatesV1CertificateSigningRequestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}