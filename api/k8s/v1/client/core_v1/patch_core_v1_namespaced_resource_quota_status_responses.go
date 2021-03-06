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

// PatchCoreV1NamespacedResourceQuotaStatusReader is a Reader for the PatchCoreV1NamespacedResourceQuotaStatus structure.
type PatchCoreV1NamespacedResourceQuotaStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1NamespacedResourceQuotaStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1NamespacedResourceQuotaStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1NamespacedResourceQuotaStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1NamespacedResourceQuotaStatusOK creates a PatchCoreV1NamespacedResourceQuotaStatusOK with default headers values
func NewPatchCoreV1NamespacedResourceQuotaStatusOK() *PatchCoreV1NamespacedResourceQuotaStatusOK {
	return &PatchCoreV1NamespacedResourceQuotaStatusOK{}
}

/*PatchCoreV1NamespacedResourceQuotaStatusOK handles this case with default header values.

OK
*/
type PatchCoreV1NamespacedResourceQuotaStatusOK struct {
	Payload *models.IoK8sAPICoreV1ResourceQuota
}

func (o *PatchCoreV1NamespacedResourceQuotaStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}/status][%d] patchCoreV1NamespacedResourceQuotaStatusOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1NamespacedResourceQuotaStatusOK) GetPayload() *models.IoK8sAPICoreV1ResourceQuota {
	return o.Payload
}

func (o *PatchCoreV1NamespacedResourceQuotaStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1NamespacedResourceQuotaStatusUnauthorized creates a PatchCoreV1NamespacedResourceQuotaStatusUnauthorized with default headers values
func NewPatchCoreV1NamespacedResourceQuotaStatusUnauthorized() *PatchCoreV1NamespacedResourceQuotaStatusUnauthorized {
	return &PatchCoreV1NamespacedResourceQuotaStatusUnauthorized{}
}

/*PatchCoreV1NamespacedResourceQuotaStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1NamespacedResourceQuotaStatusUnauthorized struct {
}

func (o *PatchCoreV1NamespacedResourceQuotaStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}/status][%d] patchCoreV1NamespacedResourceQuotaStatusUnauthorized ", 401)
}

func (o *PatchCoreV1NamespacedResourceQuotaStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
