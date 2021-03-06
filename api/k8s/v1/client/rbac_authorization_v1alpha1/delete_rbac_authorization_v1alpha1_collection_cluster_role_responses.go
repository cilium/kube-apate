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

// DeleteRbacAuthorizationV1alpha1CollectionClusterRoleReader is a Reader for the DeleteRbacAuthorizationV1alpha1CollectionClusterRole structure.
type DeleteRbacAuthorizationV1alpha1CollectionClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK creates a DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK with default headers values
func NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK() *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK {
	return &DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK{}
}

/*DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK handles this case with default header values.

OK
*/
type DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles][%d] deleteRbacAuthorizationV1alpha1CollectionClusterRoleOK  %+v", 200, o.Payload)
}

func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized creates a DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized() *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized {
	return &DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized{}
}

/*DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized struct {
}

func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles][%d] deleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized ", 401)
}

func (o *DeleteRbacAuthorizationV1alpha1CollectionClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
