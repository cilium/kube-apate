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

// ReplaceCoreV1NodeReader is a Reader for the ReplaceCoreV1Node structure.
type ReplaceCoreV1NodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoreV1NodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoreV1NodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoreV1NodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoreV1NodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoreV1NodeOK creates a ReplaceCoreV1NodeOK with default headers values
func NewReplaceCoreV1NodeOK() *ReplaceCoreV1NodeOK {
	return &ReplaceCoreV1NodeOK{}
}

/*ReplaceCoreV1NodeOK handles this case with default header values.

OK
*/
type ReplaceCoreV1NodeOK struct {
	Payload *models.IoK8sAPICoreV1Node
}

func (o *ReplaceCoreV1NodeOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/nodes/{name}][%d] replaceCoreV1NodeOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoreV1NodeOK) GetPayload() *models.IoK8sAPICoreV1Node {
	return o.Payload
}

func (o *ReplaceCoreV1NodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NodeCreated creates a ReplaceCoreV1NodeCreated with default headers values
func NewReplaceCoreV1NodeCreated() *ReplaceCoreV1NodeCreated {
	return &ReplaceCoreV1NodeCreated{}
}

/*ReplaceCoreV1NodeCreated handles this case with default header values.

Created
*/
type ReplaceCoreV1NodeCreated struct {
	Payload *models.IoK8sAPICoreV1Node
}

func (o *ReplaceCoreV1NodeCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/nodes/{name}][%d] replaceCoreV1NodeCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoreV1NodeCreated) GetPayload() *models.IoK8sAPICoreV1Node {
	return o.Payload
}

func (o *ReplaceCoreV1NodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NodeUnauthorized creates a ReplaceCoreV1NodeUnauthorized with default headers values
func NewReplaceCoreV1NodeUnauthorized() *ReplaceCoreV1NodeUnauthorized {
	return &ReplaceCoreV1NodeUnauthorized{}
}

/*ReplaceCoreV1NodeUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoreV1NodeUnauthorized struct {
}

func (o *ReplaceCoreV1NodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/nodes/{name}][%d] replaceCoreV1NodeUnauthorized ", 401)
}

func (o *ReplaceCoreV1NodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}