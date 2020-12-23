// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateStorageV1StorageClassReader is a Reader for the CreateStorageV1StorageClass structure.
type CreateStorageV1StorageClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageV1StorageClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStorageV1StorageClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateStorageV1StorageClassCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateStorageV1StorageClassAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateStorageV1StorageClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateStorageV1StorageClassOK creates a CreateStorageV1StorageClassOK with default headers values
func NewCreateStorageV1StorageClassOK() *CreateStorageV1StorageClassOK {
	return &CreateStorageV1StorageClassOK{}
}

/*CreateStorageV1StorageClassOK handles this case with default header values.

OK
*/
type CreateStorageV1StorageClassOK struct {
	Payload *models.IoK8sAPIStorageV1StorageClass
}

func (o *CreateStorageV1StorageClassOK) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1/storageclasses][%d] createStorageV1StorageClassOK  %+v", 200, o.Payload)
}

func (o *CreateStorageV1StorageClassOK) GetPayload() *models.IoK8sAPIStorageV1StorageClass {
	return o.Payload
}

func (o *CreateStorageV1StorageClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1StorageClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1StorageClassCreated creates a CreateStorageV1StorageClassCreated with default headers values
func NewCreateStorageV1StorageClassCreated() *CreateStorageV1StorageClassCreated {
	return &CreateStorageV1StorageClassCreated{}
}

/*CreateStorageV1StorageClassCreated handles this case with default header values.

Created
*/
type CreateStorageV1StorageClassCreated struct {
	Payload *models.IoK8sAPIStorageV1StorageClass
}

func (o *CreateStorageV1StorageClassCreated) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1/storageclasses][%d] createStorageV1StorageClassCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageV1StorageClassCreated) GetPayload() *models.IoK8sAPIStorageV1StorageClass {
	return o.Payload
}

func (o *CreateStorageV1StorageClassCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1StorageClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1StorageClassAccepted creates a CreateStorageV1StorageClassAccepted with default headers values
func NewCreateStorageV1StorageClassAccepted() *CreateStorageV1StorageClassAccepted {
	return &CreateStorageV1StorageClassAccepted{}
}

/*CreateStorageV1StorageClassAccepted handles this case with default header values.

Accepted
*/
type CreateStorageV1StorageClassAccepted struct {
	Payload *models.IoK8sAPIStorageV1StorageClass
}

func (o *CreateStorageV1StorageClassAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1/storageclasses][%d] createStorageV1StorageClassAccepted  %+v", 202, o.Payload)
}

func (o *CreateStorageV1StorageClassAccepted) GetPayload() *models.IoK8sAPIStorageV1StorageClass {
	return o.Payload
}

func (o *CreateStorageV1StorageClassAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1StorageClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1StorageClassUnauthorized creates a CreateStorageV1StorageClassUnauthorized with default headers values
func NewCreateStorageV1StorageClassUnauthorized() *CreateStorageV1StorageClassUnauthorized {
	return &CreateStorageV1StorageClassUnauthorized{}
}

/*CreateStorageV1StorageClassUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateStorageV1StorageClassUnauthorized struct {
}

func (o *CreateStorageV1StorageClassUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1/storageclasses][%d] createStorageV1StorageClassUnauthorized ", 401)
}

func (o *CreateStorageV1StorageClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}