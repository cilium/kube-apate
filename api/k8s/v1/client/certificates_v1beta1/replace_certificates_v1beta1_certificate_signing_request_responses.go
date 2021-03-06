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

// ReplaceCertificatesV1beta1CertificateSigningRequestReader is a Reader for the ReplaceCertificatesV1beta1CertificateSigningRequest structure.
type ReplaceCertificatesV1beta1CertificateSigningRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestOK creates a ReplaceCertificatesV1beta1CertificateSigningRequestOK with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestOK() *ReplaceCertificatesV1beta1CertificateSigningRequestOK {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestOK{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestOK handles this case with default header values.

OK
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestOK struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestOK) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}][%d] replaceCertificatesV1beta1CertificateSigningRequestOK  %+v", 200, o.Payload)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestOK) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestCreated creates a ReplaceCertificatesV1beta1CertificateSigningRequestCreated with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestCreated() *ReplaceCertificatesV1beta1CertificateSigningRequestCreated {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestCreated{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestCreated handles this case with default header values.

Created
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestCreated struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}][%d] replaceCertificatesV1beta1CertificateSigningRequestCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestCreated) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized creates a ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized() *ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized struct {
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}][%d] replaceCertificatesV1beta1CertificateSigningRequestUnauthorized ", 401)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
