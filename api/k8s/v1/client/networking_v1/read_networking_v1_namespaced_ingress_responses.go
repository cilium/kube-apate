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

// ReadNetworkingV1NamespacedIngressReader is a Reader for the ReadNetworkingV1NamespacedIngress structure.
type ReadNetworkingV1NamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadNetworkingV1NamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadNetworkingV1NamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadNetworkingV1NamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadNetworkingV1NamespacedIngressOK creates a ReadNetworkingV1NamespacedIngressOK with default headers values
func NewReadNetworkingV1NamespacedIngressOK() *ReadNetworkingV1NamespacedIngressOK {
	return &ReadNetworkingV1NamespacedIngressOK{}
}

/*ReadNetworkingV1NamespacedIngressOK handles this case with default header values.

OK
*/
type ReadNetworkingV1NamespacedIngressOK struct {
	Payload *models.IoK8sAPINetworkingV1Ingress
}

func (o *ReadNetworkingV1NamespacedIngressOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] readNetworkingV1NamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *ReadNetworkingV1NamespacedIngressOK) GetPayload() *models.IoK8sAPINetworkingV1Ingress {
	return o.Payload
}

func (o *ReadNetworkingV1NamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPINetworkingV1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadNetworkingV1NamespacedIngressUnauthorized creates a ReadNetworkingV1NamespacedIngressUnauthorized with default headers values
func NewReadNetworkingV1NamespacedIngressUnauthorized() *ReadNetworkingV1NamespacedIngressUnauthorized {
	return &ReadNetworkingV1NamespacedIngressUnauthorized{}
}

/*ReadNetworkingV1NamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadNetworkingV1NamespacedIngressUnauthorized struct {
}

func (o *ReadNetworkingV1NamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] readNetworkingV1NamespacedIngressUnauthorized ", 401)
}

func (o *ReadNetworkingV1NamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
