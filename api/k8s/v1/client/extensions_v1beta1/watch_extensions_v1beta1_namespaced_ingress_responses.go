// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package extensions_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchExtensionsV1beta1NamespacedIngressReader is a Reader for the WatchExtensionsV1beta1NamespacedIngress structure.
type WatchExtensionsV1beta1NamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchExtensionsV1beta1NamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchExtensionsV1beta1NamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchExtensionsV1beta1NamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchExtensionsV1beta1NamespacedIngressOK creates a WatchExtensionsV1beta1NamespacedIngressOK with default headers values
func NewWatchExtensionsV1beta1NamespacedIngressOK() *WatchExtensionsV1beta1NamespacedIngressOK {
	return &WatchExtensionsV1beta1NamespacedIngressOK{}
}

/*WatchExtensionsV1beta1NamespacedIngressOK handles this case with default header values.

OK
*/
type WatchExtensionsV1beta1NamespacedIngressOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchExtensionsV1beta1NamespacedIngressOK) Error() string {
	return fmt.Sprintf("[GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/ingresses/{name}][%d] watchExtensionsV1beta1NamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *WatchExtensionsV1beta1NamespacedIngressOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchExtensionsV1beta1NamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchExtensionsV1beta1NamespacedIngressUnauthorized creates a WatchExtensionsV1beta1NamespacedIngressUnauthorized with default headers values
func NewWatchExtensionsV1beta1NamespacedIngressUnauthorized() *WatchExtensionsV1beta1NamespacedIngressUnauthorized {
	return &WatchExtensionsV1beta1NamespacedIngressUnauthorized{}
}

/*WatchExtensionsV1beta1NamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchExtensionsV1beta1NamespacedIngressUnauthorized struct {
}

func (o *WatchExtensionsV1beta1NamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/ingresses/{name}][%d] watchExtensionsV1beta1NamespacedIngressUnauthorized ", 401)
}

func (o *WatchExtensionsV1beta1NamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
