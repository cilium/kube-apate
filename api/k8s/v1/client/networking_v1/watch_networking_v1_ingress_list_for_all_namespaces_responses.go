// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchNetworkingV1IngressListForAllNamespacesReader is a Reader for the WatchNetworkingV1IngressListForAllNamespaces structure.
type WatchNetworkingV1IngressListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchNetworkingV1IngressListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchNetworkingV1IngressListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchNetworkingV1IngressListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchNetworkingV1IngressListForAllNamespacesOK creates a WatchNetworkingV1IngressListForAllNamespacesOK with default headers values
func NewWatchNetworkingV1IngressListForAllNamespacesOK() *WatchNetworkingV1IngressListForAllNamespacesOK {
	return &WatchNetworkingV1IngressListForAllNamespacesOK{}
}

/*WatchNetworkingV1IngressListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchNetworkingV1IngressListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchNetworkingV1IngressListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/watch/ingresses][%d] watchNetworkingV1IngressListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchNetworkingV1IngressListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchNetworkingV1IngressListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchNetworkingV1IngressListForAllNamespacesUnauthorized creates a WatchNetworkingV1IngressListForAllNamespacesUnauthorized with default headers values
func NewWatchNetworkingV1IngressListForAllNamespacesUnauthorized() *WatchNetworkingV1IngressListForAllNamespacesUnauthorized {
	return &WatchNetworkingV1IngressListForAllNamespacesUnauthorized{}
}

/*WatchNetworkingV1IngressListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchNetworkingV1IngressListForAllNamespacesUnauthorized struct {
}

func (o *WatchNetworkingV1IngressListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1/watch/ingresses][%d] watchNetworkingV1IngressListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchNetworkingV1IngressListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
