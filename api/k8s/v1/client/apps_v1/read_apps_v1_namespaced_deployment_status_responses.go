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

// ReadAppsV1NamespacedDeploymentStatusReader is a Reader for the ReadAppsV1NamespacedDeploymentStatus structure.
type ReadAppsV1NamespacedDeploymentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadAppsV1NamespacedDeploymentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadAppsV1NamespacedDeploymentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadAppsV1NamespacedDeploymentStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadAppsV1NamespacedDeploymentStatusOK creates a ReadAppsV1NamespacedDeploymentStatusOK with default headers values
func NewReadAppsV1NamespacedDeploymentStatusOK() *ReadAppsV1NamespacedDeploymentStatusOK {
	return &ReadAppsV1NamespacedDeploymentStatusOK{}
}

/*ReadAppsV1NamespacedDeploymentStatusOK handles this case with default header values.

OK
*/
type ReadAppsV1NamespacedDeploymentStatusOK struct {
	Payload *models.IoK8sAPIAppsV1Deployment
}

func (o *ReadAppsV1NamespacedDeploymentStatusOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status][%d] readAppsV1NamespacedDeploymentStatusOK  %+v", 200, o.Payload)
}

func (o *ReadAppsV1NamespacedDeploymentStatusOK) GetPayload() *models.IoK8sAPIAppsV1Deployment {
	return o.Payload
}

func (o *ReadAppsV1NamespacedDeploymentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadAppsV1NamespacedDeploymentStatusUnauthorized creates a ReadAppsV1NamespacedDeploymentStatusUnauthorized with default headers values
func NewReadAppsV1NamespacedDeploymentStatusUnauthorized() *ReadAppsV1NamespacedDeploymentStatusUnauthorized {
	return &ReadAppsV1NamespacedDeploymentStatusUnauthorized{}
}

/*ReadAppsV1NamespacedDeploymentStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadAppsV1NamespacedDeploymentStatusUnauthorized struct {
}

func (o *ReadAppsV1NamespacedDeploymentStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status][%d] readAppsV1NamespacedDeploymentStatusUnauthorized ", 401)
}

func (o *ReadAppsV1NamespacedDeploymentStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}