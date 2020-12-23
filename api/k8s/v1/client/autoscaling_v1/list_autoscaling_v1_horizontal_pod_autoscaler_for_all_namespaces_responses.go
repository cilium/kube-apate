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

// ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesReader is a Reader for the ListAutoscalingV1HorizontalPodAutoscalerForAllNamespaces structure.
type ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK creates a ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK with default headers values
func NewListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK() *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK {
	return &ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK{}
}

/*ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK handles this case with default header values.

OK
*/
type ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK struct {
	Payload *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscalerList
}

func (o *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v1/horizontalpodautoscalers][%d] listAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK) GetPayload() *models.IoK8sAPIAutoscalingV1HorizontalPodAutoscalerList {
	return o.Payload
}

func (o *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1HorizontalPodAutoscalerList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized creates a ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized with default headers values
func NewListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized() *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized {
	return &ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized{}
}

/*ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized struct {
}

func (o *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v1/horizontalpodautoscalers][%d] listAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized ", 401)
}

func (o *ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}