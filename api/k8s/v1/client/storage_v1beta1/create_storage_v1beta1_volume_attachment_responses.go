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

// CreateStorageV1beta1VolumeAttachmentReader is a Reader for the CreateStorageV1beta1VolumeAttachment structure.
type CreateStorageV1beta1VolumeAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageV1beta1VolumeAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateStorageV1beta1VolumeAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateStorageV1beta1VolumeAttachmentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateStorageV1beta1VolumeAttachmentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateStorageV1beta1VolumeAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateStorageV1beta1VolumeAttachmentOK creates a CreateStorageV1beta1VolumeAttachmentOK with default headers values
func NewCreateStorageV1beta1VolumeAttachmentOK() *CreateStorageV1beta1VolumeAttachmentOK {
	return &CreateStorageV1beta1VolumeAttachmentOK{}
}

/*CreateStorageV1beta1VolumeAttachmentOK handles this case with default header values.

OK
*/
type CreateStorageV1beta1VolumeAttachmentOK struct {
	Payload *models.IoK8sAPIStorageV1beta1VolumeAttachment
}

func (o *CreateStorageV1beta1VolumeAttachmentOK) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/volumeattachments][%d] createStorageV1beta1VolumeAttachmentOK  %+v", 200, o.Payload)
}

func (o *CreateStorageV1beta1VolumeAttachmentOK) GetPayload() *models.IoK8sAPIStorageV1beta1VolumeAttachment {
	return o.Payload
}

func (o *CreateStorageV1beta1VolumeAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1beta1VolumeAttachmentCreated creates a CreateStorageV1beta1VolumeAttachmentCreated with default headers values
func NewCreateStorageV1beta1VolumeAttachmentCreated() *CreateStorageV1beta1VolumeAttachmentCreated {
	return &CreateStorageV1beta1VolumeAttachmentCreated{}
}

/*CreateStorageV1beta1VolumeAttachmentCreated handles this case with default header values.

Created
*/
type CreateStorageV1beta1VolumeAttachmentCreated struct {
	Payload *models.IoK8sAPIStorageV1beta1VolumeAttachment
}

func (o *CreateStorageV1beta1VolumeAttachmentCreated) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/volumeattachments][%d] createStorageV1beta1VolumeAttachmentCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageV1beta1VolumeAttachmentCreated) GetPayload() *models.IoK8sAPIStorageV1beta1VolumeAttachment {
	return o.Payload
}

func (o *CreateStorageV1beta1VolumeAttachmentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1beta1VolumeAttachmentAccepted creates a CreateStorageV1beta1VolumeAttachmentAccepted with default headers values
func NewCreateStorageV1beta1VolumeAttachmentAccepted() *CreateStorageV1beta1VolumeAttachmentAccepted {
	return &CreateStorageV1beta1VolumeAttachmentAccepted{}
}

/*CreateStorageV1beta1VolumeAttachmentAccepted handles this case with default header values.

Accepted
*/
type CreateStorageV1beta1VolumeAttachmentAccepted struct {
	Payload *models.IoK8sAPIStorageV1beta1VolumeAttachment
}

func (o *CreateStorageV1beta1VolumeAttachmentAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/volumeattachments][%d] createStorageV1beta1VolumeAttachmentAccepted  %+v", 202, o.Payload)
}

func (o *CreateStorageV1beta1VolumeAttachmentAccepted) GetPayload() *models.IoK8sAPIStorageV1beta1VolumeAttachment {
	return o.Payload
}

func (o *CreateStorageV1beta1VolumeAttachmentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageV1beta1VolumeAttachmentUnauthorized creates a CreateStorageV1beta1VolumeAttachmentUnauthorized with default headers values
func NewCreateStorageV1beta1VolumeAttachmentUnauthorized() *CreateStorageV1beta1VolumeAttachmentUnauthorized {
	return &CreateStorageV1beta1VolumeAttachmentUnauthorized{}
}

/*CreateStorageV1beta1VolumeAttachmentUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateStorageV1beta1VolumeAttachmentUnauthorized struct {
}

func (o *CreateStorageV1beta1VolumeAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/storage.k8s.io/v1beta1/volumeattachments][%d] createStorageV1beta1VolumeAttachmentUnauthorized ", 401)
}

func (o *CreateStorageV1beta1VolumeAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
