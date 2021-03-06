// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateRbacAuthorizationV1NamespacedRoleReader is a Reader for the CreateRbacAuthorizationV1NamespacedRole structure.
type CreateRbacAuthorizationV1NamespacedRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRbacAuthorizationV1NamespacedRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRbacAuthorizationV1NamespacedRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateRbacAuthorizationV1NamespacedRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateRbacAuthorizationV1NamespacedRoleAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRbacAuthorizationV1NamespacedRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRbacAuthorizationV1NamespacedRoleOK creates a CreateRbacAuthorizationV1NamespacedRoleOK with default headers values
func NewCreateRbacAuthorizationV1NamespacedRoleOK() *CreateRbacAuthorizationV1NamespacedRoleOK {
	return &CreateRbacAuthorizationV1NamespacedRoleOK{}
}

/*CreateRbacAuthorizationV1NamespacedRoleOK handles this case with default header values.

OK
*/
type CreateRbacAuthorizationV1NamespacedRoleOK struct {
	Payload *models.IoK8sAPIRbacV1Role
}

func (o *CreateRbacAuthorizationV1NamespacedRoleOK) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles][%d] createRbacAuthorizationV1NamespacedRoleOK  %+v", 200, o.Payload)
}

func (o *CreateRbacAuthorizationV1NamespacedRoleOK) GetPayload() *models.IoK8sAPIRbacV1Role {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1NamespacedRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1NamespacedRoleCreated creates a CreateRbacAuthorizationV1NamespacedRoleCreated with default headers values
func NewCreateRbacAuthorizationV1NamespacedRoleCreated() *CreateRbacAuthorizationV1NamespacedRoleCreated {
	return &CreateRbacAuthorizationV1NamespacedRoleCreated{}
}

/*CreateRbacAuthorizationV1NamespacedRoleCreated handles this case with default header values.

Created
*/
type CreateRbacAuthorizationV1NamespacedRoleCreated struct {
	Payload *models.IoK8sAPIRbacV1Role
}

func (o *CreateRbacAuthorizationV1NamespacedRoleCreated) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles][%d] createRbacAuthorizationV1NamespacedRoleCreated  %+v", 201, o.Payload)
}

func (o *CreateRbacAuthorizationV1NamespacedRoleCreated) GetPayload() *models.IoK8sAPIRbacV1Role {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1NamespacedRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1NamespacedRoleAccepted creates a CreateRbacAuthorizationV1NamespacedRoleAccepted with default headers values
func NewCreateRbacAuthorizationV1NamespacedRoleAccepted() *CreateRbacAuthorizationV1NamespacedRoleAccepted {
	return &CreateRbacAuthorizationV1NamespacedRoleAccepted{}
}

/*CreateRbacAuthorizationV1NamespacedRoleAccepted handles this case with default header values.

Accepted
*/
type CreateRbacAuthorizationV1NamespacedRoleAccepted struct {
	Payload *models.IoK8sAPIRbacV1Role
}

func (o *CreateRbacAuthorizationV1NamespacedRoleAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles][%d] createRbacAuthorizationV1NamespacedRoleAccepted  %+v", 202, o.Payload)
}

func (o *CreateRbacAuthorizationV1NamespacedRoleAccepted) GetPayload() *models.IoK8sAPIRbacV1Role {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1NamespacedRoleAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1NamespacedRoleUnauthorized creates a CreateRbacAuthorizationV1NamespacedRoleUnauthorized with default headers values
func NewCreateRbacAuthorizationV1NamespacedRoleUnauthorized() *CreateRbacAuthorizationV1NamespacedRoleUnauthorized {
	return &CreateRbacAuthorizationV1NamespacedRoleUnauthorized{}
}

/*CreateRbacAuthorizationV1NamespacedRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRbacAuthorizationV1NamespacedRoleUnauthorized struct {
}

func (o *CreateRbacAuthorizationV1NamespacedRoleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1/namespaces/{namespace}/roles][%d] createRbacAuthorizationV1NamespacedRoleUnauthorized ", 401)
}

func (o *CreateRbacAuthorizationV1NamespacedRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
