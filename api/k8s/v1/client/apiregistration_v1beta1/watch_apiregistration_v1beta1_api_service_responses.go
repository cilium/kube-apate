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

// WatchApiregistrationV1beta1APIServiceReader is a Reader for the WatchApiregistrationV1beta1APIService structure.
type WatchApiregistrationV1beta1APIServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchApiregistrationV1beta1APIServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchApiregistrationV1beta1APIServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchApiregistrationV1beta1APIServiceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchApiregistrationV1beta1APIServiceOK creates a WatchApiregistrationV1beta1APIServiceOK with default headers values
func NewWatchApiregistrationV1beta1APIServiceOK() *WatchApiregistrationV1beta1APIServiceOK {
	return &WatchApiregistrationV1beta1APIServiceOK{}
}

/*WatchApiregistrationV1beta1APIServiceOK handles this case with default header values.

OK
*/
type WatchApiregistrationV1beta1APIServiceOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchApiregistrationV1beta1APIServiceOK) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/v1beta1/watch/apiservices/{name}][%d] watchApiregistrationV1beta1ApiServiceOK  %+v", 200, o.Payload)
}

func (o *WatchApiregistrationV1beta1APIServiceOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchApiregistrationV1beta1APIServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchApiregistrationV1beta1APIServiceUnauthorized creates a WatchApiregistrationV1beta1APIServiceUnauthorized with default headers values
func NewWatchApiregistrationV1beta1APIServiceUnauthorized() *WatchApiregistrationV1beta1APIServiceUnauthorized {
	return &WatchApiregistrationV1beta1APIServiceUnauthorized{}
}

/*WatchApiregistrationV1beta1APIServiceUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchApiregistrationV1beta1APIServiceUnauthorized struct {
}

func (o *WatchApiregistrationV1beta1APIServiceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apiregistration.k8s.io/v1beta1/watch/apiservices/{name}][%d] watchApiregistrationV1beta1ApiServiceUnauthorized ", 401)
}

func (o *WatchApiregistrationV1beta1APIServiceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}