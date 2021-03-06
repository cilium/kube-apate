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

// ListCertificatesV1beta1CertificateSigningRequestReader is a Reader for the ListCertificatesV1beta1CertificateSigningRequest structure.
type ListCertificatesV1beta1CertificateSigningRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCertificatesV1beta1CertificateSigningRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCertificatesV1beta1CertificateSigningRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCertificatesV1beta1CertificateSigningRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCertificatesV1beta1CertificateSigningRequestOK creates a ListCertificatesV1beta1CertificateSigningRequestOK with default headers values
func NewListCertificatesV1beta1CertificateSigningRequestOK() *ListCertificatesV1beta1CertificateSigningRequestOK {
	return &ListCertificatesV1beta1CertificateSigningRequestOK{}
}

/*ListCertificatesV1beta1CertificateSigningRequestOK handles this case with default header values.

OK
*/
type ListCertificatesV1beta1CertificateSigningRequestOK struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequestList
}

func (o *ListCertificatesV1beta1CertificateSigningRequestOK) Error() string {
	return fmt.Sprintf("[GET /apis/certificates.k8s.io/v1beta1/certificatesigningrequests][%d] listCertificatesV1beta1CertificateSigningRequestOK  %+v", 200, o.Payload)
}

func (o *ListCertificatesV1beta1CertificateSigningRequestOK) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequestList {
	return o.Payload
}

func (o *ListCertificatesV1beta1CertificateSigningRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequestList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCertificatesV1beta1CertificateSigningRequestUnauthorized creates a ListCertificatesV1beta1CertificateSigningRequestUnauthorized with default headers values
func NewListCertificatesV1beta1CertificateSigningRequestUnauthorized() *ListCertificatesV1beta1CertificateSigningRequestUnauthorized {
	return &ListCertificatesV1beta1CertificateSigningRequestUnauthorized{}
}

/*ListCertificatesV1beta1CertificateSigningRequestUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCertificatesV1beta1CertificateSigningRequestUnauthorized struct {
}

func (o *ListCertificatesV1beta1CertificateSigningRequestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/certificates.k8s.io/v1beta1/certificatesigningrequests][%d] listCertificatesV1beta1CertificateSigningRequestUnauthorized ", 401)
}

func (o *ListCertificatesV1beta1CertificateSigningRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
