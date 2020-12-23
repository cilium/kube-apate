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

// WatchCoreV1NamespacedPodTemplateReader is a Reader for the WatchCoreV1NamespacedPodTemplate structure.
type WatchCoreV1NamespacedPodTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WatchCoreV1NamespacedPodTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWatchCoreV1NamespacedPodTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewWatchCoreV1NamespacedPodTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWatchCoreV1NamespacedPodTemplateOK creates a WatchCoreV1NamespacedPodTemplateOK with default headers values
func NewWatchCoreV1NamespacedPodTemplateOK() *WatchCoreV1NamespacedPodTemplateOK {
	return &WatchCoreV1NamespacedPodTemplateOK{}
}

/*WatchCoreV1NamespacedPodTemplateOK handles this case with default header values.

OK
*/
type WatchCoreV1NamespacedPodTemplateOK struct {
	Payload *models.IoK8sApimachineryPkgApisMetaV1WatchEvent
}

func (o *WatchCoreV1NamespacedPodTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/podtemplates/{name}][%d] watchCoreV1NamespacedPodTemplateOK  %+v", 200, o.Payload)
}

func (o *WatchCoreV1NamespacedPodTemplateOK) GetPayload() *models.IoK8sApimachineryPkgApisMetaV1WatchEvent {
	return o.Payload
}

func (o *WatchCoreV1NamespacedPodTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sApimachineryPkgApisMetaV1WatchEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWatchCoreV1NamespacedPodTemplateUnauthorized creates a WatchCoreV1NamespacedPodTemplateUnauthorized with default headers values
func NewWatchCoreV1NamespacedPodTemplateUnauthorized() *WatchCoreV1NamespacedPodTemplateUnauthorized {
	return &WatchCoreV1NamespacedPodTemplateUnauthorized{}
}

/*WatchCoreV1NamespacedPodTemplateUnauthorized handles this case with default header values.

Unauthorized
*/
type WatchCoreV1NamespacedPodTemplateUnauthorized struct {
}

func (o *WatchCoreV1NamespacedPodTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/watch/namespaces/{namespace}/podtemplates/{name}][%d] watchCoreV1NamespacedPodTemplateUnauthorized ", 401)
}

func (o *WatchCoreV1NamespacedPodTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}