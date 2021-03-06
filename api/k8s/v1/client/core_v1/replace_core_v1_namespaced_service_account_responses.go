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

// ReplaceCoreV1NamespacedServiceAccountReader is a Reader for the ReplaceCoreV1NamespacedServiceAccount structure.
type ReplaceCoreV1NamespacedServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceCoreV1NamespacedServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceCoreV1NamespacedServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceCoreV1NamespacedServiceAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceCoreV1NamespacedServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceCoreV1NamespacedServiceAccountOK creates a ReplaceCoreV1NamespacedServiceAccountOK with default headers values
func NewReplaceCoreV1NamespacedServiceAccountOK() *ReplaceCoreV1NamespacedServiceAccountOK {
	return &ReplaceCoreV1NamespacedServiceAccountOK{}
}

/*ReplaceCoreV1NamespacedServiceAccountOK handles this case with default header values.

OK
*/
type ReplaceCoreV1NamespacedServiceAccountOK struct {
	Payload *models.IoK8sAPICoreV1ServiceAccount
}

func (o *ReplaceCoreV1NamespacedServiceAccountOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/serviceaccounts/{name}][%d] replaceCoreV1NamespacedServiceAccountOK  %+v", 200, o.Payload)
}

func (o *ReplaceCoreV1NamespacedServiceAccountOK) GetPayload() *models.IoK8sAPICoreV1ServiceAccount {
	return o.Payload
}

func (o *ReplaceCoreV1NamespacedServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NamespacedServiceAccountCreated creates a ReplaceCoreV1NamespacedServiceAccountCreated with default headers values
func NewReplaceCoreV1NamespacedServiceAccountCreated() *ReplaceCoreV1NamespacedServiceAccountCreated {
	return &ReplaceCoreV1NamespacedServiceAccountCreated{}
}

/*ReplaceCoreV1NamespacedServiceAccountCreated handles this case with default header values.

Created
*/
type ReplaceCoreV1NamespacedServiceAccountCreated struct {
	Payload *models.IoK8sAPICoreV1ServiceAccount
}

func (o *ReplaceCoreV1NamespacedServiceAccountCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/serviceaccounts/{name}][%d] replaceCoreV1NamespacedServiceAccountCreated  %+v", 201, o.Payload)
}

func (o *ReplaceCoreV1NamespacedServiceAccountCreated) GetPayload() *models.IoK8sAPICoreV1ServiceAccount {
	return o.Payload
}

func (o *ReplaceCoreV1NamespacedServiceAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1ServiceAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceCoreV1NamespacedServiceAccountUnauthorized creates a ReplaceCoreV1NamespacedServiceAccountUnauthorized with default headers values
func NewReplaceCoreV1NamespacedServiceAccountUnauthorized() *ReplaceCoreV1NamespacedServiceAccountUnauthorized {
	return &ReplaceCoreV1NamespacedServiceAccountUnauthorized{}
}

/*ReplaceCoreV1NamespacedServiceAccountUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceCoreV1NamespacedServiceAccountUnauthorized struct {
}

func (o *ReplaceCoreV1NamespacedServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/namespaces/{namespace}/serviceaccounts/{name}][%d] replaceCoreV1NamespacedServiceAccountUnauthorized ", 401)
}

func (o *ReplaceCoreV1NamespacedServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
