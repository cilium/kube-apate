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

// ReadCoreV1PersistentVolumeReader is a Reader for the ReadCoreV1PersistentVolume structure.
type ReadCoreV1PersistentVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1PersistentVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1PersistentVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1PersistentVolumeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1PersistentVolumeOK creates a ReadCoreV1PersistentVolumeOK with default headers values
func NewReadCoreV1PersistentVolumeOK() *ReadCoreV1PersistentVolumeOK {
	return &ReadCoreV1PersistentVolumeOK{}
}

/*ReadCoreV1PersistentVolumeOK handles this case with default header values.

OK
*/
type ReadCoreV1PersistentVolumeOK struct {
	Payload *models.IoK8sAPICoreV1PersistentVolume
}

func (o *ReadCoreV1PersistentVolumeOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/persistentvolumes/{name}][%d] readCoreV1PersistentVolumeOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1PersistentVolumeOK) GetPayload() *models.IoK8sAPICoreV1PersistentVolume {
	return o.Payload
}

func (o *ReadCoreV1PersistentVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1PersistentVolumeUnauthorized creates a ReadCoreV1PersistentVolumeUnauthorized with default headers values
func NewReadCoreV1PersistentVolumeUnauthorized() *ReadCoreV1PersistentVolumeUnauthorized {
	return &ReadCoreV1PersistentVolumeUnauthorized{}
}

/*ReadCoreV1PersistentVolumeUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1PersistentVolumeUnauthorized struct {
}

func (o *ReadCoreV1PersistentVolumeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/persistentvolumes/{name}][%d] readCoreV1PersistentVolumeUnauthorized ", 401)
}

func (o *ReadCoreV1PersistentVolumeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
