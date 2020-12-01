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

// DeleteCoreV1CollectionNodeReader is a Reader for the DeleteCoreV1CollectionNode structure.
type DeleteCoreV1CollectionNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoreV1CollectionNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoreV1CollectionNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoreV1CollectionNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoreV1CollectionNodeOK creates a DeleteCoreV1CollectionNodeOK with default headers values
func NewDeleteCoreV1CollectionNodeOK() *DeleteCoreV1CollectionNodeOK {
	return &DeleteCoreV1CollectionNodeOK{}
}

/*DeleteCoreV1CollectionNodeOK handles this case with default header values.

OK
*/
type DeleteCoreV1CollectionNodeOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoreV1CollectionNodeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/nodes][%d] deleteCoreV1CollectionNodeOK  %+v", 200, o.Payload)
}

func (o *DeleteCoreV1CollectionNodeOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoreV1CollectionNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1CollectionNodeUnauthorized creates a DeleteCoreV1CollectionNodeUnauthorized with default headers values
func NewDeleteCoreV1CollectionNodeUnauthorized() *DeleteCoreV1CollectionNodeUnauthorized {
	return &DeleteCoreV1CollectionNodeUnauthorized{}
}

/*DeleteCoreV1CollectionNodeUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoreV1CollectionNodeUnauthorized struct {
}

func (o *DeleteCoreV1CollectionNodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/nodes][%d] deleteCoreV1CollectionNodeUnauthorized ", 401)
}

func (o *DeleteCoreV1CollectionNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
