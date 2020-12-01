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

// ReadStorageV1VolumeAttachmentReader is a Reader for the ReadStorageV1VolumeAttachment structure.
type ReadStorageV1VolumeAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadStorageV1VolumeAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadStorageV1VolumeAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadStorageV1VolumeAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadStorageV1VolumeAttachmentOK creates a ReadStorageV1VolumeAttachmentOK with default headers values
func NewReadStorageV1VolumeAttachmentOK() *ReadStorageV1VolumeAttachmentOK {
	return &ReadStorageV1VolumeAttachmentOK{}
}

/*ReadStorageV1VolumeAttachmentOK handles this case with default header values.

OK
*/
type ReadStorageV1VolumeAttachmentOK struct {
	Payload *models.IoK8sAPIStorageV1VolumeAttachment
}

func (o *ReadStorageV1VolumeAttachmentOK) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/volumeattachments/{name}][%d] readStorageV1VolumeAttachmentOK  %+v", 200, o.Payload)
}

func (o *ReadStorageV1VolumeAttachmentOK) GetPayload() *models.IoK8sAPIStorageV1VolumeAttachment {
	return o.Payload
}

func (o *ReadStorageV1VolumeAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadStorageV1VolumeAttachmentUnauthorized creates a ReadStorageV1VolumeAttachmentUnauthorized with default headers values
func NewReadStorageV1VolumeAttachmentUnauthorized() *ReadStorageV1VolumeAttachmentUnauthorized {
	return &ReadStorageV1VolumeAttachmentUnauthorized{}
}

/*ReadStorageV1VolumeAttachmentUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadStorageV1VolumeAttachmentUnauthorized struct {
}

func (o *ReadStorageV1VolumeAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/volumeattachments/{name}][%d] readStorageV1VolumeAttachmentUnauthorized ", 401)
}

func (o *ReadStorageV1VolumeAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
