// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationReader is a Reader for the ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration structure.
type ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK creates a ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK with default headers values
func NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK() *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK {
	return &ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK{}
}

/*ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK handles this case with default header values.

OK
*/
type ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK struct {
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration
}

func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name}][%d] readAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK  %+v", 200, o.Payload)
}

func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) GetPayload() *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration {
	return o.Payload
}

func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized creates a ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized() *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized {
	return &ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized{}
}

/*ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized struct {
}

func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name}][%d] readAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized ", 401)
}

func (o *ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
