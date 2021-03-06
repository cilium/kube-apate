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

// PatchRbacAuthorizationV1beta1NamespacedRoleReader is a Reader for the PatchRbacAuthorizationV1beta1NamespacedRole structure.
type PatchRbacAuthorizationV1beta1NamespacedRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRbacAuthorizationV1beta1NamespacedRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRbacAuthorizationV1beta1NamespacedRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRbacAuthorizationV1beta1NamespacedRoleOK creates a PatchRbacAuthorizationV1beta1NamespacedRoleOK with default headers values
func NewPatchRbacAuthorizationV1beta1NamespacedRoleOK() *PatchRbacAuthorizationV1beta1NamespacedRoleOK {
	return &PatchRbacAuthorizationV1beta1NamespacedRoleOK{}
}

/*PatchRbacAuthorizationV1beta1NamespacedRoleOK handles this case with default header values.

OK
*/
type PatchRbacAuthorizationV1beta1NamespacedRoleOK struct {
	Payload *models.IoK8sAPIRbacV1beta1Role
}

func (o *PatchRbacAuthorizationV1beta1NamespacedRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/roles/{name}][%d] patchRbacAuthorizationV1beta1NamespacedRoleOK  %+v", 200, o.Payload)
}

func (o *PatchRbacAuthorizationV1beta1NamespacedRoleOK) GetPayload() *models.IoK8sAPIRbacV1beta1Role {
	return o.Payload
}

func (o *PatchRbacAuthorizationV1beta1NamespacedRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized creates a PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized with default headers values
func NewPatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized() *PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized {
	return &PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized{}
}

/*PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized struct {
}

func (o *PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1beta1/namespaces/{namespace}/roles/{name}][%d] patchRbacAuthorizationV1beta1NamespacedRoleUnauthorized ", 401)
}

func (o *PatchRbacAuthorizationV1beta1NamespacedRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
