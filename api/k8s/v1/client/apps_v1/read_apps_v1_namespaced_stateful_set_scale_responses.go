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

// ReadAppsV1NamespacedStatefulSetScaleReader is a Reader for the ReadAppsV1NamespacedStatefulSetScale structure.
type ReadAppsV1NamespacedStatefulSetScaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadAppsV1NamespacedStatefulSetScaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadAppsV1NamespacedStatefulSetScaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadAppsV1NamespacedStatefulSetScaleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadAppsV1NamespacedStatefulSetScaleOK creates a ReadAppsV1NamespacedStatefulSetScaleOK with default headers values
func NewReadAppsV1NamespacedStatefulSetScaleOK() *ReadAppsV1NamespacedStatefulSetScaleOK {
	return &ReadAppsV1NamespacedStatefulSetScaleOK{}
}

/*ReadAppsV1NamespacedStatefulSetScaleOK handles this case with default header values.

OK
*/
type ReadAppsV1NamespacedStatefulSetScaleOK struct {
	Payload *models.IoK8sAPIAutoscalingV1Scale
}

func (o *ReadAppsV1NamespacedStatefulSetScaleOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/scale][%d] readAppsV1NamespacedStatefulSetScaleOK  %+v", 200, o.Payload)
}

func (o *ReadAppsV1NamespacedStatefulSetScaleOK) GetPayload() *models.IoK8sAPIAutoscalingV1Scale {
	return o.Payload
}

func (o *ReadAppsV1NamespacedStatefulSetScaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1Scale)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadAppsV1NamespacedStatefulSetScaleUnauthorized creates a ReadAppsV1NamespacedStatefulSetScaleUnauthorized with default headers values
func NewReadAppsV1NamespacedStatefulSetScaleUnauthorized() *ReadAppsV1NamespacedStatefulSetScaleUnauthorized {
	return &ReadAppsV1NamespacedStatefulSetScaleUnauthorized{}
}

/*ReadAppsV1NamespacedStatefulSetScaleUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadAppsV1NamespacedStatefulSetScaleUnauthorized struct {
}

func (o *ReadAppsV1NamespacedStatefulSetScaleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/statefulsets/{name}/scale][%d] readAppsV1NamespacedStatefulSetScaleUnauthorized ", 401)
}

func (o *ReadAppsV1NamespacedStatefulSetScaleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
