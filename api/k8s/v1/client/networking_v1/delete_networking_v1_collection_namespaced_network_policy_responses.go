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

// DeleteNetworkingV1CollectionNamespacedNetworkPolicyReader is a Reader for the DeleteNetworkingV1CollectionNamespacedNetworkPolicy structure.
type DeleteNetworkingV1CollectionNamespacedNetworkPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkingV1CollectionNamespacedNetworkPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkingV1CollectionNamespacedNetworkPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkingV1CollectionNamespacedNetworkPolicyOK creates a DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK with default headers values
func NewDeleteNetworkingV1CollectionNamespacedNetworkPolicyOK() *DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK {
	return &DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK{}
}

/*DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK handles this case with default header values.

OK
*/
type DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies][%d] deleteNetworkingV1CollectionNamespacedNetworkPolicyOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1CollectionNamespacedNetworkPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized creates a DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized with default headers values
func NewDeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized() *DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized {
	return &DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized{}
}

/*DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized struct {
}

func (o *DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/networkpolicies][%d] deleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized ", 401)
}

func (o *DeleteNetworkingV1CollectionNamespacedNetworkPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
