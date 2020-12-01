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

// WatchCoreV1NamespacedEndpointsListReader is a Reader for the WatchCoreV1NamespacedEndpointsList structure.
type WatchCoreV1NamespacedEndpointsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedEndpointsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedEndpointsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedEndpointsListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedEndpointsListOK creates a WatchCoreV1NamespacedEndpointsListOK with default headers values
func NewWatchCoreV1NamespacedEndpointsListOK() *WatchCoreV1NamespacedEndpointsListOK {
	return &WatchCoreV1NamespacedEndpointsListOK{}
}

/*WatchCoreV1NamespacedEndpointsListOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedEndpointsListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedEndpointsListOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/endpoints][%d] watchCoreV1NamespacedEndpointsListOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedEndpointsListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedEndpointsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedEndpointsListUnauthorized creates a WatchCoreV1NamespacedEndpointsListUnauthorized with default headers values
func NewWatchCoreV1NamespacedEndpointsListUnauthorized() *WatchCoreV1NamespacedEndpointsListUnauthorized {
	return &WatchCoreV1NamespacedEndpointsListUnauthorized{}
}

/*WatchCoreV1NamespacedEndpointsListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedEndpointsListUnauthorized struct {
}

func (o *WatchCoreV1NamespacedEndpointsListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/endpoints][%d] watchCoreV1NamespacedEndpointsListUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedEndpointsListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
