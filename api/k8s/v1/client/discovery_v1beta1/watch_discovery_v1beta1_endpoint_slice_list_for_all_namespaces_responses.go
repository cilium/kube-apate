// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package discovery_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesReader is a Reader for the WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces structure.
type WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK creates a WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK with default headers values
func NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK() *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK {
	return &WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK{}
}

/*WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/discovery.k8s.io/v1beta1/watch/endpointslices][%d] watchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized creates a WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized with default headers values
func NewWatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized() *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized {
	return &WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized{}
}

/*WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized struct {
}

func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/discovery.k8s.io/v1beta1/watch/endpointslices][%d] watchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
