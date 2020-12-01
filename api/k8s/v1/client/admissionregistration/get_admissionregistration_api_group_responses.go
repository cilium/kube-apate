// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package admissionregistration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// GetAdmissionregistrationAPIGroupReader is a Reader for the GetAdmissionregistrationAPIGroup structure.
type GetAdmissionregistrationAPIGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdmissionregistrationAPIGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdmissionregistrationAPIGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAdmissionregistrationAPIGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAdmissionregistrationAPIGroupOK creates a GetAdmissionregistrationAPIGroupOK with default headers values
func NewGetAdmissionregistrationAPIGroupOK() *GetAdmissionregistrationAPIGroupOK {
	return &GetAdmissionregistrationAPIGroupOK{}
}

/*GetAdmissionregistrationAPIGroupOK handles this case with default header values.

OK
*/
type GetAdmissionregistrationAPIGroupOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1APIGroup
}

func (o *GetAdmissionregistrationAPIGroupOK) Error() string {
	return fmt.Sprintf("[GET /apis/admissionregistration.k8s.io/][%d] getAdmissionregistrationApiGroupOK  %+v", 200, o.Payload)
}

func (o *GetAdmissionregistrationAPIGroupOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1APIGroup {
	return o.Payload
}

func (o *GetAdmissionregistrationAPIGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1APIGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdmissionregistrationAPIGroupUnauthorized creates a GetAdmissionregistrationAPIGroupUnauthorized with default headers values
func NewGetAdmissionregistrationAPIGroupUnauthorized() *GetAdmissionregistrationAPIGroupUnauthorized {
	return &GetAdmissionregistrationAPIGroupUnauthorized{}
}

/*GetAdmissionregistrationAPIGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAdmissionregistrationAPIGroupUnauthorized struct {
}

func (o *GetAdmissionregistrationAPIGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/admissionregistration.k8s.io/][%d] getAdmissionregistrationApiGroupUnauthorized ", 401)
}

func (o *GetAdmissionregistrationAPIGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
