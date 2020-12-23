// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package authentication_v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/kube-apate/api/k8s/v1/models"
)

// CreateAuthenticationV1beta1TokenReviewReader is a Reader for the CreateAuthenticationV1beta1TokenReview structure.
type CreateAuthenticationV1beta1TokenReviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthenticationV1beta1TokenReviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAuthenticationV1beta1TokenReviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateAuthenticationV1beta1TokenReviewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateAuthenticationV1beta1TokenReviewAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateAuthenticationV1beta1TokenReviewUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAuthenticationV1beta1TokenReviewOK creates a CreateAuthenticationV1beta1TokenReviewOK with default headers values
func NewCreateAuthenticationV1beta1TokenReviewOK() *CreateAuthenticationV1beta1TokenReviewOK {
	return &CreateAuthenticationV1beta1TokenReviewOK{}
}

/*CreateAuthenticationV1beta1TokenReviewOK handles this case with default header values.

OK
*/
type CreateAuthenticationV1beta1TokenReviewOK struct {
	Payload *models.IoK8sAPIAuthenticationV1beta1TokenReview
}

func (o *CreateAuthenticationV1beta1TokenReviewOK) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1beta1/tokenreviews][%d] createAuthenticationV1beta1TokenReviewOK  %+v", 200, o.Payload)
}

func (o *CreateAuthenticationV1beta1TokenReviewOK) GetPayload() *models.IoK8sAPIAuthenticationV1beta1TokenReview {
	return o.Payload
}

func (o *CreateAuthenticationV1beta1TokenReviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthenticationV1beta1TokenReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticationV1beta1TokenReviewCreated creates a CreateAuthenticationV1beta1TokenReviewCreated with default headers values
func NewCreateAuthenticationV1beta1TokenReviewCreated() *CreateAuthenticationV1beta1TokenReviewCreated {
	return &CreateAuthenticationV1beta1TokenReviewCreated{}
}

/*CreateAuthenticationV1beta1TokenReviewCreated handles this case with default header values.

Created
*/
type CreateAuthenticationV1beta1TokenReviewCreated struct {
	Payload *models.IoK8sAPIAuthenticationV1beta1TokenReview
}

func (o *CreateAuthenticationV1beta1TokenReviewCreated) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1beta1/tokenreviews][%d] createAuthenticationV1beta1TokenReviewCreated  %+v", 201, o.Payload)
}

func (o *CreateAuthenticationV1beta1TokenReviewCreated) GetPayload() *models.IoK8sAPIAuthenticationV1beta1TokenReview {
	return o.Payload
}

func (o *CreateAuthenticationV1beta1TokenReviewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthenticationV1beta1TokenReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticationV1beta1TokenReviewAccepted creates a CreateAuthenticationV1beta1TokenReviewAccepted with default headers values
func NewCreateAuthenticationV1beta1TokenReviewAccepted() *CreateAuthenticationV1beta1TokenReviewAccepted {
	return &CreateAuthenticationV1beta1TokenReviewAccepted{}
}

/*CreateAuthenticationV1beta1TokenReviewAccepted handles this case with default header values.

Accepted
*/
type CreateAuthenticationV1beta1TokenReviewAccepted struct {
	Payload *models.IoK8sAPIAuthenticationV1beta1TokenReview
}

func (o *CreateAuthenticationV1beta1TokenReviewAccepted) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1beta1/tokenreviews][%d] createAuthenticationV1beta1TokenReviewAccepted  %+v", 202, o.Payload)
}

func (o *CreateAuthenticationV1beta1TokenReviewAccepted) GetPayload() *models.IoK8sAPIAuthenticationV1beta1TokenReview {
	return o.Payload
}

func (o *CreateAuthenticationV1beta1TokenReviewAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IoK8sAPIAuthenticationV1beta1TokenReview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthenticationV1beta1TokenReviewUnauthorized creates a CreateAuthenticationV1beta1TokenReviewUnauthorized with default headers values
func NewCreateAuthenticationV1beta1TokenReviewUnauthorized() *CreateAuthenticationV1beta1TokenReviewUnauthorized {
	return &CreateAuthenticationV1beta1TokenReviewUnauthorized{}
}

/*CreateAuthenticationV1beta1TokenReviewUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAuthenticationV1beta1TokenReviewUnauthorized struct {
}

func (o *CreateAuthenticationV1beta1TokenReviewUnauthorized) Error() string {
	return fmt.Sprintf("[POST /apis/authentication.k8s.io/v1beta1/tokenreviews][%d] createAuthenticationV1beta1TokenReviewUnauthorized ", 401)
}

func (o *CreateAuthenticationV1beta1TokenReviewUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}