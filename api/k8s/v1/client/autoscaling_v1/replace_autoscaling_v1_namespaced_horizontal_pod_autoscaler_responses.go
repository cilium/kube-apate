// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerReader is a Reader for the ReplaceAutoscalingV1NamespacedHorizontalPodAutoscaler structure.
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK creates a ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK {
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK{}
}

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK handles this case with default header values.

OK
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK struct {
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK) Error() string {
	return fmt.Sprintf("[PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}][%d] replaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK  %+v", 200, o.Payload)
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK) GetPayload() *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler {
	return o.Payload
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated creates a ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated {
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated{}
}

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated handles this case with default header values.

Created
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated struct {
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}][%d] replaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated) GetPayload() *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler {
	return o.Payload
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized creates a ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized() *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized {
	return &ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

/*ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}][%d] replaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized ", 401)
}

func (o *ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
