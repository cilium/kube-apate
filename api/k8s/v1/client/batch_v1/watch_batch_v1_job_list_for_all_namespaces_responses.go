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

// WatchBatchV1JobListForAllNamespacesReader is a Reader for the WatchBatchV1JobListForAllNamespaces structure.
type WatchBatchV1JobListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchBatchV1JobListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchBatchV1JobListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchBatchV1JobListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchBatchV1JobListForAllNamespacesOK creates a WatchBatchV1JobListForAllNamespacesOK with default headers values
func NewWatchBatchV1JobListForAllNamespacesOK() *WatchBatchV1JobListForAllNamespacesOK {
	return &WatchBatchV1JobListForAllNamespacesOK{}
}

/*WatchBatchV1JobListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchBatchV1JobListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchBatchV1JobListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1/watch/jobs][%d] watchBatchV1JobListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchBatchV1JobListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchBatchV1JobListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchBatchV1JobListForAllNamespacesUnauthorized creates a WatchBatchV1JobListForAllNamespacesUnauthorized with default headers values
func NewWatchBatchV1JobListForAllNamespacesUnauthorized() *WatchBatchV1JobListForAllNamespacesUnauthorized {
	return &WatchBatchV1JobListForAllNamespacesUnauthorized{}
}

/*WatchBatchV1JobListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchBatchV1JobListForAllNamespacesUnauthorized struct {
}

func (o *WatchBatchV1JobListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1/watch/jobs][%d] watchBatchV1JobListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchBatchV1JobListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
