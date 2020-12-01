// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package certificates_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchCertificatesV1CertificateSigningRequestStatusReader is a Reader for the PatchCertificatesV1CertificateSigningRequestStatus structure.
type PatchCertificatesV1CertificateSigningRequestStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCertificatesV1CertificateSigningRequestStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCertificatesV1CertificateSigningRequestStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCertificatesV1CertificateSigningRequestStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCertificatesV1CertificateSigningRequestStatusOK creates a PatchCertificatesV1CertificateSigningRequestStatusOK with default headers values
func NewPatchCertificatesV1CertificateSigningRequestStatusOK() *PatchCertificatesV1CertificateSigningRequestStatusOK {
	return &PatchCertificatesV1CertificateSigningRequestStatusOK{}
}

/*PatchCertificatesV1CertificateSigningRequestStatusOK handles this case with default header values.

OK
*/
type PatchCertificatesV1CertificateSigningRequestStatusOK struct {
	Payload *models.IoK8sAPICertificatesV1CertificateSigningRequest
}

func (o *PatchCertificatesV1CertificateSigningRequestStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/status][%d] patchCertificatesV1CertificateSigningRequestStatusOK  %+v", 200, o.Payload)
}

func (o *PatchCertificatesV1CertificateSigningRequestStatusOK) GetPayload() *models.IoK8sAPICertificatesV1CertificateSigningRequest {
	return o.Payload
}

func (o *PatchCertificatesV1CertificateSigningRequestStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCertificatesV1CertificateSigningRequestStatusUnauthorized creates a PatchCertificatesV1CertificateSigningRequestStatusUnauthorized with default headers values
func NewPatchCertificatesV1CertificateSigningRequestStatusUnauthorized() *PatchCertificatesV1CertificateSigningRequestStatusUnauthorized {
	return &PatchCertificatesV1CertificateSigningRequestStatusUnauthorized{}
}

/*PatchCertificatesV1CertificateSigningRequestStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCertificatesV1CertificateSigningRequestStatusUnauthorized struct {
}

func (o *PatchCertificatesV1CertificateSigningRequestStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/certificates.k8s.io/v1/certificatesigningrequests/{name}/status][%d] patchCertificatesV1CertificateSigningRequestStatusUnauthorized ", 401)
}

func (o *PatchCertificatesV1CertificateSigningRequestStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
