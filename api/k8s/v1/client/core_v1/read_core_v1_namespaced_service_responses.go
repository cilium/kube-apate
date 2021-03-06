// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadCoreV1NamespacedServiceReader is a Reader for the ReadCoreV1NamespacedService structure.
type ReadCoreV1NamespacedServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1NamespacedServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1NamespacedServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1NamespacedServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1NamespacedServiceOK creates a ReadCoreV1NamespacedServiceOK with default headers values
func NewReadCoreV1NamespacedServiceOK() *ReadCoreV1NamespacedServiceOK {
	return &ReadCoreV1NamespacedServiceOK{}
}

/*ReadCoreV1NamespacedServiceOK handles this case with default header values.

OK
*/
type ReadCoreV1NamespacedServiceOK struct {
	Payload *models.IoK8sAPICoreV1Service
}

func (o *ReadCoreV1NamespacedServiceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/services/{name}][%d] readCoreV1NamespacedServiceOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1NamespacedServiceOK) GetPayload() *models.IoK8sAPICoreV1Service {
	return o.Payload
}

func (o *ReadCoreV1NamespacedServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1NamespacedServiceUnauthorized creates a ReadCoreV1NamespacedServiceUnauthorized with default headers values
func NewReadCoreV1NamespacedServiceUnauthorized() *ReadCoreV1NamespacedServiceUnauthorized {
	return &ReadCoreV1NamespacedServiceUnauthorized{}
}

/*ReadCoreV1NamespacedServiceUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1NamespacedServiceUnauthorized struct {
}

func (o *ReadCoreV1NamespacedServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/services/{name}][%d] readCoreV1NamespacedServiceUnauthorized ", 401)
}

func (o *ReadCoreV1NamespacedServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
