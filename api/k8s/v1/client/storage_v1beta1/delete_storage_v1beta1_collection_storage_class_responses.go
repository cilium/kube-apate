// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteStorageV1beta1CollectionStorageClassReader is a Reader for the DeleteStorageV1beta1CollectionStorageClass structure.
type DeleteStorageV1beta1CollectionStorageClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageV1beta1CollectionStorageClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteStorageV1beta1CollectionStorageClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteStorageV1beta1CollectionStorageClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteStorageV1beta1CollectionStorageClassOK creates a DeleteStorageV1beta1CollectionStorageClassOK with default headers values
func NewDeleteStorageV1beta1CollectionStorageClassOK() *DeleteStorageV1beta1CollectionStorageClassOK {
	return &DeleteStorageV1beta1CollectionStorageClassOK{}
}

/*DeleteStorageV1beta1CollectionStorageClassOK handles this case with default header values.

OK
*/
type DeleteStorageV1beta1CollectionStorageClassOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteStorageV1beta1CollectionStorageClassOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1beta1/storageclasses][%d] deleteStorageV1beta1CollectionStorageClassOK  %+v", 200, o.Payload)
}

func (o *DeleteStorageV1beta1CollectionStorageClassOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteStorageV1beta1CollectionStorageClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStorageV1beta1CollectionStorageClassUnauthorized creates a DeleteStorageV1beta1CollectionStorageClassUnauthorized with default headers values
func NewDeleteStorageV1beta1CollectionStorageClassUnauthorized() *DeleteStorageV1beta1CollectionStorageClassUnauthorized {
	return &DeleteStorageV1beta1CollectionStorageClassUnauthorized{}
}

/*DeleteStorageV1beta1CollectionStorageClassUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteStorageV1beta1CollectionStorageClassUnauthorized struct {
}

func (o *DeleteStorageV1beta1CollectionStorageClassUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/storage.k8s.io/v1beta1/storageclasses][%d] deleteStorageV1beta1CollectionStorageClassUnauthorized ", 401)
}

func (o *DeleteStorageV1beta1CollectionStorageClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
