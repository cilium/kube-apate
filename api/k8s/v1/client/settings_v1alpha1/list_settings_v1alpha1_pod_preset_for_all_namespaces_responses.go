// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package settings_v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// ListSettingsV1alpha1PodPresetForAllNamespacesReader is a Reader for the ListSettingsV1alpha1PodPresetForAllNamespaces structure.
type ListSettingsV1alpha1PodPresetForAllNamespacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSettingsV1alpha1PodPresetForAllNamespacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSettingsV1alpha1PodPresetForAllNamespacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSettingsV1alpha1PodPresetForAllNamespacesOK creates a ListSettingsV1alpha1PodPresetForAllNamespacesOK with default headers values
func NewListSettingsV1alpha1PodPresetForAllNamespacesOK() *ListSettingsV1alpha1PodPresetForAllNamespacesOK {
	return &ListSettingsV1alpha1PodPresetForAllNamespacesOK{}
}

/*ListSettingsV1alpha1PodPresetForAllNamespacesOK handles this case with default header values.

OK
*/
type ListSettingsV1alpha1PodPresetForAllNamespacesOK struct {
	Payload *models.IoK8sAPISettingsV1alpha1PodPresetList
}

func (o *ListSettingsV1alpha1PodPresetForAllNamespacesOK) Error() string {
	return fmt.Sprintf("[GET /apis/settings.k8s.io/v1alpha1/podpresets][%d] listSettingsV1alpha1PodPresetForAllNamespacesOK  %+v", 200, o.Payload)
}

func (o *ListSettingsV1alpha1PodPresetForAllNamespacesOK) GetPayload() *models.IoK8sAPISettingsV1alpha1PodPresetList {
	return o.Payload
}

func (o *ListSettingsV1alpha1PodPresetForAllNamespacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISettingsV1alpha1PodPresetList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized creates a ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized with default headers values
func NewListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized() *ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized {
	return &ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized{}
}

/*ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized struct {
}

func (o *ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apis/settings.k8s.io/v1alpha1/podpresets][%d] listSettingsV1alpha1PodPresetForAllNamespacesUnauthorized ", 401)
}

func (o *ListSettingsV1alpha1PodPresetForAllNamespacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}