// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiextensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadApiextensionsV1beta1CustomResourceDefinitionReader is a Reader for the ReadApiextensionsV1beta1CustomResourceDefinition structure.
type ReadApiextensionsV1beta1CustomResourceDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadApiextensionsV1beta1CustomResourceDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadApiextensionsV1beta1CustomResourceDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadApiextensionsV1beta1CustomResourceDefinitionOK creates a ReadApiextensionsV1beta1CustomResourceDefinitionOK with default headers values
func NewReadApiextensionsV1beta1CustomResourceDefinitionOK() *ReadApiextensionsV1beta1CustomResourceDefinitionOK {
	return &ReadApiextensionsV1beta1CustomResourceDefinitionOK{}
}

/*ReadApiextensionsV1beta1CustomResourceDefinitionOK handles this case with default header values.

OK
*/
type ReadApiextensionsV1beta1CustomResourceDefinitionOK struct {
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition
}

func (o *ReadApiextensionsV1beta1CustomResourceDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions/{name}][%d] readApiextensionsV1beta1CustomResourceDefinitionOK  %+v", 200, o.Payload)
}

func (o *ReadApiextensionsV1beta1CustomResourceDefinitionOK) GetPayload() *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition {
	return o.Payload
}

func (o *ReadApiextensionsV1beta1CustomResourceDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized creates a ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized with default headers values
func NewReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized() *ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized {
	return &ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized{}
}

/*ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized struct {
}

func (o *ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiextensions.k8s.io/v1beta1/customresourcedefinitions/{name}][%d] readApiextensionsV1beta1CustomResourceDefinitionUnauthorized ", 401)
}

func (o *ReadApiextensionsV1beta1CustomResourceDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
