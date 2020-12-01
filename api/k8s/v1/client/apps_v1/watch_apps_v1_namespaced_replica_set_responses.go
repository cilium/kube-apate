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

// WatchAppsV1NamespacedReplicaSetReader is a Reader for the WatchAppsV1NamespacedReplicaSet structure.
type WatchAppsV1NamespacedReplicaSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchAppsV1NamespacedReplicaSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchAppsV1NamespacedReplicaSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchAppsV1NamespacedReplicaSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchAppsV1NamespacedReplicaSetOK creates a WatchAppsV1NamespacedReplicaSetOK with default headers values
func NewWatchAppsV1NamespacedReplicaSetOK() *WatchAppsV1NamespacedReplicaSetOK {
	return &WatchAppsV1NamespacedReplicaSetOK{}
}

/*WatchAppsV1NamespacedReplicaSetOK handles this case with default header values.

OK
*/
type WatchAppsV1NamespacedReplicaSetOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchAppsV1NamespacedReplicaSetOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/namespaces/{namespace}/replicasets/{name}][%d] watchAppsV1NamespacedReplicaSetOK  %+v", 200, o.Payload)
}

func (o *WatchAppsV1NamespacedReplicaSetOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchAppsV1NamespacedReplicaSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchAppsV1NamespacedReplicaSetUnauthorized creates a WatchAppsV1NamespacedReplicaSetUnauthorized with default headers values
func NewWatchAppsV1NamespacedReplicaSetUnauthorized() *WatchAppsV1NamespacedReplicaSetUnauthorized {
	return &WatchAppsV1NamespacedReplicaSetUnauthorized{}
}

/*WatchAppsV1NamespacedReplicaSetUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchAppsV1NamespacedReplicaSetUnauthorized struct {
}

func (o *WatchAppsV1NamespacedReplicaSetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/watch/namespaces/{namespace}/replicasets/{name}][%d] watchAppsV1NamespacedReplicaSetUnauthorized ", 401)
}

func (o *WatchAppsV1NamespacedReplicaSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
