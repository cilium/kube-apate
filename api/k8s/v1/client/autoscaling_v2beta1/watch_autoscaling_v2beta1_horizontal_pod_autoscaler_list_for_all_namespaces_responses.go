// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package autoscaling_v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesReader is a Reader for the WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces structure.
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK creates a WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK with default headers values
func NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK() *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK {
	return &WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK{}
}

/*WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v2beta1/watch/horizontalpodautoscalers][%d] watchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized creates a WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized with default headers values
func NewWatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized() *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized {
	return &WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized{}
}

/*WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized struct {
}

func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/autoscaling/v2beta1/watch/horizontalpodautoscalers][%d] watchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
