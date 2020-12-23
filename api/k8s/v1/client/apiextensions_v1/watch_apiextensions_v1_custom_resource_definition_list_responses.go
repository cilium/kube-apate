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

// WatchApiextensionsV1CustomResourceDefinitionListReader is a Reader for the WatchApiextensionsV1CustomResourceDefinitionList structure.
type WatchApiextensionsV1CustomResourceDefinitionListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchApiextensionsV1CustomResourceDefinitionListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchApiextensionsV1CustomResourceDefinitionListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchApiextensionsV1CustomResourceDefinitionListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchApiextensionsV1CustomResourceDefinitionListOK creates a WatchApiextensionsV1CustomResourceDefinitionListOK with default headers values
func NewWatchApiextensionsV1CustomResourceDefinitionListOK() *WatchApiextensionsV1CustomResourceDefinitionListOK {
	return &WatchApiextensionsV1CustomResourceDefinitionListOK{}
}

/*WatchApiextensionsV1CustomResourceDefinitionListOK handles this case with default header values.

OK
*/
type WatchApiextensionsV1CustomResourceDefinitionListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchApiextensionsV1CustomResourceDefinitionListOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiextensions.k8s.io/v1/watch/customresourcedefinitions][%d] watchApiextensionsV1CustomResourceDefinitionListOK  %+v", 200, o.Payload)
}

func (o *WatchApiextensionsV1CustomResourceDefinitionListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchApiextensionsV1CustomResourceDefinitionListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchApiextensionsV1CustomResourceDefinitionListUnauthorized creates a WatchApiextensionsV1CustomResourceDefinitionListUnauthorized with default headers values
func NewWatchApiextensionsV1CustomResourceDefinitionListUnauthorized() *WatchApiextensionsV1CustomResourceDefinitionListUnauthorized {
	return &WatchApiextensionsV1CustomResourceDefinitionListUnauthorized{}
}

/*WatchApiextensionsV1CustomResourceDefinitionListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchApiextensionsV1CustomResourceDefinitionListUnauthorized struct {
}

func (o *WatchApiextensionsV1CustomResourceDefinitionListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiextensions.k8s.io/v1/watch/customresourcedefinitions][%d] watchApiextensionsV1CustomResourceDefinitionListUnauthorized ", 401)
}

func (o *WatchApiextensionsV1CustomResourceDefinitionListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}