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

// DeleteRbacAuthorizationV1beta1ClusterRoleBindingReader is a Reader for the DeleteRbacAuthorizationV1beta1ClusterRoleBinding structure.
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingOK creates a DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK with default headers values
func NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingOK() *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK {
	return &DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK{}
}

/*DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK handles this case with default header values.

OK
*/
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[DELETE /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings/{name}][%d] deleteRbacAuthorizationV1beta1ClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted creates a DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted with default headers values
func NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted() *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted {
	return &DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted{}
}

/*DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted handles this case with default header values.

Accepted
*/
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1Status
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted) Error() string {
	return fmt.Sprintf("[DELETE /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings/{name}][%d] deleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted  %+v", 202, o.Payload)
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1Status {
	return o.Payload
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized creates a DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized with default headers values
func NewDeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized() *DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized {
	return &DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized{}
}

/*DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized struct {
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apis/rbac.authorization.k8s.io/v1beta1/clusterrolebindings/{name}][%d] deleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized ", 401)
}

func (o *DeleteRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}