// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateSchedulingV1alpha1PriorityClassReader is a Reader for the CreateSchedulingV1alpha1PriorityClass structure.
type CreateSchedulingV1alpha1PriorityClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSchedulingV1alpha1PriorityClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSchedulingV1alpha1PriorityClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateSchedulingV1alpha1PriorityClassCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateSchedulingV1alpha1PriorityClassAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateSchedulingV1alpha1PriorityClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSchedulingV1alpha1PriorityClassOK creates a CreateSchedulingV1alpha1PriorityClassOK with default headers values
func NewCreateSchedulingV1alpha1PriorityClassOK() *CreateSchedulingV1alpha1PriorityClassOK {
	return &CreateSchedulingV1alpha1PriorityClassOK{}
}

/*CreateSchedulingV1alpha1PriorityClassOK handles this case with default header values.

OK
*/
type CreateSchedulingV1alpha1PriorityClassOK struct {
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClass
}

func (o *CreateSchedulingV1alpha1PriorityClassOK) Error() string {
	return fmt.Sprintf("[POST /apis/scheduling.k8s.io/v1alpha1/priorityclasses][%d] createSchedulingV1alpha1PriorityClassOK  %+v", 200, o.Payload)
}

func (o *CreateSchedulingV1alpha1PriorityClassOK) GetPayload() *models.IoK8sAPISchedulingV1alpha1PriorityClass {
	return o.Payload
}

func (o *CreateSchedulingV1alpha1PriorityClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISchedulingV1alpha1PriorityClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchedulingV1alpha1PriorityClassCreated creates a CreateSchedulingV1alpha1PriorityClassCreated with default headers values
func NewCreateSchedulingV1alpha1PriorityClassCreated() *CreateSchedulingV1alpha1PriorityClassCreated {
	return &CreateSchedulingV1alpha1PriorityClassCreated{}
}

/*CreateSchedulingV1alpha1PriorityClassCreated handles this case with default header values.

Created
*/
type CreateSchedulingV1alpha1PriorityClassCreated struct {
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClass
}

func (o *CreateSchedulingV1alpha1PriorityClassCreated) Error() string {
	return fmt.Sprintf("[POST /apis/scheduling.k8s.io/v1alpha1/priorityclasses][%d] createSchedulingV1alpha1PriorityClassCreated  %+v", 201, o.Payload)
}

func (o *CreateSchedulingV1alpha1PriorityClassCreated) GetPayload() *models.IoK8sAPISchedulingV1alpha1PriorityClass {
	return o.Payload
}

func (o *CreateSchedulingV1alpha1PriorityClassCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISchedulingV1alpha1PriorityClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchedulingV1alpha1PriorityClassAccepted creates a CreateSchedulingV1alpha1PriorityClassAccepted with default headers values
func NewCreateSchedulingV1alpha1PriorityClassAccepted() *CreateSchedulingV1alpha1PriorityClassAccepted {
	return &CreateSchedulingV1alpha1PriorityClassAccepted{}
}

/*CreateSchedulingV1alpha1PriorityClassAccepted handles this case with default header values.

Accepted
*/
type CreateSchedulingV1alpha1PriorityClassAccepted struct {
	Payload *models.IoK8sAPISchedulingV1alpha1PriorityClass
}

func (o *CreateSchedulingV1alpha1PriorityClassAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/scheduling.k8s.io/v1alpha1/priorityclasses][%d] createSchedulingV1alpha1PriorityClassAccepted  %+v", 202, o.Payload)
}

func (o *CreateSchedulingV1alpha1PriorityClassAccepted) GetPayload() *models.IoK8sAPISchedulingV1alpha1PriorityClass {
	return o.Payload
}

func (o *CreateSchedulingV1alpha1PriorityClassAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISchedulingV1alpha1PriorityClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSchedulingV1alpha1PriorityClassUnauthorized creates a CreateSchedulingV1alpha1PriorityClassUnauthorized with default headers values
func NewCreateSchedulingV1alpha1PriorityClassUnauthorized() *CreateSchedulingV1alpha1PriorityClassUnauthorized {
	return &CreateSchedulingV1alpha1PriorityClassUnauthorized{}
}

/*CreateSchedulingV1alpha1PriorityClassUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSchedulingV1alpha1PriorityClassUnauthorized struct {
}

func (o *CreateSchedulingV1alpha1PriorityClassUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/scheduling.k8s.io/v1alpha1/priorityclasses][%d] createSchedulingV1alpha1PriorityClassUnauthorized ", 401)
}

func (o *CreateSchedulingV1alpha1PriorityClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
