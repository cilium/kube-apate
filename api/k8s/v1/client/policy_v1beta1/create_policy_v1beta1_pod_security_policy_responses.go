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

// CreatePolicyV1beta1PodSecurityPolicyReader is a Reader for the CreatePolicyV1beta1PodSecurityPolicy structure.
type CreatePolicyV1beta1PodSecurityPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyV1beta1PodSecurityPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePolicyV1beta1PodSecurityPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreatePolicyV1beta1PodSecurityPolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreatePolicyV1beta1PodSecurityPolicyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreatePolicyV1beta1PodSecurityPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePolicyV1beta1PodSecurityPolicyOK creates a CreatePolicyV1beta1PodSecurityPolicyOK with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyOK() *CreatePolicyV1beta1PodSecurityPolicyOK {
	return &CreatePolicyV1beta1PodSecurityPolicyOK{}
}

/*CreatePolicyV1beta1PodSecurityPolicyOK handles this case with default header values.

OK
*/
type CreatePolicyV1beta1PodSecurityPolicyOK struct {
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy
}

func (o *CreatePolicyV1beta1PodSecurityPolicyOK) Error() string {
	return fmt.Sprintf("[POST /apis/policy/v1beta1/podsecuritypolicies][%d] createPolicyV1beta1PodSecurityPolicyOK  %+v", 200, o.Payload)
}

func (o *CreatePolicyV1beta1PodSecurityPolicyOK) GetPayload() *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy {
	return o.Payload
}

func (o *CreatePolicyV1beta1PodSecurityPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1PodSecurityPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyV1beta1PodSecurityPolicyCreated creates a CreatePolicyV1beta1PodSecurityPolicyCreated with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyCreated() *CreatePolicyV1beta1PodSecurityPolicyCreated {
	return &CreatePolicyV1beta1PodSecurityPolicyCreated{}
}

/*CreatePolicyV1beta1PodSecurityPolicyCreated handles this case with default header values.

Created
*/
type CreatePolicyV1beta1PodSecurityPolicyCreated struct {
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy
}

func (o *CreatePolicyV1beta1PodSecurityPolicyCreated) Error() string {
	return fmt.Sprintf("[POST /apis/policy/v1beta1/podsecuritypolicies][%d] createPolicyV1beta1PodSecurityPolicyCreated  %+v", 201, o.Payload)
}

func (o *CreatePolicyV1beta1PodSecurityPolicyCreated) GetPayload() *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy {
	return o.Payload
}

func (o *CreatePolicyV1beta1PodSecurityPolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1PodSecurityPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyV1beta1PodSecurityPolicyAccepted creates a CreatePolicyV1beta1PodSecurityPolicyAccepted with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyAccepted() *CreatePolicyV1beta1PodSecurityPolicyAccepted {
	return &CreatePolicyV1beta1PodSecurityPolicyAccepted{}
}

/*CreatePolicyV1beta1PodSecurityPolicyAccepted handles this case with default header values.

Accepted
*/
type CreatePolicyV1beta1PodSecurityPolicyAccepted struct {
	Payload *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy
}

func (o *CreatePolicyV1beta1PodSecurityPolicyAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/policy/v1beta1/podsecuritypolicies][%d] createPolicyV1beta1PodSecurityPolicyAccepted  %+v", 202, o.Payload)
}

func (o *CreatePolicyV1beta1PodSecurityPolicyAccepted) GetPayload() *models.IoK8sAPIPolicyV1beta1PodSecurityPolicy {
	return o.Payload
}

func (o *CreatePolicyV1beta1PodSecurityPolicyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIPolicyV1beta1PodSecurityPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePolicyV1beta1PodSecurityPolicyUnauthorized creates a CreatePolicyV1beta1PodSecurityPolicyUnauthorized with default headers values
func NewCreatePolicyV1beta1PodSecurityPolicyUnauthorized() *CreatePolicyV1beta1PodSecurityPolicyUnauthorized {
	return &CreatePolicyV1beta1PodSecurityPolicyUnauthorized{}
}

/*CreatePolicyV1beta1PodSecurityPolicyUnauthorized handles this case with default header values.

Unauthorized
*/
type CreatePolicyV1beta1PodSecurityPolicyUnauthorized struct {
}

func (o *CreatePolicyV1beta1PodSecurityPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/policy/v1beta1/podsecuritypolicies][%d] createPolicyV1beta1PodSecurityPolicyUnauthorized ", 401)
}

func (o *CreatePolicyV1beta1PodSecurityPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
