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

// WatchCoreV1LimitRangeListForAllNamespacesReader is a Reader for the WatchCoreV1LimitRangeListForAllNamespaces structure.
type WatchCoreV1LimitRangeListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1LimitRangeListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1LimitRangeListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1LimitRangeListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1LimitRangeListForAllNamespacesOK creates a WatchCoreV1LimitRangeListForAllNamespacesOK with default headers values
func NewWatchCoreV1LimitRangeListForAllNamespacesOK() *WatchCoreV1LimitRangeListForAllNamespacesOK {
	return &WatchCoreV1LimitRangeListForAllNamespacesOK{}
}

/*WatchCoreV1LimitRangeListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchCoreV1LimitRangeListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1LimitRangeListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/limitranges][%d] watchCoreV1LimitRangeListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1LimitRangeListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1LimitRangeListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1LimitRangeListForAllNamespacesUnauthorized creates a WatchCoreV1LimitRangeListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1LimitRangeListForAllNamespacesUnauthorized() *WatchCoreV1LimitRangeListForAllNamespacesUnauthorized {
	return &WatchCoreV1LimitRangeListForAllNamespacesUnauthorized{}
}

/*WatchCoreV1LimitRangeListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1LimitRangeListForAllNamespacesUnauthorized struct {
}

func (o *WatchCoreV1LimitRangeListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/limitranges][%d] watchCoreV1LimitRangeListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchCoreV1LimitRangeListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
