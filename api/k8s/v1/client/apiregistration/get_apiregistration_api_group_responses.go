// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetApiregistrationAPIGroupReader is a Reader for the GetApiregistrationAPIGroup structure.
type GetApiregistrationAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApiregistrationAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApiregistrationAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApiregistrationAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApiregistrationAPIGroupOK creates a GetApiregistrationAPIGroupOK with default headers values
func NewGetApiregistrationAPIGroupOK() *GetApiregistrationAPIGroupOK {
	return &GetApiregistrationAPIGroupOK{}
}

/*GetApiregistrationAPIGroupOK handles this case with default header values.

OK
*/
type GetApiregistrationAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetApiregistrationAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/][%d] getApiregistrationApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetApiregistrationAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetApiregistrationAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApiregistrationAPIGroupUnauthorized creates a GetApiregistrationAPIGroupUnauthorized with default headers values
func NewGetApiregistrationAPIGroupUnauthorized() *GetApiregistrationAPIGroupUnauthorized {
	return &GetApiregistrationAPIGroupUnauthorized{}
}

/*GetApiregistrationAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetApiregistrationAPIGroupUnauthorized struct {
}

func (o *GetApiregistrationAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/][%d] getApiregistrationApiGroupUnauthorized ", 401)
}

func (o *GetApiregistrationAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
