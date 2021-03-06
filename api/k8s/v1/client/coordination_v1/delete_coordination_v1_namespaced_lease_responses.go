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

// DeleteCoordinationV1NamespacedLeaseReader is a Reader for the DeleteCoordinationV1NamespacedLease structure.
type DeleteCoordinationV1NamespacedLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoordinationV1NamespacedLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoordinationV1NamespacedLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteCoordinationV1NamespacedLeaseAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoordinationV1NamespacedLeaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoordinationV1NamespacedLeaseOK creates a DeleteCoordinationV1NamespacedLeaseOK with default headers values
func NewDeleteCoordinationV1NamespacedLeaseOK() *DeleteCoordinationV1NamespacedLeaseOK {
	return &DeleteCoordinationV1NamespacedLeaseOK{}
}

/*DeleteCoordinationV1NamespacedLeaseOK handles this case with default header values.

OK
*/
type DeleteCoordinationV1NamespacedLeaseOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoordinationV1NamespacedLeaseOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}][%d] deleteCoordinationV1NamespacedLeaseOK  %+v", 200, o.Payload)
}

func (o *DeleteCoordinationV1NamespacedLeaseOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoordinationV1NamespacedLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoordinationV1NamespacedLeaseAccepted creates a DeleteCoordinationV1NamespacedLeaseAccepted with default headers values
func NewDeleteCoordinationV1NamespacedLeaseAccepted() *DeleteCoordinationV1NamespacedLeaseAccepted {
	return &DeleteCoordinationV1NamespacedLeaseAccepted{}
}

/*DeleteCoordinationV1NamespacedLeaseAccepted handles this case with default header values.

Accepted
*/
type DeleteCoordinationV1NamespacedLeaseAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoordinationV1NamespacedLeaseAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}][%d] deleteCoordinationV1NamespacedLeaseAccepted  %+v", 202, o.Payload)
}

func (o *DeleteCoordinationV1NamespacedLeaseAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoordinationV1NamespacedLeaseAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoordinationV1NamespacedLeaseUnauthorized creates a DeleteCoordinationV1NamespacedLeaseUnauthorized with default headers values
func NewDeleteCoordinationV1NamespacedLeaseUnauthorized() *DeleteCoordinationV1NamespacedLeaseUnauthorized {
	return &DeleteCoordinationV1NamespacedLeaseUnauthorized{}
}

/*DeleteCoordinationV1NamespacedLeaseUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoordinationV1NamespacedLeaseUnauthorized struct {
}

func (o *DeleteCoordinationV1NamespacedLeaseUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/coordination.k8s.io/v1/namespaces/{namespace}/leases/{name}][%d] deleteCoordinationV1NamespacedLeaseUnauthorized ", 401)
}

func (o *DeleteCoordinationV1NamespacedLeaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
