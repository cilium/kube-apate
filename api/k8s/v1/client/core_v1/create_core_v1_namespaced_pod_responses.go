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

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateCoreV1NamespacedPodReader is a Reader for the CreateCoreV1NamespacedPod structure.
type CreateCoreV1NamespacedPodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCoreV1NamespacedPodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCoreV1NamespacedPodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateCoreV1NamespacedPodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCoreV1NamespacedPodAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateCoreV1NamespacedPodUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCoreV1NamespacedPodOK creates a CreateCoreV1NamespacedPodOK with default headers values
func NewCreateCoreV1NamespacedPodOK() *CreateCoreV1NamespacedPodOK {
	return &CreateCoreV1NamespacedPodOK{}
}

/*CreateCoreV1NamespacedPodOK handles this case with default header values.

OK
*/
type CreateCoreV1NamespacedPodOK struct {
	Payload *models.IoK8sAPICoreV1Pod
}

func (o *CreateCoreV1NamespacedPodOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods][%d] createCoreV1NamespacedPodOK  %+v", 200, o.Payload)
}

func (o *CreateCoreV1NamespacedPodOK) GetPayload() *models.IoK8sAPICoreV1Pod {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Pod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodCreated creates a CreateCoreV1NamespacedPodCreated with default headers values
func NewCreateCoreV1NamespacedPodCreated() *CreateCoreV1NamespacedPodCreated {
	return &CreateCoreV1NamespacedPodCreated{}
}

/*CreateCoreV1NamespacedPodCreated handles this case with default header values.

Created
*/
type CreateCoreV1NamespacedPodCreated struct {
	Payload *models.IoK8sAPICoreV1Pod
}

func (o *CreateCoreV1NamespacedPodCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods][%d] createCoreV1NamespacedPodCreated  %+v", 201, o.Payload)
}

func (o *CreateCoreV1NamespacedPodCreated) GetPayload() *models.IoK8sAPICoreV1Pod {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Pod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodAccepted creates a CreateCoreV1NamespacedPodAccepted with default headers values
func NewCreateCoreV1NamespacedPodAccepted() *CreateCoreV1NamespacedPodAccepted {
	return &CreateCoreV1NamespacedPodAccepted{}
}

/*CreateCoreV1NamespacedPodAccepted handles this case with default header values.

Accepted
*/
type CreateCoreV1NamespacedPodAccepted struct {
	Payload *models.IoK8sAPICoreV1Pod
}

func (o *CreateCoreV1NamespacedPodAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods][%d] createCoreV1NamespacedPodAccepted  %+v", 202, o.Payload)
}

func (o *CreateCoreV1NamespacedPodAccepted) GetPayload() *models.IoK8sAPICoreV1Pod {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Pod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodUnauthorized creates a CreateCoreV1NamespacedPodUnauthorized with default headers values
func NewCreateCoreV1NamespacedPodUnauthorized() *CreateCoreV1NamespacedPodUnauthorized {
	return &CreateCoreV1NamespacedPodUnauthorized{}
}

/*CreateCoreV1NamespacedPodUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCoreV1NamespacedPodUnauthorized struct {
}

func (o *CreateCoreV1NamespacedPodUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods][%d] createCoreV1NamespacedPodUnauthorized ", 401)
}

func (o *CreateCoreV1NamespacedPodUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
