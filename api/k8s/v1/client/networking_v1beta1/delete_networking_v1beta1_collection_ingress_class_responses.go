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

// DeleteNetworkingV1beta1CollectionIngressClassReader is a Reader for the DeleteNetworkingV1beta1CollectionIngressClass structure.
type DeleteNetworkingV1beta1CollectionIngressClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkingV1beta1CollectionIngressClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkingV1beta1CollectionIngressClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkingV1beta1CollectionIngressClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkingV1beta1CollectionIngressClassOK creates a DeleteNetworkingV1beta1CollectionIngressClassOK with default headers values
func NewDeleteNetworkingV1beta1CollectionIngressClassOK() *DeleteNetworkingV1beta1CollectionIngressClassOK {
	return &DeleteNetworkingV1beta1CollectionIngressClassOK{}
}

/*DeleteNetworkingV1beta1CollectionIngressClassOK handles this case with default header values.

OK
*/
type DeleteNetworkingV1beta1CollectionIngressClassOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1beta1CollectionIngressClassOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1beta1/ingressclasses][%d] deleteNetworkingV1beta1CollectionIngressClassOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkingV1beta1CollectionIngressClassOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1beta1CollectionIngressClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1beta1CollectionIngressClassUnauthorized creates a DeleteNetworkingV1beta1CollectionIngressClassUnauthorized with default headers values
func NewDeleteNetworkingV1beta1CollectionIngressClassUnauthorized() *DeleteNetworkingV1beta1CollectionIngressClassUnauthorized {
	return &DeleteNetworkingV1beta1CollectionIngressClassUnauthorized{}
}

/*DeleteNetworkingV1beta1CollectionIngressClassUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNetworkingV1beta1CollectionIngressClassUnauthorized struct {
}

func (o *DeleteNetworkingV1beta1CollectionIngressClassUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1beta1/ingressclasses][%d] deleteNetworkingV1beta1CollectionIngressClassUnauthorized ", 401)
}

func (o *DeleteNetworkingV1beta1CollectionIngressClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
