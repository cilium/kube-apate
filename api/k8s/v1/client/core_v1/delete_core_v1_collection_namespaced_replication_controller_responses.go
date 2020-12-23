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

// DeleteCoreV1CollectionNamespacedReplicationControllerReader is a Reader for the DeleteCoreV1CollectionNamespacedReplicationController structure.
type DeleteCoreV1CollectionNamespacedReplicationControllerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoreV1CollectionNamespacedReplicationControllerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoreV1CollectionNamespacedReplicationControllerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoreV1CollectionNamespacedReplicationControllerOK creates a DeleteCoreV1CollectionNamespacedReplicationControllerOK with default headers values
func NewDeleteCoreV1CollectionNamespacedReplicationControllerOK() *DeleteCoreV1CollectionNamespacedReplicationControllerOK {
	return &DeleteCoreV1CollectionNamespacedReplicationControllerOK{}
}

/*DeleteCoreV1CollectionNamespacedReplicationControllerOK handles this case with default header values.

OK
*/
type DeleteCoreV1CollectionNamespacedReplicationControllerOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoreV1CollectionNamespacedReplicationControllerOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/replicationcontrollers][%d] deleteCoreV1CollectionNamespacedReplicationControllerOK  %+v", 200, o.Payload)
}

func (o *DeleteCoreV1CollectionNamespacedReplicationControllerOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoreV1CollectionNamespacedReplicationControllerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized creates a DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized with default headers values
func NewDeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized() *DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized {
	return &DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized{}
}

/*DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized struct {
}

func (o *DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/replicationcontrollers][%d] deleteCoreV1CollectionNamespacedReplicationControllerUnauthorized ", 401)
}

func (o *DeleteCoreV1CollectionNamespacedReplicationControllerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}