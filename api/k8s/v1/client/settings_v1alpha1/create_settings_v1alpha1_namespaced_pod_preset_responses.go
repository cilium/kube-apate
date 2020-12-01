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

// CreateSettingsV1alpha1NamespacedPodPresetReader is a Reader for the CreateSettingsV1alpha1NamespacedPodPreset structure.
type CreateSettingsV1alpha1NamespacedPodPresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSettingsV1alpha1NamespacedPodPresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateSettingsV1alpha1NamespacedPodPresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateSettingsV1alpha1NamespacedPodPresetCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateSettingsV1alpha1NamespacedPodPresetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateSettingsV1alpha1NamespacedPodPresetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSettingsV1alpha1NamespacedPodPresetOK creates a CreateSettingsV1alpha1NamespacedPodPresetOK with default headers values
func NewCreateSettingsV1alpha1NamespacedPodPresetOK() *CreateSettingsV1alpha1NamespacedPodPresetOK {
	return &CreateSettingsV1alpha1NamespacedPodPresetOK{}
}

/*CreateSettingsV1alpha1NamespacedPodPresetOK handles this case with default header values.

OK
*/
type CreateSettingsV1alpha1NamespacedPodPresetOK struct {
	Payload *models.IoK8sAPISettingsV1alpha1PodPreset
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetOK) Error() string {
	return fmt.Sprintf("[POST /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets][%d] createSettingsV1alpha1NamespacedPodPresetOK  %+v", 200, o.Payload)
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetOK) GetPayload() *models.IoK8sAPISettingsV1alpha1PodPreset {
	return o.Payload
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISettingsV1alpha1PodPreset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSettingsV1alpha1NamespacedPodPresetCreated creates a CreateSettingsV1alpha1NamespacedPodPresetCreated with default headers values
func NewCreateSettingsV1alpha1NamespacedPodPresetCreated() *CreateSettingsV1alpha1NamespacedPodPresetCreated {
	return &CreateSettingsV1alpha1NamespacedPodPresetCreated{}
}

/*CreateSettingsV1alpha1NamespacedPodPresetCreated handles this case with default header values.

Created
*/
type CreateSettingsV1alpha1NamespacedPodPresetCreated struct {
	Payload *models.IoK8sAPISettingsV1alpha1PodPreset
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetCreated) Error() string {
	return fmt.Sprintf("[POST /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets][%d] createSettingsV1alpha1NamespacedPodPresetCreated  %+v", 201, o.Payload)
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetCreated) GetPayload() *models.IoK8sAPISettingsV1alpha1PodPreset {
	return o.Payload
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISettingsV1alpha1PodPreset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSettingsV1alpha1NamespacedPodPresetAccepted creates a CreateSettingsV1alpha1NamespacedPodPresetAccepted with default headers values
func NewCreateSettingsV1alpha1NamespacedPodPresetAccepted() *CreateSettingsV1alpha1NamespacedPodPresetAccepted {
	return &CreateSettingsV1alpha1NamespacedPodPresetAccepted{}
}

/*CreateSettingsV1alpha1NamespacedPodPresetAccepted handles this case with default header values.

Accepted
*/
type CreateSettingsV1alpha1NamespacedPodPresetAccepted struct {
	Payload *models.IoK8sAPISettingsV1alpha1PodPreset
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets][%d] createSettingsV1alpha1NamespacedPodPresetAccepted  %+v", 202, o.Payload)
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetAccepted) GetPayload() *models.IoK8sAPISettingsV1alpha1PodPreset {
	return o.Payload
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPISettingsV1alpha1PodPreset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSettingsV1alpha1NamespacedPodPresetUnauthorized creates a CreateSettingsV1alpha1NamespacedPodPresetUnauthorized with default headers values
func NewCreateSettingsV1alpha1NamespacedPodPresetUnauthorized() *CreateSettingsV1alpha1NamespacedPodPresetUnauthorized {
	return &CreateSettingsV1alpha1NamespacedPodPresetUnauthorized{}
}

/*CreateSettingsV1alpha1NamespacedPodPresetUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateSettingsV1alpha1NamespacedPodPresetUnauthorized struct {
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/settings.k8s.io/v1alpha1/namespaces/{namespace}/podpresets][%d] createSettingsV1alpha1NamespacedPodPresetUnauthorized ", 401)
}

func (o *CreateSettingsV1alpha1NamespacedPodPresetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
