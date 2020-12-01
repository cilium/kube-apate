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

// ReplaceCoreV1PersistentVolumeReader is a Reader for the ReplaceCoreV1PersistentVolume structure.
type ReplaceCoreV1PersistentVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoreV1PersistentVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoreV1PersistentVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoreV1PersistentVolumeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoreV1PersistentVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoreV1PersistentVolumeOK creates a ReplaceCoreV1PersistentVolumeOK with default headers values
func NewReplaceCoreV1PersistentVolumeOK() *ReplaceCoreV1PersistentVolumeOK {
	return &ReplaceCoreV1PersistentVolumeOK{}
}

/*ReplaceCoreV1PersistentVolumeOK handles this case with default header values.

OK
*/
type ReplaceCoreV1PersistentVolumeOK struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *ReplaceCoreV1PersistentVolumeOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/persistentvolumes/{name}][%d] replaceCoreV1PersistentVolumeOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoreV1PersistentVolumeOK) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *ReplaceCoreV1PersistentVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1PersistentVolumeCreated creates a ReplaceCoreV1PersistentVolumeCreated with default headers values
func NewReplaceCoreV1PersistentVolumeCreated() *ReplaceCoreV1PersistentVolumeCreated {
	return &ReplaceCoreV1PersistentVolumeCreated{}
}

/*ReplaceCoreV1PersistentVolumeCreated handles this case with default header values.

Created
*/
type ReplaceCoreV1PersistentVolumeCreated struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *ReplaceCoreV1PersistentVolumeCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/persistentvolumes/{name}][%d] replaceCoreV1PersistentVolumeCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoreV1PersistentVolumeCreated) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *ReplaceCoreV1PersistentVolumeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1PersistentVolumeUnauthorized creates a ReplaceCoreV1PersistentVolumeUnauthorized with default headers values
func NewReplaceCoreV1PersistentVolumeUnauthorized() *ReplaceCoreV1PersistentVolumeUnauthorized {
	return &ReplaceCoreV1PersistentVolumeUnauthorized{}
}

/*ReplaceCoreV1PersistentVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoreV1PersistentVolumeUnauthorized struct {
}

func (o *ReplaceCoreV1PersistentVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/persistentvolumes/{name}][%d] replaceCoreV1PersistentVolumeUnauthorized ", 401)
}

func (o *ReplaceCoreV1PersistentVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
