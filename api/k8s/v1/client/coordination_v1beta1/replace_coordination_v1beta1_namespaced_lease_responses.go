// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceCoordinationV1beta1NamespacedLeaseReader is a Reader for the ReplaceCoordinationV1beta1NamespacedLease structure.
type ReplaceCoordinationV1beta1NamespacedLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoordinationV1beta1NamespacedLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoordinationV1beta1NamespacedLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoordinationV1beta1NamespacedLeaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoordinationV1beta1NamespacedLeaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoordinationV1beta1NamespacedLeaseOK creates a ReplaceCoordinationV1beta1NamespacedLeaseOK with default headers values
func NewReplaceCoordinationV1beta1NamespacedLeaseOK() *ReplaceCoordinationV1beta1NamespacedLeaseOK {
	return &ReplaceCoordinationV1beta1NamespacedLeaseOK{}
}

/*ReplaceCoordinationV1beta1NamespacedLeaseOK handles this case with default header values.

OK
*/
type ReplaceCoordinationV1beta1NamespacedLeaseOK struct {
	Payload *models.IoK8sAPICoordinationV1beta1Lease
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseOK) Error() string {
	return fmt.Sprintf("[PUT /apis/coordination.k8s.io/v1beta1/namespaces/{namespace}/leases/{name}][%d] replaceCoordinationV1beta1NamespacedLeaseOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseOK) GetPayload() *models.IoK8sAPICoordinationV1beta1Lease {
	return o.Payload
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoordinationV1beta1Lease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoordinationV1beta1NamespacedLeaseCreated creates a ReplaceCoordinationV1beta1NamespacedLeaseCreated with default headers values
func NewReplaceCoordinationV1beta1NamespacedLeaseCreated() *ReplaceCoordinationV1beta1NamespacedLeaseCreated {
	return &ReplaceCoordinationV1beta1NamespacedLeaseCreated{}
}

/*ReplaceCoordinationV1beta1NamespacedLeaseCreated handles this case with default header values.

Created
*/
type ReplaceCoordinationV1beta1NamespacedLeaseCreated struct {
	Payload *models.IoK8sAPICoordinationV1beta1Lease
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/coordination.k8s.io/v1beta1/namespaces/{namespace}/leases/{name}][%d] replaceCoordinationV1beta1NamespacedLeaseCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseCreated) GetPayload() *models.IoK8sAPICoordinationV1beta1Lease {
	return o.Payload
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoordinationV1beta1Lease)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoordinationV1beta1NamespacedLeaseUnauthorized creates a ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized with default headers values
func NewReplaceCoordinationV1beta1NamespacedLeaseUnauthorized() *ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized {
	return &ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized{}
}

/*ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized struct {
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/coordination.k8s.io/v1beta1/namespaces/{namespace}/leases/{name}][%d] replaceCoordinationV1beta1NamespacedLeaseUnauthorized ", 401)
}

func (o *ReplaceCoordinationV1beta1NamespacedLeaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
