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

// WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesReader is a Reader for the WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespaces structure.
type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK creates a WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK with default headers values
func NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK() *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK {
	return &WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK{}
}

/*WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/watch/rolebindings][%d] watchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized creates a WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized() *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized {
	return &WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized{}
}

/*WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/watch/rolebindings][%d] watchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
