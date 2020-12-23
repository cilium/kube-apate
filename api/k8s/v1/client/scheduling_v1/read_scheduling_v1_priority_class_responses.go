// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadSchedulingV1PriorityClassReader is a Reader for the ReadSchedulingV1PriorityClass structure.
type ReadSchedulingV1PriorityClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadSchedulingV1PriorityClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadSchedulingV1PriorityClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadSchedulingV1PriorityClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadSchedulingV1PriorityClassOK creates a ReadSchedulingV1PriorityClassOK with default headers values
func NewReadSchedulingV1PriorityClassOK() *ReadSchedulingV1PriorityClassOK {
	return &ReadSchedulingV1PriorityClassOK{}
}

/*ReadSchedulingV1PriorityClassOK handles this case with default header values.

OK
*/
type ReadSchedulingV1PriorityClassOK struct {
	Payload *models.IoK8sAPISchedulingV1PriorityClass
}

func (o *ReadSchedulingV1PriorityClassOK) Error() string {
	return fmt.Sprintf("[GET /apis/scheduling.k8s.io/v1/priorityclasses/{name}][%d] readSchedulingV1PriorityClassOK  %+v", 200, o.Payload)
}

func (o *ReadSchedulingV1PriorityClassOK) GetPayload() *models.IoK8sAPISchedulingV1PriorityClass {
	return o.Payload
}

func (o *ReadSchedulingV1PriorityClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISchedulingV1PriorityClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadSchedulingV1PriorityClassUnauthorized creates a ReadSchedulingV1PriorityClassUnauthorized with default headers values
func NewReadSchedulingV1PriorityClassUnauthorized() *ReadSchedulingV1PriorityClassUnauthorized {
	return &ReadSchedulingV1PriorityClassUnauthorized{}
}

/*ReadSchedulingV1PriorityClassUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadSchedulingV1PriorityClassUnauthorized struct {
}

func (o *ReadSchedulingV1PriorityClassUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/scheduling.k8s.io/v1/priorityclasses/{name}][%d] readSchedulingV1PriorityClassUnauthorized ", 401)
}

func (o *ReadSchedulingV1PriorityClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}