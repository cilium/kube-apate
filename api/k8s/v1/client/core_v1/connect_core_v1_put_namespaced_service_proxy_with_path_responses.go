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

// ConnectCoreV1PutNamespacedServiceProxyWithPathReader is a Reader for the ConnectCoreV1PutNamespacedServiceProxyWithPath structure.
type ConnectCoreV1PutNamespacedServiceProxyWithPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1PutNamespacedServiceProxyWithPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1PutNamespacedServiceProxyWithPathOK creates a ConnectCoreV1PutNamespacedServiceProxyWithPathOK with default headers values
func NewConnectCoreV1PutNamespacedServiceProxyWithPathOK() *ConnectCoreV1PutNamespacedServiceProxyWithPathOK {
	return &ConnectCoreV1PutNamespacedServiceProxyWithPathOK{}
}

/*ConnectCoreV1PutNamespacedServiceProxyWithPathOK handles this case with default header values.

OK
*/
type ConnectCoreV1PutNamespacedServiceProxyWithPathOK struct {
	Payload string
}

func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}][%d] connectCoreV1PutNamespacedServiceProxyWithPathOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized creates a ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized() *ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized {
	return &ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized{}
}

/*ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized struct {
}

func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}][%d] connectCoreV1PutNamespacedServiceProxyWithPathUnauthorized ", 401)
}

func (o *ConnectCoreV1PutNamespacedServiceProxyWithPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
