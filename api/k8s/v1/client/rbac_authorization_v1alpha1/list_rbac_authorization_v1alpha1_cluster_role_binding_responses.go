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

// ListRbacAuthorizationV1alpha1ClusterRoleBindingReader is a Reader for the ListRbacAuthorizationV1alpha1ClusterRoleBinding structure.
type ListRbacAuthorizationV1alpha1ClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRbacAuthorizationV1alpha1ClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRbacAuthorizationV1alpha1ClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRbacAuthorizationV1alpha1ClusterRoleBindingOK creates a ListRbacAuthorizationV1alpha1ClusterRoleBindingOK with default headers values
func NewListRbacAuthorizationV1alpha1ClusterRoleBindingOK() *ListRbacAuthorizationV1alpha1ClusterRoleBindingOK {
	return &ListRbacAuthorizationV1alpha1ClusterRoleBindingOK{}
}

/*ListRbacAuthorizationV1alpha1ClusterRoleBindingOK handles this case with default header values.

OK
*/
type ListRbacAuthorizationV1alpha1ClusterRoleBindingOK struct {
	Payload *models.IoK8sAPIRbacV1alpha1ClusterRoleBindingList
}

func (o *ListRbacAuthorizationV1alpha1ClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings][%d] listRbacAuthorizationV1alpha1ClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *ListRbacAuthorizationV1alpha1ClusterRoleBindingOK) GetPayload() *models.IoK8sAPIRbacV1alpha1ClusterRoleBindingList {
	return o.Payload
}

func (o *ListRbacAuthorizationV1alpha1ClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIRbacV1alpha1ClusterRoleBindingList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized creates a ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized with default headers values
func NewListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized() *ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized {
	return &ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized{}
}

/*ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized struct {
}

func (o *ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings][%d] listRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized ", 401)
}

func (o *ListRbacAuthorizationV1alpha1ClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
