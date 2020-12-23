// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateRbacAuthorizationV1alpha1NamespacedRoleBindingReader is a Reader for the CreateRbacAuthorizationV1alpha1NamespacedRoleBinding structure.
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK creates a CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK with default headers values
func NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK() *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK {
	return &CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK{}
}

/*CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK handles this case with default header values.

OK
*/
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1RoleBinding
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings][%d] createRbacAuthorizationV1alpha1NamespacedRoleBindingOK  %+v", 200, o.Payload)
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1alpha1RoleBinding {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated creates a CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated with default headers values
func NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated() *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated {
	return &CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated{}
}

/*CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated handles this case with default header values.

Created
*/
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated struct {
	Payload *models.IoK8sAPIRbacV1alpha1RoleBinding
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings][%d] createRbacAuthorizationV1alpha1NamespacedRoleBindingCreated  %+v", 201, o.Payload)
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated) GetPayload() *models.IoK8sAPIRbacV1alpha1RoleBinding {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted creates a CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted with default headers values
func NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted() *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted {
	return &CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted{}
}

/*CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted handles this case with default header values.

Accepted
*/
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted struct {
	Payload *models.IoK8sAPIRbacV1alpha1RoleBinding
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings][%d] createRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted  %+v", 202, o.Payload)
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted) GetPayload() *models.IoK8sAPIRbacV1alpha1RoleBinding {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized creates a CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized with default headers values
func NewCreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized() *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized {
	return &CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized{}
}

/*CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized struct {
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings][%d] createRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized ", 401)
}

func (o *CreateRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}