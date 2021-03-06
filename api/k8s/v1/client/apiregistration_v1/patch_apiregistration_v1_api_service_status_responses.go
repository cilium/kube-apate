// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchApiregistrationV1APIServiceStatusReader is a Reader for the PatchApiregistrationV1APIServiceStatus structure.
type PatchApiregistrationV1APIServiceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchApiregistrationV1APIServiceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApiregistrationV1APIServiceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchApiregistrationV1APIServiceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchApiregistrationV1APIServiceStatusOK creates a PatchApiregistrationV1APIServiceStatusOK with default headers values
func NewPatchApiregistrationV1APIServiceStatusOK() *PatchApiregistrationV1APIServiceStatusOK {
	return &PatchApiregistrationV1APIServiceStatusOK{}
}

/*PatchApiregistrationV1APIServiceStatusOK handles this case with default header values.

OK
*/
type PatchApiregistrationV1APIServiceStatusOK struct {
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService
}

func (o *PatchApiregistrationV1APIServiceStatusOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/apiregistration.k8s.io/v1/apiservices/{name}/status][%d] patchApiregistrationV1ApiServiceStatusOK  %+v", 200, o.Payload)
}

func (o *PatchApiregistrationV1APIServiceStatusOK) GetPayload() *models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService {
	return o.Payload
}

func (o *PatchApiregistrationV1APIServiceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sKubeAggregatorPkgApisApiregistrationV1APIService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApiregistrationV1APIServiceStatusUnauthorized creates a PatchApiregistrationV1APIServiceStatusUnauthorized with default headers values
func NewPatchApiregistrationV1APIServiceStatusUnauthorized() *PatchApiregistrationV1APIServiceStatusUnauthorized {
	return &PatchApiregistrationV1APIServiceStatusUnauthorized{}
}

/*PatchApiregistrationV1APIServiceStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchApiregistrationV1APIServiceStatusUnauthorized struct {
}

func (o *PatchApiregistrationV1APIServiceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/apiregistration.k8s.io/v1/apiservices/{name}/status][%d] patchApiregistrationV1ApiServiceStatusUnauthorized ", 401)
}

func (o *PatchApiregistrationV1APIServiceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
