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

// DeleteEventsV1NamespacedEventReader is a Reader for the DeleteEventsV1NamespacedEvent structure.
type DeleteEventsV1NamespacedEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEventsV1NamespacedEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteEventsV1NamespacedEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteEventsV1NamespacedEventAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteEventsV1NamespacedEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEventsV1NamespacedEventOK creates a DeleteEventsV1NamespacedEventOK with default headers values
func NewDeleteEventsV1NamespacedEventOK() *DeleteEventsV1NamespacedEventOK {
	return &DeleteEventsV1NamespacedEventOK{}
}

/*DeleteEventsV1NamespacedEventOK handles this case with default header values.

OK
*/
type DeleteEventsV1NamespacedEventOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteEventsV1NamespacedEventOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}][%d] deleteEventsV1NamespacedEventOK  %+v", 200, o.Payload)
}

func (o *DeleteEventsV1NamespacedEventOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteEventsV1NamespacedEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEventsV1NamespacedEventAccepted creates a DeleteEventsV1NamespacedEventAccepted with default headers values
func NewDeleteEventsV1NamespacedEventAccepted() *DeleteEventsV1NamespacedEventAccepted {
	return &DeleteEventsV1NamespacedEventAccepted{}
}

/*DeleteEventsV1NamespacedEventAccepted handles this case with default header values.

Accepted
*/
type DeleteEventsV1NamespacedEventAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteEventsV1NamespacedEventAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}][%d] deleteEventsV1NamespacedEventAccepted  %+v", 202, o.Payload)
}

func (o *DeleteEventsV1NamespacedEventAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteEventsV1NamespacedEventAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEventsV1NamespacedEventUnauthorized creates a DeleteEventsV1NamespacedEventUnauthorized with default headers values
func NewDeleteEventsV1NamespacedEventUnauthorized() *DeleteEventsV1NamespacedEventUnauthorized {
	return &DeleteEventsV1NamespacedEventUnauthorized{}
}

/*DeleteEventsV1NamespacedEventUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteEventsV1NamespacedEventUnauthorized struct {
}

func (o *DeleteEventsV1NamespacedEventUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/events.k8s.io/v1/namespaces/{namespace}/events/{name}][%d] deleteEventsV1NamespacedEventUnauthorized ", 401)
}

func (o *DeleteEventsV1NamespacedEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}