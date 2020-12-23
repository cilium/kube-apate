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

// ReadRbacAuthorizationV1alpha1NamespacedRoleBindingReader is a Reader for the ReadRbacAuthorizationV1alpha1NamespacedRoleBinding structure.
type ReadRbacAuthorizationV1alpha1NamespacedRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK creates a ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK with default headers values
func NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK() *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK {
	return &ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK{}
}

/*ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK handles this case with default header values.

OK
*/
type ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1RoleBinding
}

func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings/{name}][%d] readRbacAuthorizationV1alpha1NamespacedRoleBindingOK  %+v", 200, o.Payload)
}

func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1alpha1RoleBinding {
	return o.Payload
}

func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized creates a ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized with default headers values
func NewReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized() *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized {
	return &ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized{}
}

/*ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized struct {
}

func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1alpha1/namespaces/{namespace}/rolebindings/{name}][%d] readRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized ", 401)
}

func (o *ReadRbacAuthorizationV1alpha1NamespacedRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}