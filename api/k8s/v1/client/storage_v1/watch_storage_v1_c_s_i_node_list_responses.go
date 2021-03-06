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

// WatchStorageV1CSINodeListReader is a Reader for the WatchStorageV1CSINodeList structure.
type WatchStorageV1CSINodeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchStorageV1CSINodeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchStorageV1CSINodeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchStorageV1CSINodeListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchStorageV1CSINodeListOK creates a WatchStorageV1CSINodeListOK with default headers values
func NewWatchStorageV1CSINodeListOK() *WatchStorageV1CSINodeListOK {
	return &WatchStorageV1CSINodeListOK{}
}

/*WatchStorageV1CSINodeListOK handles this case with default header values.

OK
*/
type WatchStorageV1CSINodeListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchStorageV1CSINodeListOK) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/watch/csinodes][%d] watchStorageV1CSINodeListOK  %+v", 200, o.Payload)
}

func (o *WatchStorageV1CSINodeListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchStorageV1CSINodeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchStorageV1CSINodeListUnauthorized creates a WatchStorageV1CSINodeListUnauthorized with default headers values
func NewWatchStorageV1CSINodeListUnauthorized() *WatchStorageV1CSINodeListUnauthorized {
	return &WatchStorageV1CSINodeListUnauthorized{}
}

/*WatchStorageV1CSINodeListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchStorageV1CSINodeListUnauthorized struct {
}

func (o *WatchStorageV1CSINodeListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/watch/csinodes][%d] watchStorageV1CSINodeListUnauthorized ", 401)
}

func (o *WatchStorageV1CSINodeListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
