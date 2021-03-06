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

// DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerReader is a Reader for the DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscaler structure.
type DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK creates a DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK with default headers values
func NewDeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK() *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK {
	return &DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK{}
}

/*DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK handles this case with default header values.

OK
*/
type DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers][%d] deleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK  %+v", 200, o.Payload)
}

func (o *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized creates a DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewDeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized() *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized {
	return &DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized{}
}

/*DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized struct {
}

func (o *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers][%d] deleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized ", 401)
}

func (o *DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
