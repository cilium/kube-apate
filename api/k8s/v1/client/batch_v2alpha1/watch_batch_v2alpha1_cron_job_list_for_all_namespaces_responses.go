// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchBatchV2alpha1CronJobListForAllNamespacesReader is a Reader for the WatchBatchV2alpha1CronJobListForAllNamespaces structure.
type WatchBatchV2alpha1CronJobListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchBatchV2alpha1CronJobListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchBatchV2alpha1CronJobListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchBatchV2alpha1CronJobListForAllNamespacesOK creates a WatchBatchV2alpha1CronJobListForAllNamespacesOK with default headers values
func NewWatchBatchV2alpha1CronJobListForAllNamespacesOK() *WatchBatchV2alpha1CronJobListForAllNamespacesOK {
	return &WatchBatchV2alpha1CronJobListForAllNamespacesOK{}
}

/*WatchBatchV2alpha1CronJobListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchBatchV2alpha1CronJobListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchBatchV2alpha1CronJobListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v2alpha1/watch/cronjobs][%d] watchBatchV2alpha1CronJobListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchBatchV2alpha1CronJobListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchBatchV2alpha1CronJobListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized creates a WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized with default headers values
func NewWatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized() *WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized {
	return &WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized{}
}

/*WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized struct {
}

func (o *WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v2alpha1/watch/cronjobs][%d] watchBatchV2alpha1CronJobListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchBatchV2alpha1CronJobListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
