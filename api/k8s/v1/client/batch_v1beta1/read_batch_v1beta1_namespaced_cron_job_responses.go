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

// ReadBatchV1beta1NamespacedCronJobReader is a Reader for the ReadBatchV1beta1NamespacedCronJob structure.
type ReadBatchV1beta1NamespacedCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadBatchV1beta1NamespacedCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadBatchV1beta1NamespacedCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadBatchV1beta1NamespacedCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadBatchV1beta1NamespacedCronJobOK creates a ReadBatchV1beta1NamespacedCronJobOK with default headers values
func NewReadBatchV1beta1NamespacedCronJobOK() *ReadBatchV1beta1NamespacedCronJobOK {
	return &ReadBatchV1beta1NamespacedCronJobOK{}
}

/*ReadBatchV1beta1NamespacedCronJobOK handles this case with default header values.

OK
*/
type ReadBatchV1beta1NamespacedCronJobOK struct {
	Payload *models.IoK8sAPIBatchV1beta1CronJob
}

func (o *ReadBatchV1beta1NamespacedCronJobOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}][%d] readBatchV1beta1NamespacedCronJobOK  %+v", 200, o.Payload)
}

func (o *ReadBatchV1beta1NamespacedCronJobOK) GetPayload() *models.IoK8sAPIBatchV1beta1CronJob {
	return o.Payload
}

func (o *ReadBatchV1beta1NamespacedCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV1beta1CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadBatchV1beta1NamespacedCronJobUnauthorized creates a ReadBatchV1beta1NamespacedCronJobUnauthorized with default headers values
func NewReadBatchV1beta1NamespacedCronJobUnauthorized() *ReadBatchV1beta1NamespacedCronJobUnauthorized {
	return &ReadBatchV1beta1NamespacedCronJobUnauthorized{}
}

/*ReadBatchV1beta1NamespacedCronJobUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadBatchV1beta1NamespacedCronJobUnauthorized struct {
}

func (o *ReadBatchV1beta1NamespacedCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1beta1/namespaces/{namespace}/cronjobs/{name}][%d] readBatchV1beta1NamespacedCronJobUnauthorized ", 401)
}

func (o *ReadBatchV1beta1NamespacedCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
