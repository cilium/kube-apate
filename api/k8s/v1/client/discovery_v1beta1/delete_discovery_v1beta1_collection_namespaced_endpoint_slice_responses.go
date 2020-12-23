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

// DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceReader is a Reader for the DeleteDiscoveryV1beta1CollectionNamespacedEndpointSlice structure.
type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK creates a DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK with default headers values
func NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK() *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK {
	return &DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK{}
}

/*DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK handles this case with default header values.

OK
*/
type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices][%d] deleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK  %+v", 200, o.Payload)
}

func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized creates a DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized with default headers values
func NewDeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized() *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized {
	return &DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized{}
}

/*DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized struct {
}

func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices][%d] deleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized ", 401)
}

func (o *DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}