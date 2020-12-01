// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAuthorizationV1beta1SelfSubjectAccessReviewReader is a Reader for the CreateAuthorizationV1beta1SelfSubjectAccessReview structure.
type CreateAuthorizationV1beta1SelfSubjectAccessReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAuthorizationV1beta1SelfSubjectAccessReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAuthorizationV1beta1SelfSubjectAccessReviewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAuthorizationV1beta1SelfSubjectAccessReviewOK creates a CreateAuthorizationV1beta1SelfSubjectAccessReviewOK with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectAccessReviewOK() *CreateAuthorizationV1beta1SelfSubjectAccessReviewOK {
	return &CreateAuthorizationV1beta1SelfSubjectAccessReviewOK{}
}

/*CreateAuthorizationV1beta1SelfSubjectAccessReviewOK handles this case with default header values.

OK
*/
type CreateAuthorizationV1beta1SelfSubjectAccessReviewOK struct {
	Payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewOK) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews][%d] createAuthorizationV1beta1SelfSubjectAccessReviewOK  %+v", 200, o.Payload)
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewOK) GetPayload() *models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview {
	return o.Payload
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1beta1SelfSubjectAccessReviewCreated creates a CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectAccessReviewCreated() *CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated {
	return &CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated{}
}

/*CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated handles this case with default header values.

Created
*/
type CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated struct {
	Payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews][%d] createAuthorizationV1beta1SelfSubjectAccessReviewCreated  %+v", 201, o.Payload)
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated) GetPayload() *models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview {
	return o.Payload
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted creates a CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted() *CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted {
	return &CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted{}
}

/*CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted handles this case with default header values.

Accepted
*/
type CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted struct {
	Payload *models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews][%d] createAuthorizationV1beta1SelfSubjectAccessReviewAccepted  %+v", 202, o.Payload)
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted) GetPayload() *models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview {
	return o.Payload
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1beta1SelfSubjectAccessReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized creates a CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized with default headers values
func NewCreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized() *CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized {
	return &CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized{}
}

/*CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized struct {
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews][%d] createAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized ", 401)
}

func (o *CreateAuthorizationV1beta1SelfSubjectAccessReviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
