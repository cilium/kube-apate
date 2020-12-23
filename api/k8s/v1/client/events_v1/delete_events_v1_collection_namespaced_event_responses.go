// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// DeleteEventsV1CollectionNamespacedEventReader is a Reader for the DeleteEventsV1CollectionNamespacedEvent structure.
type DeleteEventsV1CollectionNamespacedEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEventsV1CollectionNamespacedEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEventsV1CollectionNamespacedEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEventsV1CollectionNamespacedEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEventsV1CollectionNamespacedEventOK creates a DeleteEventsV1CollectionNamespacedEventOK with default headers values
func NewDeleteEventsV1CollectionNamespacedEventOK() *DeleteEventsV1CollectionNamespacedEventOK {
	return &DeleteEventsV1CollectionNamespacedEventOK{}
}

/*DeleteEventsV1CollectionNamespacedEventOK handles this case with default header values.

OK
*/
type DeleteEventsV1CollectionNamespacedEventOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteEventsV1CollectionNamespacedEventOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events][%d] deleteEventsV1CollectionNamespacedEventOK  %+v", 200, o.Payload)
}

func (o *DeleteEventsV1CollectionNamespacedEventOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteEventsV1CollectionNamespacedEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEventsV1CollectionNamespacedEventUnauthorized creates a DeleteEventsV1CollectionNamespacedEventUnauthorized with default headers values
func NewDeleteEventsV1CollectionNamespacedEventUnauthorized() *DeleteEventsV1CollectionNamespacedEventUnauthorized {
	return &DeleteEventsV1CollectionNamespacedEventUnauthorized{}
}

/*DeleteEventsV1CollectionNamespacedEventUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteEventsV1CollectionNamespacedEventUnauthorized struct {
}

func (o *DeleteEventsV1CollectionNamespacedEventUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events][%d] deleteEventsV1CollectionNamespacedEventUnauthorized ", 401)
}

func (o *DeleteEventsV1CollectionNamespacedEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}