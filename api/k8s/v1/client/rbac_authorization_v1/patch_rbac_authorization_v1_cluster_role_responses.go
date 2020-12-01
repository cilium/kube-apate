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

// PatchRbacAuthorizationV1ClusterRoleReader is a Reader for the PatchRbacAuthorizationV1ClusterRole structure.
type PatchRbacAuthorizationV1ClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRbacAuthorizationV1ClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRbacAuthorizationV1ClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchRbacAuthorizationV1ClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRbacAuthorizationV1ClusterRoleOK creates a PatchRbacAuthorizationV1ClusterRoleOK with default headers values
func NewPatchRbacAuthorizationV1ClusterRoleOK() *PatchRbacAuthorizationV1ClusterRoleOK {
	return &PatchRbacAuthorizationV1ClusterRoleOK{}
}

/*PatchRbacAuthorizationV1ClusterRoleOK handles this case with default header values.

OK
*/
type PatchRbacAuthorizationV1ClusterRoleOK struct {
	Payload *models.IoK8sAPIRbacV1ClusterRole
}

func (o *PatchRbacAuthorizationV1ClusterRoleOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}][%d] patchRbacAuthorizationV1ClusterRoleOK  %+v", 200, o.Payload)
}

func (o *PatchRbacAuthorizationV1ClusterRoleOK) GetPayload() *models.IoK8sAPIRbacV1ClusterRole {
	return o.Payload
}

func (o *PatchRbacAuthorizationV1ClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRbacAuthorizationV1ClusterRoleUnauthorized creates a PatchRbacAuthorizationV1ClusterRoleUnauthorized with default headers values
func NewPatchRbacAuthorizationV1ClusterRoleUnauthorized() *PatchRbacAuthorizationV1ClusterRoleUnauthorized {
	return &PatchRbacAuthorizationV1ClusterRoleUnauthorized{}
}

/*PatchRbacAuthorizationV1ClusterRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchRbacAuthorizationV1ClusterRoleUnauthorized struct {
}

func (o *PatchRbacAuthorizationV1ClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}][%d] patchRbacAuthorizationV1ClusterRoleUnauthorized ", 401)
}

func (o *PatchRbacAuthorizationV1ClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
