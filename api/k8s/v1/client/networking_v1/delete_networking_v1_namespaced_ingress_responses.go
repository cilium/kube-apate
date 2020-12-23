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

// DeleteNetworkingV1NamespacedIngressReader is a Reader for the DeleteNetworkingV1NamespacedIngress structure.
type DeleteNetworkingV1NamespacedIngressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkingV1NamespacedIngressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkingV1NamespacedIngressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteNetworkingV1NamespacedIngressAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteNetworkingV1NamespacedIngressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNetworkingV1NamespacedIngressOK creates a DeleteNetworkingV1NamespacedIngressOK with default headers values
func NewDeleteNetworkingV1NamespacedIngressOK() *DeleteNetworkingV1NamespacedIngressOK {
	return &DeleteNetworkingV1NamespacedIngressOK{}
}

/*DeleteNetworkingV1NamespacedIngressOK handles this case with default header values.

OK
*/
type DeleteNetworkingV1NamespacedIngressOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1NamespacedIngressOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] deleteNetworkingV1NamespacedIngressOK  %+v", 200, o.Payload)
}

func (o *DeleteNetworkingV1NamespacedIngressOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1NamespacedIngressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1NamespacedIngressAccepted creates a DeleteNetworkingV1NamespacedIngressAccepted with default headers values
func NewDeleteNetworkingV1NamespacedIngressAccepted() *DeleteNetworkingV1NamespacedIngressAccepted {
	return &DeleteNetworkingV1NamespacedIngressAccepted{}
}

/*DeleteNetworkingV1NamespacedIngressAccepted handles this case with default header values.

Accepted
*/
type DeleteNetworkingV1NamespacedIngressAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteNetworkingV1NamespacedIngressAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] deleteNetworkingV1NamespacedIngressAccepted  %+v", 202, o.Payload)
}

func (o *DeleteNetworkingV1NamespacedIngressAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteNetworkingV1NamespacedIngressAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkingV1NamespacedIngressUnauthorized creates a DeleteNetworkingV1NamespacedIngressUnauthorized with default headers values
func NewDeleteNetworkingV1NamespacedIngressUnauthorized() *DeleteNetworkingV1NamespacedIngressUnauthorized {
	return &DeleteNetworkingV1NamespacedIngressUnauthorized{}
}

/*DeleteNetworkingV1NamespacedIngressUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteNetworkingV1NamespacedIngressUnauthorized struct {
}

func (o *DeleteNetworkingV1NamespacedIngressUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/networking.k8s.io/v1/namespaces/{namespace}/ingresses/{name}][%d] deleteNetworkingV1NamespacedIngressUnauthorized ", 401)
}

func (o *DeleteNetworkingV1NamespacedIngressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}