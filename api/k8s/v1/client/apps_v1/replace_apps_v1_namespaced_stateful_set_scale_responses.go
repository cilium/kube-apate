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

// ReplaceAppsV1NamespacedStatefulSetScaleReader is a Reader for the ReplaceAppsV1NamespacedStatefulSetScale structure.
type ReplaceAppsV1NamespacedStatefulSetScaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAppsV1NamespacedStatefulSetScaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAppsV1NamespacedStatefulSetScaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAppsV1NamespacedStatefulSetScaleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAppsV1NamespacedStatefulSetScaleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAppsV1NamespacedStatefulSetScaleOK creates a ReplaceAppsV1NamespacedStatefulSetScaleOK with default headers values
func NewReplaceAppsV1NamespacedStatefulSetScaleOK() *ReplaceAppsV1NamespacedStatefulSetScaleOK {
	return &ReplaceAppsV1NamespacedStatefulSetScaleOK{}
}

/*ReplaceAppsV1NamespacedStatefulSetScaleOK handles this case with default header values.

OK
*/
type ReplaceAppsV1NamespacedStatefulSetScaleOK struct {
	Payload *models.IoK8sAPIAutoscalingV1Scale
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleOK) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/scale][%d] replaceAppsV1NamespacedStatefulSetScaleOK  %+v", 200, o.Payload)
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleOK) GetPayload() *models.IoK8sAPIAutoscalingV1Scale {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1Scale)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedStatefulSetScaleCreated creates a ReplaceAppsV1NamespacedStatefulSetScaleCreated with default headers values
func NewReplaceAppsV1NamespacedStatefulSetScaleCreated() *ReplaceAppsV1NamespacedStatefulSetScaleCreated {
	return &ReplaceAppsV1NamespacedStatefulSetScaleCreated{}
}

/*ReplaceAppsV1NamespacedStatefulSetScaleCreated handles this case with default header values.

Created
*/
type ReplaceAppsV1NamespacedStatefulSetScaleCreated struct {
	Payload *models.IoK8sAPIAutoscalingV1Scale
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/scale][%d] replaceAppsV1NamespacedStatefulSetScaleCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleCreated) GetPayload() *models.IoK8sAPIAutoscalingV1Scale {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1Scale)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedStatefulSetScaleUnauthorized creates a ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized with default headers values
func NewReplaceAppsV1NamespacedStatefulSetScaleUnauthorized() *ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized {
	return &ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized{}
}

/*ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized struct {
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/scale][%d] replaceAppsV1NamespacedStatefulSetScaleUnauthorized ", 401)
}

func (o *ReplaceAppsV1NamespacedStatefulSetScaleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
