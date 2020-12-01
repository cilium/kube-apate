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

// WatchCertificatesV1beta1CertificateSigningRequestListReader is a Reader for the WatchCertificatesV1beta1CertificateSigningRequestList structure.
type WatchCertificatesV1beta1CertificateSigningRequestListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCertificatesV1beta1CertificateSigningRequestListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCertificatesV1beta1CertificateSigningRequestListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCertificatesV1beta1CertificateSigningRequestListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCertificatesV1beta1CertificateSigningRequestListOK creates a WatchCertificatesV1beta1CertificateSigningRequestListOK with default headers values
func NewWatchCertificatesV1beta1CertificateSigningRequestListOK() *WatchCertificatesV1beta1CertificateSigningRequestListOK {
	return &WatchCertificatesV1beta1CertificateSigningRequestListOK{}
}

/*WatchCertificatesV1beta1CertificateSigningRequestListOK handles this case with default header values.

OK
*/
type WatchCertificatesV1beta1CertificateSigningRequestListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCertificatesV1beta1CertificateSigningRequestListOK) Error() string {
	return fmt.Sprintf("[GET /apis/certificates.k8s.io/v1beta1/watch/certificatesigningrequests][%d] watchCertificatesV1beta1CertificateSigningRequestListOK  %+v", 200, o.Payload)
}

func (o *WatchCertificatesV1beta1CertificateSigningRequestListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCertificatesV1beta1CertificateSigningRequestListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCertificatesV1beta1CertificateSigningRequestListUnauthorized creates a WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized with default headers values
func NewWatchCertificatesV1beta1CertificateSigningRequestListUnauthorized() *WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized {
	return &WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized{}
}

/*WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized struct {
}

func (o *WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/certificates.k8s.io/v1beta1/watch/certificatesigningrequests][%d] watchCertificatesV1beta1CertificateSigningRequestListUnauthorized ", 401)
}

func (o *WatchCertificatesV1beta1CertificateSigningRequestListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
