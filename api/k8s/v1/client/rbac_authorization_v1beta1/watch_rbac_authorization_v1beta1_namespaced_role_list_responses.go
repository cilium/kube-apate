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

// WatchRbacAuthorizationV1beta1NamespacedRoleListReader is a Reader for the WatchRbacAuthorizationV1beta1NamespacedRoleList structure.
type WatchRbacAuthorizationV1beta1NamespacedRoleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1beta1NamespacedRoleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleListOK creates a WatchRbacAuthorizationV1beta1NamespacedRoleListOK with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleListOK() *WatchRbacAuthorizationV1beta1NamespacedRoleListOK {
	return &WatchRbacAuthorizationV1beta1NamespacedRoleListOK{}
}

/*WatchRbacAuthorizationV1beta1NamespacedRoleListOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/watch/namespaces/{namespace}/roles][%d] watchRbacAuthorizationV1beta1NamespacedRoleListOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized creates a WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized() *WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized {
	return &WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized{}
}

/*WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/watch/namespaces/{namespace}/roles][%d] watchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1beta1NamespacedRoleListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
