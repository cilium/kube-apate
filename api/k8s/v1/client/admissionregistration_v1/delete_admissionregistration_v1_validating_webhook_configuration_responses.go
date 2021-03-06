// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteAdmissionregistrationV1ValidatingWebhookConfigurationReader is a Reader for the DeleteAdmissionregistrationV1ValidatingWebhookConfiguration structure.
type DeleteAdmissionregistrationV1ValidatingWebhookConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK creates a DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK with default headers values
func NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK() *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK {
	return &DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK{}
}

/*DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK handles this case with default header values.

OK
*/
type DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}][%d] deleteAdmissionregistrationV1ValidatingWebhookConfigurationOK  %+v", 200, o.Payload)
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted creates a DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted with default headers values
func NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted() *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted {
	return &DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted{}
}

/*DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted handles this case with default header values.

Accepted
*/
type DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}][%d] deleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted  %+v", 202, o.Payload)
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized creates a DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized with default headers values
func NewDeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized() *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized {
	return &DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized{}
}

/*DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized struct {
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/admissionregistration.k8s.io/v1/validatingwebhookconfigurations/{name}][%d] deleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized ", 401)
}

func (o *DeleteAdmissionregistrationV1ValidatingWebhookConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
