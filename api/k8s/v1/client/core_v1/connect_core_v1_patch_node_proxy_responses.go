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

// ConnectCoreV1PatchNodeProxyReader is a Reader for the ConnectCoreV1PatchNodeProxy structure.
type ConnectCoreV1PatchNodeProxyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1PatchNodeProxyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1PatchNodeProxyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1PatchNodeProxyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1PatchNodeProxyOK creates a ConnectCoreV1PatchNodeProxyOK with default headers values
func NewConnectCoreV1PatchNodeProxyOK() *ConnectCoreV1PatchNodeProxyOK {
	return &ConnectCoreV1PatchNodeProxyOK{}
}

/*ConnectCoreV1PatchNodeProxyOK handles this case with default header values.

OK
*/
type ConnectCoreV1PatchNodeProxyOK struct {
	Payload string
}

func (o *ConnectCoreV1PatchNodeProxyOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/nodes/{name}/proxy][%d] connectCoreV1PatchNodeProxyOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1PatchNodeProxyOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1PatchNodeProxyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1PatchNodeProxyUnauthorized creates a ConnectCoreV1PatchNodeProxyUnauthorized with default headers values
func NewConnectCoreV1PatchNodeProxyUnauthorized() *ConnectCoreV1PatchNodeProxyUnauthorized {
	return &ConnectCoreV1PatchNodeProxyUnauthorized{}
}

/*ConnectCoreV1PatchNodeProxyUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1PatchNodeProxyUnauthorized struct {
}

func (o *ConnectCoreV1PatchNodeProxyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/nodes/{name}/proxy][%d] connectCoreV1PatchNodeProxyUnauthorized ", 401)
}

func (o *ConnectCoreV1PatchNodeProxyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
