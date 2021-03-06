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

// ConnectCoreV1PatchNamespacedServiceProxyWithPathReader is a Reader for the ConnectCoreV1PatchNamespacedServiceProxyWithPath structure.
type ConnectCoreV1PatchNamespacedServiceProxyWithPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1PatchNamespacedServiceProxyWithPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1PatchNamespacedServiceProxyWithPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1PatchNamespacedServiceProxyWithPathOK creates a ConnectCoreV1PatchNamespacedServiceProxyWithPathOK with default headers values
func NewConnectCoreV1PatchNamespacedServiceProxyWithPathOK() *ConnectCoreV1PatchNamespacedServiceProxyWithPathOK {
	return &ConnectCoreV1PatchNamespacedServiceProxyWithPathOK{}
}

/*ConnectCoreV1PatchNamespacedServiceProxyWithPathOK handles this case with default header values.

OK
*/
type ConnectCoreV1PatchNamespacedServiceProxyWithPathOK struct {
	Payload string
}

func (o *ConnectCoreV1PatchNamespacedServiceProxyWithPathOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}][%d] connectCoreV1PatchNamespacedServiceProxyWithPathOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1PatchNamespacedServiceProxyWithPathOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1PatchNamespacedServiceProxyWithPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized creates a ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized() *ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized {
	return &ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized{}
}

/*ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized struct {
}

func (o *ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/services/{name}/proxy/{path}][%d] connectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized ", 401)
}

func (o *ConnectCoreV1PatchNamespacedServiceProxyWithPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
