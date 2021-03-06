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

// DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationReader is a Reader for the DeleteAdmissionregistrationV1beta1MutatingWebhookConfiguration structure.
type DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK creates a DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK with default headers values
func NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK() *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK {
	return &DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK{}
}

/*DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK handles this case with default header values.

OK
*/
type DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name}][%d] deleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK  %+v", 200, o.Payload)
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted creates a DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted with default headers values
func NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted() *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted {
	return &DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted{}
}

/*DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted handles this case with default header values.

Accepted
*/
type DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name}][%d] deleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted  %+v", 202, o.Payload)
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized creates a DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized with default headers values
func NewDeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized() *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized {
	return &DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized{}
}

/*DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized struct {
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/admissionregistration.k8s.io/v1beta1/mutatingwebhookconfigurations/{name}][%d] deleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized ", 401)
}

func (o *DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
