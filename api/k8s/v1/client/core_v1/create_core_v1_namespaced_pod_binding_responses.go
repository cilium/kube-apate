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

// CreateCoreV1NamespacedPodBindingReader is a Reader for the CreateCoreV1NamespacedPodBinding structure.
type CreateCoreV1NamespacedPodBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCoreV1NamespacedPodBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCoreV1NamespacedPodBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateCoreV1NamespacedPodBindingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCoreV1NamespacedPodBindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateCoreV1NamespacedPodBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCoreV1NamespacedPodBindingOK creates a CreateCoreV1NamespacedPodBindingOK with default headers values
func NewCreateCoreV1NamespacedPodBindingOK() *CreateCoreV1NamespacedPodBindingOK {
	return &CreateCoreV1NamespacedPodBindingOK{}
}

/*CreateCoreV1NamespacedPodBindingOK handles this case with default header values.

OK
*/
type CreateCoreV1NamespacedPodBindingOK struct {
	Payload *models.IoK8sAPICoreV1Binding
}

func (o *CreateCoreV1NamespacedPodBindingOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/binding][%d] createCoreV1NamespacedPodBindingOK  %+v", 200, o.Payload)
}

func (o *CreateCoreV1NamespacedPodBindingOK) GetPayload() *models.IoK8sAPICoreV1Binding {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Binding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodBindingCreated creates a CreateCoreV1NamespacedPodBindingCreated with default headers values
func NewCreateCoreV1NamespacedPodBindingCreated() *CreateCoreV1NamespacedPodBindingCreated {
	return &CreateCoreV1NamespacedPodBindingCreated{}
}

/*CreateCoreV1NamespacedPodBindingCreated handles this case with default header values.

Created
*/
type CreateCoreV1NamespacedPodBindingCreated struct {
	Payload *models.IoK8sAPICoreV1Binding
}

func (o *CreateCoreV1NamespacedPodBindingCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/binding][%d] createCoreV1NamespacedPodBindingCreated  %+v", 201, o.Payload)
}

func (o *CreateCoreV1NamespacedPodBindingCreated) GetPayload() *models.IoK8sAPICoreV1Binding {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodBindingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Binding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodBindingAccepted creates a CreateCoreV1NamespacedPodBindingAccepted with default headers values
func NewCreateCoreV1NamespacedPodBindingAccepted() *CreateCoreV1NamespacedPodBindingAccepted {
	return &CreateCoreV1NamespacedPodBindingAccepted{}
}

/*CreateCoreV1NamespacedPodBindingAccepted handles this case with default header values.

Accepted
*/
type CreateCoreV1NamespacedPodBindingAccepted struct {
	Payload *models.IoK8sAPICoreV1Binding
}

func (o *CreateCoreV1NamespacedPodBindingAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/binding][%d] createCoreV1NamespacedPodBindingAccepted  %+v", 202, o.Payload)
}

func (o *CreateCoreV1NamespacedPodBindingAccepted) GetPayload() *models.IoK8sAPICoreV1Binding {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodBindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Binding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodBindingUnauthorized creates a CreateCoreV1NamespacedPodBindingUnauthorized with default headers values
func NewCreateCoreV1NamespacedPodBindingUnauthorized() *CreateCoreV1NamespacedPodBindingUnauthorized {
	return &CreateCoreV1NamespacedPodBindingUnauthorized{}
}

/*CreateCoreV1NamespacedPodBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCoreV1NamespacedPodBindingUnauthorized struct {
}

func (o *CreateCoreV1NamespacedPodBindingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/binding][%d] createCoreV1NamespacedPodBindingUnauthorized ", 401)
}

func (o *CreateCoreV1NamespacedPodBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}