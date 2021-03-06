// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchAppsV1NamespacedDeploymentScaleReader is a Reader for the PatchAppsV1NamespacedDeploymentScale structure.
type PatchAppsV1NamespacedDeploymentScaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAppsV1NamespacedDeploymentScaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAppsV1NamespacedDeploymentScaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAppsV1NamespacedDeploymentScaleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAppsV1NamespacedDeploymentScaleOK creates a PatchAppsV1NamespacedDeploymentScaleOK with default headers values
func NewPatchAppsV1NamespacedDeploymentScaleOK() *PatchAppsV1NamespacedDeploymentScaleOK {
	return &PatchAppsV1NamespacedDeploymentScaleOK{}
}

/*PatchAppsV1NamespacedDeploymentScaleOK handles this case with default header values.

OK
*/
type PatchAppsV1NamespacedDeploymentScaleOK struct {
	Payload *models.IoK8sAPIAutoscalingV1Scale
}

func (o *PatchAppsV1NamespacedDeploymentScaleOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/apps/v1/namespaces/{namespace}/deployments/{name}/scale][%d] patchAppsV1NamespacedDeploymentScaleOK  %+v", 200, o.Payload)
}

func (o *PatchAppsV1NamespacedDeploymentScaleOK) GetPayload() *models.IoK8sAPIAutoscalingV1Scale {
	return o.Payload
}

func (o *PatchAppsV1NamespacedDeploymentScaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAutoscalingV1Scale)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAppsV1NamespacedDeploymentScaleUnauthorized creates a PatchAppsV1NamespacedDeploymentScaleUnauthorized with default headers values
func NewPatchAppsV1NamespacedDeploymentScaleUnauthorized() *PatchAppsV1NamespacedDeploymentScaleUnauthorized {
	return &PatchAppsV1NamespacedDeploymentScaleUnauthorized{}
}

/*PatchAppsV1NamespacedDeploymentScaleUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchAppsV1NamespacedDeploymentScaleUnauthorized struct {
}

func (o *PatchAppsV1NamespacedDeploymentScaleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/apps/v1/namespaces/{namespace}/deployments/{name}/scale][%d] patchAppsV1NamespacedDeploymentScaleUnauthorized ", 401)
}

func (o *PatchAppsV1NamespacedDeploymentScaleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
