// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchCoordinationV1NamespacedLeaseReader is a Reader for the PatchCoordinationV1NamespacedLease structure.
type PatchCoordinationV1NamespacedLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchCoordinationV1NamespacedLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchCoordinationV1NamespacedLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchCoordinationV1NamespacedLeaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchCoordinationV1NamespacedLeaseOK creates a PatchCoordinationV1NamespacedLeaseOK with default headers values
func NewPatchCoordinationV1NamespacedLeaseOK() *PatchCoordinationV1NamespacedLeaseOK {
	return &PatchCoordinationV1NamespacedLeaseOK{}
}

/*PatchCoordinationV1NamespacedLeaseOK handles this case with default header values.

OK
*/
type PatchCoordinationV1NamespacedLeaseOK struct {
	Payload *models.IoK8sAPICoordinationV1Lease
}

func (o *PatchCoordinationV1NamespacedLeaseOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}][%d] patchCoordinationV1NamespacedLeaseOK  %+v", 200, o.Payload)
}

func (o *PatchCoordinationV1NamespacedLeaseOK) GetPayload() *models.IoK8sAPICoordinationV1Lease {
	return o.Payload
}

func (o *PatchCoordinationV1NamespacedLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoordinationV1Lease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCoordinationV1NamespacedLeaseUnauthorized creates a PatchCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewPatchCoordinationV1NamespacedLeaseUnauthorized() *PatchCoordinationV1NamespacedLeaseUnauthorized {
	return &PatchCoordinationV1NamespacedLeaseUnauthorized{}
}

/*PatchCoordinationV1NamespacedLeaseUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchCoordinationV1NamespacedLeaseUnauthorized struct {
}

func (o *PatchCoordinationV1NamespacedLeaseUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}][%d] patchCoordinationV1NamespacedLeaseUnauthorized ", 401)
}

func (o *PatchCoordinationV1NamespacedLeaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
