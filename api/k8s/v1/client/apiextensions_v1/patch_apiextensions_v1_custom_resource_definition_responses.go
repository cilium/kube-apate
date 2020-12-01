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

// PatchApiextensionsV1CustomResourceDefinitionReader is a Reader for the PatchApiextensionsV1CustomResourceDefinition structure.
type PatchApiextensionsV1CustomResourceDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchApiextensionsV1CustomResourceDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApiextensionsV1CustomResourceDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchApiextensionsV1CustomResourceDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchApiextensionsV1CustomResourceDefinitionOK creates a PatchApiextensionsV1CustomResourceDefinitionOK with default headers values
func NewPatchApiextensionsV1CustomResourceDefinitionOK() *PatchApiextensionsV1CustomResourceDefinitionOK {
	return &PatchApiextensionsV1CustomResourceDefinitionOK{}
}

/*PatchApiextensionsV1CustomResourceDefinitionOK handles this case with default header values.

OK
*/
type PatchApiextensionsV1CustomResourceDefinitionOK struct {
	Payload *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition
}

func (o *PatchApiextensionsV1CustomResourceDefinitionOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}][%d] patchApiextensionsV1CustomResourceDefinitionOK  %+v", 200, o.Payload)
}

func (o *PatchApiextensionsV1CustomResourceDefinitionOK) GetPayload() *models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition {
	return o.Payload
}

func (o *PatchApiextensionsV1CustomResourceDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApiextensionsApiserverPkgApisApiextensionsV1CustomResourceDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApiextensionsV1CustomResourceDefinitionUnauthorized creates a PatchApiextensionsV1CustomResourceDefinitionUnauthorized with default headers values
func NewPatchApiextensionsV1CustomResourceDefinitionUnauthorized() *PatchApiextensionsV1CustomResourceDefinitionUnauthorized {
	return &PatchApiextensionsV1CustomResourceDefinitionUnauthorized{}
}

/*PatchApiextensionsV1CustomResourceDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchApiextensionsV1CustomResourceDefinitionUnauthorized struct {
}

func (o *PatchApiextensionsV1CustomResourceDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/apiextensions.k8s.io/v1/customresourcedefinitions/{name}][%d] patchApiextensionsV1CustomResourceDefinitionUnauthorized ", 401)
}

func (o *PatchApiextensionsV1CustomResourceDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
