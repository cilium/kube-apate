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

// WatchRbacAuthorizationV1NamespacedRoleBindingReader is a Reader for the WatchRbacAuthorizationV1NamespacedRoleBinding structure.
type WatchRbacAuthorizationV1NamespacedRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1NamespacedRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1NamespacedRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingOK creates a WatchRbacAuthorizationV1NamespacedRoleBindingOK with default headers values
func NewWatchRbacAuthorizationV1NamespacedRoleBindingOK() *WatchRbacAuthorizationV1NamespacedRoleBindingOK {
	return &WatchRbacAuthorizationV1NamespacedRoleBindingOK{}
}

/*WatchRbacAuthorizationV1NamespacedRoleBindingOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1NamespacedRoleBindingOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/namespaces/{namespace}/rolebindings/{name}][%d] watchRbacAuthorizationV1NamespacedRoleBindingOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized creates a WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized with default headers values
func NewWatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized() *WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized {
	return &WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized{}
}

/*WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/namespaces/{namespace}/rolebindings/{name}][%d] watchRbacAuthorizationV1NamespacedRoleBindingUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1NamespacedRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
