// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceCertificatesV1beta1CertificateSigningRequestStatusReader is a Reader for the ReplaceCertificatesV1beta1CertificateSigningRequestStatus structure.
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestStatusOK creates a ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestStatusOK() *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK handles this case with default header values.

OK
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}/status][%d] replaceCertificatesV1beta1CertificateSigningRequestStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated creates a ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated() *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated handles this case with default header values.

Created
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}/status][%d] replaceCertificatesV1beta1CertificateSigningRequestStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized creates a ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized() *ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized struct {
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}/status][%d] replaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized ", 401)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
