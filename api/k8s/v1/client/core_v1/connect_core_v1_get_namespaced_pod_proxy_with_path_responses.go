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

// ConnectCoreV1GetNamespacedPodProxyWithPathReader is a Reader for the ConnectCoreV1GetNamespacedPodProxyWithPath structure.
type ConnectCoreV1GetNamespacedPodProxyWithPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1GetNamespacedPodProxyWithPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1GetNamespacedPodProxyWithPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1GetNamespacedPodProxyWithPathOK creates a ConnectCoreV1GetNamespacedPodProxyWithPathOK with default headers values
func NewConnectCoreV1GetNamespacedPodProxyWithPathOK() *ConnectCoreV1GetNamespacedPodProxyWithPathOK {
	return &ConnectCoreV1GetNamespacedPodProxyWithPathOK{}
}

/*ConnectCoreV1GetNamespacedPodProxyWithPathOK handles this case with default header values.

OK
*/
type ConnectCoreV1GetNamespacedPodProxyWithPathOK struct {
	Payload string
}

func (o *ConnectCoreV1GetNamespacedPodProxyWithPathOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}][%d] connectCoreV1GetNamespacedPodProxyWithPathOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1GetNamespacedPodProxyWithPathOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1GetNamespacedPodProxyWithPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized creates a ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized() *ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized {
	return &ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized{}
}

/*ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized struct {
}

func (o *ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}][%d] connectCoreV1GetNamespacedPodProxyWithPathUnauthorized ", 401)
}

func (o *ConnectCoreV1GetNamespacedPodProxyWithPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
