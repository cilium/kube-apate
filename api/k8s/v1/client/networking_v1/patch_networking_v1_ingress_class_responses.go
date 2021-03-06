// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchNetworkingV1IngressClassReader is a Reader for the PatchNetworkingV1IngressClass structure.
type PatchNetworkingV1IngressClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchNetworkingV1IngressClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchNetworkingV1IngressClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchNetworkingV1IngressClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchNetworkingV1IngressClassOK creates a PatchNetworkingV1IngressClassOK with default headers values
func NewPatchNetworkingV1IngressClassOK() *PatchNetworkingV1IngressClassOK {
	return &PatchNetworkingV1IngressClassOK{}
}

/*PatchNetworkingV1IngressClassOK handles this case with default header values.

OK
*/
type PatchNetworkingV1IngressClassOK struct {
	Payload *models.IoK8sAPINetworkingV1IngressClass
}

func (o *PatchNetworkingV1IngressClassOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] patchNetworkingV1IngressClassOK  %+v", 200, o.Payload)
}

func (o *PatchNetworkingV1IngressClassOK) GetPayload() *models.IoK8sAPINetworkingV1IngressClass {
	return o.Payload
}

func (o *PatchNetworkingV1IngressClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1IngressClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchNetworkingV1IngressClassUnauthorized creates a PatchNetworkingV1IngressClassUnauthorized with default headers values
func NewPatchNetworkingV1IngressClassUnauthorized() *PatchNetworkingV1IngressClassUnauthorized {
	return &PatchNetworkingV1IngressClassUnauthorized{}
}

/*PatchNetworkingV1IngressClassUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchNetworkingV1IngressClassUnauthorized struct {
}

func (o *PatchNetworkingV1IngressClassUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] patchNetworkingV1IngressClassUnauthorized ", 401)
}

func (o *PatchNetworkingV1IngressClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
