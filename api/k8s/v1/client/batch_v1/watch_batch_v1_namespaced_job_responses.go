// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchBatchV1NamespacedJobReader is a Reader for the WatchBatchV1NamespacedJob structure.
type WatchBatchV1NamespacedJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchBatchV1NamespacedJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchBatchV1NamespacedJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchBatchV1NamespacedJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchBatchV1NamespacedJobOK creates a WatchBatchV1NamespacedJobOK with default headers values
func NewWatchBatchV1NamespacedJobOK() *WatchBatchV1NamespacedJobOK {
	return &WatchBatchV1NamespacedJobOK{}
}

/*WatchBatchV1NamespacedJobOK handles this case with default header values.

OK
*/
type WatchBatchV1NamespacedJobOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchBatchV1NamespacedJobOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1/watch/namespaces/{namespace}/jobs/{name}][%d] watchBatchV1NamespacedJobOK  %+v", 200, o.Payload)
}

func (o *WatchBatchV1NamespacedJobOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchBatchV1NamespacedJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchBatchV1NamespacedJobUnauthorized creates a WatchBatchV1NamespacedJobUnauthorized with default headers values
func NewWatchBatchV1NamespacedJobUnauthorized() *WatchBatchV1NamespacedJobUnauthorized {
	return &WatchBatchV1NamespacedJobUnauthorized{}
}

/*WatchBatchV1NamespacedJobUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchBatchV1NamespacedJobUnauthorized struct {
}

func (o *WatchBatchV1NamespacedJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1/watch/namespaces/{namespace}/jobs/{name}][%d] watchBatchV1NamespacedJobUnauthorized ", 401)
}

func (o *WatchBatchV1NamespacedJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}