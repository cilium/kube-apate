// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteSchedulingV1beta1PriorityClassReader is a Reader for the DeleteSchedulingV1beta1PriorityClass structure.
type DeleteSchedulingV1beta1PriorityClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSchedulingV1beta1PriorityClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSchedulingV1beta1PriorityClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteSchedulingV1beta1PriorityClassAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteSchedulingV1beta1PriorityClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSchedulingV1beta1PriorityClassOK creates a DeleteSchedulingV1beta1PriorityClassOK with default headers values
func NewDeleteSchedulingV1beta1PriorityClassOK() *DeleteSchedulingV1beta1PriorityClassOK {
	return &DeleteSchedulingV1beta1PriorityClassOK{}
}

/*DeleteSchedulingV1beta1PriorityClassOK handles this case with default header values.

OK
*/
type DeleteSchedulingV1beta1PriorityClassOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteSchedulingV1beta1PriorityClassOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/scheduling.k8s.io/v1beta1/priorityclasses/{name}][%d] deleteSchedulingV1beta1PriorityClassOK  %+v", 200, o.Payload)
}

func (o *DeleteSchedulingV1beta1PriorityClassOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteSchedulingV1beta1PriorityClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSchedulingV1beta1PriorityClassAccepted creates a DeleteSchedulingV1beta1PriorityClassAccepted with default headers values
func NewDeleteSchedulingV1beta1PriorityClassAccepted() *DeleteSchedulingV1beta1PriorityClassAccepted {
	return &DeleteSchedulingV1beta1PriorityClassAccepted{}
}

/*DeleteSchedulingV1beta1PriorityClassAccepted handles this case with default header values.

Accepted
*/
type DeleteSchedulingV1beta1PriorityClassAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteSchedulingV1beta1PriorityClassAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/scheduling.k8s.io/v1beta1/priorityclasses/{name}][%d] deleteSchedulingV1beta1PriorityClassAccepted  %+v", 202, o.Payload)
}

func (o *DeleteSchedulingV1beta1PriorityClassAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteSchedulingV1beta1PriorityClassAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSchedulingV1beta1PriorityClassUnauthorized creates a DeleteSchedulingV1beta1PriorityClassUnauthorized with default headers values
func NewDeleteSchedulingV1beta1PriorityClassUnauthorized() *DeleteSchedulingV1beta1PriorityClassUnauthorized {
	return &DeleteSchedulingV1beta1PriorityClassUnauthorized{}
}

/*DeleteSchedulingV1beta1PriorityClassUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteSchedulingV1beta1PriorityClassUnauthorized struct {
}

func (o *DeleteSchedulingV1beta1PriorityClassUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/scheduling.k8s.io/v1beta1/priorityclasses/{name}][%d] deleteSchedulingV1beta1PriorityClassUnauthorized ", 401)
}

func (o *DeleteSchedulingV1beta1PriorityClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}