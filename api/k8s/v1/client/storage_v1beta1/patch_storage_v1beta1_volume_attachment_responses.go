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

// PatchStorageV1beta1VolumeAttachmentReader is a Reader for the PatchStorageV1beta1VolumeAttachment structure.
type PatchStorageV1beta1VolumeAttachmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageV1beta1VolumeAttachmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchStorageV1beta1VolumeAttachmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchStorageV1beta1VolumeAttachmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchStorageV1beta1VolumeAttachmentOK creates a PatchStorageV1beta1VolumeAttachmentOK with default headers values
func NewPatchStorageV1beta1VolumeAttachmentOK() *PatchStorageV1beta1VolumeAttachmentOK {
	return &PatchStorageV1beta1VolumeAttachmentOK{}
}

/*PatchStorageV1beta1VolumeAttachmentOK handles this case with default header values.

OK
*/
type PatchStorageV1beta1VolumeAttachmentOK struct {
	Payload *models.IoK8sAPIStorageV1beta1VolumeAttachment
}

func (o *PatchStorageV1beta1VolumeAttachmentOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/storage.k8s.io/v1beta1/volumeattachments/{name}][%d] patchStorageV1beta1VolumeAttachmentOK  %+v", 200, o.Payload)
}

func (o *PatchStorageV1beta1VolumeAttachmentOK) GetPayload() *models.IoK8sAPIStorageV1beta1VolumeAttachment {
	return o.Payload
}

func (o *PatchStorageV1beta1VolumeAttachmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1VolumeAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStorageV1beta1VolumeAttachmentUnauthorized creates a PatchStorageV1beta1VolumeAttachmentUnauthorized with default headers values
func NewPatchStorageV1beta1VolumeAttachmentUnauthorized() *PatchStorageV1beta1VolumeAttachmentUnauthorized {
	return &PatchStorageV1beta1VolumeAttachmentUnauthorized{}
}

/*PatchStorageV1beta1VolumeAttachmentUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchStorageV1beta1VolumeAttachmentUnauthorized struct {
}

func (o *PatchStorageV1beta1VolumeAttachmentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/storage.k8s.io/v1beta1/volumeattachments/{name}][%d] patchStorageV1beta1VolumeAttachmentUnauthorized ", 401)
}

func (o *PatchStorageV1beta1VolumeAttachmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
