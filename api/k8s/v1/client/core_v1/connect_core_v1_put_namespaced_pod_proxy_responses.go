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
)

// ConnectCoreV1PutNamespacedPodProxyReader is a Reader for the ConnectCoreV1PutNamespacedPodProxy structure.
type ConnectCoreV1PutNamespacedPodProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1PutNamespacedPodProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1PutNamespacedPodProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1PutNamespacedPodProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1PutNamespacedPodProxyOK creates a ConnectCoreV1PutNamespacedPodProxyOK with default headers values
func NewConnectCoreV1PutNamespacedPodProxyOK() *ConnectCoreV1PutNamespacedPodProxyOK {
	return &ConnectCoreV1PutNamespacedPodProxyOK{}
}

/*ConnectCoreV1PutNamespacedPodProxyOK handles this case with default header values.

OK
*/
type ConnectCoreV1PutNamespacedPodProxyOK struct {
	Payload string
}

func (o *ConnectCoreV1PutNamespacedPodProxyOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/pods/{name}/proxy][%d] connectCoreV1PutNamespacedPodProxyOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1PutNamespacedPodProxyOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1PutNamespacedPodProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1PutNamespacedPodProxyUnauthorized creates a ConnectCoreV1PutNamespacedPodProxyUnauthorized with default headers values
func NewConnectCoreV1PutNamespacedPodProxyUnauthorized() *ConnectCoreV1PutNamespacedPodProxyUnauthorized {
	return &ConnectCoreV1PutNamespacedPodProxyUnauthorized{}
}

/*ConnectCoreV1PutNamespacedPodProxyUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1PutNamespacedPodProxyUnauthorized struct {
}

func (o *ConnectCoreV1PutNamespacedPodProxyUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/pods/{name}/proxy][%d] connectCoreV1PutNamespacedPodProxyUnauthorized ", 401)
}

func (o *ConnectCoreV1PutNamespacedPodProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}