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

// WatchAppsV1ControllerRevisionListForAllNamespacesReader is a Reader for the WatchAppsV1ControllerRevisionListForAllNamespaces structure.
type WatchAppsV1ControllerRevisionListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAppsV1ControllerRevisionListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAppsV1ControllerRevisionListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAppsV1ControllerRevisionListForAllNamespacesOK creates a WatchAppsV1ControllerRevisionListForAllNamespacesOK with default headers values
func NewWatchAppsV1ControllerRevisionListForAllNamespacesOK() *WatchAppsV1ControllerRevisionListForAllNamespacesOK {
	return &WatchAppsV1ControllerRevisionListForAllNamespacesOK{}
}

/*WatchAppsV1ControllerRevisionListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchAppsV1ControllerRevisionListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAppsV1ControllerRevisionListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/controllerrevisions][%d] watchAppsV1ControllerRevisionListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchAppsV1ControllerRevisionListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAppsV1ControllerRevisionListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized creates a WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized with default headers values
func NewWatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized() *WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized {
	return &WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized{}
}

/*WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized struct {
}

func (o *WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/controllerrevisions][%d] watchAppsV1ControllerRevisionListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchAppsV1ControllerRevisionListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
