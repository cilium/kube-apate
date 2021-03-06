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

// ReplaceApiextensionsV1CustomResourceDefinitionReader is a Reader for the ReplaceApiextensionsV1CustomResourceDefinition structure.
type ReplaceApiextensionsV1CustomResourceDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceApiextensionsV1CustomResourceDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceApiextensionsV1CustomResourceDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceApiextensionsV1CustomResourceDefinitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceApiextensionsV1CustomResourceDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceApiextensionsV1CustomResourceDefinitionOK creates a ReplaceApiextensionsV1CustomResourceDefinitionOK with default headers values
func NewReplaceApiextensionsV1CustomResourceDefinitionOK() *ReplaceApiextensionsV1CustomResourceDefinitionOK {
	return &ReplaceApiextensionsV1CustomResourceDefinitionOK{}
}

/*ReplaceApiextensionsV1CustomResourceDefinitionOK handles this case with default header values.

OK
*/
type ReplaceApiextensionsV1CustomResourceDefinitionOK struct {
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionOK) Error() string {
	return fmt.Sprintf("[PUT /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}][%d] replaceApiextensionsV1CustomResourceDefinitionOK  %+v", 200, o.Payload)
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionOK) GetPayload() *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition {
	return o.Payload
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceApiextensionsV1CustomResourceDefinitionCreated creates a ReplaceApiextensionsV1CustomResourceDefinitionCreated with default headers values
func NewReplaceApiextensionsV1CustomResourceDefinitionCreated() *ReplaceApiextensionsV1CustomResourceDefinitionCreated {
	return &ReplaceApiextensionsV1CustomResourceDefinitionCreated{}
}

/*ReplaceApiextensionsV1CustomResourceDefinitionCreated handles this case with default header values.

Created
*/
type ReplaceApiextensionsV1CustomResourceDefinitionCreated struct {
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}][%d] replaceApiextensionsV1CustomResourceDefinitionCreated  %+v", 201, o.Payload)
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionCreated) GetPayload() *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition {
	return o.Payload
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceApiextensionsV1CustomResourceDefinitionUnauthorized creates a ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized with default headers values
func NewReplaceApiextensionsV1CustomResourceDefinitionUnauthorized() *ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized {
	return &ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized{}
}

/*ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized struct {
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}][%d] replaceApiextensionsV1CustomResourceDefinitionUnauthorized ", 401)
}

func (o *ReplaceApiextensionsV1CustomResourceDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
