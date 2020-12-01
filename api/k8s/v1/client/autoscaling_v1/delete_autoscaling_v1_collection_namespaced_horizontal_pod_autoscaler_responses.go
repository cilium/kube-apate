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

// DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerReader is a Reader for the DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscaler structure.
type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK creates a DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK with default headers values
func NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK() *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK {
	return &DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK{}
}

/*DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK handles this case with default header values.

OK
*/
type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers][%d] deleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK  %+v", 200, o.Payload)
}

func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized creates a DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewDeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized() *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized {
	return &DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized{}
}

/*DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized struct {
}

func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers][%d] deleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized ", 401)
}

func (o *DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
