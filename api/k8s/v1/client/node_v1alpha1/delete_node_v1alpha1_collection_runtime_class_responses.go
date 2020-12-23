// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package node_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteNodeV1alpha1CollectionRuntimeClassReader is a Reader for the DeleteNodeV1alpha1CollectionRuntimeClass structure.
type DeleteNodeV1alpha1CollectionRuntimeClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNodeV1alpha1CollectionRuntimeClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNodeV1alpha1CollectionRuntimeClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNodeV1alpha1CollectionRuntimeClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNodeV1alpha1CollectionRuntimeClassOK creates a DeleteNodeV1alpha1CollectionRuntimeClassOK with default headers values
func NewDeleteNodeV1alpha1CollectionRuntimeClassOK() *DeleteNodeV1alpha1CollectionRuntimeClassOK {
	return &DeleteNodeV1alpha1CollectionRuntimeClassOK{}
}

/*DeleteNodeV1alpha1CollectionRuntimeClassOK handles this case with default header values.

OK
*/
type DeleteNodeV1alpha1CollectionRuntimeClassOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNodeV1alpha1CollectionRuntimeClassOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/node.k8s.io/v1alpha1/runtimeclasses][%d] deleteNodeV1alpha1CollectionRuntimeClassOK  %+v", 200, o.Payload)
}

func (o *DeleteNodeV1alpha1CollectionRuntimeClassOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNodeV1alpha1CollectionRuntimeClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNodeV1alpha1CollectionRuntimeClassUnauthorized creates a DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized with default headers values
func NewDeleteNodeV1alpha1CollectionRuntimeClassUnauthorized() *DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized {
	return &DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized{}
}

/*DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized struct {
}

func (o *DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/node.k8s.io/v1alpha1/runtimeclasses][%d] deleteNodeV1alpha1CollectionRuntimeClassUnauthorized ", 401)
}

func (o *DeleteNodeV1alpha1CollectionRuntimeClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}