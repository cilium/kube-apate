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

// CreateRbacAuthorizationV1beta1ClusterRoleBindingReader is a Reader for the CreateRbacAuthorizationV1beta1ClusterRoleBinding structure.
type CreateRbacAuthorizationV1beta1ClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRbacAuthorizationV1beta1ClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateRbacAuthorizationV1beta1ClusterRoleBindingCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingOK creates a CreateRbacAuthorizationV1beta1ClusterRoleBindingOK with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingOK() *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK {
	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingOK{}
}

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingOK handles this case with default header values.

OK
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings][%d] createRbacAuthorizationV1beta1ClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1beta1ClusterRoleBinding {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingCreated creates a CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingCreated() *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated {
	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated{}
}

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated handles this case with default header values.

Created
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated struct {
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings][%d] createRbacAuthorizationV1beta1ClusterRoleBindingCreated  %+v", 201, o.Payload)
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated) GetPayload() *models.IoK8sAPIRbacV1beta1ClusterRoleBinding {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted creates a CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted() *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted {
	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted{}
}

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted handles this case with default header values.

Accepted
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted struct {
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleBinding
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings][%d] createRbacAuthorizationV1beta1ClusterRoleBindingAccepted  %+v", 202, o.Payload)
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted) GetPayload() *models.IoK8sAPIRbacV1beta1ClusterRoleBinding {
	return o.Payload
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1ClusterRoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized creates a CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized with default headers values
func NewCreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized() *CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized {
	return &CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized{}
}

/*CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized struct {
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings][%d] createRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized ", 401)
}

func (o *CreateRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
