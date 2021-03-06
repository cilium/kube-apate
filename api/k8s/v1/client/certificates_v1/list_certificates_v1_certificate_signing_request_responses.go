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

// ListCertificatesV1CertificateSigningRequestReader is a Reader for the ListCertificatesV1CertificateSigningRequest structure.
type ListCertificatesV1CertificateSigningRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCertificatesV1CertificateSigningRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCertificatesV1CertificateSigningRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCertificatesV1CertificateSigningRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCertificatesV1CertificateSigningRequestOK creates a ListCertificatesV1CertificateSigningRequestOK with default headers values
func NewListCertificatesV1CertificateSigningRequestOK() *ListCertificatesV1CertificateSigningRequestOK {
	return &ListCertificatesV1CertificateSigningRequestOK{}
}

/*ListCertificatesV1CertificateSigningRequestOK handles this case with default header values.

OK
*/
type ListCertificatesV1CertificateSigningRequestOK struct {
	Payload *models.IoK8sAPICertificatesV1CertificateSigningRequestList
}

func (o *ListCertificatesV1CertificateSigningRequestOK) Error() string {
	return fmt.Sprintf("[GET /apis/certificates.k8s.io/v1/certificatesigningrequests][%d] listCertificatesV1CertificateSigningRequestOK  %+v", 200, o.Payload)
}

func (o *ListCertificatesV1CertificateSigningRequestOK) GetPayload() *models.IoK8sAPICertificatesV1CertificateSigningRequestList {
	return o.Payload
}

func (o *ListCertificatesV1CertificateSigningRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1CertificateSigningRequestList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCertificatesV1CertificateSigningRequestUnauthorized creates a ListCertificatesV1CertificateSigningRequestUnauthorized with default headers values
func NewListCertificatesV1CertificateSigningRequestUnauthorized() *ListCertificatesV1CertificateSigningRequestUnauthorized {
	return &ListCertificatesV1CertificateSigningRequestUnauthorized{}
}

/*ListCertificatesV1CertificateSigningRequestUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCertificatesV1CertificateSigningRequestUnauthorized struct {
}

func (o *ListCertificatesV1CertificateSigningRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/certificates.k8s.io/v1/certificatesigningrequests][%d] listCertificatesV1CertificateSigningRequestUnauthorized ", 401)
}

func (o *ListCertificatesV1CertificateSigningRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
