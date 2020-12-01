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

// CreateAppsV1NamespacedDeploymentReader is a Reader for the CreateAppsV1NamespacedDeployment structure.
type CreateAppsV1NamespacedDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppsV1NamespacedDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAppsV1NamespacedDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAppsV1NamespacedDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAppsV1NamespacedDeploymentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAppsV1NamespacedDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAppsV1NamespacedDeploymentOK creates a CreateAppsV1NamespacedDeploymentOK with default headers values
func NewCreateAppsV1NamespacedDeploymentOK() *CreateAppsV1NamespacedDeploymentOK {
	return &CreateAppsV1NamespacedDeploymentOK{}
}

/*CreateAppsV1NamespacedDeploymentOK handles this case with default header values.

OK
*/
type CreateAppsV1NamespacedDeploymentOK struct {
	Payload *models.IoK8sAPIAppsV1Deployment
}

func (o *CreateAppsV1NamespacedDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/deployments][%d] createAppsV1NamespacedDeploymentOK  %+v", 200, o.Payload)
}

func (o *CreateAppsV1NamespacedDeploymentOK) GetPayload() *models.IoK8sAPIAppsV1Deployment {
	return o.Payload
}

func (o *CreateAppsV1NamespacedDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppsV1NamespacedDeploymentCreated creates a CreateAppsV1NamespacedDeploymentCreated with default headers values
func NewCreateAppsV1NamespacedDeploymentCreated() *CreateAppsV1NamespacedDeploymentCreated {
	return &CreateAppsV1NamespacedDeploymentCreated{}
}

/*CreateAppsV1NamespacedDeploymentCreated handles this case with default header values.

Created
*/
type CreateAppsV1NamespacedDeploymentCreated struct {
	Payload *models.IoK8sAPIAppsV1Deployment
}

func (o *CreateAppsV1NamespacedDeploymentCreated) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/deployments][%d] createAppsV1NamespacedDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateAppsV1NamespacedDeploymentCreated) GetPayload() *models.IoK8sAPIAppsV1Deployment {
	return o.Payload
}

func (o *CreateAppsV1NamespacedDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppsV1NamespacedDeploymentAccepted creates a CreateAppsV1NamespacedDeploymentAccepted with default headers values
func NewCreateAppsV1NamespacedDeploymentAccepted() *CreateAppsV1NamespacedDeploymentAccepted {
	return &CreateAppsV1NamespacedDeploymentAccepted{}
}

/*CreateAppsV1NamespacedDeploymentAccepted handles this case with default header values.

Accepted
*/
type CreateAppsV1NamespacedDeploymentAccepted struct {
	Payload *models.IoK8sAPIAppsV1Deployment
}

func (o *CreateAppsV1NamespacedDeploymentAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/deployments][%d] createAppsV1NamespacedDeploymentAccepted  %+v", 202, o.Payload)
}

func (o *CreateAppsV1NamespacedDeploymentAccepted) GetPayload() *models.IoK8sAPIAppsV1Deployment {
	return o.Payload
}

func (o *CreateAppsV1NamespacedDeploymentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1Deployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppsV1NamespacedDeploymentUnauthorized creates a CreateAppsV1NamespacedDeploymentUnauthorized with default headers values
func NewCreateAppsV1NamespacedDeploymentUnauthorized() *CreateAppsV1NamespacedDeploymentUnauthorized {
	return &CreateAppsV1NamespacedDeploymentUnauthorized{}
}

/*CreateAppsV1NamespacedDeploymentUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAppsV1NamespacedDeploymentUnauthorized struct {
}

func (o *CreateAppsV1NamespacedDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/deployments][%d] createAppsV1NamespacedDeploymentUnauthorized ", 401)
}

func (o *CreateAppsV1NamespacedDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
