// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package batch_v2alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchBatchV2alpha1NamespacedCronJobReader is a Reader for the PatchBatchV2alpha1NamespacedCronJob structure.
type PatchBatchV2alpha1NamespacedCronJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBatchV2alpha1NamespacedCronJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchBatchV2alpha1NamespacedCronJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchBatchV2alpha1NamespacedCronJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchBatchV2alpha1NamespacedCronJobOK creates a PatchBatchV2alpha1NamespacedCronJobOK with default headers values
func NewPatchBatchV2alpha1NamespacedCronJobOK() *PatchBatchV2alpha1NamespacedCronJobOK {
	return &PatchBatchV2alpha1NamespacedCronJobOK{}
}

/*PatchBatchV2alpha1NamespacedCronJobOK handles this case with default header values.

OK
*/
type PatchBatchV2alpha1NamespacedCronJobOK struct {
	Payload *models.IoK8sAPIBatchV2alpha1CronJob
}

func (o *PatchBatchV2alpha1NamespacedCronJobOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}][%d] patchBatchV2alpha1NamespacedCronJobOK  %+v", 200, o.Payload)
}

func (o *PatchBatchV2alpha1NamespacedCronJobOK) GetPayload() *models.IoK8sAPIBatchV2alpha1CronJob {
	return o.Payload
}

func (o *PatchBatchV2alpha1NamespacedCronJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIBatchV2alpha1CronJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBatchV2alpha1NamespacedCronJobUnauthorized creates a PatchBatchV2alpha1NamespacedCronJobUnauthorized with default headers values
func NewPatchBatchV2alpha1NamespacedCronJobUnauthorized() *PatchBatchV2alpha1NamespacedCronJobUnauthorized {
	return &PatchBatchV2alpha1NamespacedCronJobUnauthorized{}
}

/*PatchBatchV2alpha1NamespacedCronJobUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchBatchV2alpha1NamespacedCronJobUnauthorized struct {
}

func (o *PatchBatchV2alpha1NamespacedCronJobUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/batch/v2alpha1/namespaces/{namespace}/cronjobs/{name}][%d] patchBatchV2alpha1NamespacedCronJobUnauthorized ", 401)
}

func (o *PatchBatchV2alpha1NamespacedCronJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
