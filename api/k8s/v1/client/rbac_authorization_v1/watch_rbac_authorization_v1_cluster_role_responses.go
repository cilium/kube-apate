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

// WatchRbacAuthorizationV1ClusterRoleReader is a Reader for the WatchRbacAuthorizationV1ClusterRole structure.
type WatchRbacAuthorizationV1ClusterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchRbacAuthorizationV1ClusterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchRbacAuthorizationV1ClusterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchRbacAuthorizationV1ClusterRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchRbacAuthorizationV1ClusterRoleOK creates a WatchRbacAuthorizationV1ClusterRoleOK with default headers values
func NewWatchRbacAuthorizationV1ClusterRoleOK() *WatchRbacAuthorizationV1ClusterRoleOK {
	return &WatchRbacAuthorizationV1ClusterRoleOK{}
}

/*WatchRbacAuthorizationV1ClusterRoleOK handles this case with default header values.

OK
*/
type WatchRbacAuthorizationV1ClusterRoleOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchRbacAuthorizationV1ClusterRoleOK) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/clusterroles/{name}][%d] watchRbacAuthorizationV1ClusterRoleOK  %+v", 200, o.Payload)
}

func (o *WatchRbacAuthorizationV1ClusterRoleOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchRbacAuthorizationV1ClusterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchRbacAuthorizationV1ClusterRoleUnauthorized creates a WatchRbacAuthorizationV1ClusterRoleUnauthorized with default headers values
func NewWatchRbacAuthorizationV1ClusterRoleUnauthorized() *WatchRbacAuthorizationV1ClusterRoleUnauthorized {
	return &WatchRbacAuthorizationV1ClusterRoleUnauthorized{}
}

/*WatchRbacAuthorizationV1ClusterRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchRbacAuthorizationV1ClusterRoleUnauthorized struct {
}

func (o *WatchRbacAuthorizationV1ClusterRoleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/rbac.authorization.k8s.io/v1/watch/clusterroles/{name}][%d] watchRbacAuthorizationV1ClusterRoleUnauthorized ", 401)
}

func (o *WatchRbacAuthorizationV1ClusterRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
