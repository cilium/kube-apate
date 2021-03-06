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

// CreateCertificatesV1beta1CertificateSigningRequestReader is a Reader for the CreateCertificatesV1beta1CertificateSigningRequest structure.
type CreateCertificatesV1beta1CertificateSigningRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCertificatesV1beta1CertificateSigningRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCertificatesV1beta1CertificateSigningRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateCertificatesV1beta1CertificateSigningRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCertificatesV1beta1CertificateSigningRequestAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateCertificatesV1beta1CertificateSigningRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCertificatesV1beta1CertificateSigningRequestOK creates a CreateCertificatesV1beta1CertificateSigningRequestOK with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestOK() *CreateCertificatesV1beta1CertificateSigningRequestOK {
	return &CreateCertificatesV1beta1CertificateSigningRequestOK{}
}

/*CreateCertificatesV1beta1CertificateSigningRequestOK handles this case with default header values.

OK
*/
type CreateCertificatesV1beta1CertificateSigningRequestOK struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestOK) Error() string {
	return fmt.Sprintf("[POST /apis/certificates.k8s.io/v1beta1/certificatesigningrequests][%d] createCertificatesV1beta1CertificateSigningRequestOK  %+v", 200, o.Payload)
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestOK) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCertificatesV1beta1CertificateSigningRequestCreated creates a CreateCertificatesV1beta1CertificateSigningRequestCreated with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestCreated() *CreateCertificatesV1beta1CertificateSigningRequestCreated {
	return &CreateCertificatesV1beta1CertificateSigningRequestCreated{}
}

/*CreateCertificatesV1beta1CertificateSigningRequestCreated handles this case with default header values.

Created
*/
type CreateCertificatesV1beta1CertificateSigningRequestCreated struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestCreated) Error() string {
	return fmt.Sprintf("[POST /apis/certificates.k8s.io/v1beta1/certificatesigningrequests][%d] createCertificatesV1beta1CertificateSigningRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestCreated) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCertificatesV1beta1CertificateSigningRequestAccepted creates a CreateCertificatesV1beta1CertificateSigningRequestAccepted with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestAccepted() *CreateCertificatesV1beta1CertificateSigningRequestAccepted {
	return &CreateCertificatesV1beta1CertificateSigningRequestAccepted{}
}

/*CreateCertificatesV1beta1CertificateSigningRequestAccepted handles this case with default header values.

Accepted
*/
type CreateCertificatesV1beta1CertificateSigningRequestAccepted struct {
	Payload *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/certificates.k8s.io/v1beta1/certificatesigningrequests][%d] createCertificatesV1beta1CertificateSigningRequestAccepted  %+v", 202, o.Payload)
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestAccepted) GetPayload() *models.IoK8sAPICertificatesV1beta1CertificateSigningRequest {
	return o.Payload
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICertificatesV1beta1CertificateSigningRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCertificatesV1beta1CertificateSigningRequestUnauthorized creates a CreateCertificatesV1beta1CertificateSigningRequestUnauthorized with default headers values
func NewCreateCertificatesV1beta1CertificateSigningRequestUnauthorized() *CreateCertificatesV1beta1CertificateSigningRequestUnauthorized {
	return &CreateCertificatesV1beta1CertificateSigningRequestUnauthorized{}
}

/*CreateCertificatesV1beta1CertificateSigningRequestUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCertificatesV1beta1CertificateSigningRequestUnauthorized struct {
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/certificates.k8s.io/v1beta1/certificatesigningrequests][%d] createCertificatesV1beta1CertificateSigningRequestUnauthorized ", 401)
}

func (o *CreateCertificatesV1beta1CertificateSigningRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
