// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package scheduling_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchSchedulingV1PriorityClassListReader is a Reader for the WatchSchedulingV1PriorityClassList structure.
type WatchSchedulingV1PriorityClassListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchSchedulingV1PriorityClassListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchSchedulingV1PriorityClassListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchSchedulingV1PriorityClassListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchSchedulingV1PriorityClassListOK creates a WatchSchedulingV1PriorityClassListOK with default headers values
func NewWatchSchedulingV1PriorityClassListOK() *WatchSchedulingV1PriorityClassListOK {
	return &WatchSchedulingV1PriorityClassListOK{}
}

/*WatchSchedulingV1PriorityClassListOK handles this case with default header values.

OK
*/
type WatchSchedulingV1PriorityClassListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchSchedulingV1PriorityClassListOK) Error() string {
	return fmt.Sprintf("[GET /apis/scheduling.k8s.io/v1/watch/priorityclasses][%d] watchSchedulingV1PriorityClassListOK  %+v", 200, o.Payload)
}

func (o *WatchSchedulingV1PriorityClassListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchSchedulingV1PriorityClassListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchSchedulingV1PriorityClassListUnauthorized creates a WatchSchedulingV1PriorityClassListUnauthorized with default headers values
func NewWatchSchedulingV1PriorityClassListUnauthorized() *WatchSchedulingV1PriorityClassListUnauthorized {
	return &WatchSchedulingV1PriorityClassListUnauthorized{}
}

/*WatchSchedulingV1PriorityClassListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchSchedulingV1PriorityClassListUnauthorized struct {
}

func (o *WatchSchedulingV1PriorityClassListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/scheduling.k8s.io/v1/watch/priorityclasses][%d] watchSchedulingV1PriorityClassListUnauthorized ", 401)
}

func (o *WatchSchedulingV1PriorityClassListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
