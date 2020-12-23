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

// PatchCoreV1PersistentVolumeStatusReader is a Reader for the PatchCoreV1PersistentVolumeStatus structure.
type PatchCoreV1PersistentVolumeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1PersistentVolumeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1PersistentVolumeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1PersistentVolumeStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1PersistentVolumeStatusOK creates a PatchCoreV1PersistentVolumeStatusOK with default headers values
func NewPatchCoreV1PersistentVolumeStatusOK() *PatchCoreV1PersistentVolumeStatusOK {
	return &PatchCoreV1PersistentVolumeStatusOK{}
}

/*PatchCoreV1PersistentVolumeStatusOK handles this case with default header values.

OK
*/
type PatchCoreV1PersistentVolumeStatusOK struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *PatchCoreV1PersistentVolumeStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/persistentvolumes/{name}/status][%d] patchCoreV1PersistentVolumeStatusOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1PersistentVolumeStatusOK) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *PatchCoreV1PersistentVolumeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1PersistentVolumeStatusUnauthorized creates a PatchCoreV1PersistentVolumeStatusUnauthorized with default headers values
func NewPatchCoreV1PersistentVolumeStatusUnauthorized() *PatchCoreV1PersistentVolumeStatusUnauthorized {
	return &PatchCoreV1PersistentVolumeStatusUnauthorized{}
}

/*PatchCoreV1PersistentVolumeStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1PersistentVolumeStatusUnauthorized struct {
}

func (o *PatchCoreV1PersistentVolumeStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/persistentvolumes/{name}/status][%d] patchCoreV1PersistentVolumeStatusUnauthorized ", 401)
}

func (o *PatchCoreV1PersistentVolumeStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}