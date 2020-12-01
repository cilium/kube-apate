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

// WatchCoreV1NamespacedSecretReader is a Reader for the WatchCoreV1NamespacedSecret structure.
type WatchCoreV1NamespacedSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedSecretUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedSecretOK creates a WatchCoreV1NamespacedSecretOK with default headers values
func NewWatchCoreV1NamespacedSecretOK() *WatchCoreV1NamespacedSecretOK {
	return &WatchCoreV1NamespacedSecretOK{}
}

/*WatchCoreV1NamespacedSecretOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedSecretOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedSecretOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/secrets/{name}][%d] watchCoreV1NamespacedSecretOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedSecretOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedSecretUnauthorized creates a WatchCoreV1NamespacedSecretUnauthorized with default headers values
func NewWatchCoreV1NamespacedSecretUnauthorized() *WatchCoreV1NamespacedSecretUnauthorized {
	return &WatchCoreV1NamespacedSecretUnauthorized{}
}

/*WatchCoreV1NamespacedSecretUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedSecretUnauthorized struct {
}

func (o *WatchCoreV1NamespacedSecretUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/secrets/{name}][%d] watchCoreV1NamespacedSecretUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedSecretUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
