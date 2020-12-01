// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchAppsV1NamespacedDeploymentReader is a Reader for the WatchAppsV1NamespacedDeployment structure.
type WatchAppsV1NamespacedDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAppsV1NamespacedDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAppsV1NamespacedDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAppsV1NamespacedDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAppsV1NamespacedDeploymentOK creates a WatchAppsV1NamespacedDeploymentOK with default headers values
func NewWatchAppsV1NamespacedDeploymentOK() *WatchAppsV1NamespacedDeploymentOK {
	return &WatchAppsV1NamespacedDeploymentOK{}
}

/*WatchAppsV1NamespacedDeploymentOK handles this case with default header values.

OK
*/
type WatchAppsV1NamespacedDeploymentOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAppsV1NamespacedDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/namespaces/{namespace}/deployments/{name}][%d] watchAppsV1NamespacedDeploymentOK  %+v", 200, o.Payload)
}

func (o *WatchAppsV1NamespacedDeploymentOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAppsV1NamespacedDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAppsV1NamespacedDeploymentUnauthorized creates a WatchAppsV1NamespacedDeploymentUnauthorized with default headers values
func NewWatchAppsV1NamespacedDeploymentUnauthorized() *WatchAppsV1NamespacedDeploymentUnauthorized {
	return &WatchAppsV1NamespacedDeploymentUnauthorized{}
}

/*WatchAppsV1NamespacedDeploymentUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAppsV1NamespacedDeploymentUnauthorized struct {
}

func (o *WatchAppsV1NamespacedDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/namespaces/{namespace}/deployments/{name}][%d] watchAppsV1NamespacedDeploymentUnauthorized ", 401)
}

func (o *WatchAppsV1NamespacedDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
