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

// ReplaceCoreV1PersistentVolumeStatusReader is a Reader for the ReplaceCoreV1PersistentVolumeStatus structure.
type ReplaceCoreV1PersistentVolumeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoreV1PersistentVolumeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoreV1PersistentVolumeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoreV1PersistentVolumeStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoreV1PersistentVolumeStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoreV1PersistentVolumeStatusOK creates a ReplaceCoreV1PersistentVolumeStatusOK with default headers values
func NewReplaceCoreV1PersistentVolumeStatusOK() *ReplaceCoreV1PersistentVolumeStatusOK {
	return &ReplaceCoreV1PersistentVolumeStatusOK{}
}

/*ReplaceCoreV1PersistentVolumeStatusOK handles this case with default header values.

OK
*/
type ReplaceCoreV1PersistentVolumeStatusOK struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *ReplaceCoreV1PersistentVolumeStatusOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/persistentvolumes/{name}/status][%d] replaceCoreV1PersistentVolumeStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoreV1PersistentVolumeStatusOK) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *ReplaceCoreV1PersistentVolumeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1PersistentVolumeStatusCreated creates a ReplaceCoreV1PersistentVolumeStatusCreated with default headers values
func NewReplaceCoreV1PersistentVolumeStatusCreated() *ReplaceCoreV1PersistentVolumeStatusCreated {
	return &ReplaceCoreV1PersistentVolumeStatusCreated{}
}

/*ReplaceCoreV1PersistentVolumeStatusCreated handles this case with default header values.

Created
*/
type ReplaceCoreV1PersistentVolumeStatusCreated struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *ReplaceCoreV1PersistentVolumeStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/persistentvolumes/{name}/status][%d] replaceCoreV1PersistentVolumeStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoreV1PersistentVolumeStatusCreated) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *ReplaceCoreV1PersistentVolumeStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1PersistentVolumeStatusUnauthorized creates a ReplaceCoreV1PersistentVolumeStatusUnauthorized with default headers values
func NewReplaceCoreV1PersistentVolumeStatusUnauthorized() *ReplaceCoreV1PersistentVolumeStatusUnauthorized {
	return &ReplaceCoreV1PersistentVolumeStatusUnauthorized{}
}

/*ReplaceCoreV1PersistentVolumeStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoreV1PersistentVolumeStatusUnauthorized struct {
}

func (o *ReplaceCoreV1PersistentVolumeStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/persistentvolumes/{name}/status][%d] replaceCoreV1PersistentVolumeStatusUnauthorized ", 401)
}

func (o *ReplaceCoreV1PersistentVolumeStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
