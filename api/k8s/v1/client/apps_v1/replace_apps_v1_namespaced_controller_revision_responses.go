// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceAppsV1NamespacedControllerRevisionReader is a Reader for the ReplaceAppsV1NamespacedControllerRevision structure.
type ReplaceAppsV1NamespacedControllerRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAppsV1NamespacedControllerRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAppsV1NamespacedControllerRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceAppsV1NamespacedControllerRevisionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceAppsV1NamespacedControllerRevisionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceAppsV1NamespacedControllerRevisionOK creates a ReplaceAppsV1NamespacedControllerRevisionOK with default headers values
func NewReplaceAppsV1NamespacedControllerRevisionOK() *ReplaceAppsV1NamespacedControllerRevisionOK {
	return &ReplaceAppsV1NamespacedControllerRevisionOK{}
}

/*ReplaceAppsV1NamespacedControllerRevisionOK handles this case with default header values.

OK
*/
type ReplaceAppsV1NamespacedControllerRevisionOK struct {
	Payload *models.IoK8sAPIAppsV1ControllerRevision
}

func (o *ReplaceAppsV1NamespacedControllerRevisionOK) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/controllerrevisions/{name}][%d] replaceAppsV1NamespacedControllerRevisionOK  %+v", 200, o.Payload)
}

func (o *ReplaceAppsV1NamespacedControllerRevisionOK) GetPayload() *models.IoK8sAPIAppsV1ControllerRevision {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedControllerRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1ControllerRevision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedControllerRevisionCreated creates a ReplaceAppsV1NamespacedControllerRevisionCreated with default headers values
func NewReplaceAppsV1NamespacedControllerRevisionCreated() *ReplaceAppsV1NamespacedControllerRevisionCreated {
	return &ReplaceAppsV1NamespacedControllerRevisionCreated{}
}

/*ReplaceAppsV1NamespacedControllerRevisionCreated handles this case with default header values.

Created
*/
type ReplaceAppsV1NamespacedControllerRevisionCreated struct {
	Payload *models.IoK8sAPIAppsV1ControllerRevision
}

func (o *ReplaceAppsV1NamespacedControllerRevisionCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/controllerrevisions/{name}][%d] replaceAppsV1NamespacedControllerRevisionCreated  %+v", 201, o.Payload)
}

func (o *ReplaceAppsV1NamespacedControllerRevisionCreated) GetPayload() *models.IoK8sAPIAppsV1ControllerRevision {
	return o.Payload
}

func (o *ReplaceAppsV1NamespacedControllerRevisionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1ControllerRevision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceAppsV1NamespacedControllerRevisionUnauthorized creates a ReplaceAppsV1NamespacedControllerRevisionUnauthorized with default headers values
func NewReplaceAppsV1NamespacedControllerRevisionUnauthorized() *ReplaceAppsV1NamespacedControllerRevisionUnauthorized {
	return &ReplaceAppsV1NamespacedControllerRevisionUnauthorized{}
}

/*ReplaceAppsV1NamespacedControllerRevisionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceAppsV1NamespacedControllerRevisionUnauthorized struct {
}

func (o *ReplaceAppsV1NamespacedControllerRevisionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/apps/v1/namespaces/{namespace}/controllerrevisions/{name}][%d] replaceAppsV1NamespacedControllerRevisionUnauthorized ", 401)
}

func (o *ReplaceAppsV1NamespacedControllerRevisionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
