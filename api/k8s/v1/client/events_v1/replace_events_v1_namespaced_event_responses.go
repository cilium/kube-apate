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

// ReplaceEventsV1NamespacedEventReader is a Reader for the ReplaceEventsV1NamespacedEvent structure.
type ReplaceEventsV1NamespacedEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceEventsV1NamespacedEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceEventsV1NamespacedEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceEventsV1NamespacedEventCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceEventsV1NamespacedEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceEventsV1NamespacedEventOK creates a ReplaceEventsV1NamespacedEventOK with default headers values
func NewReplaceEventsV1NamespacedEventOK() *ReplaceEventsV1NamespacedEventOK {
	return &ReplaceEventsV1NamespacedEventOK{}
}

/*ReplaceEventsV1NamespacedEventOK handles this case with default header values.

OK
*/
type ReplaceEventsV1NamespacedEventOK struct {
	Payload *models.IoK8sAPIEventsV1Event
}

func (o *ReplaceEventsV1NamespacedEventOK) Error() string {
	return fmt.Sprintf("[PUT /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}][%d] replaceEventsV1NamespacedEventOK  %+v", 200, o.Payload)
}

func (o *ReplaceEventsV1NamespacedEventOK) GetPayload() *models.IoK8sAPIEventsV1Event {
	return o.Payload
}

func (o *ReplaceEventsV1NamespacedEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIEventsV1Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceEventsV1NamespacedEventCreated creates a ReplaceEventsV1NamespacedEventCreated with default headers values
func NewReplaceEventsV1NamespacedEventCreated() *ReplaceEventsV1NamespacedEventCreated {
	return &ReplaceEventsV1NamespacedEventCreated{}
}

/*ReplaceEventsV1NamespacedEventCreated handles this case with default header values.

Created
*/
type ReplaceEventsV1NamespacedEventCreated struct {
	Payload *models.IoK8sAPIEventsV1Event
}

func (o *ReplaceEventsV1NamespacedEventCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}][%d] replaceEventsV1NamespacedEventCreated  %+v", 201, o.Payload)
}

func (o *ReplaceEventsV1NamespacedEventCreated) GetPayload() *models.IoK8sAPIEventsV1Event {
	return o.Payload
}

func (o *ReplaceEventsV1NamespacedEventCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIEventsV1Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceEventsV1NamespacedEventUnauthorized creates a ReplaceEventsV1NamespacedEventUnauthorized with default headers values
func NewReplaceEventsV1NamespacedEventUnauthorized() *ReplaceEventsV1NamespacedEventUnauthorized {
	return &ReplaceEventsV1NamespacedEventUnauthorized{}
}

/*ReplaceEventsV1NamespacedEventUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceEventsV1NamespacedEventUnauthorized struct {
}

func (o *ReplaceEventsV1NamespacedEventUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}][%d] replaceEventsV1NamespacedEventUnauthorized ", 401)
}

func (o *ReplaceEventsV1NamespacedEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}