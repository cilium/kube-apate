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

// ReplaceRbacAuthorizationV1ClusterRoleReader is a Reader for the ReplaceRbacAuthorizationV1ClusterRole structure.
type ReplaceRbacAuthorizationV1ClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceRbacAuthorizationV1ClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceRbacAuthorizationV1ClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceRbacAuthorizationV1ClusterRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceRbacAuthorizationV1ClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceRbacAuthorizationV1ClusterRoleOK creates a ReplaceRbacAuthorizationV1ClusterRoleOK with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleOK() *ReplaceRbacAuthorizationV1ClusterRoleOK {
	return &ReplaceRbacAuthorizationV1ClusterRoleOK{}
}

/*ReplaceRbacAuthorizationV1ClusterRoleOK handles this case with default header values.

OK
*/
type ReplaceRbacAuthorizationV1ClusterRoleOK struct {
	Payload *models.IoK8sAPIRbacV1ClusterRole
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleOK) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}][%d] replaceRbacAuthorizationV1ClusterRoleOK  %+v", 200, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleOK) GetPayload() *models.IoK8sAPIRbacV1ClusterRole {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1ClusterRoleCreated creates a ReplaceRbacAuthorizationV1ClusterRoleCreated with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleCreated() *ReplaceRbacAuthorizationV1ClusterRoleCreated {
	return &ReplaceRbacAuthorizationV1ClusterRoleCreated{}
}

/*ReplaceRbacAuthorizationV1ClusterRoleCreated handles this case with default header values.

Created
*/
type ReplaceRbacAuthorizationV1ClusterRoleCreated struct {
	Payload *models.IoK8sAPIRbacV1ClusterRole
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleCreated) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}][%d] replaceRbacAuthorizationV1ClusterRoleCreated  %+v", 201, o.Payload)
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleCreated) GetPayload() *models.IoK8sAPIRbacV1ClusterRole {
	return o.Payload
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1ClusterRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceRbacAuthorizationV1ClusterRoleUnauthorized creates a ReplaceRbacAuthorizationV1ClusterRoleUnauthorized with default headers values
func NewReplaceRbacAuthorizationV1ClusterRoleUnauthorized() *ReplaceRbacAuthorizationV1ClusterRoleUnauthorized {
	return &ReplaceRbacAuthorizationV1ClusterRoleUnauthorized{}
}

/*ReplaceRbacAuthorizationV1ClusterRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceRbacAuthorizationV1ClusterRoleUnauthorized struct {
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /apis/rbac.authorization.k8s.io/v1/clusterroles/{name}][%d] replaceRbacAuthorizationV1ClusterRoleUnauthorized ", 401)
}

func (o *ReplaceRbacAuthorizationV1ClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
