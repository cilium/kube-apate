// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package cilium

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListApisCiliumIoV2CiliumIdentityReader is a Reader for the ListApisCiliumIoV2CiliumIdentity structure.
type ListApisCiliumIoV2CiliumIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListApisCiliumIoV2CiliumIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListApisCiliumIoV2CiliumIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListApisCiliumIoV2CiliumIdentityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListApisCiliumIoV2CiliumIdentityOK creates a ListApisCiliumIoV2CiliumIdentityOK with default headers values
func NewListApisCiliumIoV2CiliumIdentityOK() *ListApisCiliumIoV2CiliumIdentityOK {
	return &ListApisCiliumIoV2CiliumIdentityOK{}
}

/*ListApisCiliumIoV2CiliumIdentityOK handles this case with default header values.

OK
*/
type ListApisCiliumIoV2CiliumIdentityOK struct {
	Payload interface{}
}

func (o *ListApisCiliumIoV2CiliumIdentityOK) Error() string {
	return fmt.Sprintf("[GET /apis/cilium.io/v2/ciliumidentities][%d] listApisCiliumIoV2CiliumIdentityOK  %+v", 200, o.Payload)
}

func (o *ListApisCiliumIoV2CiliumIdentityOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ListApisCiliumIoV2CiliumIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListApisCiliumIoV2CiliumIdentityUnauthorized creates a ListApisCiliumIoV2CiliumIdentityUnauthorized with default headers values
func NewListApisCiliumIoV2CiliumIdentityUnauthorized() *ListApisCiliumIoV2CiliumIdentityUnauthorized {
	return &ListApisCiliumIoV2CiliumIdentityUnauthorized{}
}

/*ListApisCiliumIoV2CiliumIdentityUnauthorized handles this case with default header values.

Unauthorized
*/
type ListApisCiliumIoV2CiliumIdentityUnauthorized struct {
}

func (o *ListApisCiliumIoV2CiliumIdentityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/cilium.io/v2/ciliumidentities][%d] listApisCiliumIoV2CiliumIdentityUnauthorized ", 401)
}

func (o *ListApisCiliumIoV2CiliumIdentityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}