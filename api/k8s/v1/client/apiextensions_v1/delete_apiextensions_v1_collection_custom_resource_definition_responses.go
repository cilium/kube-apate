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

// DeleteApiextensionsV1CollectionCustomResourceDefinitionReader is a Reader for the DeleteApiextensionsV1CollectionCustomResourceDefinition structure.
type DeleteApiextensionsV1CollectionCustomResourceDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApiextensionsV1CollectionCustomResourceDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApiextensionsV1CollectionCustomResourceDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteApiextensionsV1CollectionCustomResourceDefinitionOK creates a DeleteApiextensionsV1CollectionCustomResourceDefinitionOK with default headers values
func NewDeleteApiextensionsV1CollectionCustomResourceDefinitionOK() *DeleteApiextensionsV1CollectionCustomResourceDefinitionOK {
	return &DeleteApiextensionsV1CollectionCustomResourceDefinitionOK{}
}

/*DeleteApiextensionsV1CollectionCustomResourceDefinitionOK handles this case with default header values.

OK
*/
type DeleteApiextensionsV1CollectionCustomResourceDefinitionOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteApiextensionsV1CollectionCustomResourceDefinitionOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/apiextensions.k8s.io/v1/customresourcedefinitions][%d] deleteApiextensionsV1CollectionCustomResourceDefinitionOK  %+v", 200, o.Payload)
}

func (o *DeleteApiextensionsV1CollectionCustomResourceDefinitionOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteApiextensionsV1CollectionCustomResourceDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized creates a DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized with default headers values
func NewDeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized() *DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized {
	return &DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized{}
}

/*DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized struct {
}

func (o *DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/apiextensions.k8s.io/v1/customresourcedefinitions][%d] deleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized ", 401)
}

func (o *DeleteApiextensionsV1CollectionCustomResourceDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
