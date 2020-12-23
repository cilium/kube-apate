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

// WatchCoreV1NamespacedLimitRangeReader is a Reader for the WatchCoreV1NamespacedLimitRange structure.
type WatchCoreV1NamespacedLimitRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedLimitRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedLimitRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedLimitRangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedLimitRangeOK creates a WatchCoreV1NamespacedLimitRangeOK with default headers values
func NewWatchCoreV1NamespacedLimitRangeOK() *WatchCoreV1NamespacedLimitRangeOK {
	return &WatchCoreV1NamespacedLimitRangeOK{}
}

/*WatchCoreV1NamespacedLimitRangeOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedLimitRangeOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedLimitRangeOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/limitranges/{name}][%d] watchCoreV1NamespacedLimitRangeOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedLimitRangeOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedLimitRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedLimitRangeUnauthorized creates a WatchCoreV1NamespacedLimitRangeUnauthorized with default headers values
func NewWatchCoreV1NamespacedLimitRangeUnauthorized() *WatchCoreV1NamespacedLimitRangeUnauthorized {
	return &WatchCoreV1NamespacedLimitRangeUnauthorized{}
}

/*WatchCoreV1NamespacedLimitRangeUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedLimitRangeUnauthorized struct {
}

func (o *WatchCoreV1NamespacedLimitRangeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/limitranges/{name}][%d] watchCoreV1NamespacedLimitRangeUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedLimitRangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}