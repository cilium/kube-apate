// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerReader is a Reader for the CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscaler structure.
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK creates a CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK {
	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK{}
}

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK handles this case with default header values.

OK
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK struct {
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) Error() string {
	return fmt.Sprintf("[POST /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers][%d] createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK  %+v", 200, o.Payload)
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) GetPayload() *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler {
	return o.Payload
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated creates a CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated {
	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated{}
}

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated handles this case with default header values.

Created
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated struct {
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) Error() string {
	return fmt.Sprintf("[POST /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers][%d] createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated  %+v", 201, o.Payload)
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) GetPayload() *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler {
	return o.Payload
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted creates a CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted {
	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted{}
}

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted handles this case with default header values.

Accepted
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted struct {
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers][%d] createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted  %+v", 202, o.Payload)
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) GetPayload() *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler {
	return o.Payload
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized creates a CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewCreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized() *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized {
	return &CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized{}
}

/*CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers][%d] createAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized ", 401)
}

func (o *CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
