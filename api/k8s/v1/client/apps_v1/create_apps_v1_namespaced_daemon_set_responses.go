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

// CreateAppsV1NamespacedDaemonSetReader is a Reader for the CreateAppsV1NamespacedDaemonSet structure.
type CreateAppsV1NamespacedDaemonSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppsV1NamespacedDaemonSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAppsV1NamespacedDaemonSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAppsV1NamespacedDaemonSetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAppsV1NamespacedDaemonSetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAppsV1NamespacedDaemonSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAppsV1NamespacedDaemonSetOK creates a CreateAppsV1NamespacedDaemonSetOK with default headers values
func NewCreateAppsV1NamespacedDaemonSetOK() *CreateAppsV1NamespacedDaemonSetOK {
	return &CreateAppsV1NamespacedDaemonSetOK{}
}

/*CreateAppsV1NamespacedDaemonSetOK handles this case with default header values.

OK
*/
type CreateAppsV1NamespacedDaemonSetOK struct {
	Payload *models.IoK8sAPIAppsV1DaemonSet
}

func (o *CreateAppsV1NamespacedDaemonSetOK) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/daemonsets][%d] createAppsV1NamespacedDaemonSetOK  %+v", 200, o.Payload)
}

func (o *CreateAppsV1NamespacedDaemonSetOK) GetPayload() *models.IoK8sAPIAppsV1DaemonSet {
	return o.Payload
}

func (o *CreateAppsV1NamespacedDaemonSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppsV1NamespacedDaemonSetCreated creates a CreateAppsV1NamespacedDaemonSetCreated with default headers values
func NewCreateAppsV1NamespacedDaemonSetCreated() *CreateAppsV1NamespacedDaemonSetCreated {
	return &CreateAppsV1NamespacedDaemonSetCreated{}
}

/*CreateAppsV1NamespacedDaemonSetCreated handles this case with default header values.

Created
*/
type CreateAppsV1NamespacedDaemonSetCreated struct {
	Payload *models.IoK8sAPIAppsV1DaemonSet
}

func (o *CreateAppsV1NamespacedDaemonSetCreated) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/daemonsets][%d] createAppsV1NamespacedDaemonSetCreated  %+v", 201, o.Payload)
}

func (o *CreateAppsV1NamespacedDaemonSetCreated) GetPayload() *models.IoK8sAPIAppsV1DaemonSet {
	return o.Payload
}

func (o *CreateAppsV1NamespacedDaemonSetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppsV1NamespacedDaemonSetAccepted creates a CreateAppsV1NamespacedDaemonSetAccepted with default headers values
func NewCreateAppsV1NamespacedDaemonSetAccepted() *CreateAppsV1NamespacedDaemonSetAccepted {
	return &CreateAppsV1NamespacedDaemonSetAccepted{}
}

/*CreateAppsV1NamespacedDaemonSetAccepted handles this case with default header values.

Accepted
*/
type CreateAppsV1NamespacedDaemonSetAccepted struct {
	Payload *models.IoK8sAPIAppsV1DaemonSet
}

func (o *CreateAppsV1NamespacedDaemonSetAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/daemonsets][%d] createAppsV1NamespacedDaemonSetAccepted  %+v", 202, o.Payload)
}

func (o *CreateAppsV1NamespacedDaemonSetAccepted) GetPayload() *models.IoK8sAPIAppsV1DaemonSet {
	return o.Payload
}

func (o *CreateAppsV1NamespacedDaemonSetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppsV1NamespacedDaemonSetUnauthorized creates a CreateAppsV1NamespacedDaemonSetUnauthorized with default headers values
func NewCreateAppsV1NamespacedDaemonSetUnauthorized() *CreateAppsV1NamespacedDaemonSetUnauthorized {
	return &CreateAppsV1NamespacedDaemonSetUnauthorized{}
}

/*CreateAppsV1NamespacedDaemonSetUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAppsV1NamespacedDaemonSetUnauthorized struct {
}

func (o *CreateAppsV1NamespacedDaemonSetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/apps/v1/namespaces/{namespace}/daemonsets][%d] createAppsV1NamespacedDaemonSetUnauthorized ", 401)
}

func (o *CreateAppsV1NamespacedDaemonSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}