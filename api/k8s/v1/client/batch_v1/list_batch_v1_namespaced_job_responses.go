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

// ListBatchV1NamespacedJobReader is a Reader for the ListBatchV1NamespacedJob structure.
type ListBatchV1NamespacedJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBatchV1NamespacedJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBatchV1NamespacedJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListBatchV1NamespacedJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListBatchV1NamespacedJobOK creates a ListBatchV1NamespacedJobOK with default headers values
func NewListBatchV1NamespacedJobOK() *ListBatchV1NamespacedJobOK {
	return &ListBatchV1NamespacedJobOK{}
}

/*ListBatchV1NamespacedJobOK handles this case with default header values.

OK
*/
type ListBatchV1NamespacedJobOK struct {
	Payload *models.IoK8sAPIBatchV1JobList
}

func (o *ListBatchV1NamespacedJobOK) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1/namespaces/{namespace}/jobs][%d] listBatchV1NamespacedJobOK  %+v", 200, o.Payload)
}

func (o *ListBatchV1NamespacedJobOK) GetPayload() *models.IoK8sAPIBatchV1JobList {
	return o.Payload
}

func (o *ListBatchV1NamespacedJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV1JobList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBatchV1NamespacedJobUnauthorized creates a ListBatchV1NamespacedJobUnauthorized with default headers values
func NewListBatchV1NamespacedJobUnauthorized() *ListBatchV1NamespacedJobUnauthorized {
	return &ListBatchV1NamespacedJobUnauthorized{}
}

/*ListBatchV1NamespacedJobUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBatchV1NamespacedJobUnauthorized struct {
}

func (o *ListBatchV1NamespacedJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/batch/v1/namespaces/{namespace}/jobs][%d] listBatchV1NamespacedJobUnauthorized ", 401)
}

func (o *ListBatchV1NamespacedJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
