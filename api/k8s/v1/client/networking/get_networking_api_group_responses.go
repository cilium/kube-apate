// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetNetworkingAPIGroupReader is a Reader for the GetNetworkingAPIGroup structure.
type GetNetworkingAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkingAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkingAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNetworkingAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkingAPIGroupOK creates a GetNetworkingAPIGroupOK with default headers values
func NewGetNetworkingAPIGroupOK() *GetNetworkingAPIGroupOK {
	return &GetNetworkingAPIGroupOK{}
}

/*GetNetworkingAPIGroupOK handles this case with default header values.

OK
*/
type GetNetworkingAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetNetworkingAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/][%d] getNetworkingApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetNetworkingAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetNetworkingAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkingAPIGroupUnauthorized creates a GetNetworkingAPIGroupUnauthorized with default headers values
func NewGetNetworkingAPIGroupUnauthorized() *GetNetworkingAPIGroupUnauthorized {
	return &GetNetworkingAPIGroupUnauthorized{}
}

/*GetNetworkingAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNetworkingAPIGroupUnauthorized struct {
}

func (o *GetNetworkingAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/networking.k8s.io/][%d] getNetworkingApiGroupUnauthorized ", 401)
}

func (o *GetNetworkingAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}