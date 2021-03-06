// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetStorageAPIGroupReader is a Reader for the GetStorageAPIGroup structure.
type GetStorageAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetStorageAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStorageAPIGroupOK creates a GetStorageAPIGroupOK with default headers values
func NewGetStorageAPIGroupOK() *GetStorageAPIGroupOK {
	return &GetStorageAPIGroupOK{}
}

/*GetStorageAPIGroupOK handles this case with default header values.

OK
*/
type GetStorageAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetStorageAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/][%d] getStorageApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetStorageAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetStorageAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageAPIGroupUnauthorized creates a GetStorageAPIGroupUnauthorized with default headers values
func NewGetStorageAPIGroupUnauthorized() *GetStorageAPIGroupUnauthorized {
	return &GetStorageAPIGroupUnauthorized{}
}

/*GetStorageAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetStorageAPIGroupUnauthorized struct {
}

func (o *GetStorageAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/][%d] getStorageApiGroupUnauthorized ", 401)
}

func (o *GetStorageAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
