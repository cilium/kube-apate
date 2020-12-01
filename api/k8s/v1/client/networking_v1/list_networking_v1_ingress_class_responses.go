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

// ListNetworkingV1IngressClassReader is a Reader for the ListNetworkingV1IngressClass structure.
type ListNetworkingV1IngressClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNetworkingV1IngressClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNetworkingV1IngressClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNetworkingV1IngressClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListNetworkingV1IngressClassOK creates a ListNetworkingV1IngressClassOK with default headers values
func NewListNetworkingV1IngressClassOK() *ListNetworkingV1IngressClassOK {
	return &ListNetworkingV1IngressClassOK{}
}

/*ListNetworkingV1IngressClassOK handles this case with default header values.

OK
*/
type ListNetworkingV1IngressClassOK struct {
	Payload *models.IoK8sAPINetworkingV1IngressClassList
}

func (o *ListNetworkingV1IngressClassOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/ingressclasses][%d] listNetworkingV1IngressClassOK  %+v", 200, o.Payload)
}

func (o *ListNetworkingV1IngressClassOK) GetPayload() *models.IoK8sAPINetworkingV1IngressClassList {
	return o.Payload
}

func (o *ListNetworkingV1IngressClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1IngressClassList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNetworkingV1IngressClassUnauthorized creates a ListNetworkingV1IngressClassUnauthorized with default headers values
func NewListNetworkingV1IngressClassUnauthorized() *ListNetworkingV1IngressClassUnauthorized {
	return &ListNetworkingV1IngressClassUnauthorized{}
}

/*ListNetworkingV1IngressClassUnauthorized handles this case with default header values.

Unauthorized
*/
type ListNetworkingV1IngressClassUnauthorized struct {
}

func (o *ListNetworkingV1IngressClassUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/ingressclasses][%d] listNetworkingV1IngressClassUnauthorized ", 401)
}

func (o *ListNetworkingV1IngressClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
