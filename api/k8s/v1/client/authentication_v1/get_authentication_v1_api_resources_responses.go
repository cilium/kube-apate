// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetAuthenticationV1APIResourcesReader is a Reader for the GetAuthenticationV1APIResources structure.
type GetAuthenticationV1APIResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthenticationV1APIResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthenticationV1APIResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAuthenticationV1APIResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthenticationV1APIResourcesOK creates a GetAuthenticationV1APIResourcesOK with default headers values
func NewGetAuthenticationV1APIResourcesOK() *GetAuthenticationV1APIResourcesOK {
	return &GetAuthenticationV1APIResourcesOK{}
}

/*GetAuthenticationV1APIResourcesOK handles this case with default header values.

OK
*/
type GetAuthenticationV1APIResourcesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIResourceList
}

func (o *GetAuthenticationV1APIResourcesOK) Error() string {
	return fmt.Sprintf("[GET /apis/authentication.k8s.io/v1/][%d] getAuthenticationV1ApiResourcesOK  %+v", 200, o.Payload)
}

func (o *GetAuthenticationV1APIResourcesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIResourceList {
	return o.Payload
}

func (o *GetAuthenticationV1APIResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIResourceList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthenticationV1APIResourcesUnauthorized creates a GetAuthenticationV1APIResourcesUnauthorized with default headers values
func NewGetAuthenticationV1APIResourcesUnauthorized() *GetAuthenticationV1APIResourcesUnauthorized {
	return &GetAuthenticationV1APIResourcesUnauthorized{}
}

/*GetAuthenticationV1APIResourcesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAuthenticationV1APIResourcesUnauthorized struct {
}

func (o *GetAuthenticationV1APIResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/authentication.k8s.io/v1/][%d] getAuthenticationV1ApiResourcesUnauthorized ", 401)
}

func (o *GetAuthenticationV1APIResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}