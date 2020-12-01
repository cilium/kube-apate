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

// CreateCoreV1NamespacedResourceQuotaReader is a Reader for the CreateCoreV1NamespacedResourceQuota structure.
type CreateCoreV1NamespacedResourceQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCoreV1NamespacedResourceQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCoreV1NamespacedResourceQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateCoreV1NamespacedResourceQuotaCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCoreV1NamespacedResourceQuotaAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateCoreV1NamespacedResourceQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCoreV1NamespacedResourceQuotaOK creates a CreateCoreV1NamespacedResourceQuotaOK with default headers values
func NewCreateCoreV1NamespacedResourceQuotaOK() *CreateCoreV1NamespacedResourceQuotaOK {
	return &CreateCoreV1NamespacedResourceQuotaOK{}
}

/*CreateCoreV1NamespacedResourceQuotaOK handles this case with default header values.

OK
*/
type CreateCoreV1NamespacedResourceQuotaOK struct {
	Payload *models.IoK8sAPICoreV1ResourceQuota
}

func (o *CreateCoreV1NamespacedResourceQuotaOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/resourcequotas][%d] createCoreV1NamespacedResourceQuotaOK  %+v", 200, o.Payload)
}

func (o *CreateCoreV1NamespacedResourceQuotaOK) GetPayload() *models.IoK8sAPICoreV1ResourceQuota {
	return o.Payload
}

func (o *CreateCoreV1NamespacedResourceQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedResourceQuotaCreated creates a CreateCoreV1NamespacedResourceQuotaCreated with default headers values
func NewCreateCoreV1NamespacedResourceQuotaCreated() *CreateCoreV1NamespacedResourceQuotaCreated {
	return &CreateCoreV1NamespacedResourceQuotaCreated{}
}

/*CreateCoreV1NamespacedResourceQuotaCreated handles this case with default header values.

Created
*/
type CreateCoreV1NamespacedResourceQuotaCreated struct {
	Payload *models.IoK8sAPICoreV1ResourceQuota
}

func (o *CreateCoreV1NamespacedResourceQuotaCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/resourcequotas][%d] createCoreV1NamespacedResourceQuotaCreated  %+v", 201, o.Payload)
}

func (o *CreateCoreV1NamespacedResourceQuotaCreated) GetPayload() *models.IoK8sAPICoreV1ResourceQuota {
	return o.Payload
}

func (o *CreateCoreV1NamespacedResourceQuotaCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedResourceQuotaAccepted creates a CreateCoreV1NamespacedResourceQuotaAccepted with default headers values
func NewCreateCoreV1NamespacedResourceQuotaAccepted() *CreateCoreV1NamespacedResourceQuotaAccepted {
	return &CreateCoreV1NamespacedResourceQuotaAccepted{}
}

/*CreateCoreV1NamespacedResourceQuotaAccepted handles this case with default header values.

Accepted
*/
type CreateCoreV1NamespacedResourceQuotaAccepted struct {
	Payload *models.IoK8sAPICoreV1ResourceQuota
}

func (o *CreateCoreV1NamespacedResourceQuotaAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/resourcequotas][%d] createCoreV1NamespacedResourceQuotaAccepted  %+v", 202, o.Payload)
}

func (o *CreateCoreV1NamespacedResourceQuotaAccepted) GetPayload() *models.IoK8sAPICoreV1ResourceQuota {
	return o.Payload
}

func (o *CreateCoreV1NamespacedResourceQuotaAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedResourceQuotaUnauthorized creates a CreateCoreV1NamespacedResourceQuotaUnauthorized with default headers values
func NewCreateCoreV1NamespacedResourceQuotaUnauthorized() *CreateCoreV1NamespacedResourceQuotaUnauthorized {
	return &CreateCoreV1NamespacedResourceQuotaUnauthorized{}
}

/*CreateCoreV1NamespacedResourceQuotaUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCoreV1NamespacedResourceQuotaUnauthorized struct {
}

func (o *CreateCoreV1NamespacedResourceQuotaUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/resourcequotas][%d] createCoreV1NamespacedResourceQuotaUnauthorized ", 401)
}

func (o *CreateCoreV1NamespacedResourceQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
