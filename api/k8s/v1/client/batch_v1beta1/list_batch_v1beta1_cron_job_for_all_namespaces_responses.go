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

// ListBatchV1beta1CronJobForAllNamespacesReader is a Reader for the ListBatchV1beta1CronJobForAllNamespaces structure.
type ListBatchV1beta1CronJobForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBatchV1beta1CronJobForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBatchV1beta1CronJobForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListBatchV1beta1CronJobForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListBatchV1beta1CronJobForAllNamespacesOK creates a ListBatchV1beta1CronJobForAllNamespacesOK with default headers values
func NewListBatchV1beta1CronJobForAllNamespacesOK() *ListBatchV1beta1CronJobForAllNamespacesOK {
	return &ListBatchV1beta1CronJobForAllNamespacesOK{}
}

/*ListBatchV1beta1CronJobForAllNamespacesOK handles this case with default header values.

OK
*/
type ListBatchV1beta1CronJobForAllNamespacesOK struct {
	Payload *models.IoK8sAPIBatchV1beta1CronJobList
}

func (o *ListBatchV1beta1CronJobForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1beta1/cronjobs][%d] listBatchV1beta1CronJobForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListBatchV1beta1CronJobForAllNamespacesOK) GetPayload() *models.IoK8sAPIBatchV1beta1CronJobList {
	return o.Payload
}

func (o *ListBatchV1beta1CronJobForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV1beta1CronJobList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBatchV1beta1CronJobForAllNamespacesUnauthorized creates a ListBatchV1beta1CronJobForAllNamespacesUnauthorized with default headers values
func NewListBatchV1beta1CronJobForAllNamespacesUnauthorized() *ListBatchV1beta1CronJobForAllNamespacesUnauthorized {
	return &ListBatchV1beta1CronJobForAllNamespacesUnauthorized{}
}

/*ListBatchV1beta1CronJobForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBatchV1beta1CronJobForAllNamespacesUnauthorized struct {
}

func (o *ListBatchV1beta1CronJobForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1beta1/cronjobs][%d] listBatchV1beta1CronJobForAllNamespacesUnauthorized ", 401)
}

func (o *ListBatchV1beta1CronJobForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}