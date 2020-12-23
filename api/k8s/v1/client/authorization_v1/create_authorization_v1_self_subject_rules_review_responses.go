// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authorization_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAuthorizationV1SelfSubjectRulesReviewReader is a Reader for the CreateAuthorizationV1SelfSubjectRulesReview structure.
type CreateAuthorizationV1SelfSubjectRulesReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthorizationV1SelfSubjectRulesReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAuthorizationV1SelfSubjectRulesReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAuthorizationV1SelfSubjectRulesReviewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAuthorizationV1SelfSubjectRulesReviewAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAuthorizationV1SelfSubjectRulesReviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAuthorizationV1SelfSubjectRulesReviewOK creates a CreateAuthorizationV1SelfSubjectRulesReviewOK with default headers values
func NewCreateAuthorizationV1SelfSubjectRulesReviewOK() *CreateAuthorizationV1SelfSubjectRulesReviewOK {
	return &CreateAuthorizationV1SelfSubjectRulesReviewOK{}
}

/*CreateAuthorizationV1SelfSubjectRulesReviewOK handles this case with default header values.

OK
*/
type CreateAuthorizationV1SelfSubjectRulesReviewOK struct {
	Payload *models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewOK) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1/selfsubjectrulesreviews][%d] createAuthorizationV1SelfSubjectRulesReviewOK  %+v", 200, o.Payload)
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewOK) GetPayload() *models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview {
	return o.Payload
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1SelfSubjectRulesReviewCreated creates a CreateAuthorizationV1SelfSubjectRulesReviewCreated with default headers values
func NewCreateAuthorizationV1SelfSubjectRulesReviewCreated() *CreateAuthorizationV1SelfSubjectRulesReviewCreated {
	return &CreateAuthorizationV1SelfSubjectRulesReviewCreated{}
}

/*CreateAuthorizationV1SelfSubjectRulesReviewCreated handles this case with default header values.

Created
*/
type CreateAuthorizationV1SelfSubjectRulesReviewCreated struct {
	Payload *models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewCreated) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1/selfsubjectrulesreviews][%d] createAuthorizationV1SelfSubjectRulesReviewCreated  %+v", 201, o.Payload)
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewCreated) GetPayload() *models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview {
	return o.Payload
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1SelfSubjectRulesReviewAccepted creates a CreateAuthorizationV1SelfSubjectRulesReviewAccepted with default headers values
func NewCreateAuthorizationV1SelfSubjectRulesReviewAccepted() *CreateAuthorizationV1SelfSubjectRulesReviewAccepted {
	return &CreateAuthorizationV1SelfSubjectRulesReviewAccepted{}
}

/*CreateAuthorizationV1SelfSubjectRulesReviewAccepted handles this case with default header values.

Accepted
*/
type CreateAuthorizationV1SelfSubjectRulesReviewAccepted struct {
	Payload *models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1/selfsubjectrulesreviews][%d] createAuthorizationV1SelfSubjectRulesReviewAccepted  %+v", 202, o.Payload)
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewAccepted) GetPayload() *models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview {
	return o.Payload
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthorizationV1SelfSubjectRulesReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthorizationV1SelfSubjectRulesReviewUnauthorized creates a CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized with default headers values
func NewCreateAuthorizationV1SelfSubjectRulesReviewUnauthorized() *CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized {
	return &CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized{}
}

/*CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized struct {
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/authorization.k8s.io/v1/selfsubjectrulesreviews][%d] createAuthorizationV1SelfSubjectRulesReviewUnauthorized ", 401)
}

func (o *CreateAuthorizationV1SelfSubjectRulesReviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}