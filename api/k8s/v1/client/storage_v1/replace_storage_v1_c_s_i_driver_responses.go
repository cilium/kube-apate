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

// ReplaceStorageV1CSIDriverReader is a Reader for the ReplaceStorageV1CSIDriver structure.
type ReplaceStorageV1CSIDriverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStorageV1CSIDriverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceStorageV1CSIDriverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceStorageV1CSIDriverCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceStorageV1CSIDriverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceStorageV1CSIDriverOK creates a ReplaceStorageV1CSIDriverOK with default headers values
func NewReplaceStorageV1CSIDriverOK() *ReplaceStorageV1CSIDriverOK {
	return &ReplaceStorageV1CSIDriverOK{}
}

/*ReplaceStorageV1CSIDriverOK handles this case with default header values.

OK
*/
type ReplaceStorageV1CSIDriverOK struct {
	Payload *models.IoK8sAPIStorageV1CSIDriver
}

func (o *ReplaceStorageV1CSIDriverOK) Error() string {
	return fmt.Sprintf("[PUT /apis/storage.k8s.io/v1/csidrivers/{name}][%d] replaceStorageV1CSIDriverOK  %+v", 200, o.Payload)
}

func (o *ReplaceStorageV1CSIDriverOK) GetPayload() *models.IoK8sAPIStorageV1CSIDriver {
	return o.Payload
}

func (o *ReplaceStorageV1CSIDriverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1CSIDriver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageV1CSIDriverCreated creates a ReplaceStorageV1CSIDriverCreated with default headers values
func NewReplaceStorageV1CSIDriverCreated() *ReplaceStorageV1CSIDriverCreated {
	return &ReplaceStorageV1CSIDriverCreated{}
}

/*ReplaceStorageV1CSIDriverCreated handles this case with default header values.

Created
*/
type ReplaceStorageV1CSIDriverCreated struct {
	Payload *models.IoK8sAPIStorageV1CSIDriver
}

func (o *ReplaceStorageV1CSIDriverCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/storage.k8s.io/v1/csidrivers/{name}][%d] replaceStorageV1CSIDriverCreated  %+v", 201, o.Payload)
}

func (o *ReplaceStorageV1CSIDriverCreated) GetPayload() *models.IoK8sAPIStorageV1CSIDriver {
	return o.Payload
}

func (o *ReplaceStorageV1CSIDriverCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1CSIDriver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageV1CSIDriverUnauthorized creates a ReplaceStorageV1CSIDriverUnauthorized with default headers values
func NewReplaceStorageV1CSIDriverUnauthorized() *ReplaceStorageV1CSIDriverUnauthorized {
	return &ReplaceStorageV1CSIDriverUnauthorized{}
}

/*ReplaceStorageV1CSIDriverUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceStorageV1CSIDriverUnauthorized struct {
}

func (o *ReplaceStorageV1CSIDriverUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/storage.k8s.io/v1/csidrivers/{name}][%d] replaceStorageV1CSIDriverUnauthorized ", 401)
}

func (o *ReplaceStorageV1CSIDriverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}