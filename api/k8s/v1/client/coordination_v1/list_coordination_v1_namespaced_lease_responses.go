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

// ListCoordinationV1NamespacedLeaseReader is a Reader for the ListCoordinationV1NamespacedLease structure.
type ListCoordinationV1NamespacedLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCoordinationV1NamespacedLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCoordinationV1NamespacedLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCoordinationV1NamespacedLeaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCoordinationV1NamespacedLeaseOK creates a ListCoordinationV1NamespacedLeaseOK with default headers values
func NewListCoordinationV1NamespacedLeaseOK() *ListCoordinationV1NamespacedLeaseOK {
	return &ListCoordinationV1NamespacedLeaseOK{}
}

/*ListCoordinationV1NamespacedLeaseOK handles this case with default header values.

OK
*/
type ListCoordinationV1NamespacedLeaseOK struct {
	Payload *models.IoK8sAPICoordinationV1LeaseList
}

func (o *ListCoordinationV1NamespacedLeaseOK) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases][%d] listCoordinationV1NamespacedLeaseOK  %+v", 200, o.Payload)
}

func (o *ListCoordinationV1NamespacedLeaseOK) GetPayload() *models.IoK8sAPICoordinationV1LeaseList {
	return o.Payload
}

func (o *ListCoordinationV1NamespacedLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoordinationV1LeaseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCoordinationV1NamespacedLeaseUnauthorized creates a ListCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewListCoordinationV1NamespacedLeaseUnauthorized() *ListCoordinationV1NamespacedLeaseUnauthorized {
	return &ListCoordinationV1NamespacedLeaseUnauthorized{}
}

/*ListCoordinationV1NamespacedLeaseUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCoordinationV1NamespacedLeaseUnauthorized struct {
}

func (o *ListCoordinationV1NamespacedLeaseUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases][%d] listCoordinationV1NamespacedLeaseUnauthorized ", 401)
}

func (o *ListCoordinationV1NamespacedLeaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
