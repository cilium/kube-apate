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

// ReplaceRbacAuthorizationV1alpha1NamespacedRoleReader is a Reader for the ReplaceRbacAuthorizationV1alpha1NamespacedRole structure.
type ReplaceRbacAuthorizationV1alpha1NamespacedRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRbacAuthorizationV1alpha1NamespacedRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceRbacAuthorizationV1alpha1NamespacedRoleOK creates a ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK with default headers values
func NewReplaceRbacAuthorizationV1alpha1NamespacedRoleOK() *ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK {
	return &ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK{}
}

/*ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK handles this case with default header values.

OK
*/
type ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1Role
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles/{name}][%d] replaceRbacAuthorizationV1alpha1NamespacedRoleOK  %+v", 200, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK) GetPayload() *models.IoK8sAPIRbacV1alpha1Role {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated creates a ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated with default headers values
func NewReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated() *ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated {
	return &ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated{}
}

/*ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated handles this case with default header values.

Created
*/
type ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated struct {
	Payload *models.IoK8sAPIRbacV1alpha1Role
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles/{name}][%d] replaceRbacAuthorizationV1alpha1NamespacedRoleCreated  %+v", 201, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated) GetPayload() *models.IoK8sAPIRbacV1alpha1Role {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized creates a ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized() *ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized {
	return &ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized{}
}

/*ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized struct {
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles/{name}][%d] replaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized ", 401)
}

func (o *ReplaceRbacAuthorizationV1alpha1NamespacedRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}