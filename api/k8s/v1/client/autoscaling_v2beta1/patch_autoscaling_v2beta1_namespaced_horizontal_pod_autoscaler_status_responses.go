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

// PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusReader is a Reader for the PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus structure.
type PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK creates a PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK with default headers values
func NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK() *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK {
	return &PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK{}
}

/*PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK handles this case with default header values.

OK
*/
type PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK struct {
	Payload *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler
}

func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status][%d] patchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK  %+v", 200, o.Payload)
}

func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK) GetPayload() *models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler {
	return o.Payload
}

func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV2beta1HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized creates a PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized with default headers values
func NewPatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized() *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized {
	return &PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized{}
}

/*PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized struct {
}

func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/autoscaling/v2beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status][%d] patchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized ", 401)
}

func (o *PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
