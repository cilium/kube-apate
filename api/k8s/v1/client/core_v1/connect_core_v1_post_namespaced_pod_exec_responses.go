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

// ConnectCoreV1PostNamespacedPodExecReader is a Reader for the ConnectCoreV1PostNamespacedPodExec structure.
type ConnectCoreV1PostNamespacedPodExecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1PostNamespacedPodExecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1PostNamespacedPodExecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1PostNamespacedPodExecUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1PostNamespacedPodExecOK creates a ConnectCoreV1PostNamespacedPodExecOK with default headers values
func NewConnectCoreV1PostNamespacedPodExecOK() *ConnectCoreV1PostNamespacedPodExecOK {
	return &ConnectCoreV1PostNamespacedPodExecOK{}
}

/*ConnectCoreV1PostNamespacedPodExecOK handles this case with default header values.

OK
*/
type ConnectCoreV1PostNamespacedPodExecOK struct {
	Payload string
}

func (o *ConnectCoreV1PostNamespacedPodExecOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/exec][%d] connectCoreV1PostNamespacedPodExecOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1PostNamespacedPodExecOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1PostNamespacedPodExecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1PostNamespacedPodExecUnauthorized creates a ConnectCoreV1PostNamespacedPodExecUnauthorized with default headers values
func NewConnectCoreV1PostNamespacedPodExecUnauthorized() *ConnectCoreV1PostNamespacedPodExecUnauthorized {
	return &ConnectCoreV1PostNamespacedPodExecUnauthorized{}
}

/*ConnectCoreV1PostNamespacedPodExecUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1PostNamespacedPodExecUnauthorized struct {
}

func (o *ConnectCoreV1PostNamespacedPodExecUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/exec][%d] connectCoreV1PostNamespacedPodExecUnauthorized ", 401)
}

func (o *ConnectCoreV1PostNamespacedPodExecUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}