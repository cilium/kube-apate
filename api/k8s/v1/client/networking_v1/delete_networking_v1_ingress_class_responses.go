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

// DeleteNetworkingV1IngressClassReader is a Reader for the DeleteNetworkingV1IngressClass structure.
type DeleteNetworkingV1IngressClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkingV1IngressClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkingV1IngressClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteNetworkingV1IngressClassAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkingV1IngressClassUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkingV1IngressClassOK creates a DeleteNetworkingV1IngressClassOK with default headers values
func NewDeleteNetworkingV1IngressClassOK() *DeleteNetworkingV1IngressClassOK {
	return &DeleteNetworkingV1IngressClassOK{}
}

/*DeleteNetworkingV1IngressClassOK handles this case with default header values.

OK
*/
type DeleteNetworkingV1IngressClassOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1IngressClassOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] deleteNetworkingV1IngressClassOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkingV1IngressClassOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1IngressClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1IngressClassAccepted creates a DeleteNetworkingV1IngressClassAccepted with default headers values
func NewDeleteNetworkingV1IngressClassAccepted() *DeleteNetworkingV1IngressClassAccepted {
	return &DeleteNetworkingV1IngressClassAccepted{}
}

/*DeleteNetworkingV1IngressClassAccepted handles this case with default header values.

Accepted
*/
type DeleteNetworkingV1IngressClassAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1IngressClassAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] deleteNetworkingV1IngressClassAccepted  %+v", 202, o.Payload)
}

func (o *DeleteNetworkingV1IngressClassAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1IngressClassAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1IngressClassUnauthorized creates a DeleteNetworkingV1IngressClassUnauthorized with default headers values
func NewDeleteNetworkingV1IngressClassUnauthorized() *DeleteNetworkingV1IngressClassUnauthorized {
	return &DeleteNetworkingV1IngressClassUnauthorized{}
}

/*DeleteNetworkingV1IngressClassUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNetworkingV1IngressClassUnauthorized struct {
}

func (o *DeleteNetworkingV1IngressClassUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/ingressclasses/{name}][%d] deleteNetworkingV1IngressClassUnauthorized ", 401)
}

func (o *DeleteNetworkingV1IngressClassUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
