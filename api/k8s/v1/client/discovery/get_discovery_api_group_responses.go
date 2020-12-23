// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetDiscoveryAPIGroupReader is a Reader for the GetDiscoveryAPIGroup structure.
type GetDiscoveryAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoveryAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoveryAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetDiscoveryAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoveryAPIGroupOK creates a GetDiscoveryAPIGroupOK with default headers values
func NewGetDiscoveryAPIGroupOK() *GetDiscoveryAPIGroupOK {
	return &GetDiscoveryAPIGroupOK{}
}

/*GetDiscoveryAPIGroupOK handles this case with default header values.

OK
*/
type GetDiscoveryAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetDiscoveryAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/discovery.k8s.io/][%d] getDiscoveryApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetDiscoveryAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetDiscoveryAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoveryAPIGroupUnauthorized creates a GetDiscoveryAPIGroupUnauthorized with default headers values
func NewGetDiscoveryAPIGroupUnauthorized() *GetDiscoveryAPIGroupUnauthorized {
	return &GetDiscoveryAPIGroupUnauthorized{}
}

/*GetDiscoveryAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDiscoveryAPIGroupUnauthorized struct {
}

func (o *GetDiscoveryAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/discovery.k8s.io/][%d] getDiscoveryApiGroupUnauthorized ", 401)
}

func (o *GetDiscoveryAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}