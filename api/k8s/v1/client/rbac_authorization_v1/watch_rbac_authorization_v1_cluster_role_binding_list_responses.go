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

// WatchRbacAuthorizationV1ClusterRoleBindingListReader is a Reader for the WatchRbacAuthorizationV1ClusterRoleBindingList structure.
type WatchRbacAuthorizationV1ClusterRoleBindingListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1ClusterRoleBindingListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1ClusterRoleBindingListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1ClusterRoleBindingListOK creates a WatchRbacAuthorizationV1ClusterRoleBindingListOK with default headers values
func NewWatchRbacAuthorizationV1ClusterRoleBindingListOK() *WatchRbacAuthorizationV1ClusterRoleBindingListOK {
	return &WatchRbacAuthorizationV1ClusterRoleBindingListOK{}
}

/*WatchRbacAuthorizationV1ClusterRoleBindingListOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1ClusterRoleBindingListOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1ClusterRoleBindingListOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/clusterrolebindings][%d] watchRbacAuthorizationV1ClusterRoleBindingListOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1ClusterRoleBindingListOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1ClusterRoleBindingListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized creates a WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized with default headers values
func NewWatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized() *WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized {
	return &WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized{}
}

/*WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/clusterrolebindings][%d] watchRbacAuthorizationV1ClusterRoleBindingListUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1ClusterRoleBindingListUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}