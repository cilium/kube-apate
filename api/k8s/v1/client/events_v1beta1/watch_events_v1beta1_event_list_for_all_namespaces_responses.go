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

// WatchEventsV1beta1EventListForAllNamespacesReader is a Reader for the WatchEventsV1beta1EventListForAllNamespaces structure.
type WatchEventsV1beta1EventListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchEventsV1beta1EventListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchEventsV1beta1EventListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchEventsV1beta1EventListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchEventsV1beta1EventListForAllNamespacesOK creates a WatchEventsV1beta1EventListForAllNamespacesOK with default headers values
func NewWatchEventsV1beta1EventListForAllNamespacesOK() *WatchEventsV1beta1EventListForAllNamespacesOK {
	return &WatchEventsV1beta1EventListForAllNamespacesOK{}
}

/*WatchEventsV1beta1EventListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchEventsV1beta1EventListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchEventsV1beta1EventListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/events.k8s.io/v1beta1/watch/events][%d] watchEventsV1beta1EventListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchEventsV1beta1EventListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchEventsV1beta1EventListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchEventsV1beta1EventListForAllNamespacesUnauthorized creates a WatchEventsV1beta1EventListForAllNamespacesUnauthorized with default headers values
func NewWatchEventsV1beta1EventListForAllNamespacesUnauthorized() *WatchEventsV1beta1EventListForAllNamespacesUnauthorized {
	return &WatchEventsV1beta1EventListForAllNamespacesUnauthorized{}
}

/*WatchEventsV1beta1EventListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchEventsV1beta1EventListForAllNamespacesUnauthorized struct {
}

func (o *WatchEventsV1beta1EventListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/events.k8s.io/v1beta1/watch/events][%d] watchEventsV1beta1EventListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchEventsV1beta1EventListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}