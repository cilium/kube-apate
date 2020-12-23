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

// ConnectCoreV1PatchNamespacedPodProxyWithPathReader is a Reader for the ConnectCoreV1PatchNamespacedPodProxyWithPath structure.
type ConnectCoreV1PatchNamespacedPodProxyWithPathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectCoreV1PatchNamespacedPodProxyWithPathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConnectCoreV1PatchNamespacedPodProxyWithPathOK creates a ConnectCoreV1PatchNamespacedPodProxyWithPathOK with default headers values
func NewConnectCoreV1PatchNamespacedPodProxyWithPathOK() *ConnectCoreV1PatchNamespacedPodProxyWithPathOK {
	return &ConnectCoreV1PatchNamespacedPodProxyWithPathOK{}
}

/*ConnectCoreV1PatchNamespacedPodProxyWithPathOK handles this case with default header values.

OK
*/
type ConnectCoreV1PatchNamespacedPodProxyWithPathOK struct {
	Payload string
}

func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}][%d] connectCoreV1PatchNamespacedPodProxyWithPathOK  %+v", 200, o.Payload)
}

func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathOK) GetPayload() string {
	return o.Payload
}

func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized creates a ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized with default headers values
func NewConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized() *ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized {
	return &ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized{}
}

/*ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized handles this case with default header values.

Unauthorized
*/
type ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized struct {
}

func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/pods/{name}/proxy/{path}][%d] connectCoreV1PatchNamespacedPodProxyWithPathUnauthorized ", 401)
}

func (o *ConnectCoreV1PatchNamespacedPodProxyWithPathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}