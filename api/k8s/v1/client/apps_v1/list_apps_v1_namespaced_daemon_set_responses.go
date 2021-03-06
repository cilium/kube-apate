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

// ListAppsV1NamespacedDaemonSetReader is a Reader for the ListAppsV1NamespacedDaemonSet structure.
type ListAppsV1NamespacedDaemonSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAppsV1NamespacedDaemonSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAppsV1NamespacedDaemonSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListAppsV1NamespacedDaemonSetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAppsV1NamespacedDaemonSetOK creates a ListAppsV1NamespacedDaemonSetOK with default headers values
func NewListAppsV1NamespacedDaemonSetOK() *ListAppsV1NamespacedDaemonSetOK {
	return &ListAppsV1NamespacedDaemonSetOK{}
}

/*ListAppsV1NamespacedDaemonSetOK handles this case with default header values.

OK
*/
type ListAppsV1NamespacedDaemonSetOK struct {
	Payload *models.IoK8sAPIAppsV1DaemonSetList
}

func (o *ListAppsV1NamespacedDaemonSetOK) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/daemonsets][%d] listAppsV1NamespacedDaemonSetOK  %+v", 200, o.Payload)
}

func (o *ListAppsV1NamespacedDaemonSetOK) GetPayload() *models.IoK8sAPIAppsV1DaemonSetList {
	return o.Payload
}

func (o *ListAppsV1NamespacedDaemonSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAppsV1DaemonSetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAppsV1NamespacedDaemonSetUnauthorized creates a ListAppsV1NamespacedDaemonSetUnauthorized with default headers values
func NewListAppsV1NamespacedDaemonSetUnauthorized() *ListAppsV1NamespacedDaemonSetUnauthorized {
	return &ListAppsV1NamespacedDaemonSetUnauthorized{}
}

/*ListAppsV1NamespacedDaemonSetUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAppsV1NamespacedDaemonSetUnauthorized struct {
}

func (o *ListAppsV1NamespacedDaemonSetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/apps/v1/namespaces/{namespace}/daemonsets][%d] listAppsV1NamespacedDaemonSetUnauthorized ", 401)
}

func (o *ListAppsV1NamespacedDaemonSetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
