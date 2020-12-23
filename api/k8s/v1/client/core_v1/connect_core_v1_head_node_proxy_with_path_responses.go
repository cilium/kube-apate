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

// ConnectCoreV1HeadNodeProxyWithPathReader is a Reader for the ConnectCoreV1HeadNodeProxyWithPath structure.
type ConnectCoreV1HeadNodeProxyWithPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1HeadNodeProxyWithPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1HeadNodeProxyWithPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1HeadNodeProxyWithPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1HeadNodeProxyWithPathOK creates a ConnectCoreV1HeadNodeProxyWithPathOK with default headers values
func NewConnectCoreV1HeadNodeProxyWithPathOK() *ConnectCoreV1HeadNodeProxyWithPathOK {
	return &ConnectCoreV1HeadNodeProxyWithPathOK{}
}

/*ConnectCoreV1HeadNodeProxyWithPathOK handles this case with default header values.

OK
*/
type ConnectCoreV1HeadNodeProxyWithPathOK struct {
	Payload string
}

func (o *ConnectCoreV1HeadNodeProxyWithPathOK) Error() string {
	return fmt.Sprintf("[HEAD /api/v1/nodes/{name}/proxy/{path}][%d] connectCoreV1HeadNodeProxyWithPathOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1HeadNodeProxyWithPathOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1HeadNodeProxyWithPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1HeadNodeProxyWithPathUnauthorized creates a ConnectCoreV1HeadNodeProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1HeadNodeProxyWithPathUnauthorized() *ConnectCoreV1HeadNodeProxyWithPathUnauthorized {
	return &ConnectCoreV1HeadNodeProxyWithPathUnauthorized{}
}

/*ConnectCoreV1HeadNodeProxyWithPathUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1HeadNodeProxyWithPathUnauthorized struct {
}

func (o *ConnectCoreV1HeadNodeProxyWithPathUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /api/v1/nodes/{name}/proxy/{path}][%d] connectCoreV1HeadNodeProxyWithPathUnauthorized ", 401)
}

func (o *ConnectCoreV1HeadNodeProxyWithPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}