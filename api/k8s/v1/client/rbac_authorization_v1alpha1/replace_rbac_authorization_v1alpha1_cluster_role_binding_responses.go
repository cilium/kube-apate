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

// ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingReader is a Reader for the ReplaceRbacAuthorizationV1alpha1ClusterRoleBinding structure.
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK creates a ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK with default headers values
func NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK() *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK {
	return &ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK{}
}

/*ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK handles this case with default header values.

OK
*/
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}][%d] replaceRbacAuthorizationV1alpha1ClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated creates a ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated with default headers values
func NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated() *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated {
	return &ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated{}
}

/*ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated handles this case with default header values.

Created
*/
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated struct {
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}][%d] replaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated  %+v", 201, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated) GetPayload() *models.IoK8sAPIRbacV1alpha1ClusterRoleBinding {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized creates a ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized() *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized {
	return &ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized{}
}

/*ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized struct {
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}][%d] replaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized ", 401)
}

func (o *ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
