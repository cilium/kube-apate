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

// ReadNetworkingV1IngressClassReader is a Reader for the ReadNetworkingV1IngressClass structure.
type ReadNetworkingV1IngressClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadNetworkingV1IngressClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadNetworkingV1IngressClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadNetworkingV1IngressClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadNetworkingV1IngressClassOK creates a ReadNetworkingV1IngressClassOK with default headers values
func NewReadNetworkingV1IngressClassOK() *ReadNetworkingV1IngressClassOK {
	return &ReadNetworkingV1IngressClassOK{}
}

/*ReadNetworkingV1IngressClassOK handles this case with default header values.

OK
*/
type ReadNetworkingV1IngressClassOK struct {
	Payload *models.IoK8sAPINetworkingV1IngressClass
}

func (o *ReadNetworkingV1IngressClassOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] readNetworkingV1IngressClassOK  %+v", 200, o.Payload)
}

func (o *ReadNetworkingV1IngressClassOK) GetPayload() *models.IoK8sAPINetworkingV1IngressClass {
	return o.Payload
}

func (o *ReadNetworkingV1IngressClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1IngressClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNetworkingV1IngressClassUnauthorized creates a ReadNetworkingV1IngressClassUnauthorized with default headers values
func NewReadNetworkingV1IngressClassUnauthorized() *ReadNetworkingV1IngressClassUnauthorized {
	return &ReadNetworkingV1IngressClassUnauthorized{}
}

/*ReadNetworkingV1IngressClassUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadNetworkingV1IngressClassUnauthorized struct {
}

func (o *ReadNetworkingV1IngressClassUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] readNetworkingV1IngressClassUnauthorized ", 401)
}

func (o *ReadNetworkingV1IngressClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
