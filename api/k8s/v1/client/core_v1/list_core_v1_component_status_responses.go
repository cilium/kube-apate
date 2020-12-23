// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListCoreV1ComponentStatusReader is a Reader for the ListCoreV1ComponentStatus structure.
type ListCoreV1ComponentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCoreV1ComponentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCoreV1ComponentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCoreV1ComponentStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCoreV1ComponentStatusOK creates a ListCoreV1ComponentStatusOK with default headers values
func NewListCoreV1ComponentStatusOK() *ListCoreV1ComponentStatusOK {
	return &ListCoreV1ComponentStatusOK{}
}

/*ListCoreV1ComponentStatusOK handles this case with default header values.

OK
*/
type ListCoreV1ComponentStatusOK struct {
	Payload *models.IoK8sAPICoreV1ComponentStatusList
}

func (o *ListCoreV1ComponentStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/componentstatuses][%d] listCoreV1ComponentStatusOK  %+v", 200, o.Payload)
}

func (o *ListCoreV1ComponentStatusOK) GetPayload() *models.IoK8sAPICoreV1ComponentStatusList {
	return o.Payload
}

func (o *ListCoreV1ComponentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ComponentStatusList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCoreV1ComponentStatusUnauthorized creates a ListCoreV1ComponentStatusUnauthorized with default headers values
func NewListCoreV1ComponentStatusUnauthorized() *ListCoreV1ComponentStatusUnauthorized {
	return &ListCoreV1ComponentStatusUnauthorized{}
}

/*ListCoreV1ComponentStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCoreV1ComponentStatusUnauthorized struct {
}

func (o *ListCoreV1ComponentStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/componentstatuses][%d] listCoreV1ComponentStatusUnauthorized ", 401)
}

func (o *ListCoreV1ComponentStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}