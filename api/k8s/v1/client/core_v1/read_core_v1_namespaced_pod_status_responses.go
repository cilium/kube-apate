// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReadCoreV1NamespacedPodStatusReader is a Reader for the ReadCoreV1NamespacedPodStatus structure.
type ReadCoreV1NamespacedPodStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadCoreV1NamespacedPodStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadCoreV1NamespacedPodStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadCoreV1NamespacedPodStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadCoreV1NamespacedPodStatusOK creates a ReadCoreV1NamespacedPodStatusOK with default headers values
func NewReadCoreV1NamespacedPodStatusOK() *ReadCoreV1NamespacedPodStatusOK {
	return &ReadCoreV1NamespacedPodStatusOK{}
}

/*ReadCoreV1NamespacedPodStatusOK handles this case with default header values.

OK
*/
type ReadCoreV1NamespacedPodStatusOK struct {
	Payload *models.IoK8sAPICoreV1Pod
}

func (o *ReadCoreV1NamespacedPodStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/pods/{name}/status][%d] readCoreV1NamespacedPodStatusOK  %+v", 200, o.Payload)
}

func (o *ReadCoreV1NamespacedPodStatusOK) GetPayload() *models.IoK8sAPICoreV1Pod {
	return o.Payload
}

func (o *ReadCoreV1NamespacedPodStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Pod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadCoreV1NamespacedPodStatusUnauthorized creates a ReadCoreV1NamespacedPodStatusUnauthorized with default headers values
func NewReadCoreV1NamespacedPodStatusUnauthorized() *ReadCoreV1NamespacedPodStatusUnauthorized {
	return &ReadCoreV1NamespacedPodStatusUnauthorized{}
}

/*ReadCoreV1NamespacedPodStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadCoreV1NamespacedPodStatusUnauthorized struct {
}

func (o *ReadCoreV1NamespacedPodStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/namespaces/{namespace}/pods/{name}/status][%d] readCoreV1NamespacedPodStatusUnauthorized ", 401)
}

func (o *ReadCoreV1NamespacedPodStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
