// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package coordination_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListCoordinationV1beta1LeaseForAllNamespacesReader is a Reader for the ListCoordinationV1beta1LeaseForAllNamespaces structure.
type ListCoordinationV1beta1LeaseForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCoordinationV1beta1LeaseForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCoordinationV1beta1LeaseForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListCoordinationV1beta1LeaseForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListCoordinationV1beta1LeaseForAllNamespacesOK creates a ListCoordinationV1beta1LeaseForAllNamespacesOK with default headers values
func NewListCoordinationV1beta1LeaseForAllNamespacesOK() *ListCoordinationV1beta1LeaseForAllNamespacesOK {
	return &ListCoordinationV1beta1LeaseForAllNamespacesOK{}
}

/*ListCoordinationV1beta1LeaseForAllNamespacesOK handles this case with default header values.

OK
*/
type ListCoordinationV1beta1LeaseForAllNamespacesOK struct {
	Payload *models.IoK8sAPICoordinationV1beta1LeaseList
}

func (o *ListCoordinationV1beta1LeaseForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1beta1/leases][%d] listCoordinationV1beta1LeaseForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListCoordinationV1beta1LeaseForAllNamespacesOK) GetPayload() *models.IoK8sAPICoordinationV1beta1LeaseList {
	return o.Payload
}

func (o *ListCoordinationV1beta1LeaseForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoordinationV1beta1LeaseList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCoordinationV1beta1LeaseForAllNamespacesUnauthorized creates a ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized with default headers values
func NewListCoordinationV1beta1LeaseForAllNamespacesUnauthorized() *ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized {
	return &ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized{}
}

/*ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized struct {
}

func (o *ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1beta1/leases][%d] listCoordinationV1beta1LeaseForAllNamespacesUnauthorized ", 401)
}

func (o *ListCoordinationV1beta1LeaseForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
