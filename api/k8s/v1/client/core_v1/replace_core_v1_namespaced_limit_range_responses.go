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

// ReplaceCoreV1NamespacedLimitRangeReader is a Reader for the ReplaceCoreV1NamespacedLimitRange structure.
type ReplaceCoreV1NamespacedLimitRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoreV1NamespacedLimitRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoreV1NamespacedLimitRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoreV1NamespacedLimitRangeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoreV1NamespacedLimitRangeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoreV1NamespacedLimitRangeOK creates a ReplaceCoreV1NamespacedLimitRangeOK with default headers values
func NewReplaceCoreV1NamespacedLimitRangeOK() *ReplaceCoreV1NamespacedLimitRangeOK {
	return &ReplaceCoreV1NamespacedLimitRangeOK{}
}

/*ReplaceCoreV1NamespacedLimitRangeOK handles this case with default header values.

OK
*/
type ReplaceCoreV1NamespacedLimitRangeOK struct {
	Payload *models.IoK8sAPICoreV1LimitRange
}

func (o *ReplaceCoreV1NamespacedLimitRangeOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/limitranges/{name}][%d] replaceCoreV1NamespacedLimitRangeOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoreV1NamespacedLimitRangeOK) GetPayload() *models.IoK8sAPICoreV1LimitRange {
	return o.Payload
}

func (o *ReplaceCoreV1NamespacedLimitRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1LimitRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NamespacedLimitRangeCreated creates a ReplaceCoreV1NamespacedLimitRangeCreated with default headers values
func NewReplaceCoreV1NamespacedLimitRangeCreated() *ReplaceCoreV1NamespacedLimitRangeCreated {
	return &ReplaceCoreV1NamespacedLimitRangeCreated{}
}

/*ReplaceCoreV1NamespacedLimitRangeCreated handles this case with default header values.

Created
*/
type ReplaceCoreV1NamespacedLimitRangeCreated struct {
	Payload *models.IoK8sAPICoreV1LimitRange
}

func (o *ReplaceCoreV1NamespacedLimitRangeCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/limitranges/{name}][%d] replaceCoreV1NamespacedLimitRangeCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoreV1NamespacedLimitRangeCreated) GetPayload() *models.IoK8sAPICoreV1LimitRange {
	return o.Payload
}

func (o *ReplaceCoreV1NamespacedLimitRangeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1LimitRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NamespacedLimitRangeUnauthorized creates a ReplaceCoreV1NamespacedLimitRangeUnauthorized with default headers values
func NewReplaceCoreV1NamespacedLimitRangeUnauthorized() *ReplaceCoreV1NamespacedLimitRangeUnauthorized {
	return &ReplaceCoreV1NamespacedLimitRangeUnauthorized{}
}

/*ReplaceCoreV1NamespacedLimitRangeUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoreV1NamespacedLimitRangeUnauthorized struct {
}

func (o *ReplaceCoreV1NamespacedLimitRangeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/limitranges/{name}][%d] replaceCoreV1NamespacedLimitRangeUnauthorized ", 401)
}

func (o *ReplaceCoreV1NamespacedLimitRangeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
