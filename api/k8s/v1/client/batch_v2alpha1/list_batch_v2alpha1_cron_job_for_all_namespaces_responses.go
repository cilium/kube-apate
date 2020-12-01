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

// ListBatchV2alpha1CronJobForAllNamespacesReader is a Reader for the ListBatchV2alpha1CronJobForAllNamespaces structure.
type ListBatchV2alpha1CronJobForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBatchV2alpha1CronJobForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBatchV2alpha1CronJobForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListBatchV2alpha1CronJobForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListBatchV2alpha1CronJobForAllNamespacesOK creates a ListBatchV2alpha1CronJobForAllNamespacesOK with default headers values
func NewListBatchV2alpha1CronJobForAllNamespacesOK() *ListBatchV2alpha1CronJobForAllNamespacesOK {
	return &ListBatchV2alpha1CronJobForAllNamespacesOK{}
}

/*ListBatchV2alpha1CronJobForAllNamespacesOK handles this case with default header values.

OK
*/
type ListBatchV2alpha1CronJobForAllNamespacesOK struct {
	Payload *models.IoK8sAPIBatchV2alpha1CronJobList
}

func (o *ListBatchV2alpha1CronJobForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v2alpha1/cronjobs][%d] listBatchV2alpha1CronJobForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListBatchV2alpha1CronJobForAllNamespacesOK) GetPayload() *models.IoK8sAPIBatchV2alpha1CronJobList {
	return o.Payload
}

func (o *ListBatchV2alpha1CronJobForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV2alpha1CronJobList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBatchV2alpha1CronJobForAllNamespacesUnauthorized creates a ListBatchV2alpha1CronJobForAllNamespacesUnauthorized with default headers values
func NewListBatchV2alpha1CronJobForAllNamespacesUnauthorized() *ListBatchV2alpha1CronJobForAllNamespacesUnauthorized {
	return &ListBatchV2alpha1CronJobForAllNamespacesUnauthorized{}
}

/*ListBatchV2alpha1CronJobForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBatchV2alpha1CronJobForAllNamespacesUnauthorized struct {
}

func (o *ListBatchV2alpha1CronJobForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v2alpha1/cronjobs][%d] listBatchV2alpha1CronJobForAllNamespacesUnauthorized ", 401)
}

func (o *ListBatchV2alpha1CronJobForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
