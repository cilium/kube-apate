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

// DeleteDiscoveryV1beta1NamespacedEndpointSliceReader is a Reader for the DeleteDiscoveryV1beta1NamespacedEndpointSlice structure.
type DeleteDiscoveryV1beta1NamespacedEndpointSliceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteDiscoveryV1beta1NamespacedEndpointSliceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteDiscoveryV1beta1NamespacedEndpointSliceOK creates a DeleteDiscoveryV1beta1NamespacedEndpointSliceOK with default headers values
func NewDeleteDiscoveryV1beta1NamespacedEndpointSliceOK() *DeleteDiscoveryV1beta1NamespacedEndpointSliceOK {
	return &DeleteDiscoveryV1beta1NamespacedEndpointSliceOK{}
}

/*DeleteDiscoveryV1beta1NamespacedEndpointSliceOK handles this case with default header values.

OK
*/
type DeleteDiscoveryV1beta1NamespacedEndpointSliceOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}][%d] deleteDiscoveryV1beta1NamespacedEndpointSliceOK  %+v", 200, o.Payload)
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted creates a DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted with default headers values
func NewDeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted() *DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted {
	return &DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted{}
}

/*DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted handles this case with default header values.

Accepted
*/
type DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}][%d] deleteDiscoveryV1beta1NamespacedEndpointSliceAccepted  %+v", 202, o.Payload)
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized creates a DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized with default headers values
func NewDeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized() *DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized {
	return &DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized{}
}

/*DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized struct {
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{name}][%d] deleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized ", 401)
}

func (o *DeleteDiscoveryV1beta1NamespacedEndpointSliceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
