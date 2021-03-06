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

// PatchCoreV1NamespacedPersistentVolumeClaimStatusReader is a Reader for the PatchCoreV1NamespacedPersistentVolumeClaimStatus structure.
type PatchCoreV1NamespacedPersistentVolumeClaimStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1NamespacedPersistentVolumeClaimStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1NamespacedPersistentVolumeClaimStatusOK creates a PatchCoreV1NamespacedPersistentVolumeClaimStatusOK with default headers values
func NewPatchCoreV1NamespacedPersistentVolumeClaimStatusOK() *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK {
	return &PatchCoreV1NamespacedPersistentVolumeClaimStatusOK{}
}

/*PatchCoreV1NamespacedPersistentVolumeClaimStatusOK handles this case with default header values.

OK
*/
type PatchCoreV1NamespacedPersistentVolumeClaimStatusOK struct {
	Payload *models.IoK8sAPICoreV1PersistentVolumeClaim
}

func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status][%d] patchCoreV1NamespacedPersistentVolumeClaimStatusOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK) GetPayload() *models.IoK8sAPICoreV1PersistentVolumeClaim {
	return o.Payload
}

func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PersistentVolumeClaim)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized creates a PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized with default headers values
func NewPatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized() *PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized {
	return &PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized{}
}

/*PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized struct {
}

func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status][%d] patchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized ", 401)
}

func (o *PatchCoreV1NamespacedPersistentVolumeClaimStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
