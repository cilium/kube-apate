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

// WatchAppsV1NamespacedDaemonSetReader is a Reader for the WatchAppsV1NamespacedDaemonSet structure.
type WatchAppsV1NamespacedDaemonSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAppsV1NamespacedDaemonSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAppsV1NamespacedDaemonSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAppsV1NamespacedDaemonSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAppsV1NamespacedDaemonSetOK creates a WatchAppsV1NamespacedDaemonSetOK with default headers values
func NewWatchAppsV1NamespacedDaemonSetOK() *WatchAppsV1NamespacedDaemonSetOK {
	return &WatchAppsV1NamespacedDaemonSetOK{}
}

/*WatchAppsV1NamespacedDaemonSetOK handles this case with default header values.

OK
*/
type WatchAppsV1NamespacedDaemonSetOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAppsV1NamespacedDaemonSetOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/namespaces/{namespace}/daemonsets/{name}][%d] watchAppsV1NamespacedDaemonSetOK  %+v", 200, o.Payload)
}

func (o *WatchAppsV1NamespacedDaemonSetOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAppsV1NamespacedDaemonSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAppsV1NamespacedDaemonSetUnauthorized creates a WatchAppsV1NamespacedDaemonSetUnauthorized with default headers values
func NewWatchAppsV1NamespacedDaemonSetUnauthorized() *WatchAppsV1NamespacedDaemonSetUnauthorized {
	return &WatchAppsV1NamespacedDaemonSetUnauthorized{}
}

/*WatchAppsV1NamespacedDaemonSetUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAppsV1NamespacedDaemonSetUnauthorized struct {
}

func (o *WatchAppsV1NamespacedDaemonSetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/namespaces/{namespace}/daemonsets/{name}][%d] watchAppsV1NamespacedDaemonSetUnauthorized ", 401)
}

func (o *WatchAppsV1NamespacedDaemonSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
