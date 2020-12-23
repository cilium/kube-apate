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

// ReadDiscoveryV1beta1NamespacedEndpointSliceReader is a Reader for the ReadDiscoveryV1beta1NamespacedEndpointSlice structure.
type ReadDiscoveryV1beta1NamespacedEndpointSliceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadDiscoveryV1beta1NamespacedEndpointSliceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadDiscoveryV1beta1NamespacedEndpointSliceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadDiscoveryV1beta1NamespacedEndpointSliceOK creates a ReadDiscoveryV1beta1NamespacedEndpointSliceOK with default headers values
func NewReadDiscoveryV1beta1NamespacedEndpointSliceOK() *ReadDiscoveryV1beta1NamespacedEndpointSliceOK {
	return &ReadDiscoveryV1beta1NamespacedEndpointSliceOK{}
}

/*ReadDiscoveryV1beta1NamespacedEndpointSliceOK handles this case with default header values.

OK
*/
type ReadDiscoveryV1beta1NamespacedEndpointSliceOK struct {
	Payload *models.IoK8sAPIDiscoveryV1beta1EndpointSlice
}

func (o *ReadDiscoveryV1beta1NamespacedEndpointSliceOK) Error() string {
	return fmt.Sprintf("[GET /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}][%d] readDiscoveryV1beta1NamespacedEndpointSliceOK  %+v", 200, o.Payload)
}

func (o *ReadDiscoveryV1beta1NamespacedEndpointSliceOK) GetPayload() *models.IoK8sAPIDiscoveryV1beta1EndpointSlice {
	return o.Payload
}

func (o *ReadDiscoveryV1beta1NamespacedEndpointSliceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIDiscoveryV1beta1EndpointSlice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized creates a ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized with default headers values
func NewReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized() *ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized {
	return &ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized{}
}

/*ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized struct {
}

func (o *ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}][%d] readDiscoveryV1beta1NamespacedEndpointSliceUnauthorized ", 401)
}

func (o *ReadDiscoveryV1beta1NamespacedEndpointSliceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}