// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListApiextensionsV1CustomResourceDefinitionReader is a Reader for the ListApiextensionsV1CustomResourceDefinition structure.
type ListApiextensionsV1CustomResourceDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListApiextensionsV1CustomResourceDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListApiextensionsV1CustomResourceDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListApiextensionsV1CustomResourceDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListApiextensionsV1CustomResourceDefinitionOK creates a ListApiextensionsV1CustomResourceDefinitionOK with default headers values
func NewListApiextensionsV1CustomResourceDefinitionOK() *ListApiextensionsV1CustomResourceDefinitionOK {
	return &ListApiextensionsV1CustomResourceDefinitionOK{}
}

/*ListApiextensionsV1CustomResourceDefinitionOK handles this case with default header values.

OK
*/
type ListApiextensionsV1CustomResourceDefinitionOK struct {
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionList
}

func (o *ListApiextensionsV1CustomResourceDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiextensions.k8s.io/v1/customresourcedefinitions][%d] listApiextensionsV1CustomResourceDefinitionOK  %+v", 200, o.Payload)
}

func (o *ListApiextensionsV1CustomResourceDefinitionOK) GetPayload() *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionList {
	return o.Payload
}

func (o *ListApiextensionsV1CustomResourceDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinitionList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListApiextensionsV1CustomResourceDefinitionUnauthorized creates a ListApiextensionsV1CustomResourceDefinitionUnauthorized with default headers values
func NewListApiextensionsV1CustomResourceDefinitionUnauthorized() *ListApiextensionsV1CustomResourceDefinitionUnauthorized {
	return &ListApiextensionsV1CustomResourceDefinitionUnauthorized{}
}

/*ListApiextensionsV1CustomResourceDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListApiextensionsV1CustomResourceDefinitionUnauthorized struct {
}

func (o *ListApiextensionsV1CustomResourceDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiextensions.k8s.io/v1/customresourcedefinitions][%d] listApiextensionsV1CustomResourceDefinitionUnauthorized ", 401)
}

func (o *ListApiextensionsV1CustomResourceDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
