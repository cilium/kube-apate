// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateCoreV1NamespacedServiceReader is a Reader for the CreateCoreV1NamespacedService structure.
type CreateCoreV1NamespacedServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCoreV1NamespacedServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCoreV1NamespacedServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateCoreV1NamespacedServiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCoreV1NamespacedServiceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateCoreV1NamespacedServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCoreV1NamespacedServiceOK creates a CreateCoreV1NamespacedServiceOK with default headers values
func NewCreateCoreV1NamespacedServiceOK() *CreateCoreV1NamespacedServiceOK {
	return &CreateCoreV1NamespacedServiceOK{}
}

/*CreateCoreV1NamespacedServiceOK handles this case with default header values.

OK
*/
type CreateCoreV1NamespacedServiceOK struct {
	Payload *models.IoK8sAPICoreV1Service
}

func (o *CreateCoreV1NamespacedServiceOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/services][%d] createCoreV1NamespacedServiceOK  %+v", 200, o.Payload)
}

func (o *CreateCoreV1NamespacedServiceOK) GetPayload() *models.IoK8sAPICoreV1Service {
	return o.Payload
}

func (o *CreateCoreV1NamespacedServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedServiceCreated creates a CreateCoreV1NamespacedServiceCreated with default headers values
func NewCreateCoreV1NamespacedServiceCreated() *CreateCoreV1NamespacedServiceCreated {
	return &CreateCoreV1NamespacedServiceCreated{}
}

/*CreateCoreV1NamespacedServiceCreated handles this case with default header values.

Created
*/
type CreateCoreV1NamespacedServiceCreated struct {
	Payload *models.IoK8sAPICoreV1Service
}

func (o *CreateCoreV1NamespacedServiceCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/services][%d] createCoreV1NamespacedServiceCreated  %+v", 201, o.Payload)
}

func (o *CreateCoreV1NamespacedServiceCreated) GetPayload() *models.IoK8sAPICoreV1Service {
	return o.Payload
}

func (o *CreateCoreV1NamespacedServiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedServiceAccepted creates a CreateCoreV1NamespacedServiceAccepted with default headers values
func NewCreateCoreV1NamespacedServiceAccepted() *CreateCoreV1NamespacedServiceAccepted {
	return &CreateCoreV1NamespacedServiceAccepted{}
}

/*CreateCoreV1NamespacedServiceAccepted handles this case with default header values.

Accepted
*/
type CreateCoreV1NamespacedServiceAccepted struct {
	Payload *models.IoK8sAPICoreV1Service
}

func (o *CreateCoreV1NamespacedServiceAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/services][%d] createCoreV1NamespacedServiceAccepted  %+v", 202, o.Payload)
}

func (o *CreateCoreV1NamespacedServiceAccepted) GetPayload() *models.IoK8sAPICoreV1Service {
	return o.Payload
}

func (o *CreateCoreV1NamespacedServiceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedServiceUnauthorized creates a CreateCoreV1NamespacedServiceUnauthorized with default headers values
func NewCreateCoreV1NamespacedServiceUnauthorized() *CreateCoreV1NamespacedServiceUnauthorized {
	return &CreateCoreV1NamespacedServiceUnauthorized{}
}

/*CreateCoreV1NamespacedServiceUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCoreV1NamespacedServiceUnauthorized struct {
}

func (o *CreateCoreV1NamespacedServiceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/services][%d] createCoreV1NamespacedServiceUnauthorized ", 401)
}

func (o *CreateCoreV1NamespacedServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
