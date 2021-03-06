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

// ReadCoreV1NamespaceReader is a Reader for the ReadCoreV1Namespace structure.
type ReadCoreV1NamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1NamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1NamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1NamespaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1NamespaceOK creates a ReadCoreV1NamespaceOK with default headers values
func NewReadCoreV1NamespaceOK() *ReadCoreV1NamespaceOK {
	return &ReadCoreV1NamespaceOK{}
}

/*ReadCoreV1NamespaceOK handles this case with default header values.

OK
*/
type ReadCoreV1NamespaceOK struct {
	Payload *models.IoK8sAPICoreV1Namespace
}

func (o *ReadCoreV1NamespaceOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{name}][%d] readCoreV1NamespaceOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1NamespaceOK) GetPayload() *models.IoK8sAPICoreV1Namespace {
	return o.Payload
}

func (o *ReadCoreV1NamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Namespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1NamespaceUnauthorized creates a ReadCoreV1NamespaceUnauthorized with default headers values
func NewReadCoreV1NamespaceUnauthorized() *ReadCoreV1NamespaceUnauthorized {
	return &ReadCoreV1NamespaceUnauthorized{}
}

/*ReadCoreV1NamespaceUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1NamespaceUnauthorized struct {
}

func (o *ReadCoreV1NamespaceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{name}][%d] readCoreV1NamespaceUnauthorized ", 401)
}

func (o *ReadCoreV1NamespaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
