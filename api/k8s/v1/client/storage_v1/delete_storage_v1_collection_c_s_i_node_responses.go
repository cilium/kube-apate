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

// DeleteStorageV1CollectionCSINodeReader is a Reader for the DeleteStorageV1CollectionCSINode structure.
type DeleteStorageV1CollectionCSINodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageV1CollectionCSINodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteStorageV1CollectionCSINodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteStorageV1CollectionCSINodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteStorageV1CollectionCSINodeOK creates a DeleteStorageV1CollectionCSINodeOK with default headers values
func NewDeleteStorageV1CollectionCSINodeOK() *DeleteStorageV1CollectionCSINodeOK {
	return &DeleteStorageV1CollectionCSINodeOK{}
}

/*DeleteStorageV1CollectionCSINodeOK handles this case with default header values.

OK
*/
type DeleteStorageV1CollectionCSINodeOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteStorageV1CollectionCSINodeOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1/csinodes][%d] deleteStorageV1CollectionCSINodeOK  %+v", 200, o.Payload)
}

func (o *DeleteStorageV1CollectionCSINodeOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteStorageV1CollectionCSINodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStorageV1CollectionCSINodeUnauthorized creates a DeleteStorageV1CollectionCSINodeUnauthorized with default headers values
func NewDeleteStorageV1CollectionCSINodeUnauthorized() *DeleteStorageV1CollectionCSINodeUnauthorized {
	return &DeleteStorageV1CollectionCSINodeUnauthorized{}
}

/*DeleteStorageV1CollectionCSINodeUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteStorageV1CollectionCSINodeUnauthorized struct {
}

func (o *DeleteStorageV1CollectionCSINodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1/csinodes][%d] deleteStorageV1CollectionCSINodeUnauthorized ", 401)
}

func (o *DeleteStorageV1CollectionCSINodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
