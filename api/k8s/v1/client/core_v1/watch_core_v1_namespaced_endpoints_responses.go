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

// WatchCoreV1NamespacedEndpointsReader is a Reader for the WatchCoreV1NamespacedEndpoints structure.
type WatchCoreV1NamespacedEndpointsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedEndpointsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedEndpointsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedEndpointsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedEndpointsOK creates a WatchCoreV1NamespacedEndpointsOK with default headers values
func NewWatchCoreV1NamespacedEndpointsOK() *WatchCoreV1NamespacedEndpointsOK {
	return &WatchCoreV1NamespacedEndpointsOK{}
}

/*WatchCoreV1NamespacedEndpointsOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedEndpointsOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedEndpointsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/endpoints/{name}][%d] watchCoreV1NamespacedEndpointsOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedEndpointsOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedEndpointsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedEndpointsUnauthorized creates a WatchCoreV1NamespacedEndpointsUnauthorized with default headers values
func NewWatchCoreV1NamespacedEndpointsUnauthorized() *WatchCoreV1NamespacedEndpointsUnauthorized {
	return &WatchCoreV1NamespacedEndpointsUnauthorized{}
}

/*WatchCoreV1NamespacedEndpointsUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedEndpointsUnauthorized struct {
}

func (o *WatchCoreV1NamespacedEndpointsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/endpoints/{name}][%d] watchCoreV1NamespacedEndpointsUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedEndpointsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
