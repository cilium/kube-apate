// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateStorageV1beta1CSINodeReader is a Reader for the CreateStorageV1beta1CSINode structure.
type CreateStorageV1beta1CSINodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageV1beta1CSINodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStorageV1beta1CSINodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateStorageV1beta1CSINodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateStorageV1beta1CSINodeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateStorageV1beta1CSINodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateStorageV1beta1CSINodeOK creates a CreateStorageV1beta1CSINodeOK with default headers values
func NewCreateStorageV1beta1CSINodeOK() *CreateStorageV1beta1CSINodeOK {
	return &CreateStorageV1beta1CSINodeOK{}
}

/*CreateStorageV1beta1CSINodeOK handles this case with default header values.

OK
*/
type CreateStorageV1beta1CSINodeOK struct {
	Payload *models.IoK8sAPIStorageV1beta1CSINode
}

func (o *CreateStorageV1beta1CSINodeOK) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/csinodes][%d] createStorageV1beta1CSINodeOK  %+v", 200, o.Payload)
}

func (o *CreateStorageV1beta1CSINodeOK) GetPayload() *models.IoK8sAPIStorageV1beta1CSINode {
	return o.Payload
}

func (o *CreateStorageV1beta1CSINodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1CSINode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1beta1CSINodeCreated creates a CreateStorageV1beta1CSINodeCreated with default headers values
func NewCreateStorageV1beta1CSINodeCreated() *CreateStorageV1beta1CSINodeCreated {
	return &CreateStorageV1beta1CSINodeCreated{}
}

/*CreateStorageV1beta1CSINodeCreated handles this case with default header values.

Created
*/
type CreateStorageV1beta1CSINodeCreated struct {
	Payload *models.IoK8sAPIStorageV1beta1CSINode
}

func (o *CreateStorageV1beta1CSINodeCreated) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/csinodes][%d] createStorageV1beta1CSINodeCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageV1beta1CSINodeCreated) GetPayload() *models.IoK8sAPIStorageV1beta1CSINode {
	return o.Payload
}

func (o *CreateStorageV1beta1CSINodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1CSINode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1beta1CSINodeAccepted creates a CreateStorageV1beta1CSINodeAccepted with default headers values
func NewCreateStorageV1beta1CSINodeAccepted() *CreateStorageV1beta1CSINodeAccepted {
	return &CreateStorageV1beta1CSINodeAccepted{}
}

/*CreateStorageV1beta1CSINodeAccepted handles this case with default header values.

Accepted
*/
type CreateStorageV1beta1CSINodeAccepted struct {
	Payload *models.IoK8sAPIStorageV1beta1CSINode
}

func (o *CreateStorageV1beta1CSINodeAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/csinodes][%d] createStorageV1beta1CSINodeAccepted  %+v", 202, o.Payload)
}

func (o *CreateStorageV1beta1CSINodeAccepted) GetPayload() *models.IoK8sAPIStorageV1beta1CSINode {
	return o.Payload
}

func (o *CreateStorageV1beta1CSINodeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1CSINode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1beta1CSINodeUnauthorized creates a CreateStorageV1beta1CSINodeUnauthorized with default headers values
func NewCreateStorageV1beta1CSINodeUnauthorized() *CreateStorageV1beta1CSINodeUnauthorized {
	return &CreateStorageV1beta1CSINodeUnauthorized{}
}

/*CreateStorageV1beta1CSINodeUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateStorageV1beta1CSINodeUnauthorized struct {
}

func (o *CreateStorageV1beta1CSINodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/csinodes][%d] createStorageV1beta1CSINodeUnauthorized ", 401)
}

func (o *CreateStorageV1beta1CSINodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}