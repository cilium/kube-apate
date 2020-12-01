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

// ListRbacAuthorizationV1beta1ClusterRoleReader is a Reader for the ListRbacAuthorizationV1beta1ClusterRole structure.
type ListRbacAuthorizationV1beta1ClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRbacAuthorizationV1beta1ClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRbacAuthorizationV1beta1ClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRbacAuthorizationV1beta1ClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRbacAuthorizationV1beta1ClusterRoleOK creates a ListRbacAuthorizationV1beta1ClusterRoleOK with default headers values
func NewListRbacAuthorizationV1beta1ClusterRoleOK() *ListRbacAuthorizationV1beta1ClusterRoleOK {
	return &ListRbacAuthorizationV1beta1ClusterRoleOK{}
}

/*ListRbacAuthorizationV1beta1ClusterRoleOK handles this case with default header values.

OK
*/
type ListRbacAuthorizationV1beta1ClusterRoleOK struct {
	Payload *models.IoK8sAPIRbacV1beta1ClusterRoleList
}

func (o *ListRbacAuthorizationV1beta1ClusterRoleOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/clusterroles][%d] listRbacAuthorizationV1beta1ClusterRoleOK  %+v", 200, o.Payload)
}

func (o *ListRbacAuthorizationV1beta1ClusterRoleOK) GetPayload() *models.IoK8sAPIRbacV1beta1ClusterRoleList {
	return o.Payload
}

func (o *ListRbacAuthorizationV1beta1ClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1beta1ClusterRoleList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRbacAuthorizationV1beta1ClusterRoleUnauthorized creates a ListRbacAuthorizationV1beta1ClusterRoleUnauthorized with default headers values
func NewListRbacAuthorizationV1beta1ClusterRoleUnauthorized() *ListRbacAuthorizationV1beta1ClusterRoleUnauthorized {
	return &ListRbacAuthorizationV1beta1ClusterRoleUnauthorized{}
}

/*ListRbacAuthorizationV1beta1ClusterRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRbacAuthorizationV1beta1ClusterRoleUnauthorized struct {
}

func (o *ListRbacAuthorizationV1beta1ClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/clusterroles][%d] listRbacAuthorizationV1beta1ClusterRoleUnauthorized ", 401)
}

func (o *ListRbacAuthorizationV1beta1ClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
