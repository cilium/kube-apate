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

// DeleteCoreV1NamespacedPodTemplateReader is a Reader for the DeleteCoreV1NamespacedPodTemplate structure.
type DeleteCoreV1NamespacedPodTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoreV1NamespacedPodTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoreV1NamespacedPodTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteCoreV1NamespacedPodTemplateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCoreV1NamespacedPodTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoreV1NamespacedPodTemplateOK creates a DeleteCoreV1NamespacedPodTemplateOK with default headers values
func NewDeleteCoreV1NamespacedPodTemplateOK() *DeleteCoreV1NamespacedPodTemplateOK {
	return &DeleteCoreV1NamespacedPodTemplateOK{}
}

/*DeleteCoreV1NamespacedPodTemplateOK handles this case with default header values.

OK
*/
type DeleteCoreV1NamespacedPodTemplateOK struct {
	Payload *models.IoK8sAPICoreV1PodTemplate
}

func (o *DeleteCoreV1NamespacedPodTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/podtemplates/{name}][%d] deleteCoreV1NamespacedPodTemplateOK  %+v", 200, o.Payload)
}

func (o *DeleteCoreV1NamespacedPodTemplateOK) GetPayload() *models.IoK8sAPICoreV1PodTemplate {
	return o.Payload
}

func (o *DeleteCoreV1NamespacedPodTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PodTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1NamespacedPodTemplateAccepted creates a DeleteCoreV1NamespacedPodTemplateAccepted with default headers values
func NewDeleteCoreV1NamespacedPodTemplateAccepted() *DeleteCoreV1NamespacedPodTemplateAccepted {
	return &DeleteCoreV1NamespacedPodTemplateAccepted{}
}

/*DeleteCoreV1NamespacedPodTemplateAccepted handles this case with default header values.

Accepted
*/
type DeleteCoreV1NamespacedPodTemplateAccepted struct {
	Payload *models.IoK8sAPICoreV1PodTemplate
}

func (o *DeleteCoreV1NamespacedPodTemplateAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/podtemplates/{name}][%d] deleteCoreV1NamespacedPodTemplateAccepted  %+v", 202, o.Payload)
}

func (o *DeleteCoreV1NamespacedPodTemplateAccepted) GetPayload() *models.IoK8sAPICoreV1PodTemplate {
	return o.Payload
}

func (o *DeleteCoreV1NamespacedPodTemplateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPICoreV1PodTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoreV1NamespacedPodTemplateUnauthorized creates a DeleteCoreV1NamespacedPodTemplateUnauthorized with default headers values
func NewDeleteCoreV1NamespacedPodTemplateUnauthorized() *DeleteCoreV1NamespacedPodTemplateUnauthorized {
	return &DeleteCoreV1NamespacedPodTemplateUnauthorized{}
}

/*DeleteCoreV1NamespacedPodTemplateUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCoreV1NamespacedPodTemplateUnauthorized struct {
}

func (o *DeleteCoreV1NamespacedPodTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/namespaces/{namespace}/podtemplates/{name}][%d] deleteCoreV1NamespacedPodTemplateUnauthorized ", 401)
}

func (o *DeleteCoreV1NamespacedPodTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
