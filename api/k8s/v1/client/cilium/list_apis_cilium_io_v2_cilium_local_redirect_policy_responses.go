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

// ListApisCiliumIoV2CiliumLocalRedirectPolicyReader is a Reader for the ListApisCiliumIoV2CiliumLocalRedirectPolicy structure.
type ListApisCiliumIoV2CiliumLocalRedirectPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListApisCiliumIoV2CiliumLocalRedirectPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListApisCiliumIoV2CiliumLocalRedirectPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListApisCiliumIoV2CiliumLocalRedirectPolicyOK creates a ListApisCiliumIoV2CiliumLocalRedirectPolicyOK with default headers values
func NewListApisCiliumIoV2CiliumLocalRedirectPolicyOK() *ListApisCiliumIoV2CiliumLocalRedirectPolicyOK {
	return &ListApisCiliumIoV2CiliumLocalRedirectPolicyOK{}
}

/*ListApisCiliumIoV2CiliumLocalRedirectPolicyOK handles this case with default header values.

OK
*/
type ListApisCiliumIoV2CiliumLocalRedirectPolicyOK struct {
	Payload interface{}
}

func (o *ListApisCiliumIoV2CiliumLocalRedirectPolicyOK) Error() string {
	return fmt.Sprintf("[GET /apis/cilium.io/v2/ciliumlocalredirectpolicy][%d] listApisCiliumIoV2CiliumLocalRedirectPolicyOK  %+v", 200, o.Payload)
}

func (o *ListApisCiliumIoV2CiliumLocalRedirectPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ListApisCiliumIoV2CiliumLocalRedirectPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized creates a ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized with default headers values
func NewListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized() *ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized {
	return &ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized{}
}

/*ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized struct {
}

func (o *ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/cilium.io/v2/ciliumlocalredirectpolicy][%d] listApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized ", 401)
}

func (o *ListApisCiliumIoV2CiliumLocalRedirectPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
