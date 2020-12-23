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

// PatchPolicyV1beta1PodSecurityPolicyReader is a Reader for the PatchPolicyV1beta1PodSecurityPolicy structure.
type PatchPolicyV1beta1PodSecurityPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPolicyV1beta1PodSecurityPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchPolicyV1beta1PodSecurityPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchPolicyV1beta1PodSecurityPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchPolicyV1beta1PodSecurityPolicyOK creates a PatchPolicyV1beta1PodSecurityPolicyOK with default headers values
func NewPatchPolicyV1beta1PodSecurityPolicyOK() *PatchPolicyV1beta1PodSecurityPolicyOK {
	return &PatchPolicyV1beta1PodSecurityPolicyOK{}
}

/*PatchPolicyV1beta1PodSecurityPolicyOK handles this case with default header values.

OK
*/
type PatchPolicyV1beta1PodSecurityPolicyOK struct {
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy
}

func (o *PatchPolicyV1beta1PodSecurityPolicyOK) Error() string {
	return fmt.Sprintf("[PATCH /apis/policy/v1beta1/podsecuritypolicies/{name}][%d] patchPolicyV1beta1PodSecurityPolicyOK  %+v", 200, o.Payload)
}

func (o *PatchPolicyV1beta1PodSecurityPolicyOK) GetPayload() *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy {
	return o.Payload
}

func (o *PatchPolicyV1beta1PodSecurityPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1PodSecurityPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPolicyV1beta1PodSecurityPolicyUnauthorized creates a PatchPolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewPatchPolicyV1beta1PodSecurityPolicyUnauthorized() *PatchPolicyV1beta1PodSecurityPolicyUnauthorized {
	return &PatchPolicyV1beta1PodSecurityPolicyUnauthorized{}
}

/*PatchPolicyV1beta1PodSecurityPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchPolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

func (o *PatchPolicyV1beta1PodSecurityPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /apis/policy/v1beta1/podsecuritypolicies/{name}][%d] patchPolicyV1beta1PodSecurityPolicyUnauthorized ", 401)
}

func (o *PatchPolicyV1beta1PodSecurityPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}