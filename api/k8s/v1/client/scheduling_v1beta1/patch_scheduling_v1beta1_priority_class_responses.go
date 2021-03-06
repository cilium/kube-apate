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

// PatchSchedulingV1beta1PriorityClassReader is a Reader for the PatchSchedulingV1beta1PriorityClass structure.
type PatchSchedulingV1beta1PriorityClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSchedulingV1beta1PriorityClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSchedulingV1beta1PriorityClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchSchedulingV1beta1PriorityClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchSchedulingV1beta1PriorityClassOK creates a PatchSchedulingV1beta1PriorityClassOK with default headers values
func NewPatchSchedulingV1beta1PriorityClassOK() *PatchSchedulingV1beta1PriorityClassOK {
	return &PatchSchedulingV1beta1PriorityClassOK{}
}

/*PatchSchedulingV1beta1PriorityClassOK handles this case with default header values.

OK
*/
type PatchSchedulingV1beta1PriorityClassOK struct {
	Payload *models.IoK8sAPISchedulingV1beta1PriorityClass
}

func (o *PatchSchedulingV1beta1PriorityClassOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/scheduling.k8s.io/v1beta1/priorityclasses/{name}][%d] patchSchedulingV1beta1PriorityClassOK  %+v", 200, o.Payload)
}

func (o *PatchSchedulingV1beta1PriorityClassOK) GetPayload() *models.IoK8sAPISchedulingV1beta1PriorityClass {
	return o.Payload
}

func (o *PatchSchedulingV1beta1PriorityClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISchedulingV1beta1PriorityClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSchedulingV1beta1PriorityClassUnauthorized creates a PatchSchedulingV1beta1PriorityClassUnauthorized with default headers values
func NewPatchSchedulingV1beta1PriorityClassUnauthorized() *PatchSchedulingV1beta1PriorityClassUnauthorized {
	return &PatchSchedulingV1beta1PriorityClassUnauthorized{}
}

/*PatchSchedulingV1beta1PriorityClassUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchSchedulingV1beta1PriorityClassUnauthorized struct {
}

func (o *PatchSchedulingV1beta1PriorityClassUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/scheduling.k8s.io/v1beta1/priorityclasses/{name}][%d] patchSchedulingV1beta1PriorityClassUnauthorized ", 401)
}

func (o *PatchSchedulingV1beta1PriorityClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
