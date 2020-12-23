// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListNetworkingV1beta1IngressForAllNamespacesReader is a Reader for the ListNetworkingV1beta1IngressForAllNamespaces structure.
type ListNetworkingV1beta1IngressForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNetworkingV1beta1IngressForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNetworkingV1beta1IngressForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNetworkingV1beta1IngressForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListNetworkingV1beta1IngressForAllNamespacesOK creates a ListNetworkingV1beta1IngressForAllNamespacesOK with default headers values
func NewListNetworkingV1beta1IngressForAllNamespacesOK() *ListNetworkingV1beta1IngressForAllNamespacesOK {
	return &ListNetworkingV1beta1IngressForAllNamespacesOK{}
}

/*ListNetworkingV1beta1IngressForAllNamespacesOK handles this case with default header values.

OK
*/
type ListNetworkingV1beta1IngressForAllNamespacesOK struct {
	Payload *models.IoK8sAPINetworkingV1beta1IngressList
}

func (o *ListNetworkingV1beta1IngressForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1beta1/ingresses][%d] listNetworkingV1beta1IngressForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListNetworkingV1beta1IngressForAllNamespacesOK) GetPayload() *models.IoK8sAPINetworkingV1beta1IngressList {
	return o.Payload
}

func (o *ListNetworkingV1beta1IngressForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1beta1IngressList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNetworkingV1beta1IngressForAllNamespacesUnauthorized creates a ListNetworkingV1beta1IngressForAllNamespacesUnauthorized with default headers values
func NewListNetworkingV1beta1IngressForAllNamespacesUnauthorized() *ListNetworkingV1beta1IngressForAllNamespacesUnauthorized {
	return &ListNetworkingV1beta1IngressForAllNamespacesUnauthorized{}
}

/*ListNetworkingV1beta1IngressForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListNetworkingV1beta1IngressForAllNamespacesUnauthorized struct {
}

func (o *ListNetworkingV1beta1IngressForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1beta1/ingresses][%d] listNetworkingV1beta1IngressForAllNamespacesUnauthorized ", 401)
}

func (o *ListNetworkingV1beta1IngressForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}