// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetApiregistrationV1APIResourcesReader is a Reader for the GetApiregistrationV1APIResources structure.
type GetApiregistrationV1APIResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApiregistrationV1APIResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApiregistrationV1APIResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApiregistrationV1APIResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApiregistrationV1APIResourcesOK creates a GetApiregistrationV1APIResourcesOK with default headers values
func NewGetApiregistrationV1APIResourcesOK() *GetApiregistrationV1APIResourcesOK {
	return &GetApiregistrationV1APIResourcesOK{}
}

/*GetApiregistrationV1APIResourcesOK handles this case with default header values.

OK
*/
type GetApiregistrationV1APIResourcesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList
}

func (o *GetApiregistrationV1APIResourcesOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/v1/][%d] getApiregistrationV1ApiResourcesOK  %+v", 200, o.Payload)
}

func (o *GetApiregistrationV1APIResourcesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIResourceList {
	return o.Payload
}

func (o *GetApiregistrationV1APIResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIResourceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApiregistrationV1APIResourcesUnauthorized creates a GetApiregistrationV1APIResourcesUnauthorized with default headers values
func NewGetApiregistrationV1APIResourcesUnauthorized() *GetApiregistrationV1APIResourcesUnauthorized {
	return &GetApiregistrationV1APIResourcesUnauthorized{}
}

/*GetApiregistrationV1APIResourcesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetApiregistrationV1APIResourcesUnauthorized struct {
}

func (o *GetApiregistrationV1APIResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/v1/][%d] getApiregistrationV1ApiResourcesUnauthorized ", 401)
}

func (o *GetApiregistrationV1APIResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
