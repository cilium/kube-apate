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

// PatchRbacAuthorizationV1alpha1ClusterRoleReader is a Reader for the PatchRbacAuthorizationV1alpha1ClusterRole structure.
type PatchRbacAuthorizationV1alpha1ClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRbacAuthorizationV1alpha1ClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRbacAuthorizationV1alpha1ClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRbacAuthorizationV1alpha1ClusterRoleOK creates a PatchRbacAuthorizationV1alpha1ClusterRoleOK with default headers values
func NewPatchRbacAuthorizationV1alpha1ClusterRoleOK() *PatchRbacAuthorizationV1alpha1ClusterRoleOK {
	return &PatchRbacAuthorizationV1alpha1ClusterRoleOK{}
}

/*PatchRbacAuthorizationV1alpha1ClusterRoleOK handles this case with default header values.

OK
*/
type PatchRbacAuthorizationV1alpha1ClusterRoleOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRole
}

func (o *PatchRbacAuthorizationV1alpha1ClusterRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/{name}][%d] patchRbacAuthorizationV1alpha1ClusterRoleOK  %+v", 200, o.Payload)
}

func (o *PatchRbacAuthorizationV1alpha1ClusterRoleOK) GetPayload() *models.IoK8sAPIRbacV1alpha1ClusterRole {
	return o.Payload
}

func (o *PatchRbacAuthorizationV1alpha1ClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized creates a PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized with default headers values
func NewPatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized() *PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized {
	return &PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized{}
}

/*PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized struct {
}

func (o *PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/{name}][%d] patchRbacAuthorizationV1alpha1ClusterRoleUnauthorized ", 401)
}

func (o *PatchRbacAuthorizationV1alpha1ClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}