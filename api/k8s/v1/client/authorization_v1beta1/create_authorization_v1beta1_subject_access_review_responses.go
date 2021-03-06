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

// CreateAuthorizationV1beta1SubjectAccessReviewReader is a Reader for the CreateAuthorizationV1beta1SubjectAccessReview structure.
type CreateAuthorizationV1beta1SubjectAccessReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthorizationV1beta1SubjectAccessReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAuthorizationV1beta1SubjectAccessReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAuthorizationV1beta1SubjectAccessReviewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAuthorizationV1beta1SubjectAccessReviewAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAuthorizationV1beta1SubjectAccessReviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAuthorizationV1beta1SubjectAccessReviewOK creates a CreateAuthorizationV1beta1SubjectAccessReviewOK with default headers values
func NewCreateAuthorizationV1beta1SubjectAccessReviewOK() *CreateAuthorizationV1beta1SubjectAccessReviewOK {
	return &CreateAuthorizationV1beta1SubjectAccessReviewOK{}
}

/*CreateAuthorizationV1beta1SubjectAccessReviewOK handles this case with default header values.

OK
*/
type CreateAuthorizationV1beta1SubjectAccessReviewOK struct {
	Payload *models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewOK) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/subjectaccessreviews][%d] createAuthorizationV1beta1SubjectAccessReviewOK  %+v", 200, o.Payload)
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewOK) GetPayload() *models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview {
	return o.Payload
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1beta1SubjectAccessReviewCreated creates a CreateAuthorizationV1beta1SubjectAccessReviewCreated with default headers values
func NewCreateAuthorizationV1beta1SubjectAccessReviewCreated() *CreateAuthorizationV1beta1SubjectAccessReviewCreated {
	return &CreateAuthorizationV1beta1SubjectAccessReviewCreated{}
}

/*CreateAuthorizationV1beta1SubjectAccessReviewCreated handles this case with default header values.

Created
*/
type CreateAuthorizationV1beta1SubjectAccessReviewCreated struct {
	Payload *models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewCreated) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/subjectaccessreviews][%d] createAuthorizationV1beta1SubjectAccessReviewCreated  %+v", 201, o.Payload)
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewCreated) GetPayload() *models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview {
	return o.Payload
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1beta1SubjectAccessReviewAccepted creates a CreateAuthorizationV1beta1SubjectAccessReviewAccepted with default headers values
func NewCreateAuthorizationV1beta1SubjectAccessReviewAccepted() *CreateAuthorizationV1beta1SubjectAccessReviewAccepted {
	return &CreateAuthorizationV1beta1SubjectAccessReviewAccepted{}
}

/*CreateAuthorizationV1beta1SubjectAccessReviewAccepted handles this case with default header values.

Accepted
*/
type CreateAuthorizationV1beta1SubjectAccessReviewAccepted struct {
	Payload *models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/subjectaccessreviews][%d] createAuthorizationV1beta1SubjectAccessReviewAccepted  %+v", 202, o.Payload)
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewAccepted) GetPayload() *models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview {
	return o.Payload
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1beta1SubjectAccessReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1beta1SubjectAccessReviewUnauthorized creates a CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized with default headers values
func NewCreateAuthorizationV1beta1SubjectAccessReviewUnauthorized() *CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized {
	return &CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized{}
}

/*CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized struct {
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1beta1/subjectaccessreviews][%d] createAuthorizationV1beta1SubjectAccessReviewUnauthorized ", 401)
}

func (o *CreateAuthorizationV1beta1SubjectAccessReviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
