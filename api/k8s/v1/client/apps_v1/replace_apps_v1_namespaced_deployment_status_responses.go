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

// ReplaceAppsV1NamespacedDeploymentStatusReader is a Reader for the ReplaceAppsV1NamespacedDeploymentStatus structure.
type ReplaceAppsV1NamespacedDeploymentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAppsV1NamespacedDeploymentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAppsV1NamespacedDeploymentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAppsV1NamespacedDeploymentStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAppsV1NamespacedDeploymentStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAppsV1NamespacedDeploymentStatusOK creates a ReplaceAppsV1NamespacedDeploymentStatusOK with default headers values
func NewReplaceAppsV1NamespacedDeploymentStatusOK() *ReplaceAppsV1NamespacedDeploymentStatusOK {
	return &ReplaceAppsV1NamespacedDeploymentStatusOK{}
}

/*ReplaceAppsV1NamespacedDeploymentStatusOK handles this case with default header values.

OK
*/
type ReplaceAppsV1NamespacedDeploymentStatusOK struct {
	Payload *models.IoK8sAPIAppsV1Deployment
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusOK) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status][%d] replaceAppsV1NamespacedDeploymentStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusOK) GetPayload() *models.IoK8sAPIAppsV1Deployment {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedDeploymentStatusCreated creates a ReplaceAppsV1NamespacedDeploymentStatusCreated with default headers values
func NewReplaceAppsV1NamespacedDeploymentStatusCreated() *ReplaceAppsV1NamespacedDeploymentStatusCreated {
	return &ReplaceAppsV1NamespacedDeploymentStatusCreated{}
}

/*ReplaceAppsV1NamespacedDeploymentStatusCreated handles this case with default header values.

Created
*/
type ReplaceAppsV1NamespacedDeploymentStatusCreated struct {
	Payload *models.IoK8sAPIAppsV1Deployment
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status][%d] replaceAppsV1NamespacedDeploymentStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusCreated) GetPayload() *models.IoK8sAPIAppsV1Deployment {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedDeploymentStatusUnauthorized creates a ReplaceAppsV1NamespacedDeploymentStatusUnauthorized with default headers values
func NewReplaceAppsV1NamespacedDeploymentStatusUnauthorized() *ReplaceAppsV1NamespacedDeploymentStatusUnauthorized {
	return &ReplaceAppsV1NamespacedDeploymentStatusUnauthorized{}
}

/*ReplaceAppsV1NamespacedDeploymentStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAppsV1NamespacedDeploymentStatusUnauthorized struct {
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/deployments/{name}/status][%d] replaceAppsV1NamespacedDeploymentStatusUnauthorized ", 401)
}

func (o *ReplaceAppsV1NamespacedDeploymentStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
