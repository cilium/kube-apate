// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package rbac_authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingReader is a Reader for the ReplaceRbacAuthorizationV1beta1NamespacedRoleBinding structure.
type ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK creates a ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK with default headers values
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK() *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK {
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK{}
}

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK handles this case with default header values.

OK
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1beta1RoleBinding
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}][%d] replaceRbacAuthorizationV1beta1NamespacedRoleBindingOK  %+v", 200, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1beta1RoleBinding {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated creates a ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated with default headers values
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated() *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated {
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated{}
}

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated handles this case with default header values.

Created
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated struct {
	Payload *models.IoK8sAPIRbacV1beta1RoleBinding
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}][%d] replaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated  %+v", 201, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated) GetPayload() *models.IoK8sAPIRbacV1beta1RoleBinding {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized creates a ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized() *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized {
	return &ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized{}
}

/*ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized struct {
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/rolebindings/{name}][%d] replaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized ", 401)
}

func (o *ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}