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

// ReplaceCoreV1NamespacedPodStatusReader is a Reader for the ReplaceCoreV1NamespacedPodStatus structure.
type ReplaceCoreV1NamespacedPodStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoreV1NamespacedPodStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoreV1NamespacedPodStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoreV1NamespacedPodStatusCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoreV1NamespacedPodStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoreV1NamespacedPodStatusOK creates a ReplaceCoreV1NamespacedPodStatusOK with default headers values
func NewReplaceCoreV1NamespacedPodStatusOK() *ReplaceCoreV1NamespacedPodStatusOK {
	return &ReplaceCoreV1NamespacedPodStatusOK{}
}

/*ReplaceCoreV1NamespacedPodStatusOK handles this case with default header values.

OK
*/
type ReplaceCoreV1NamespacedPodStatusOK struct {
	Payload *models.IoK8sAPICoreV1Pod
}

func (o *ReplaceCoreV1NamespacedPodStatusOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/pods/{name}/status][%d] replaceCoreV1NamespacedPodStatusOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoreV1NamespacedPodStatusOK) GetPayload() *models.IoK8sAPICoreV1Pod {
	return o.Payload
}

func (o *ReplaceCoreV1NamespacedPodStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Pod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NamespacedPodStatusCreated creates a ReplaceCoreV1NamespacedPodStatusCreated with default headers values
func NewReplaceCoreV1NamespacedPodStatusCreated() *ReplaceCoreV1NamespacedPodStatusCreated {
	return &ReplaceCoreV1NamespacedPodStatusCreated{}
}

/*ReplaceCoreV1NamespacedPodStatusCreated handles this case with default header values.

Created
*/
type ReplaceCoreV1NamespacedPodStatusCreated struct {
	Payload *models.IoK8sAPICoreV1Pod
}

func (o *ReplaceCoreV1NamespacedPodStatusCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/pods/{name}/status][%d] replaceCoreV1NamespacedPodStatusCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoreV1NamespacedPodStatusCreated) GetPayload() *models.IoK8sAPICoreV1Pod {
	return o.Payload
}

func (o *ReplaceCoreV1NamespacedPodStatusCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1Pod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NamespacedPodStatusUnauthorized creates a ReplaceCoreV1NamespacedPodStatusUnauthorized with default headers values
func NewReplaceCoreV1NamespacedPodStatusUnauthorized() *ReplaceCoreV1NamespacedPodStatusUnauthorized {
	return &ReplaceCoreV1NamespacedPodStatusUnauthorized{}
}

/*ReplaceCoreV1NamespacedPodStatusUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoreV1NamespacedPodStatusUnauthorized struct {
}

func (o *ReplaceCoreV1NamespacedPodStatusUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/pods/{name}/status][%d] replaceCoreV1NamespacedPodStatusUnauthorized ", 401)
}

func (o *ReplaceCoreV1NamespacedPodStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
