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

// DeleteCoreV1CollectionNamespacedSecretReader is a Reader for the DeleteCoreV1CollectionNamespacedSecret structure.
type DeleteCoreV1CollectionNamespacedSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoreV1CollectionNamespacedSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoreV1CollectionNamespacedSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoreV1CollectionNamespacedSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoreV1CollectionNamespacedSecretOK creates a DeleteCoreV1CollectionNamespacedSecretOK with default headers values
func NewDeleteCoreV1CollectionNamespacedSecretOK() *DeleteCoreV1CollectionNamespacedSecretOK {
	return &DeleteCoreV1CollectionNamespacedSecretOK{}
}

/*DeleteCoreV1CollectionNamespacedSecretOK handles this case with default header values.

OK
*/
type DeleteCoreV1CollectionNamespacedSecretOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoreV1CollectionNamespacedSecretOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/secrets][%d] deleteCoreV1CollectionNamespacedSecretOK  %+v", 200, o.Payload)
}

func (o *DeleteCoreV1CollectionNamespacedSecretOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoreV1CollectionNamespacedSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1CollectionNamespacedSecretUnauthorized creates a DeleteCoreV1CollectionNamespacedSecretUnauthorized with default headers values
func NewDeleteCoreV1CollectionNamespacedSecretUnauthorized() *DeleteCoreV1CollectionNamespacedSecretUnauthorized {
	return &DeleteCoreV1CollectionNamespacedSecretUnauthorized{}
}

/*DeleteCoreV1CollectionNamespacedSecretUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoreV1CollectionNamespacedSecretUnauthorized struct {
}

func (o *DeleteCoreV1CollectionNamespacedSecretUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/secrets][%d] deleteCoreV1CollectionNamespacedSecretUnauthorized ", 401)
}

func (o *DeleteCoreV1CollectionNamespacedSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
