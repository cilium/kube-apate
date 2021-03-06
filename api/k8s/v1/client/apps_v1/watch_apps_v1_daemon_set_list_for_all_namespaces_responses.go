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

// WatchAppsV1DaemonSetListForAllNamespacesReader is a Reader for the WatchAppsV1DaemonSetListForAllNamespaces structure.
type WatchAppsV1DaemonSetListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAppsV1DaemonSetListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAppsV1DaemonSetListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAppsV1DaemonSetListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAppsV1DaemonSetListForAllNamespacesOK creates a WatchAppsV1DaemonSetListForAllNamespacesOK with default headers values
func NewWatchAppsV1DaemonSetListForAllNamespacesOK() *WatchAppsV1DaemonSetListForAllNamespacesOK {
	return &WatchAppsV1DaemonSetListForAllNamespacesOK{}
}

/*WatchAppsV1DaemonSetListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchAppsV1DaemonSetListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAppsV1DaemonSetListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/daemonsets][%d] watchAppsV1DaemonSetListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchAppsV1DaemonSetListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAppsV1DaemonSetListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAppsV1DaemonSetListForAllNamespacesUnauthorized creates a WatchAppsV1DaemonSetListForAllNamespacesUnauthorized with default headers values
func NewWatchAppsV1DaemonSetListForAllNamespacesUnauthorized() *WatchAppsV1DaemonSetListForAllNamespacesUnauthorized {
	return &WatchAppsV1DaemonSetListForAllNamespacesUnauthorized{}
}

/*WatchAppsV1DaemonSetListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAppsV1DaemonSetListForAllNamespacesUnauthorized struct {
}

func (o *WatchAppsV1DaemonSetListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/daemonsets][%d] watchAppsV1DaemonSetListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchAppsV1DaemonSetListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
