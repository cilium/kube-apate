// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchStorageV1CSINodeReader is a Reader for the PatchStorageV1CSINode structure.
type PatchStorageV1CSINodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStorageV1CSINodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchStorageV1CSINodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchStorageV1CSINodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchStorageV1CSINodeOK creates a PatchStorageV1CSINodeOK with default headers values
func NewPatchStorageV1CSINodeOK() *PatchStorageV1CSINodeOK {
	return &PatchStorageV1CSINodeOK{}
}

/*PatchStorageV1CSINodeOK handles this case with default header values.

OK
*/
type PatchStorageV1CSINodeOK struct {
	Payload *models.IoK8sAPIStorageV1CSINode
}

func (o *PatchStorageV1CSINodeOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/storage.k8s.io/v1/csinodes/{name}][%d] patchStorageV1CSINodeOK  %+v", 200, o.Payload)
}

func (o *PatchStorageV1CSINodeOK) GetPayload() *models.IoK8sAPIStorageV1CSINode {
	return o.Payload
}

func (o *PatchStorageV1CSINodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1CSINode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStorageV1CSINodeUnauthorized creates a PatchStorageV1CSINodeUnauthorized with default headers values
func NewPatchStorageV1CSINodeUnauthorized() *PatchStorageV1CSINodeUnauthorized {
	return &PatchStorageV1CSINodeUnauthorized{}
}

/*PatchStorageV1CSINodeUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchStorageV1CSINodeUnauthorized struct {
}

func (o *PatchStorageV1CSINodeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/storage.k8s.io/v1/csinodes/{name}][%d] patchStorageV1CSINodeUnauthorized ", 401)
}

func (o *PatchStorageV1CSINodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
