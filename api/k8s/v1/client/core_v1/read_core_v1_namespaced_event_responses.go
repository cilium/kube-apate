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

// ReadCoreV1NamespacedEventReader is a Reader for the ReadCoreV1NamespacedEvent structure.
type ReadCoreV1NamespacedEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1NamespacedEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1NamespacedEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1NamespacedEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1NamespacedEventOK creates a ReadCoreV1NamespacedEventOK with default headers values
func NewReadCoreV1NamespacedEventOK() *ReadCoreV1NamespacedEventOK {
	return &ReadCoreV1NamespacedEventOK{}
}

/*ReadCoreV1NamespacedEventOK handles this case with default header values.

OK
*/
type ReadCoreV1NamespacedEventOK struct {
	Payload *models.IoK8sAPICoreV1Event
}

func (o *ReadCoreV1NamespacedEventOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/events/{name}][%d] readCoreV1NamespacedEventOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1NamespacedEventOK) GetPayload() *models.IoK8sAPICoreV1Event {
	return o.Payload
}

func (o *ReadCoreV1NamespacedEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1NamespacedEventUnauthorized creates a ReadCoreV1NamespacedEventUnauthorized with default headers values
func NewReadCoreV1NamespacedEventUnauthorized() *ReadCoreV1NamespacedEventUnauthorized {
	return &ReadCoreV1NamespacedEventUnauthorized{}
}

/*ReadCoreV1NamespacedEventUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1NamespacedEventUnauthorized struct {
}

func (o *ReadCoreV1NamespacedEventUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/events/{name}][%d] readCoreV1NamespacedEventUnauthorized ", 401)
}

func (o *ReadCoreV1NamespacedEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
