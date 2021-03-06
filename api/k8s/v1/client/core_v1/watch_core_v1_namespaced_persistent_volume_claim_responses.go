// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// WatchCoreV1NamespacedPersistentVolumeClaimReader is a Reader for the WatchCoreV1NamespacedPersistentVolumeClaim structure.
type WatchCoreV1NamespacedPersistentVolumeClaimReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedPersistentVolumeClaimReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedPersistentVolumeClaimOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedPersistentVolumeClaimUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedPersistentVolumeClaimOK creates a WatchCoreV1NamespacedPersistentVolumeClaimOK with default headers values
func NewWatchCoreV1NamespacedPersistentVolumeClaimOK() *WatchCoreV1NamespacedPersistentVolumeClaimOK {
	return &WatchCoreV1NamespacedPersistentVolumeClaimOK{}
}

/*WatchCoreV1NamespacedPersistentVolumeClaimOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedPersistentVolumeClaimOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaimOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/persistentvolumeclaims/{name}][%d] watchCoreV1NamespacedPersistentVolumeClaimOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaimOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaimOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedPersistentVolumeClaimUnauthorized creates a WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized with default headers values
func NewWatchCoreV1NamespacedPersistentVolumeClaimUnauthorized() *WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized {
	return &WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized{}
}

/*WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized struct {
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/persistentvolumeclaims/{name}][%d] watchCoreV1NamespacedPersistentVolumeClaimUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedPersistentVolumeClaimUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
