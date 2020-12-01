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

// WatchAutoscalingV1NamespacedHorizontalPodAutoscalerReader is a Reader for the WatchAutoscalingV1NamespacedHorizontalPodAutoscaler structure.
type WatchAutoscalingV1NamespacedHorizontalPodAutoscalerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK creates a WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK with default headers values
func NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK() *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK {
	return &WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK{}
}

/*WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK handles this case with default header values.

OK
*/
type WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v1/watch/namespaces/{namespace}/horizontalpodautoscalers/{name}][%d] watchAutoscalingV1NamespacedHorizontalPodAutoscalerOK  %+v", 200, o.Payload)
}

func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized creates a WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized with default headers values
func NewWatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized() *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized {
	return &WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized{}
}

/*WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized struct {
}

func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v1/watch/namespaces/{namespace}/horizontalpodautoscalers/{name}][%d] watchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized ", 401)
}

func (o *WatchAutoscalingV1NamespacedHorizontalPodAutoscalerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
