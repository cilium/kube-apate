// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadAppsV1NamespacedDaemonSetStatusReader is a Reader for the ReadAppsV1NamespacedDaemonSetStatus structure.
type ReadAppsV1NamespacedDaemonSetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadAppsV1NamespacedDaemonSetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadAppsV1NamespacedDaemonSetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadAppsV1NamespacedDaemonSetStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadAppsV1NamespacedDaemonSetStatusOK creates a ReadAppsV1NamespacedDaemonSetStatusOK with default headers values
func NewReadAppsV1NamespacedDaemonSetStatusOK() *ReadAppsV1NamespacedDaemonSetStatusOK {
	return &ReadAppsV1NamespacedDaemonSetStatusOK{}
}

/*ReadAppsV1NamespacedDaemonSetStatusOK handles this case with default header values.

OK
*/
type ReadAppsV1NamespacedDaemonSetStatusOK struct {
	Payload *models.IoK8sAPIAppsV1DaemonSet
}

func (o *ReadAppsV1NamespacedDaemonSetStatusOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status][%d] readAppsV1NamespacedDaemonSetStatusOK  %+v", 200, o.Payload)
}

func (o *ReadAppsV1NamespacedDaemonSetStatusOK) GetPayload() *models.IoK8sAPIAppsV1DaemonSet {
	return o.Payload
}

func (o *ReadAppsV1NamespacedDaemonSetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadAppsV1NamespacedDaemonSetStatusUnauthorized creates a ReadAppsV1NamespacedDaemonSetStatusUnauthorized with default headers values
func NewReadAppsV1NamespacedDaemonSetStatusUnauthorized() *ReadAppsV1NamespacedDaemonSetStatusUnauthorized {
	return &ReadAppsV1NamespacedDaemonSetStatusUnauthorized{}
}

/*ReadAppsV1NamespacedDaemonSetStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadAppsV1NamespacedDaemonSetStatusUnauthorized struct {
}

func (o *ReadAppsV1NamespacedDaemonSetStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status][%d] readAppsV1NamespacedDaemonSetStatusUnauthorized ", 401)
}

func (o *ReadAppsV1NamespacedDaemonSetStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
