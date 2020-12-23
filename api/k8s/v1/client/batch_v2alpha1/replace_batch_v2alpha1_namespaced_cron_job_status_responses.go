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

// ReplaceBatchV2alpha1NamespacedCronJobStatusReader is a Reader for the ReplaceBatchV2alpha1NamespacedCronJobStatus structure.
type ReplaceBatchV2alpha1NamespacedCronJobStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceBatchV2alpha1NamespacedCronJobStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceBatchV2alpha1NamespacedCronJobStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceBatchV2alpha1NamespacedCronJobStatusOK creates a ReplaceBatchV2alpha1NamespacedCronJobStatusOK with default headers values
func NewReplaceBatchV2alpha1NamespacedCronJobStatusOK() *ReplaceBatchV2alpha1NamespacedCronJobStatusOK {
	return &ReplaceBatchV2alpha1NamespacedCronJobStatusOK{}
}

/*ReplaceBatchV2alpha1NamespacedCronJobStatusOK handles this case with default header values.

OK
*/
type ReplaceBatchV2alpha1NamespacedCronJobStatusOK struct {
	Payload *models.IoK8sAPIBatchV2alpha1CronJob
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusOK) Error() string {
	return fmt.Sprintf("[PUT /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status][%d] replaceBatchV2alpha1NamespacedCronJobStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusOK) GetPayload() *models.IoK8sAPIBatchV2alpha1CronJob {
	return o.Payload
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV2alpha1CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBatchV2alpha1NamespacedCronJobStatusCreated creates a ReplaceBatchV2alpha1NamespacedCronJobStatusCreated with default headers values
func NewReplaceBatchV2alpha1NamespacedCronJobStatusCreated() *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated {
	return &ReplaceBatchV2alpha1NamespacedCronJobStatusCreated{}
}

/*ReplaceBatchV2alpha1NamespacedCronJobStatusCreated handles this case with default header values.

Created
*/
type ReplaceBatchV2alpha1NamespacedCronJobStatusCreated struct {
	Payload *models.IoK8sAPIBatchV2alpha1CronJob
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status][%d] replaceBatchV2alpha1NamespacedCronJobStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated) GetPayload() *models.IoK8sAPIBatchV2alpha1CronJob {
	return o.Payload
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV2alpha1CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized creates a ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized with default headers values
func NewReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized() *ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized {
	return &ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized{}
}

/*ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized struct {
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}/status][%d] replaceBatchV2alpha1NamespacedCronJobStatusUnauthorized ", 401)
}

func (o *ReplaceBatchV2alpha1NamespacedCronJobStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}