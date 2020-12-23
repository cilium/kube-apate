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

// ReplaceNetworkingV1NamespacedIngressReader is a Reader for the ReplaceNetworkingV1NamespacedIngress structure.
type ReplaceNetworkingV1NamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceNetworkingV1NamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceNetworkingV1NamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceNetworkingV1NamespacedIngressCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceNetworkingV1NamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceNetworkingV1NamespacedIngressOK creates a ReplaceNetworkingV1NamespacedIngressOK with default headers values
func NewReplaceNetworkingV1NamespacedIngressOK() *ReplaceNetworkingV1NamespacedIngressOK {
	return &ReplaceNetworkingV1NamespacedIngressOK{}
}

/*ReplaceNetworkingV1NamespacedIngressOK handles this case with default header values.

OK
*/
type ReplaceNetworkingV1NamespacedIngressOK struct {
	Payload *models.IoK8sAPINetworkingV1Ingress
}

func (o *ReplaceNetworkingV1NamespacedIngressOK) Error() string {
	return fmt.Sprintf("[PUT /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] replaceNetworkingV1NamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *ReplaceNetworkingV1NamespacedIngressOK) GetPayload() *models.IoK8sAPINetworkingV1Ingress {
	return o.Payload
}

func (o *ReplaceNetworkingV1NamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNetworkingV1NamespacedIngressCreated creates a ReplaceNetworkingV1NamespacedIngressCreated with default headers values
func NewReplaceNetworkingV1NamespacedIngressCreated() *ReplaceNetworkingV1NamespacedIngressCreated {
	return &ReplaceNetworkingV1NamespacedIngressCreated{}
}

/*ReplaceNetworkingV1NamespacedIngressCreated handles this case with default header values.

Created
*/
type ReplaceNetworkingV1NamespacedIngressCreated struct {
	Payload *models.IoK8sAPINetworkingV1Ingress
}

func (o *ReplaceNetworkingV1NamespacedIngressCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] replaceNetworkingV1NamespacedIngressCreated  %+v", 201, o.Payload)
}

func (o *ReplaceNetworkingV1NamespacedIngressCreated) GetPayload() *models.IoK8sAPINetworkingV1Ingress {
	return o.Payload
}

func (o *ReplaceNetworkingV1NamespacedIngressCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNetworkingV1NamespacedIngressUnauthorized creates a ReplaceNetworkingV1NamespacedIngressUnauthorized with default headers values
func NewReplaceNetworkingV1NamespacedIngressUnauthorized() *ReplaceNetworkingV1NamespacedIngressUnauthorized {
	return &ReplaceNetworkingV1NamespacedIngressUnauthorized{}
}

/*ReplaceNetworkingV1NamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceNetworkingV1NamespacedIngressUnauthorized struct {
}

func (o *ReplaceNetworkingV1NamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] replaceNetworkingV1NamespacedIngressUnauthorized ", 401)
}

func (o *ReplaceNetworkingV1NamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}