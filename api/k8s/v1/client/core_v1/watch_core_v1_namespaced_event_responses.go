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

// WatchCoreV1NamespacedEventReader is a Reader for the WatchCoreV1NamespacedEvent structure.
type WatchCoreV1NamespacedEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedEventOK creates a WatchCoreV1NamespacedEventOK with default headers values
func NewWatchCoreV1NamespacedEventOK() *WatchCoreV1NamespacedEventOK {
	return &WatchCoreV1NamespacedEventOK{}
}

/*WatchCoreV1NamespacedEventOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedEventOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedEventOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/events/{name}][%d] watchCoreV1NamespacedEventOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedEventOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedEventUnauthorized creates a WatchCoreV1NamespacedEventUnauthorized with default headers values
func NewWatchCoreV1NamespacedEventUnauthorized() *WatchCoreV1NamespacedEventUnauthorized {
	return &WatchCoreV1NamespacedEventUnauthorized{}
}

/*WatchCoreV1NamespacedEventUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedEventUnauthorized struct {
}

func (o *WatchCoreV1NamespacedEventUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/events/{name}][%d] watchCoreV1NamespacedEventUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
