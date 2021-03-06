// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteStorageV1alpha1VolumeAttachmentReader is a Reader for the DeleteStorageV1alpha1VolumeAttachment structure.
type DeleteStorageV1alpha1VolumeAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageV1alpha1VolumeAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteStorageV1alpha1VolumeAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteStorageV1alpha1VolumeAttachmentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteStorageV1alpha1VolumeAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteStorageV1alpha1VolumeAttachmentOK creates a DeleteStorageV1alpha1VolumeAttachmentOK with default headers values
func NewDeleteStorageV1alpha1VolumeAttachmentOK() *DeleteStorageV1alpha1VolumeAttachmentOK {
	return &DeleteStorageV1alpha1VolumeAttachmentOK{}
}

/*DeleteStorageV1alpha1VolumeAttachmentOK handles this case with default header values.

OK
*/
type DeleteStorageV1alpha1VolumeAttachmentOK struct {
	Payload *models.IoK8sAPIStorageV1alpha1VolumeAttachment
}

func (o *DeleteStorageV1alpha1VolumeAttachmentOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1alpha1/volumeattachments/{name}][%d] deleteStorageV1alpha1VolumeAttachmentOK  %+v", 200, o.Payload)
}

func (o *DeleteStorageV1alpha1VolumeAttachmentOK) GetPayload() *models.IoK8sAPIStorageV1alpha1VolumeAttachment {
	return o.Payload
}

func (o *DeleteStorageV1alpha1VolumeAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1alpha1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStorageV1alpha1VolumeAttachmentAccepted creates a DeleteStorageV1alpha1VolumeAttachmentAccepted with default headers values
func NewDeleteStorageV1alpha1VolumeAttachmentAccepted() *DeleteStorageV1alpha1VolumeAttachmentAccepted {
	return &DeleteStorageV1alpha1VolumeAttachmentAccepted{}
}

/*DeleteStorageV1alpha1VolumeAttachmentAccepted handles this case with default header values.

Accepted
*/
type DeleteStorageV1alpha1VolumeAttachmentAccepted struct {
	Payload *models.IoK8sAPIStorageV1alpha1VolumeAttachment
}

func (o *DeleteStorageV1alpha1VolumeAttachmentAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1alpha1/volumeattachments/{name}][%d] deleteStorageV1alpha1VolumeAttachmentAccepted  %+v", 202, o.Payload)
}

func (o *DeleteStorageV1alpha1VolumeAttachmentAccepted) GetPayload() *models.IoK8sAPIStorageV1alpha1VolumeAttachment {
	return o.Payload
}

func (o *DeleteStorageV1alpha1VolumeAttachmentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1alpha1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStorageV1alpha1VolumeAttachmentUnauthorized creates a DeleteStorageV1alpha1VolumeAttachmentUnauthorized with default headers values
func NewDeleteStorageV1alpha1VolumeAttachmentUnauthorized() *DeleteStorageV1alpha1VolumeAttachmentUnauthorized {
	return &DeleteStorageV1alpha1VolumeAttachmentUnauthorized{}
}

/*DeleteStorageV1alpha1VolumeAttachmentUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteStorageV1alpha1VolumeAttachmentUnauthorized struct {
}

func (o *DeleteStorageV1alpha1VolumeAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1alpha1/volumeattachments/{name}][%d] deleteStorageV1alpha1VolumeAttachmentUnauthorized ", 401)
}

func (o *DeleteStorageV1alpha1VolumeAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
