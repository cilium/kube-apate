// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchCoreV1NamespacedReplicationControllerStatusReader is a Reader for the PatchCoreV1NamespacedReplicationControllerStatus structure.
type PatchCoreV1NamespacedReplicationControllerStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1NamespacedReplicationControllerStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1NamespacedReplicationControllerStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1NamespacedReplicationControllerStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1NamespacedReplicationControllerStatusOK creates a PatchCoreV1NamespacedReplicationControllerStatusOK with default headers values
func NewPatchCoreV1NamespacedReplicationControllerStatusOK() *PatchCoreV1NamespacedReplicationControllerStatusOK {
	return &PatchCoreV1NamespacedReplicationControllerStatusOK{}
}

/*PatchCoreV1NamespacedReplicationControllerStatusOK handles this case with default header values.

OK
*/
type PatchCoreV1NamespacedReplicationControllerStatusOK struct {
	Payload *models.IoK8sAPICoreV1ReplicationController
}

func (o *PatchCoreV1NamespacedReplicationControllerStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status][%d] patchCoreV1NamespacedReplicationControllerStatusOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1NamespacedReplicationControllerStatusOK) GetPayload() *models.IoK8sAPICoreV1ReplicationController {
	return o.Payload
}

func (o *PatchCoreV1NamespacedReplicationControllerStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ReplicationController)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1NamespacedReplicationControllerStatusUnauthorized creates a PatchCoreV1NamespacedReplicationControllerStatusUnauthorized with default headers values
func NewPatchCoreV1NamespacedReplicationControllerStatusUnauthorized() *PatchCoreV1NamespacedReplicationControllerStatusUnauthorized {
	return &PatchCoreV1NamespacedReplicationControllerStatusUnauthorized{}
}

/*PatchCoreV1NamespacedReplicationControllerStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1NamespacedReplicationControllerStatusUnauthorized struct {
}

func (o *PatchCoreV1NamespacedReplicationControllerStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status][%d] patchCoreV1NamespacedReplicationControllerStatusUnauthorized ", 401)
}

func (o *PatchCoreV1NamespacedReplicationControllerStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}