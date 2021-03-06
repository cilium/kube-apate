// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceNodeV1alpha1RuntimeClassReader is a Reader for the ReplaceNodeV1alpha1RuntimeClass structure.
type ReplaceNodeV1alpha1RuntimeClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceNodeV1alpha1RuntimeClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceNodeV1alpha1RuntimeClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceNodeV1alpha1RuntimeClassCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceNodeV1alpha1RuntimeClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceNodeV1alpha1RuntimeClassOK creates a ReplaceNodeV1alpha1RuntimeClassOK with default headers values
func NewReplaceNodeV1alpha1RuntimeClassOK() *ReplaceNodeV1alpha1RuntimeClassOK {
	return &ReplaceNodeV1alpha1RuntimeClassOK{}
}

/*ReplaceNodeV1alpha1RuntimeClassOK handles this case with default header values.

OK
*/
type ReplaceNodeV1alpha1RuntimeClassOK struct {
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClass
}

func (o *ReplaceNodeV1alpha1RuntimeClassOK) Error() string {
	return fmt.Sprintf("[PUT /apis/node.k8s.io/v1alpha1/runtimeclasses/{name}][%d] replaceNodeV1alpha1RuntimeClassOK  %+v", 200, o.Payload)
}

func (o *ReplaceNodeV1alpha1RuntimeClassOK) GetPayload() *models.IoK8sAPINodeV1alpha1RuntimeClass {
	return o.Payload
}

func (o *ReplaceNodeV1alpha1RuntimeClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINodeV1alpha1RuntimeClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNodeV1alpha1RuntimeClassCreated creates a ReplaceNodeV1alpha1RuntimeClassCreated with default headers values
func NewReplaceNodeV1alpha1RuntimeClassCreated() *ReplaceNodeV1alpha1RuntimeClassCreated {
	return &ReplaceNodeV1alpha1RuntimeClassCreated{}
}

/*ReplaceNodeV1alpha1RuntimeClassCreated handles this case with default header values.

Created
*/
type ReplaceNodeV1alpha1RuntimeClassCreated struct {
	Payload *models.IoK8sAPINodeV1alpha1RuntimeClass
}

func (o *ReplaceNodeV1alpha1RuntimeClassCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/node.k8s.io/v1alpha1/runtimeclasses/{name}][%d] replaceNodeV1alpha1RuntimeClassCreated  %+v", 201, o.Payload)
}

func (o *ReplaceNodeV1alpha1RuntimeClassCreated) GetPayload() *models.IoK8sAPINodeV1alpha1RuntimeClass {
	return o.Payload
}

func (o *ReplaceNodeV1alpha1RuntimeClassCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINodeV1alpha1RuntimeClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNodeV1alpha1RuntimeClassUnauthorized creates a ReplaceNodeV1alpha1RuntimeClassUnauthorized with default headers values
func NewReplaceNodeV1alpha1RuntimeClassUnauthorized() *ReplaceNodeV1alpha1RuntimeClassUnauthorized {
	return &ReplaceNodeV1alpha1RuntimeClassUnauthorized{}
}

/*ReplaceNodeV1alpha1RuntimeClassUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceNodeV1alpha1RuntimeClassUnauthorized struct {
}

func (o *ReplaceNodeV1alpha1RuntimeClassUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/node.k8s.io/v1alpha1/runtimeclasses/{name}][%d] replaceNodeV1alpha1RuntimeClassUnauthorized ", 401)
}

func (o *ReplaceNodeV1alpha1RuntimeClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
