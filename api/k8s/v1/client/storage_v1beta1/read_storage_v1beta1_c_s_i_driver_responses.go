// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadStorageV1beta1CSIDriverReader is a Reader for the ReadStorageV1beta1CSIDriver structure.
type ReadStorageV1beta1CSIDriverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadStorageV1beta1CSIDriverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadStorageV1beta1CSIDriverOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadStorageV1beta1CSIDriverUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadStorageV1beta1CSIDriverOK creates a ReadStorageV1beta1CSIDriverOK with default headers values
func NewReadStorageV1beta1CSIDriverOK() *ReadStorageV1beta1CSIDriverOK {
	return &ReadStorageV1beta1CSIDriverOK{}
}

/*ReadStorageV1beta1CSIDriverOK handles this case with default header values.

OK
*/
type ReadStorageV1beta1CSIDriverOK struct {
	Payload *models.IoK8sAPIStorageV1beta1CSIDriver
}

func (o *ReadStorageV1beta1CSIDriverOK) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1beta1/csidrivers/{name}][%d] readStorageV1beta1CSIDriverOK  %+v", 200, o.Payload)
}

func (o *ReadStorageV1beta1CSIDriverOK) GetPayload() *models.IoK8sAPIStorageV1beta1CSIDriver {
	return o.Payload
}

func (o *ReadStorageV1beta1CSIDriverOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIStorageV1beta1CSIDriver)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadStorageV1beta1CSIDriverUnauthorized creates a ReadStorageV1beta1CSIDriverUnauthorized with default headers values
func NewReadStorageV1beta1CSIDriverUnauthorized() *ReadStorageV1beta1CSIDriverUnauthorized {
	return &ReadStorageV1beta1CSIDriverUnauthorized{}
}

/*ReadStorageV1beta1CSIDriverUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadStorageV1beta1CSIDriverUnauthorized struct {
}

func (o *ReadStorageV1beta1CSIDriverUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/storage.k8s.io/v1beta1/csidrivers/{name}][%d] readStorageV1beta1CSIDriverUnauthorized ", 401)
}

func (o *ReadStorageV1beta1CSIDriverUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
