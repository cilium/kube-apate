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

// PatchCoreV1NamespacedConfigMapReader is a Reader for the PatchCoreV1NamespacedConfigMap structure.
type PatchCoreV1NamespacedConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1NamespacedConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1NamespacedConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1NamespacedConfigMapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1NamespacedConfigMapOK creates a PatchCoreV1NamespacedConfigMapOK with default headers values
func NewPatchCoreV1NamespacedConfigMapOK() *PatchCoreV1NamespacedConfigMapOK {
	return &PatchCoreV1NamespacedConfigMapOK{}
}

/*PatchCoreV1NamespacedConfigMapOK handles this case with default header values.

OK
*/
type PatchCoreV1NamespacedConfigMapOK struct {
	Payload *models.IoK8sAPICoreV1ConfigMap
}

func (o *PatchCoreV1NamespacedConfigMapOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/configmaps/{name}][%d] patchCoreV1NamespacedConfigMapOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1NamespacedConfigMapOK) GetPayload() *models.IoK8sAPICoreV1ConfigMap {
	return o.Payload
}

func (o *PatchCoreV1NamespacedConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ConfigMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1NamespacedConfigMapUnauthorized creates a PatchCoreV1NamespacedConfigMapUnauthorized with default headers values
func NewPatchCoreV1NamespacedConfigMapUnauthorized() *PatchCoreV1NamespacedConfigMapUnauthorized {
	return &PatchCoreV1NamespacedConfigMapUnauthorized{}
}

/*PatchCoreV1NamespacedConfigMapUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1NamespacedConfigMapUnauthorized struct {
}

func (o *PatchCoreV1NamespacedConfigMapUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{namespace}/configmaps/{name}][%d] patchCoreV1NamespacedConfigMapUnauthorized ", 401)
}

func (o *PatchCoreV1NamespacedConfigMapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
