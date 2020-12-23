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

// WatchCoreV1PodTemplateListForAllNamespacesReader is a Reader for the WatchCoreV1PodTemplateListForAllNamespaces structure.
type WatchCoreV1PodTemplateListForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1PodTemplateListForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1PodTemplateListForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1PodTemplateListForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1PodTemplateListForAllNamespacesOK creates a WatchCoreV1PodTemplateListForAllNamespacesOK with default headers values
func NewWatchCoreV1PodTemplateListForAllNamespacesOK() *WatchCoreV1PodTemplateListForAllNamespacesOK {
	return &WatchCoreV1PodTemplateListForAllNamespacesOK{}
}

/*WatchCoreV1PodTemplateListForAllNamespacesOK handles this case with default header values.

OK
*/
type WatchCoreV1PodTemplateListForAllNamespacesOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1PodTemplateListForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/podtemplates][%d] watchCoreV1PodTemplateListForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1PodTemplateListForAllNamespacesOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1PodTemplateListForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1PodTemplateListForAllNamespacesUnauthorized creates a WatchCoreV1PodTemplateListForAllNamespacesUnauthorized with default headers values
func NewWatchCoreV1PodTemplateListForAllNamespacesUnauthorized() *WatchCoreV1PodTemplateListForAllNamespacesUnauthorized {
	return &WatchCoreV1PodTemplateListForAllNamespacesUnauthorized{}
}

/*WatchCoreV1PodTemplateListForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1PodTemplateListForAllNamespacesUnauthorized struct {
}

func (o *WatchCoreV1PodTemplateListForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/podtemplates][%d] watchCoreV1PodTemplateListForAllNamespacesUnauthorized ", 401)
}

func (o *WatchCoreV1PodTemplateListForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}