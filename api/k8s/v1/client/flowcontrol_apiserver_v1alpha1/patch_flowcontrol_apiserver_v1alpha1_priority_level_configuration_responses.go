// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package flowcontrol_apiserver_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationReader is a Reader for the PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration structure.
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK creates a PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK with default headers values
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK() *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK {
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK{}
}

/*PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK handles this case with default header values.

OK
*/
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK struct {
	Payload *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration
}

func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}][%d] patchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK  %+v", 200, o.Payload)
}

func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) GetPayload() *models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration {
	return o.Payload
}

func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized creates a PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized with default headers values
func NewPatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized() *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized {
	return &PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized{}
}

/*PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized struct {
}

func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/flowcontrol.apiserver.k8s.io/v1alpha1/prioritylevelconfigurations/{name}][%d] patchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized ", 401)
}

func (o *PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
