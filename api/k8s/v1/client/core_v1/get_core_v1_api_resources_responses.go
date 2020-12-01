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

// GetCoreV1APIResourcesReader is a Reader for the GetCoreV1APIResources structure.
type GetCoreV1APIResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoreV1APIResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoreV1APIResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCoreV1APIResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoreV1APIResourcesOK creates a GetCoreV1APIResourcesOK with default headers values
func NewGetCoreV1APIResourcesOK() *GetCoreV1APIResourcesOK {
	return &GetCoreV1APIResourcesOK{}
}

/*GetCoreV1APIResourcesOK handles this case with default header values.

OK
*/
type GetCoreV1APIResourcesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList
}

func (o *GetCoreV1APIResourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/][%d] getCoreV1ApiResourcesOK  %+v", 200, o.Payload)
}

func (o *GetCoreV1APIResourcesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIResourceList {
	return o.Payload
}

func (o *GetCoreV1APIResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIResourceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoreV1APIResourcesUnauthorized creates a GetCoreV1APIResourcesUnauthorized with default headers values
func NewGetCoreV1APIResourcesUnauthorized() *GetCoreV1APIResourcesUnauthorized {
	return &GetCoreV1APIResourcesUnauthorized{}
}

/*GetCoreV1APIResourcesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCoreV1APIResourcesUnauthorized struct {
}

func (o *GetCoreV1APIResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/][%d] getCoreV1ApiResourcesUnauthorized ", 401)
}

func (o *GetCoreV1APIResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
