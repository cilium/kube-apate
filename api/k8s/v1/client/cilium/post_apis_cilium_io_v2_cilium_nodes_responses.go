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

// PostApisCiliumIoV2CiliumNodesReader is a Reader for the PostApisCiliumIoV2CiliumNodes structure.
type PostApisCiliumIoV2CiliumNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApisCiliumIoV2CiliumNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApisCiliumIoV2CiliumNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPostApisCiliumIoV2CiliumNodesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostApisCiliumIoV2CiliumNodesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostApisCiliumIoV2CiliumNodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostApisCiliumIoV2CiliumNodesOK creates a PostApisCiliumIoV2CiliumNodesOK with default headers values
func NewPostApisCiliumIoV2CiliumNodesOK() *PostApisCiliumIoV2CiliumNodesOK {
	return &PostApisCiliumIoV2CiliumNodesOK{}
}

/*PostApisCiliumIoV2CiliumNodesOK handles this case with default header values.

OK
*/
type PostApisCiliumIoV2CiliumNodesOK struct {
	Payload interface{}
}

func (o *PostApisCiliumIoV2CiliumNodesOK) Error() string {
	return fmt.Sprintf("[POST /apis/cilium.io/v2/ciliumnodes][%d] postApisCiliumIoV2CiliumNodesOK  %+v", 200, o.Payload)
}

func (o *PostApisCiliumIoV2CiliumNodesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostApisCiliumIoV2CiliumNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApisCiliumIoV2CiliumNodesCreated creates a PostApisCiliumIoV2CiliumNodesCreated with default headers values
func NewPostApisCiliumIoV2CiliumNodesCreated() *PostApisCiliumIoV2CiliumNodesCreated {
	return &PostApisCiliumIoV2CiliumNodesCreated{}
}

/*PostApisCiliumIoV2CiliumNodesCreated handles this case with default header values.

Created
*/
type PostApisCiliumIoV2CiliumNodesCreated struct {
	Payload interface{}
}

func (o *PostApisCiliumIoV2CiliumNodesCreated) Error() string {
	return fmt.Sprintf("[POST /apis/cilium.io/v2/ciliumnodes][%d] postApisCiliumIoV2CiliumNodesCreated  %+v", 201, o.Payload)
}

func (o *PostApisCiliumIoV2CiliumNodesCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *PostApisCiliumIoV2CiliumNodesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApisCiliumIoV2CiliumNodesAccepted creates a PostApisCiliumIoV2CiliumNodesAccepted with default headers values
func NewPostApisCiliumIoV2CiliumNodesAccepted() *PostApisCiliumIoV2CiliumNodesAccepted {
	return &PostApisCiliumIoV2CiliumNodesAccepted{}
}

/*PostApisCiliumIoV2CiliumNodesAccepted handles this case with default header values.

Accepted
*/
type PostApisCiliumIoV2CiliumNodesAccepted struct {
	Payload interface{}
}

func (o *PostApisCiliumIoV2CiliumNodesAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/cilium.io/v2/ciliumnodes][%d] postApisCiliumIoV2CiliumNodesAccepted  %+v", 202, o.Payload)
}

func (o *PostApisCiliumIoV2CiliumNodesAccepted) GetPayload() interface{} {
	return o.Payload
}

func (o *PostApisCiliumIoV2CiliumNodesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApisCiliumIoV2CiliumNodesUnauthorized creates a PostApisCiliumIoV2CiliumNodesUnauthorized with default headers values
func NewPostApisCiliumIoV2CiliumNodesUnauthorized() *PostApisCiliumIoV2CiliumNodesUnauthorized {
	return &PostApisCiliumIoV2CiliumNodesUnauthorized{}
}

/*PostApisCiliumIoV2CiliumNodesUnauthorized handles this case with default header values.

Unauthorized
*/
type PostApisCiliumIoV2CiliumNodesUnauthorized struct {
}

func (o *PostApisCiliumIoV2CiliumNodesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/cilium.io/v2/ciliumnodes][%d] postApisCiliumIoV2CiliumNodesUnauthorized ", 401)
}

func (o *PostApisCiliumIoV2CiliumNodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
