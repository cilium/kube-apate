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

// DeleteCoreV1NamespacedSecretReader is a Reader for the DeleteCoreV1NamespacedSecret structure.
type DeleteCoreV1NamespacedSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoreV1NamespacedSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoreV1NamespacedSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteCoreV1NamespacedSecretAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoreV1NamespacedSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoreV1NamespacedSecretOK creates a DeleteCoreV1NamespacedSecretOK with default headers values
func NewDeleteCoreV1NamespacedSecretOK() *DeleteCoreV1NamespacedSecretOK {
	return &DeleteCoreV1NamespacedSecretOK{}
}

/*DeleteCoreV1NamespacedSecretOK handles this case with default header values.

OK
*/
type DeleteCoreV1NamespacedSecretOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoreV1NamespacedSecretOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/secrets/{name}][%d] deleteCoreV1NamespacedSecretOK  %+v", 200, o.Payload)
}

func (o *DeleteCoreV1NamespacedSecretOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoreV1NamespacedSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1NamespacedSecretAccepted creates a DeleteCoreV1NamespacedSecretAccepted with default headers values
func NewDeleteCoreV1NamespacedSecretAccepted() *DeleteCoreV1NamespacedSecretAccepted {
	return &DeleteCoreV1NamespacedSecretAccepted{}
}

/*DeleteCoreV1NamespacedSecretAccepted handles this case with default header values.

Accepted
*/
type DeleteCoreV1NamespacedSecretAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteCoreV1NamespacedSecretAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/secrets/{name}][%d] deleteCoreV1NamespacedSecretAccepted  %+v", 202, o.Payload)
}

func (o *DeleteCoreV1NamespacedSecretAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteCoreV1NamespacedSecretAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1NamespacedSecretUnauthorized creates a DeleteCoreV1NamespacedSecretUnauthorized with default headers values
func NewDeleteCoreV1NamespacedSecretUnauthorized() *DeleteCoreV1NamespacedSecretUnauthorized {
	return &DeleteCoreV1NamespacedSecretUnauthorized{}
}

/*DeleteCoreV1NamespacedSecretUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoreV1NamespacedSecretUnauthorized struct {
}

func (o *DeleteCoreV1NamespacedSecretUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/secrets/{name}][%d] deleteCoreV1NamespacedSecretUnauthorized ", 401)
}

func (o *DeleteCoreV1NamespacedSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
