// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package events_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchEventsV1beta1NamespacedEventReader is a Reader for the PatchEventsV1beta1NamespacedEvent structure.
type PatchEventsV1beta1NamespacedEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEventsV1beta1NamespacedEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEventsV1beta1NamespacedEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEventsV1beta1NamespacedEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchEventsV1beta1NamespacedEventOK creates a PatchEventsV1beta1NamespacedEventOK with default headers values
func NewPatchEventsV1beta1NamespacedEventOK() *PatchEventsV1beta1NamespacedEventOK {
	return &PatchEventsV1beta1NamespacedEventOK{}
}

/*PatchEventsV1beta1NamespacedEventOK handles this case with default header values.

OK
*/
type PatchEventsV1beta1NamespacedEventOK struct {
	Payload *models.IoK8sAPIEventsV1beta1Event
}

func (o *PatchEventsV1beta1NamespacedEventOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/events.k8s.io/v1beta1/namespaces/{namespace}/events/{name}][%d] patchEventsV1beta1NamespacedEventOK  %+v", 200, o.Payload)
}

func (o *PatchEventsV1beta1NamespacedEventOK) GetPayload() *models.IoK8sAPIEventsV1beta1Event {
	return o.Payload
}

func (o *PatchEventsV1beta1NamespacedEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIEventsV1beta1Event)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEventsV1beta1NamespacedEventUnauthorized creates a PatchEventsV1beta1NamespacedEventUnauthorized with default headers values
func NewPatchEventsV1beta1NamespacedEventUnauthorized() *PatchEventsV1beta1NamespacedEventUnauthorized {
	return &PatchEventsV1beta1NamespacedEventUnauthorized{}
}

/*PatchEventsV1beta1NamespacedEventUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchEventsV1beta1NamespacedEventUnauthorized struct {
}

func (o *PatchEventsV1beta1NamespacedEventUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/events.k8s.io/v1beta1/namespaces/{namespace}/events/{name}][%d] patchEventsV1beta1NamespacedEventUnauthorized ", 401)
}

func (o *PatchEventsV1beta1NamespacedEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}