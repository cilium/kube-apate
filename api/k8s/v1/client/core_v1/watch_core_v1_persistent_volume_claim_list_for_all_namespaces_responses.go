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

// WatchCoreV1PersistentVolumeClaimListForAllNamespacesReader is a Reader for the WatchCoreV1PersistentVolumeClaimListForAllNamespaces structure.
type WatchCoreV1PersistentVolumeClaimListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesOK creates a WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK with default headers values
func NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesOK() *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK {
	return &WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK{}
}

/*WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/persistentvolumeclaims][%d] watchCoreV1PersistentVolumeClaimListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized creates a WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized() *WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized {
	return &WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized{}
}

/*WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized struct {
}

func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/persistentvolumeclaims][%d] watchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchCoreV1PersistentVolumeClaimListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
