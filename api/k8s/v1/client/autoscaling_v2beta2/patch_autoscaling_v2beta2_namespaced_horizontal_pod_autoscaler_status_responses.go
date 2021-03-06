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

// PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusReader is a Reader for the PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus structure.
type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK creates a PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK with default headers values
func NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK() *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK {
	return &PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK{}
}

/*PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK handles this case with default header values.

OK
*/
type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK struct {
	Payload *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler
}

func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers/{name}/status][%d] patchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK  %+v", 200, o.Payload)
}

func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK) GetPayload() *models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler {
	return o.Payload
}

func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV2beta2HorizontalPodAutoscaler)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized creates a PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized with default headers values
func NewPatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized() *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized {
	return &PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized{}
}

/*PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized struct {
}

func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/autoscaling/v2beta2/namespaces/{namespace}/horizontalpodautoscalers/{name}/status][%d] patchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized ", 401)
}

func (o *PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
