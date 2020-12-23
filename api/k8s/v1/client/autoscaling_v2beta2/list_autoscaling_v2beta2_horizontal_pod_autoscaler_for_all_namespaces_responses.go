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

// ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesReader is a Reader for the ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces structure.
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK creates a ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK with default headers values
func NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK() *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK {
	return &ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK{}
}

/*ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK handles this case with default header values.

OK
*/
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK struct {
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerList
}

func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v2beta2/horizontalpodautoscalers][%d] listAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK) GetPayload() *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerList {
	return o.Payload
}

func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscalerList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized creates a ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized with default headers values
func NewListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized() *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized {
	return &ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized{}
}

/*ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized struct {
}

func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v2beta2/horizontalpodautoscalers][%d] listAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized ", 401)
}

func (o *ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}