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

// WatchRbacAuthorizationV1NamespacedRoleBindingListReader is a Reader for the WatchRbacAuthorizationV1NamespacedRoleBindingList structure.
type WatchRbacAuthorizationV1NamespacedRoleBindingListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1NamespacedRoleBindingListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1NamespacedRoleBindingListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingListOK creates a WatchRbacAuthorizationV1NamespacedRoleBindingListOK with default headers values
func NewWatchRbacAuthorizationV1NamespacedRoleBindingListOK() *WatchRbacAuthorizationV1NamespacedRoleBindingListOK {
	return &WatchRbacAuthorizationV1NamespacedRoleBindingListOK{}
}

/*WatchRbacAuthorizationV1NamespacedRoleBindingListOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1NamespacedRoleBindingListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingListOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/namespaces/{namespace}/rolebindings][%d] watchRbacAuthorizationV1NamespacedRoleBindingListOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized creates a WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized with default headers values
func NewWatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized() *WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized {
	return &WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized{}
}

/*WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/namespaces/{namespace}/rolebindings][%d] watchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
