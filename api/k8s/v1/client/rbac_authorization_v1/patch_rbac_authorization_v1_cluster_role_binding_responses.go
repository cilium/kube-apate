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

// PatchRbacAuthorizationV1ClusterRoleBindingReader is a Reader for the PatchRbacAuthorizationV1ClusterRoleBinding structure.
type PatchRbacAuthorizationV1ClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRbacAuthorizationV1ClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRbacAuthorizationV1ClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchRbacAuthorizationV1ClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRbacAuthorizationV1ClusterRoleBindingOK creates a PatchRbacAuthorizationV1ClusterRoleBindingOK with default headers values
func NewPatchRbacAuthorizationV1ClusterRoleBindingOK() *PatchRbacAuthorizationV1ClusterRoleBindingOK {
	return &PatchRbacAuthorizationV1ClusterRoleBindingOK{}
}

/*PatchRbacAuthorizationV1ClusterRoleBindingOK handles this case with default header values.

OK
*/
type PatchRbacAuthorizationV1ClusterRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1ClusterRoleBinding
}

func (o *PatchRbacAuthorizationV1ClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name}][%d] patchRbacAuthorizationV1ClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *PatchRbacAuthorizationV1ClusterRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1ClusterRoleBinding {
	return o.Payload
}

func (o *PatchRbacAuthorizationV1ClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRbacAuthorizationV1ClusterRoleBindingUnauthorized creates a PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized with default headers values
func NewPatchRbacAuthorizationV1ClusterRoleBindingUnauthorized() *PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized {
	return &PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized{}
}

/*PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized struct {
}

func (o *PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/{name}][%d] patchRbacAuthorizationV1ClusterRoleBindingUnauthorized ", 401)
}

func (o *PatchRbacAuthorizationV1ClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}