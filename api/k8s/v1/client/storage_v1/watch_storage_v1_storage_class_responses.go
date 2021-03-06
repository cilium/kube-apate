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

// WatchStorageV1StorageClassReader is a Reader for the WatchStorageV1StorageClass structure.
type WatchStorageV1StorageClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchStorageV1StorageClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchStorageV1StorageClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchStorageV1StorageClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchStorageV1StorageClassOK creates a WatchStorageV1StorageClassOK with default headers values
func NewWatchStorageV1StorageClassOK() *WatchStorageV1StorageClassOK {
	return &WatchStorageV1StorageClassOK{}
}

/*WatchStorageV1StorageClassOK handles this case with default header values.

OK
*/
type WatchStorageV1StorageClassOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchStorageV1StorageClassOK) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/watch/storageclasses/{name}][%d] watchStorageV1StorageClassOK  %+v", 200, o.Payload)
}

func (o *WatchStorageV1StorageClassOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchStorageV1StorageClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchStorageV1StorageClassUnauthorized creates a WatchStorageV1StorageClassUnauthorized with default headers values
func NewWatchStorageV1StorageClassUnauthorized() *WatchStorageV1StorageClassUnauthorized {
	return &WatchStorageV1StorageClassUnauthorized{}
}

/*WatchStorageV1StorageClassUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchStorageV1StorageClassUnauthorized struct {
}

func (o *WatchStorageV1StorageClassUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/watch/storageclasses/{name}][%d] watchStorageV1StorageClassUnauthorized ", 401)
}

func (o *WatchStorageV1StorageClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
