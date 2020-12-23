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

// GetCoordinationV1beta1APIResourcesReader is a Reader for the GetCoordinationV1beta1APIResources structure.
type GetCoordinationV1beta1APIResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoordinationV1beta1APIResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoordinationV1beta1APIResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCoordinationV1beta1APIResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoordinationV1beta1APIResourcesOK creates a GetCoordinationV1beta1APIResourcesOK with default headers values
func NewGetCoordinationV1beta1APIResourcesOK() *GetCoordinationV1beta1APIResourcesOK {
	return &GetCoordinationV1beta1APIResourcesOK{}
}

/*GetCoordinationV1beta1APIResourcesOK handles this case with default header values.

OK
*/
type GetCoordinationV1beta1APIResourcesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList
}

func (o *GetCoordinationV1beta1APIResourcesOK) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1beta1/][%d] getCoordinationV1beta1ApiResourcesOK  %+v", 200, o.Payload)
}

func (o *GetCoordinationV1beta1APIResourcesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIResourceList {
	return o.Payload
}

func (o *GetCoordinationV1beta1APIResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIResourceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoordinationV1beta1APIResourcesUnauthorized creates a GetCoordinationV1beta1APIResourcesUnauthorized with default headers values
func NewGetCoordinationV1beta1APIResourcesUnauthorized() *GetCoordinationV1beta1APIResourcesUnauthorized {
	return &GetCoordinationV1beta1APIResourcesUnauthorized{}
}

/*GetCoordinationV1beta1APIResourcesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCoordinationV1beta1APIResourcesUnauthorized struct {
}

func (o *GetCoordinationV1beta1APIResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/coordination.k8s.io/v1beta1/][%d] getCoordinationV1beta1ApiResourcesUnauthorized ", 401)
}

func (o *GetCoordinationV1beta1APIResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}