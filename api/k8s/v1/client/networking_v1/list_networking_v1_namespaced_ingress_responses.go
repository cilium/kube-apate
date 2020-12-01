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

// ListNetworkingV1NamespacedIngressReader is a Reader for the ListNetworkingV1NamespacedIngress structure.
type ListNetworkingV1NamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNetworkingV1NamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNetworkingV1NamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListNetworkingV1NamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListNetworkingV1NamespacedIngressOK creates a ListNetworkingV1NamespacedIngressOK with default headers values
func NewListNetworkingV1NamespacedIngressOK() *ListNetworkingV1NamespacedIngressOK {
	return &ListNetworkingV1NamespacedIngressOK{}
}

/*ListNetworkingV1NamespacedIngressOK handles this case with default header values.

OK
*/
type ListNetworkingV1NamespacedIngressOK struct {
	Payload *models.IoK8sAPINetworkingV1IngressList
}

func (o *ListNetworkingV1NamespacedIngressOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses][%d] listNetworkingV1NamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *ListNetworkingV1NamespacedIngressOK) GetPayload() *models.IoK8sAPINetworkingV1IngressList {
	return o.Payload
}

func (o *ListNetworkingV1NamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1IngressList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNetworkingV1NamespacedIngressUnauthorized creates a ListNetworkingV1NamespacedIngressUnauthorized with default headers values
func NewListNetworkingV1NamespacedIngressUnauthorized() *ListNetworkingV1NamespacedIngressUnauthorized {
	return &ListNetworkingV1NamespacedIngressUnauthorized{}
}

/*ListNetworkingV1NamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type ListNetworkingV1NamespacedIngressUnauthorized struct {
}

func (o *ListNetworkingV1NamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses][%d] listNetworkingV1NamespacedIngressUnauthorized ", 401)
}

func (o *ListNetworkingV1NamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
