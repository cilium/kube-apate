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

// CreateCoreV1NamespacedPodEvictionReader is a Reader for the CreateCoreV1NamespacedPodEviction structure.
type CreateCoreV1NamespacedPodEvictionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCoreV1NamespacedPodEvictionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCoreV1NamespacedPodEvictionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateCoreV1NamespacedPodEvictionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateCoreV1NamespacedPodEvictionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateCoreV1NamespacedPodEvictionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateCoreV1NamespacedPodEvictionOK creates a CreateCoreV1NamespacedPodEvictionOK with default headers values
func NewCreateCoreV1NamespacedPodEvictionOK() *CreateCoreV1NamespacedPodEvictionOK {
	return &CreateCoreV1NamespacedPodEvictionOK{}
}

/*CreateCoreV1NamespacedPodEvictionOK handles this case with default header values.

OK
*/
type CreateCoreV1NamespacedPodEvictionOK struct {
	Payload *models.IoK8sAPIPolicyV1beta1Eviction
}

func (o *CreateCoreV1NamespacedPodEvictionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/eviction][%d] createCoreV1NamespacedPodEvictionOK  %+v", 200, o.Payload)
}

func (o *CreateCoreV1NamespacedPodEvictionOK) GetPayload() *models.IoK8sAPIPolicyV1beta1Eviction {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodEvictionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1Eviction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodEvictionCreated creates a CreateCoreV1NamespacedPodEvictionCreated with default headers values
func NewCreateCoreV1NamespacedPodEvictionCreated() *CreateCoreV1NamespacedPodEvictionCreated {
	return &CreateCoreV1NamespacedPodEvictionCreated{}
}

/*CreateCoreV1NamespacedPodEvictionCreated handles this case with default header values.

Created
*/
type CreateCoreV1NamespacedPodEvictionCreated struct {
	Payload *models.IoK8sAPIPolicyV1beta1Eviction
}

func (o *CreateCoreV1NamespacedPodEvictionCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/eviction][%d] createCoreV1NamespacedPodEvictionCreated  %+v", 201, o.Payload)
}

func (o *CreateCoreV1NamespacedPodEvictionCreated) GetPayload() *models.IoK8sAPIPolicyV1beta1Eviction {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodEvictionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1Eviction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodEvictionAccepted creates a CreateCoreV1NamespacedPodEvictionAccepted with default headers values
func NewCreateCoreV1NamespacedPodEvictionAccepted() *CreateCoreV1NamespacedPodEvictionAccepted {
	return &CreateCoreV1NamespacedPodEvictionAccepted{}
}

/*CreateCoreV1NamespacedPodEvictionAccepted handles this case with default header values.

Accepted
*/
type CreateCoreV1NamespacedPodEvictionAccepted struct {
	Payload *models.IoK8sAPIPolicyV1beta1Eviction
}

func (o *CreateCoreV1NamespacedPodEvictionAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/eviction][%d] createCoreV1NamespacedPodEvictionAccepted  %+v", 202, o.Payload)
}

func (o *CreateCoreV1NamespacedPodEvictionAccepted) GetPayload() *models.IoK8sAPIPolicyV1beta1Eviction {
	return o.Payload
}

func (o *CreateCoreV1NamespacedPodEvictionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1Eviction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCoreV1NamespacedPodEvictionUnauthorized creates a CreateCoreV1NamespacedPodEvictionUnauthorized with default headers values
func NewCreateCoreV1NamespacedPodEvictionUnauthorized() *CreateCoreV1NamespacedPodEvictionUnauthorized {
	return &CreateCoreV1NamespacedPodEvictionUnauthorized{}
}

/*CreateCoreV1NamespacedPodEvictionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCoreV1NamespacedPodEvictionUnauthorized struct {
}

func (o *CreateCoreV1NamespacedPodEvictionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/namespaces/{namespace}/pods/{name}/eviction][%d] createCoreV1NamespacedPodEvictionUnauthorized ", 401)
}

func (o *CreateCoreV1NamespacedPodEvictionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}