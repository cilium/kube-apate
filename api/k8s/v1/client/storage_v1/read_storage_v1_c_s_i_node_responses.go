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

// ReadStorageV1CSINodeReader is a Reader for the ReadStorageV1CSINode structure.
type ReadStorageV1CSINodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadStorageV1CSINodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadStorageV1CSINodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadStorageV1CSINodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadStorageV1CSINodeOK creates a ReadStorageV1CSINodeOK with default headers values
func NewReadStorageV1CSINodeOK() *ReadStorageV1CSINodeOK {
	return &ReadStorageV1CSINodeOK{}
}

/*ReadStorageV1CSINodeOK handles this case with default header values.

OK
*/
type ReadStorageV1CSINodeOK struct {
	Payload *models.IoK8sAPIStorageV1CSINode
}

func (o *ReadStorageV1CSINodeOK) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/csinodes/{name}][%d] readStorageV1CSINodeOK  %+v", 200, o.Payload)
}

func (o *ReadStorageV1CSINodeOK) GetPayload() *models.IoK8sAPIStorageV1CSINode {
	return o.Payload
}

func (o *ReadStorageV1CSINodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1CSINode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadStorageV1CSINodeUnauthorized creates a ReadStorageV1CSINodeUnauthorized with default headers values
func NewReadStorageV1CSINodeUnauthorized() *ReadStorageV1CSINodeUnauthorized {
	return &ReadStorageV1CSINodeUnauthorized{}
}

/*ReadStorageV1CSINodeUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadStorageV1CSINodeUnauthorized struct {
}

func (o *ReadStorageV1CSINodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1/csinodes/{name}][%d] readStorageV1CSINodeUnauthorized ", 401)
}

func (o *ReadStorageV1CSINodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
