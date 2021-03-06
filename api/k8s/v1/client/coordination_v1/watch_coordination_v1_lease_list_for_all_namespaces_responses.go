// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchCoordinationV1LeaseListForAllNamespacesReader is a Reader for the WatchCoordinationV1LeaseListForAllNamespaces structure.
type WatchCoordinationV1LeaseListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoordinationV1LeaseListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoordinationV1LeaseListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoordinationV1LeaseListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoordinationV1LeaseListForAllNamespacesOK creates a WatchCoordinationV1LeaseListForAllNamespacesOK with default headers values
func NewWatchCoordinationV1LeaseListForAllNamespacesOK() *WatchCoordinationV1LeaseListForAllNamespacesOK {
	return &WatchCoordinationV1LeaseListForAllNamespacesOK{}
}

/*WatchCoordinationV1LeaseListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchCoordinationV1LeaseListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoordinationV1LeaseListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1/watch/leases][%d] watchCoordinationV1LeaseListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchCoordinationV1LeaseListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoordinationV1LeaseListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoordinationV1LeaseListForAllNamespacesUnauthorized creates a WatchCoordinationV1LeaseListForAllNamespacesUnauthorized with default headers values
func NewWatchCoordinationV1LeaseListForAllNamespacesUnauthorized() *WatchCoordinationV1LeaseListForAllNamespacesUnauthorized {
	return &WatchCoordinationV1LeaseListForAllNamespacesUnauthorized{}
}

/*WatchCoordinationV1LeaseListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoordinationV1LeaseListForAllNamespacesUnauthorized struct {
}

func (o *WatchCoordinationV1LeaseListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1/watch/leases][%d] watchCoordinationV1LeaseListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchCoordinationV1LeaseListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
