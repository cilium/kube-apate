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

// ReplaceAppsV1NamespacedStatefulSetStatusReader is a Reader for the ReplaceAppsV1NamespacedStatefulSetStatus structure.
type ReplaceAppsV1NamespacedStatefulSetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAppsV1NamespacedStatefulSetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAppsV1NamespacedStatefulSetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAppsV1NamespacedStatefulSetStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAppsV1NamespacedStatefulSetStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAppsV1NamespacedStatefulSetStatusOK creates a ReplaceAppsV1NamespacedStatefulSetStatusOK with default headers values
func NewReplaceAppsV1NamespacedStatefulSetStatusOK() *ReplaceAppsV1NamespacedStatefulSetStatusOK {
	return &ReplaceAppsV1NamespacedStatefulSetStatusOK{}
}

/*ReplaceAppsV1NamespacedStatefulSetStatusOK handles this case with default header values.

OK
*/
type ReplaceAppsV1NamespacedStatefulSetStatusOK struct {
	Payload *models.IoK8sAPIAppsV1StatefulSet
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusOK) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status][%d] replaceAppsV1NamespacedStatefulSetStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusOK) GetPayload() *models.IoK8sAPIAppsV1StatefulSet {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1StatefulSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedStatefulSetStatusCreated creates a ReplaceAppsV1NamespacedStatefulSetStatusCreated with default headers values
func NewReplaceAppsV1NamespacedStatefulSetStatusCreated() *ReplaceAppsV1NamespacedStatefulSetStatusCreated {
	return &ReplaceAppsV1NamespacedStatefulSetStatusCreated{}
}

/*ReplaceAppsV1NamespacedStatefulSetStatusCreated handles this case with default header values.

Created
*/
type ReplaceAppsV1NamespacedStatefulSetStatusCreated struct {
	Payload *models.IoK8sAPIAppsV1StatefulSet
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status][%d] replaceAppsV1NamespacedStatefulSetStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusCreated) GetPayload() *models.IoK8sAPIAppsV1StatefulSet {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1StatefulSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedStatefulSetStatusUnauthorized creates a ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized with default headers values
func NewReplaceAppsV1NamespacedStatefulSetStatusUnauthorized() *ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized {
	return &ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized{}
}

/*ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized struct {
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/status][%d] replaceAppsV1NamespacedStatefulSetStatusUnauthorized ", 401)
}

func (o *ReplaceAppsV1NamespacedStatefulSetStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
