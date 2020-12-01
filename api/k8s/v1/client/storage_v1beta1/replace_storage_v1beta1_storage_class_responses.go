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

// ReplaceStorageV1beta1StorageClassReader is a Reader for the ReplaceStorageV1beta1StorageClass structure.
type ReplaceStorageV1beta1StorageClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStorageV1beta1StorageClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceStorageV1beta1StorageClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceStorageV1beta1StorageClassCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceStorageV1beta1StorageClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceStorageV1beta1StorageClassOK creates a ReplaceStorageV1beta1StorageClassOK with default headers values
func NewReplaceStorageV1beta1StorageClassOK() *ReplaceStorageV1beta1StorageClassOK {
	return &ReplaceStorageV1beta1StorageClassOK{}
}

/*ReplaceStorageV1beta1StorageClassOK handles this case with default header values.

OK
*/
type ReplaceStorageV1beta1StorageClassOK struct {
	Payload *models.IoK8sAPIStorageV1beta1StorageClass
}

func (o *ReplaceStorageV1beta1StorageClassOK) Error() string {
	return fmt.Sprintf("[PUT /apis/storage.k8s.io/v1beta1/storageclasses/{name}][%d] replaceStorageV1beta1StorageClassOK  %+v", 200, o.Payload)
}

func (o *ReplaceStorageV1beta1StorageClassOK) GetPayload() *models.IoK8sAPIStorageV1beta1StorageClass {
	return o.Payload
}

func (o *ReplaceStorageV1beta1StorageClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1StorageClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageV1beta1StorageClassCreated creates a ReplaceStorageV1beta1StorageClassCreated with default headers values
func NewReplaceStorageV1beta1StorageClassCreated() *ReplaceStorageV1beta1StorageClassCreated {
	return &ReplaceStorageV1beta1StorageClassCreated{}
}

/*ReplaceStorageV1beta1StorageClassCreated handles this case with default header values.

Created
*/
type ReplaceStorageV1beta1StorageClassCreated struct {
	Payload *models.IoK8sAPIStorageV1beta1StorageClass
}

func (o *ReplaceStorageV1beta1StorageClassCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/storage.k8s.io/v1beta1/storageclasses/{name}][%d] replaceStorageV1beta1StorageClassCreated  %+v", 201, o.Payload)
}

func (o *ReplaceStorageV1beta1StorageClassCreated) GetPayload() *models.IoK8sAPIStorageV1beta1StorageClass {
	return o.Payload
}

func (o *ReplaceStorageV1beta1StorageClassCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1StorageClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageV1beta1StorageClassUnauthorized creates a ReplaceStorageV1beta1StorageClassUnauthorized with default headers values
func NewReplaceStorageV1beta1StorageClassUnauthorized() *ReplaceStorageV1beta1StorageClassUnauthorized {
	return &ReplaceStorageV1beta1StorageClassUnauthorized{}
}

/*ReplaceStorageV1beta1StorageClassUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceStorageV1beta1StorageClassUnauthorized struct {
}

func (o *ReplaceStorageV1beta1StorageClassUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/storage.k8s.io/v1beta1/storageclasses/{name}][%d] replaceStorageV1beta1StorageClassUnauthorized ", 401)
}

func (o *ReplaceStorageV1beta1StorageClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
