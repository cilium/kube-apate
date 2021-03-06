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

// ConnectCoreV1DeleteNamespacedPodProxyWithPathReader is a Reader for the ConnectCoreV1DeleteNamespacedPodProxyWithPath structure.
type ConnectCoreV1DeleteNamespacedPodProxyWithPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1DeleteNamespacedPodProxyWithPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1DeleteNamespacedPodProxyWithPathOK creates a ConnectCoreV1DeleteNamespacedPodProxyWithPathOK with default headers values
func NewConnectCoreV1DeleteNamespacedPodProxyWithPathOK() *ConnectCoreV1DeleteNamespacedPodProxyWithPathOK {
	return &ConnectCoreV1DeleteNamespacedPodProxyWithPathOK{}
}

/*ConnectCoreV1DeleteNamespacedPodProxyWithPathOK handles this case with default header values.

OK
*/
type ConnectCoreV1DeleteNamespacedPodProxyWithPathOK struct {
	Payload string
}

func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}][%d] connectCoreV1DeleteNamespacedPodProxyWithPathOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized creates a ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized() *ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized {
	return &ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized{}
}

/*ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized struct {
}

func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}][%d] connectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized ", 401)
}

func (o *ConnectCoreV1DeleteNamespacedPodProxyWithPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
