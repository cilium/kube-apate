// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceBatchV1beta1NamespacedCronJobReader is a Reader for the ReplaceBatchV1beta1NamespacedCronJob structure.
type ReplaceBatchV1beta1NamespacedCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceBatchV1beta1NamespacedCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceBatchV1beta1NamespacedCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceBatchV1beta1NamespacedCronJobCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceBatchV1beta1NamespacedCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceBatchV1beta1NamespacedCronJobOK creates a ReplaceBatchV1beta1NamespacedCronJobOK with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobOK() *ReplaceBatchV1beta1NamespacedCronJobOK {
	return &ReplaceBatchV1beta1NamespacedCronJobOK{}
}

/*ReplaceBatchV1beta1NamespacedCronJobOK handles this case with default header values.

OK
*/
type ReplaceBatchV1beta1NamespacedCronJobOK struct {
	Payload *models.IoK8sAPIBatchV1beta1CronJob
}

func (o *ReplaceBatchV1beta1NamespacedCronJobOK) Error() string {
	return fmt.Sprintf("[PUT /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}][%d] replaceBatchV1beta1NamespacedCronJobOK  %+v", 200, o.Payload)
}

func (o *ReplaceBatchV1beta1NamespacedCronJobOK) GetPayload() *models.IoK8sAPIBatchV1beta1CronJob {
	return o.Payload
}

func (o *ReplaceBatchV1beta1NamespacedCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV1beta1CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBatchV1beta1NamespacedCronJobCreated creates a ReplaceBatchV1beta1NamespacedCronJobCreated with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobCreated() *ReplaceBatchV1beta1NamespacedCronJobCreated {
	return &ReplaceBatchV1beta1NamespacedCronJobCreated{}
}

/*ReplaceBatchV1beta1NamespacedCronJobCreated handles this case with default header values.

Created
*/
type ReplaceBatchV1beta1NamespacedCronJobCreated struct {
	Payload *models.IoK8sAPIBatchV1beta1CronJob
}

func (o *ReplaceBatchV1beta1NamespacedCronJobCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}][%d] replaceBatchV1beta1NamespacedCronJobCreated  %+v", 201, o.Payload)
}

func (o *ReplaceBatchV1beta1NamespacedCronJobCreated) GetPayload() *models.IoK8sAPIBatchV1beta1CronJob {
	return o.Payload
}

func (o *ReplaceBatchV1beta1NamespacedCronJobCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV1beta1CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceBatchV1beta1NamespacedCronJobUnauthorized creates a ReplaceBatchV1beta1NamespacedCronJobUnauthorized with default headers values
func NewReplaceBatchV1beta1NamespacedCronJobUnauthorized() *ReplaceBatchV1beta1NamespacedCronJobUnauthorized {
	return &ReplaceBatchV1beta1NamespacedCronJobUnauthorized{}
}

/*ReplaceBatchV1beta1NamespacedCronJobUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceBatchV1beta1NamespacedCronJobUnauthorized struct {
}

func (o *ReplaceBatchV1beta1NamespacedCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}][%d] replaceBatchV1beta1NamespacedCronJobUnauthorized ", 401)
}

func (o *ReplaceBatchV1beta1NamespacedCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
