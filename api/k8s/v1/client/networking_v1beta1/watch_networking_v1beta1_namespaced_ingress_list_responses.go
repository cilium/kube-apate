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

// WatchNetworkingV1beta1NamespacedIngressListReader is a Reader for the WatchNetworkingV1beta1NamespacedIngressList structure.
type WatchNetworkingV1beta1NamespacedIngressListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchNetworkingV1beta1NamespacedIngressListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchNetworkingV1beta1NamespacedIngressListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchNetworkingV1beta1NamespacedIngressListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchNetworkingV1beta1NamespacedIngressListOK creates a WatchNetworkingV1beta1NamespacedIngressListOK with default headers values
func NewWatchNetworkingV1beta1NamespacedIngressListOK() *WatchNetworkingV1beta1NamespacedIngressListOK {
	return &WatchNetworkingV1beta1NamespacedIngressListOK{}
}

/*WatchNetworkingV1beta1NamespacedIngressListOK handles this case with default header values.

OK
*/
type WatchNetworkingV1beta1NamespacedIngressListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchNetworkingV1beta1NamespacedIngressListOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1beta1/watch/namespaces/{namespace}/ingresses][%d] watchNetworkingV1beta1NamespacedIngressListOK  %+v", 200, o.Payload)
}

func (o *WatchNetworkingV1beta1NamespacedIngressListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchNetworkingV1beta1NamespacedIngressListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchNetworkingV1beta1NamespacedIngressListUnauthorized creates a WatchNetworkingV1beta1NamespacedIngressListUnauthorized with default headers values
func NewWatchNetworkingV1beta1NamespacedIngressListUnauthorized() *WatchNetworkingV1beta1NamespacedIngressListUnauthorized {
	return &WatchNetworkingV1beta1NamespacedIngressListUnauthorized{}
}

/*WatchNetworkingV1beta1NamespacedIngressListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchNetworkingV1beta1NamespacedIngressListUnauthorized struct {
}

func (o *WatchNetworkingV1beta1NamespacedIngressListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/v1beta1/watch/namespaces/{namespace}/ingresses][%d] watchNetworkingV1beta1NamespacedIngressListUnauthorized ", 401)
}

func (o *WatchNetworkingV1beta1NamespacedIngressListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
