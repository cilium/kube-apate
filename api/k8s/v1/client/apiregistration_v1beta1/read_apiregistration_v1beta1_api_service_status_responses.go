// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadApiregistrationV1beta1APIServiceStatusReader is a Reader for the ReadApiregistrationV1beta1APIServiceStatus structure.
type ReadApiregistrationV1beta1APIServiceStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadApiregistrationV1beta1APIServiceStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadApiregistrationV1beta1APIServiceStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadApiregistrationV1beta1APIServiceStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadApiregistrationV1beta1APIServiceStatusOK creates a ReadApiregistrationV1beta1APIServiceStatusOK with default headers values
func NewReadApiregistrationV1beta1APIServiceStatusOK() *ReadApiregistrationV1beta1APIServiceStatusOK {
	return &ReadApiregistrationV1beta1APIServiceStatusOK{}
}

/*ReadApiregistrationV1beta1APIServiceStatusOK handles this case with default header values.

OK
*/
type ReadApiregistrationV1beta1APIServiceStatusOK struct {
	Payload *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService
}

func (o *ReadApiregistrationV1beta1APIServiceStatusOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/v1beta1/apiservices/{name}/status][%d] readApiregistrationV1beta1ApiServiceStatusOK  %+v", 200, o.Payload)
}

func (o *ReadApiregistrationV1beta1APIServiceStatusOK) GetPayload() *models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService {
	return o.Payload
}

func (o *ReadApiregistrationV1beta1APIServiceStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadApiregistrationV1beta1APIServiceStatusUnauthorized creates a ReadApiregistrationV1beta1APIServiceStatusUnauthorized with default headers values
func NewReadApiregistrationV1beta1APIServiceStatusUnauthorized() *ReadApiregistrationV1beta1APIServiceStatusUnauthorized {
	return &ReadApiregistrationV1beta1APIServiceStatusUnauthorized{}
}

/*ReadApiregistrationV1beta1APIServiceStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadApiregistrationV1beta1APIServiceStatusUnauthorized struct {
}

func (o *ReadApiregistrationV1beta1APIServiceStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/v1beta1/apiservices/{name}/status][%d] readApiregistrationV1beta1ApiServiceStatusUnauthorized ", 401)
}

func (o *ReadApiregistrationV1beta1APIServiceStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
