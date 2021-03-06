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

// ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationReader is a Reader for the ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfiguration structure.
type ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK creates a ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK with default headers values
func NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK() *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK {
	return &ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK{}
}

/*ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK handles this case with default header values.

OK
*/
type ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK struct {
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name}][%d] replaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK  %+v", 200, o.Payload)
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) GetPayload() *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration {
	return o.Payload
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated creates a ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated with default headers values
func NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated() *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated {
	return &ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated{}
}

/*ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated handles this case with default header values.

Created
*/
type ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated struct {
	Payload *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name}][%d] replaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated) GetPayload() *models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration {
	return o.Payload
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAdmissionregistrationV1beta1ValidatingWebhookConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized creates a ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized() *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized {
	return &ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized{}
}

/*ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized struct {
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/admissionregistration.k8s.io/v1beta1/validatingwebhookconfigurations/{name}][%d] replaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized ", 401)
}

func (o *ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
