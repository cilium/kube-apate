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

// PatchCoreV1PersistentVolumeReader is a Reader for the PatchCoreV1PersistentVolume structure.
type PatchCoreV1PersistentVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1PersistentVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1PersistentVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1PersistentVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1PersistentVolumeOK creates a PatchCoreV1PersistentVolumeOK with default headers values
func NewPatchCoreV1PersistentVolumeOK() *PatchCoreV1PersistentVolumeOK {
	return &PatchCoreV1PersistentVolumeOK{}
}

/*PatchCoreV1PersistentVolumeOK handles this case with default header values.

OK
*/
type PatchCoreV1PersistentVolumeOK struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *PatchCoreV1PersistentVolumeOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/persistentvolumes/{name}][%d] patchCoreV1PersistentVolumeOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1PersistentVolumeOK) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *PatchCoreV1PersistentVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1PersistentVolumeUnauthorized creates a PatchCoreV1PersistentVolumeUnauthorized with default headers values
func NewPatchCoreV1PersistentVolumeUnauthorized() *PatchCoreV1PersistentVolumeUnauthorized {
	return &PatchCoreV1PersistentVolumeUnauthorized{}
}

/*PatchCoreV1PersistentVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1PersistentVolumeUnauthorized struct {
}

func (o *PatchCoreV1PersistentVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/persistentvolumes/{name}][%d] patchCoreV1PersistentVolumeUnauthorized ", 401)
}

func (o *PatchCoreV1PersistentVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}