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

// PatchCoreV1NamespaceStatusReader is a Reader for the PatchCoreV1NamespaceStatus structure.
type PatchCoreV1NamespaceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoreV1NamespaceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoreV1NamespaceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoreV1NamespaceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoreV1NamespaceStatusOK creates a PatchCoreV1NamespaceStatusOK with default headers values
func NewPatchCoreV1NamespaceStatusOK() *PatchCoreV1NamespaceStatusOK {
	return &PatchCoreV1NamespaceStatusOK{}
}

/*PatchCoreV1NamespaceStatusOK handles this case with default header values.

OK
*/
type PatchCoreV1NamespaceStatusOK struct {
	Payload *models.IoK8sAPICoreV1Namespace
}

func (o *PatchCoreV1NamespaceStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{name}/status][%d] patchCoreV1NamespaceStatusOK  %+v", 200, o.Payload)
}

func (o *PatchCoreV1NamespaceStatusOK) GetPayload() *models.IoK8sAPICoreV1Namespace {
	return o.Payload
}

func (o *PatchCoreV1NamespaceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Namespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoreV1NamespaceStatusUnauthorized creates a PatchCoreV1NamespaceStatusUnauthorized with default headers values
func NewPatchCoreV1NamespaceStatusUnauthorized() *PatchCoreV1NamespaceStatusUnauthorized {
	return &PatchCoreV1NamespaceStatusUnauthorized{}
}

/*PatchCoreV1NamespaceStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoreV1NamespaceStatusUnauthorized struct {
}

func (o *PatchCoreV1NamespaceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/namespaces/{name}/status][%d] patchCoreV1NamespaceStatusUnauthorized ", 401)
}

func (o *PatchCoreV1NamespaceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
