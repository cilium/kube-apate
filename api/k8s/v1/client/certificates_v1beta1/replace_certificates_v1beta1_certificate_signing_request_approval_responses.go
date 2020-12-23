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

// ReplaceCertificatesV1beta1CertificateSigningRequestApprovalReader is a Reader for the ReplaceCertificatesV1beta1CertificateSigningRequestApproval structure.
type ReplaceCertificatesV1beta1CertificateSigningRequestApprovalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK creates a ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK() *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK handles this case with default header values.

OK
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}/approval][%d] replaceCertificatesV1beta1CertificateSigningRequestApprovalOK  %+v", 200, o.Payload)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated creates a ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated() *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated handles this case with default header values.

Created
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}/approval][%d] replaceCertificatesV1beta1CertificateSigningRequestApprovalCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized creates a ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized with default headers values
func NewReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized() *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized {
	return &ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized{}
}

/*ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized struct {
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/certificates.k8s.io/v1beta1/certificatesigningrequests/{name}/approval][%d] replaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized ", 401)
}

func (o *ReplaceCertificatesV1beta1CertificateSigningRequestApprovalUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}