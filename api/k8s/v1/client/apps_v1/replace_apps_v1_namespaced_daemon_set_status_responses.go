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

// ReplaceAppsV1NamespacedDaemonSetStatusReader is a Reader for the ReplaceAppsV1NamespacedDaemonSetStatus structure.
type ReplaceAppsV1NamespacedDaemonSetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAppsV1NamespacedDaemonSetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAppsV1NamespacedDaemonSetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAppsV1NamespacedDaemonSetStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAppsV1NamespacedDaemonSetStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAppsV1NamespacedDaemonSetStatusOK creates a ReplaceAppsV1NamespacedDaemonSetStatusOK with default headers values
func NewReplaceAppsV1NamespacedDaemonSetStatusOK() *ReplaceAppsV1NamespacedDaemonSetStatusOK {
	return &ReplaceAppsV1NamespacedDaemonSetStatusOK{}
}

/*ReplaceAppsV1NamespacedDaemonSetStatusOK handles this case with default header values.

OK
*/
type ReplaceAppsV1NamespacedDaemonSetStatusOK struct {
	Payload *models.IoK8sAPIAppsV1DaemonSet
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusOK) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status][%d] replaceAppsV1NamespacedDaemonSetStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusOK) GetPayload() *models.IoK8sAPIAppsV1DaemonSet {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedDaemonSetStatusCreated creates a ReplaceAppsV1NamespacedDaemonSetStatusCreated with default headers values
func NewReplaceAppsV1NamespacedDaemonSetStatusCreated() *ReplaceAppsV1NamespacedDaemonSetStatusCreated {
	return &ReplaceAppsV1NamespacedDaemonSetStatusCreated{}
}

/*ReplaceAppsV1NamespacedDaemonSetStatusCreated handles this case with default header values.

Created
*/
type ReplaceAppsV1NamespacedDaemonSetStatusCreated struct {
	Payload *models.IoK8sAPIAppsV1DaemonSet
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status][%d] replaceAppsV1NamespacedDaemonSetStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusCreated) GetPayload() *models.IoK8sAPIAppsV1DaemonSet {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedDaemonSetStatusUnauthorized creates a ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized with default headers values
func NewReplaceAppsV1NamespacedDaemonSetStatusUnauthorized() *ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized {
	return &ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized{}
}

/*ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized struct {
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/daemonsets/{name}/status][%d] replaceAppsV1NamespacedDaemonSetStatusUnauthorized ", 401)
}

func (o *ReplaceAppsV1NamespacedDaemonSetStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}