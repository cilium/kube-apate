// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetAuthenticationAPIGroupReader is a Reader for the GetAuthenticationAPIGroup structure.
type GetAuthenticationAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthenticationAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthenticationAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAuthenticationAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthenticationAPIGroupOK creates a GetAuthenticationAPIGroupOK with default headers values
func NewGetAuthenticationAPIGroupOK() *GetAuthenticationAPIGroupOK {
	return &GetAuthenticationAPIGroupOK{}
}

/*GetAuthenticationAPIGroupOK handles this case with default header values.

OK
*/
type GetAuthenticationAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetAuthenticationAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/authentication.k8s.io/][%d] getAuthenticationApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetAuthenticationAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetAuthenticationAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthenticationAPIGroupUnauthorized creates a GetAuthenticationAPIGroupUnauthorized with default headers values
func NewGetAuthenticationAPIGroupUnauthorized() *GetAuthenticationAPIGroupUnauthorized {
	return &GetAuthenticationAPIGroupUnauthorized{}
}

/*GetAuthenticationAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAuthenticationAPIGroupUnauthorized struct {
}

func (o *GetAuthenticationAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/authentication.k8s.io/][%d] getAuthenticationApiGroupUnauthorized ", 401)
}

func (o *GetAuthenticationAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
