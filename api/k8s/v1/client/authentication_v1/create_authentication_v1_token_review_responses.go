// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAuthenticationV1TokenReviewReader is a Reader for the CreateAuthenticationV1TokenReview structure.
type CreateAuthenticationV1TokenReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthenticationV1TokenReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAuthenticationV1TokenReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAuthenticationV1TokenReviewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAuthenticationV1TokenReviewAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAuthenticationV1TokenReviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAuthenticationV1TokenReviewOK creates a CreateAuthenticationV1TokenReviewOK with default headers values
func NewCreateAuthenticationV1TokenReviewOK() *CreateAuthenticationV1TokenReviewOK {
	return &CreateAuthenticationV1TokenReviewOK{}
}

/*CreateAuthenticationV1TokenReviewOK handles this case with default header values.

OK
*/
type CreateAuthenticationV1TokenReviewOK struct {
	Payload *models.IoK8sAPIAuthenticationV1TokenReview
}

func (o *CreateAuthenticationV1TokenReviewOK) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1/tokenreviews][%d] createAuthenticationV1TokenReviewOK  %+v", 200, o.Payload)
}

func (o *CreateAuthenticationV1TokenReviewOK) GetPayload() *models.IoK8sAPIAuthenticationV1TokenReview {
	return o.Payload
}

func (o *CreateAuthenticationV1TokenReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthenticationV1TokenReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticationV1TokenReviewCreated creates a CreateAuthenticationV1TokenReviewCreated with default headers values
func NewCreateAuthenticationV1TokenReviewCreated() *CreateAuthenticationV1TokenReviewCreated {
	return &CreateAuthenticationV1TokenReviewCreated{}
}

/*CreateAuthenticationV1TokenReviewCreated handles this case with default header values.

Created
*/
type CreateAuthenticationV1TokenReviewCreated struct {
	Payload *models.IoK8sAPIAuthenticationV1TokenReview
}

func (o *CreateAuthenticationV1TokenReviewCreated) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1/tokenreviews][%d] createAuthenticationV1TokenReviewCreated  %+v", 201, o.Payload)
}

func (o *CreateAuthenticationV1TokenReviewCreated) GetPayload() *models.IoK8sAPIAuthenticationV1TokenReview {
	return o.Payload
}

func (o *CreateAuthenticationV1TokenReviewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthenticationV1TokenReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticationV1TokenReviewAccepted creates a CreateAuthenticationV1TokenReviewAccepted with default headers values
func NewCreateAuthenticationV1TokenReviewAccepted() *CreateAuthenticationV1TokenReviewAccepted {
	return &CreateAuthenticationV1TokenReviewAccepted{}
}

/*CreateAuthenticationV1TokenReviewAccepted handles this case with default header values.

Accepted
*/
type CreateAuthenticationV1TokenReviewAccepted struct {
	Payload *models.IoK8sAPIAuthenticationV1TokenReview
}

func (o *CreateAuthenticationV1TokenReviewAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1/tokenreviews][%d] createAuthenticationV1TokenReviewAccepted  %+v", 202, o.Payload)
}

func (o *CreateAuthenticationV1TokenReviewAccepted) GetPayload() *models.IoK8sAPIAuthenticationV1TokenReview {
	return o.Payload
}

func (o *CreateAuthenticationV1TokenReviewAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthenticationV1TokenReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticationV1TokenReviewUnauthorized creates a CreateAuthenticationV1TokenReviewUnauthorized with default headers values
func NewCreateAuthenticationV1TokenReviewUnauthorized() *CreateAuthenticationV1TokenReviewUnauthorized {
	return &CreateAuthenticationV1TokenReviewUnauthorized{}
}

/*CreateAuthenticationV1TokenReviewUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAuthenticationV1TokenReviewUnauthorized struct {
}

func (o *CreateAuthenticationV1TokenReviewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1/tokenreviews][%d] createAuthenticationV1TokenReviewUnauthorized ", 401)
}

func (o *CreateAuthenticationV1TokenReviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
