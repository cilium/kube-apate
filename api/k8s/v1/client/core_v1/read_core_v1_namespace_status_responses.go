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

// ReadCoreV1NamespaceStatusReader is a Reader for the ReadCoreV1NamespaceStatus structure.
type ReadCoreV1NamespaceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1NamespaceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1NamespaceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1NamespaceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1NamespaceStatusOK creates a ReadCoreV1NamespaceStatusOK with default headers values
func NewReadCoreV1NamespaceStatusOK() *ReadCoreV1NamespaceStatusOK {
	return &ReadCoreV1NamespaceStatusOK{}
}

/*ReadCoreV1NamespaceStatusOK handles this case with default header values.

OK
*/
type ReadCoreV1NamespaceStatusOK struct {
	Payload *models.IoK8sAPICoreV1Namespace
}

func (o *ReadCoreV1NamespaceStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{name}/status][%d] readCoreV1NamespaceStatusOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1NamespaceStatusOK) GetPayload() *models.IoK8sAPICoreV1Namespace {
	return o.Payload
}

func (o *ReadCoreV1NamespaceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Namespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1NamespaceStatusUnauthorized creates a ReadCoreV1NamespaceStatusUnauthorized with default headers values
func NewReadCoreV1NamespaceStatusUnauthorized() *ReadCoreV1NamespaceStatusUnauthorized {
	return &ReadCoreV1NamespaceStatusUnauthorized{}
}

/*ReadCoreV1NamespaceStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1NamespaceStatusUnauthorized struct {
}

func (o *ReadCoreV1NamespaceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{name}/status][%d] readCoreV1NamespaceStatusUnauthorized ", 401)
}

func (o *ReadCoreV1NamespaceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}