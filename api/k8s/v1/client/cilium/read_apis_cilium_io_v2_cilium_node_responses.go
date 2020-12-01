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

// ReadApisCiliumIoV2CiliumNodeReader is a Reader for the ReadApisCiliumIoV2CiliumNode structure.
type ReadApisCiliumIoV2CiliumNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadApisCiliumIoV2CiliumNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadApisCiliumIoV2CiliumNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadApisCiliumIoV2CiliumNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadApisCiliumIoV2CiliumNodeOK creates a ReadApisCiliumIoV2CiliumNodeOK with default headers values
func NewReadApisCiliumIoV2CiliumNodeOK() *ReadApisCiliumIoV2CiliumNodeOK {
	return &ReadApisCiliumIoV2CiliumNodeOK{}
}

/*ReadApisCiliumIoV2CiliumNodeOK handles this case with default header values.

OK
*/
type ReadApisCiliumIoV2CiliumNodeOK struct {
	Payload interface{}
}

func (o *ReadApisCiliumIoV2CiliumNodeOK) Error() string {
	return fmt.Sprintf("[GET /apis/cilium.io/v2/ciliumnodes/{name}][%d] readApisCiliumIoV2CiliumNodeOK  %+v", 200, o.Payload)
}

func (o *ReadApisCiliumIoV2CiliumNodeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReadApisCiliumIoV2CiliumNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadApisCiliumIoV2CiliumNodeUnauthorized creates a ReadApisCiliumIoV2CiliumNodeUnauthorized with default headers values
func NewReadApisCiliumIoV2CiliumNodeUnauthorized() *ReadApisCiliumIoV2CiliumNodeUnauthorized {
	return &ReadApisCiliumIoV2CiliumNodeUnauthorized{}
}

/*ReadApisCiliumIoV2CiliumNodeUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadApisCiliumIoV2CiliumNodeUnauthorized struct {
}

func (o *ReadApisCiliumIoV2CiliumNodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/cilium.io/v2/ciliumnodes/{name}][%d] readApisCiliumIoV2CiliumNodeUnauthorized ", 401)
}

func (o *ReadApisCiliumIoV2CiliumNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
