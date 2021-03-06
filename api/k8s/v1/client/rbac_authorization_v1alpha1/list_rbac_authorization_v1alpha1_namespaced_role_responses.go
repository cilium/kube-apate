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

// ListRbacAuthorizationV1alpha1NamespacedRoleReader is a Reader for the ListRbacAuthorizationV1alpha1NamespacedRole structure.
type ListRbacAuthorizationV1alpha1NamespacedRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRbacAuthorizationV1alpha1NamespacedRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRbacAuthorizationV1alpha1NamespacedRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRbacAuthorizationV1alpha1NamespacedRoleOK creates a ListRbacAuthorizationV1alpha1NamespacedRoleOK with default headers values
func NewListRbacAuthorizationV1alpha1NamespacedRoleOK() *ListRbacAuthorizationV1alpha1NamespacedRoleOK {
	return &ListRbacAuthorizationV1alpha1NamespacedRoleOK{}
}

/*ListRbacAuthorizationV1alpha1NamespacedRoleOK handles this case with default header values.

OK
*/
type ListRbacAuthorizationV1alpha1NamespacedRoleOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1RoleList
}

func (o *ListRbacAuthorizationV1alpha1NamespacedRoleOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles][%d] listRbacAuthorizationV1alpha1NamespacedRoleOK  %+v", 200, o.Payload)
}

func (o *ListRbacAuthorizationV1alpha1NamespacedRoleOK) GetPayload() *models.IoK8sAPIRbacV1alpha1RoleList {
	return o.Payload
}

func (o *ListRbacAuthorizationV1alpha1NamespacedRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1RoleList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized creates a ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized with default headers values
func NewListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized() *ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized {
	return &ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized{}
}

/*ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized struct {
}

func (o *ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/roles][%d] listRbacAuthorizationV1alpha1NamespacedRoleUnauthorized ", 401)
}

func (o *ListRbacAuthorizationV1alpha1NamespacedRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
