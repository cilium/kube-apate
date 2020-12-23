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

// DeleteCoordinationV1CollectionNamespacedLeaseReader is a Reader for the DeleteCoordinationV1CollectionNamespacedLease structure.
type DeleteCoordinationV1CollectionNamespacedLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoordinationV1CollectionNamespacedLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoordinationV1CollectionNamespacedLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoordinationV1CollectionNamespacedLeaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoordinationV1CollectionNamespacedLeaseOK creates a DeleteCoordinationV1CollectionNamespacedLeaseOK with default headers values
func NewDeleteCoordinationV1CollectionNamespacedLeaseOK() *DeleteCoordinationV1CollectionNamespacedLeaseOK {
	return &DeleteCoordinationV1CollectionNamespacedLeaseOK{}
}

/*DeleteCoordinationV1CollectionNamespacedLeaseOK handles this case with default header values.

OK
*/
type DeleteCoordinationV1CollectionNamespacedLeaseOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoordinationV1CollectionNamespacedLeaseOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases][%d] deleteCoordinationV1CollectionNamespacedLeaseOK  %+v", 200, o.Payload)
}

func (o *DeleteCoordinationV1CollectionNamespacedLeaseOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoordinationV1CollectionNamespacedLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoordinationV1CollectionNamespacedLeaseUnauthorized creates a DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized with default headers values
func NewDeleteCoordinationV1CollectionNamespacedLeaseUnauthorized() *DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized {
	return &DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized{}
}

/*DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized struct {
}

func (o *DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases][%d] deleteCoordinationV1CollectionNamespacedLeaseUnauthorized ", 401)
}

func (o *DeleteCoordinationV1CollectionNamespacedLeaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}