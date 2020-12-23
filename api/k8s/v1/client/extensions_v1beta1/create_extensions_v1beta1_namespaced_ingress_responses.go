// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateExtensionsV1beta1NamespacedIngressReader is a Reader for the CreateExtensionsV1beta1NamespacedIngress structure.
type CreateExtensionsV1beta1NamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExtensionsV1beta1NamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateExtensionsV1beta1NamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateExtensionsV1beta1NamespacedIngressCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateExtensionsV1beta1NamespacedIngressAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateExtensionsV1beta1NamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateExtensionsV1beta1NamespacedIngressOK creates a CreateExtensionsV1beta1NamespacedIngressOK with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressOK() *CreateExtensionsV1beta1NamespacedIngressOK {
	return &CreateExtensionsV1beta1NamespacedIngressOK{}
}

/*CreateExtensionsV1beta1NamespacedIngressOK handles this case with default header values.

OK
*/
type CreateExtensionsV1beta1NamespacedIngressOK struct {
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress
}

func (o *CreateExtensionsV1beta1NamespacedIngressOK) Error() string {
	return fmt.Sprintf("[POST /apis/extensions/v1beta1/namespaces/{namespace}/ingresses][%d] createExtensionsV1beta1NamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *CreateExtensionsV1beta1NamespacedIngressOK) GetPayload() *models.IoK8sAPIExtensionsV1beta1Ingress {
	return o.Payload
}

func (o *CreateExtensionsV1beta1NamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIExtensionsV1beta1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionsV1beta1NamespacedIngressCreated creates a CreateExtensionsV1beta1NamespacedIngressCreated with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressCreated() *CreateExtensionsV1beta1NamespacedIngressCreated {
	return &CreateExtensionsV1beta1NamespacedIngressCreated{}
}

/*CreateExtensionsV1beta1NamespacedIngressCreated handles this case with default header values.

Created
*/
type CreateExtensionsV1beta1NamespacedIngressCreated struct {
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress
}

func (o *CreateExtensionsV1beta1NamespacedIngressCreated) Error() string {
	return fmt.Sprintf("[POST /apis/extensions/v1beta1/namespaces/{namespace}/ingresses][%d] createExtensionsV1beta1NamespacedIngressCreated  %+v", 201, o.Payload)
}

func (o *CreateExtensionsV1beta1NamespacedIngressCreated) GetPayload() *models.IoK8sAPIExtensionsV1beta1Ingress {
	return o.Payload
}

func (o *CreateExtensionsV1beta1NamespacedIngressCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIExtensionsV1beta1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionsV1beta1NamespacedIngressAccepted creates a CreateExtensionsV1beta1NamespacedIngressAccepted with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressAccepted() *CreateExtensionsV1beta1NamespacedIngressAccepted {
	return &CreateExtensionsV1beta1NamespacedIngressAccepted{}
}

/*CreateExtensionsV1beta1NamespacedIngressAccepted handles this case with default header values.

Accepted
*/
type CreateExtensionsV1beta1NamespacedIngressAccepted struct {
	Payload *models.IoK8sAPIExtensionsV1beta1Ingress
}

func (o *CreateExtensionsV1beta1NamespacedIngressAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/extensions/v1beta1/namespaces/{namespace}/ingresses][%d] createExtensionsV1beta1NamespacedIngressAccepted  %+v", 202, o.Payload)
}

func (o *CreateExtensionsV1beta1NamespacedIngressAccepted) GetPayload() *models.IoK8sAPIExtensionsV1beta1Ingress {
	return o.Payload
}

func (o *CreateExtensionsV1beta1NamespacedIngressAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIExtensionsV1beta1Ingress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionsV1beta1NamespacedIngressUnauthorized creates a CreateExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewCreateExtensionsV1beta1NamespacedIngressUnauthorized() *CreateExtensionsV1beta1NamespacedIngressUnauthorized {
	return &CreateExtensionsV1beta1NamespacedIngressUnauthorized{}
}

/*CreateExtensionsV1beta1NamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

func (o *CreateExtensionsV1beta1NamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/extensions/v1beta1/namespaces/{namespace}/ingresses][%d] createExtensionsV1beta1NamespacedIngressUnauthorized ", 401)
}

func (o *CreateExtensionsV1beta1NamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}