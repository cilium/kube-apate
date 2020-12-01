// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// PatchPolicyV1beta1NamespacedPodDisruptionBudgetReader is a Reader for the PatchPolicyV1beta1NamespacedPodDisruptionBudget structure.
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetOK creates a PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK with default headers values
func NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetOK() *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK {
	return &PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK{}
}

/*PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK handles this case with default header values.

OK
*/
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK struct {
	Payload *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}][%d] patchPolicyV1beta1NamespacedPodDisruptionBudgetOK  %+v", 200, o.Payload)
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK) GetPayload() *models.IoK8sAPIPolicyV1beta1PodDisruptionBudget {
	return o.Payload
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1PodDisruptionBudget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized creates a PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized with default headers values
func NewPatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized() *PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized {
	return &PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized{}
}

/*PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized struct {
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/policy/v1beta1/namespaces/{namespace}/poddisruptionbudgets/{name}][%d] patchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized ", 401)
}

func (o *PatchPolicyV1beta1NamespacedPodDisruptionBudgetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
