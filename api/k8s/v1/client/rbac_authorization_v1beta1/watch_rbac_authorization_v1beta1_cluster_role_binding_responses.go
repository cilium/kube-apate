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

// WatchRbacAuthorizationV1beta1ClusterRoleBindingReader is a Reader for the WatchRbacAuthorizationV1beta1ClusterRoleBinding structure.
type WatchRbacAuthorizationV1beta1ClusterRoleBindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1beta1ClusterRoleBindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1beta1ClusterRoleBindingOK creates a WatchRbacAuthorizationV1beta1ClusterRoleBindingOK with default headers values
func NewWatchRbacAuthorizationV1beta1ClusterRoleBindingOK() *WatchRbacAuthorizationV1beta1ClusterRoleBindingOK {
	return &WatchRbacAuthorizationV1beta1ClusterRoleBindingOK{}
}

/*WatchRbacAuthorizationV1beta1ClusterRoleBindingOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1beta1ClusterRoleBindingOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/watch/clusterrolebindings/{name}][%d] watchRbacAuthorizationV1beta1ClusterRoleBindingOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized creates a WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized with default headers values
func NewWatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized() *WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized {
	return &WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized{}
}

/*WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1beta1/watch/clusterrolebindings/{name}][%d] watchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1beta1ClusterRoleBindingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
