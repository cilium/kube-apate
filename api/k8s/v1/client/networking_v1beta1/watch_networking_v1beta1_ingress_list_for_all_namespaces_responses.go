// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchNetworkingV1beta1IngressListForAllNamespacesReader is a Reader for the WatchNetworkingV1beta1IngressListForAllNamespaces structure.
type WatchNetworkingV1beta1IngressListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchNetworkingV1beta1IngressListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchNetworkingV1beta1IngressListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchNetworkingV1beta1IngressListForAllNamespacesOK creates a WatchNetworkingV1beta1IngressListForAllNamespacesOK with default headers values
func NewWatchNetworkingV1beta1IngressListForAllNamespacesOK() *WatchNetworkingV1beta1IngressListForAllNamespacesOK {
	return &WatchNetworkingV1beta1IngressListForAllNamespacesOK{}
}

/*WatchNetworkingV1beta1IngressListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchNetworkingV1beta1IngressListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchNetworkingV1beta1IngressListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1beta1/watch/ingresses][%d] watchNetworkingV1beta1IngressListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchNetworkingV1beta1IngressListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchNetworkingV1beta1IngressListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized creates a WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized with default headers values
func NewWatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized() *WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized {
	return &WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized{}
}

/*WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized struct {
}

func (o *WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1beta1/watch/ingresses][%d] watchNetworkingV1beta1IngressListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchNetworkingV1beta1IngressListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
