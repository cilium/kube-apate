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

// ConnectCoreV1HeadNodeProxyReader is a Reader for the ConnectCoreV1HeadNodeProxy structure.
type ConnectCoreV1HeadNodeProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1HeadNodeProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1HeadNodeProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1HeadNodeProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1HeadNodeProxyOK creates a ConnectCoreV1HeadNodeProxyOK with default headers values
func NewConnectCoreV1HeadNodeProxyOK() *ConnectCoreV1HeadNodeProxyOK {
	return &ConnectCoreV1HeadNodeProxyOK{}
}

/*ConnectCoreV1HeadNodeProxyOK handles this case with default header values.

OK
*/
type ConnectCoreV1HeadNodeProxyOK struct {
	Payload string
}

func (o *ConnectCoreV1HeadNodeProxyOK) Error() string {
	return fmt.Sprintf("[HEAD /api/v1/nodes/{name}/proxy][%d] connectCoreV1HeadNodeProxyOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1HeadNodeProxyOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1HeadNodeProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1HeadNodeProxyUnauthorized creates a ConnectCoreV1HeadNodeProxyUnauthorized with default headers values
func NewConnectCoreV1HeadNodeProxyUnauthorized() *ConnectCoreV1HeadNodeProxyUnauthorized {
	return &ConnectCoreV1HeadNodeProxyUnauthorized{}
}

/*ConnectCoreV1HeadNodeProxyUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1HeadNodeProxyUnauthorized struct {
}

func (o *ConnectCoreV1HeadNodeProxyUnauthorized) Error() string {
	return fmt.Sprintf("[HEAD /api/v1/nodes/{name}/proxy][%d] connectCoreV1HeadNodeProxyUnauthorized ", 401)
}

func (o *ConnectCoreV1HeadNodeProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
