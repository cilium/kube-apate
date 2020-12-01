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

// DeleteNetworkingV1CollectionNamespacedIngressReader is a Reader for the DeleteNetworkingV1CollectionNamespacedIngress structure.
type DeleteNetworkingV1CollectionNamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkingV1CollectionNamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkingV1CollectionNamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkingV1CollectionNamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkingV1CollectionNamespacedIngressOK creates a DeleteNetworkingV1CollectionNamespacedIngressOK with default headers values
func NewDeleteNetworkingV1CollectionNamespacedIngressOK() *DeleteNetworkingV1CollectionNamespacedIngressOK {
	return &DeleteNetworkingV1CollectionNamespacedIngressOK{}
}

/*DeleteNetworkingV1CollectionNamespacedIngressOK handles this case with default header values.

OK
*/
type DeleteNetworkingV1CollectionNamespacedIngressOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1CollectionNamespacedIngressOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses][%d] deleteNetworkingV1CollectionNamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkingV1CollectionNamespacedIngressOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1CollectionNamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1CollectionNamespacedIngressUnauthorized creates a DeleteNetworkingV1CollectionNamespacedIngressUnauthorized with default headers values
func NewDeleteNetworkingV1CollectionNamespacedIngressUnauthorized() *DeleteNetworkingV1CollectionNamespacedIngressUnauthorized {
	return &DeleteNetworkingV1CollectionNamespacedIngressUnauthorized{}
}

/*DeleteNetworkingV1CollectionNamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNetworkingV1CollectionNamespacedIngressUnauthorized struct {
}

func (o *DeleteNetworkingV1CollectionNamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses][%d] deleteNetworkingV1CollectionNamespacedIngressUnauthorized ", 401)
}

func (o *DeleteNetworkingV1CollectionNamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
