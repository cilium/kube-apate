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

// PatchAppsV1NamespacedControllerRevisionReader is a Reader for the PatchAppsV1NamespacedControllerRevision structure.
type PatchAppsV1NamespacedControllerRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAppsV1NamespacedControllerRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAppsV1NamespacedControllerRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchAppsV1NamespacedControllerRevisionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAppsV1NamespacedControllerRevisionOK creates a PatchAppsV1NamespacedControllerRevisionOK with default headers values
func NewPatchAppsV1NamespacedControllerRevisionOK() *PatchAppsV1NamespacedControllerRevisionOK {
	return &PatchAppsV1NamespacedControllerRevisionOK{}
}

/*PatchAppsV1NamespacedControllerRevisionOK handles this case with default header values.

OK
*/
type PatchAppsV1NamespacedControllerRevisionOK struct {
	Payload *models.IoK8sAPIAppsV1ControllerRevision
}

func (o *PatchAppsV1NamespacedControllerRevisionOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/apps/v1/namespaces/{namespace}/controllerrevisions/{name}][%d] patchAppsV1NamespacedControllerRevisionOK  %+v", 200, o.Payload)
}

func (o *PatchAppsV1NamespacedControllerRevisionOK) GetPayload() *models.IoK8sAPIAppsV1ControllerRevision {
	return o.Payload
}

func (o *PatchAppsV1NamespacedControllerRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1ControllerRevision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAppsV1NamespacedControllerRevisionUnauthorized creates a PatchAppsV1NamespacedControllerRevisionUnauthorized with default headers values
func NewPatchAppsV1NamespacedControllerRevisionUnauthorized() *PatchAppsV1NamespacedControllerRevisionUnauthorized {
	return &PatchAppsV1NamespacedControllerRevisionUnauthorized{}
}

/*PatchAppsV1NamespacedControllerRevisionUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchAppsV1NamespacedControllerRevisionUnauthorized struct {
}

func (o *PatchAppsV1NamespacedControllerRevisionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/apps/v1/namespaces/{namespace}/controllerrevisions/{name}][%d] patchAppsV1NamespacedControllerRevisionUnauthorized ", 401)
}

func (o *PatchAppsV1NamespacedControllerRevisionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
