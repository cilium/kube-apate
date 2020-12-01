// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAuthenticationV1beta1TokenReviewStatus TokenReviewStatus is the result of the token authentication request.
//
// swagger:model io.k8s.api.authentication.v1beta1.TokenReviewStatus
type IoK8sAPIAuthenticationV1beta1TokenReviewStatus struct {

	// Audiences are audience identifiers chosen by the authenticator that are compatible with both the TokenReview and token. An identifier is any identifier in the intersection of the TokenReviewSpec audiences and the token's audiences. A client of the TokenReview API that sets the spec.audiences field should validate that a compatible audience identifier is returned in the status.audiences field to ensure that the TokenReview server is audience aware. If a TokenReview returns an empty status.audience field where status.authenticated is "true", the token is valid against the audience of the Kubernetes API server.
	Audiences []string `json:"audiences"`

	// Authenticated indicates that the token was associated with a known user.
	Authenticated bool `json:"authenticated,omitempty"`

	// Error indicates that the token couldn't be checked
	Error string `json:"error,omitempty"`

	// User is the UserInfo associated with the provided token.
	User *IoK8sAPIAuthenticationV1beta1UserInfo `json:"user,omitempty"`
}

// Validate validates this io k8s api authentication v1beta1 token review status
func (m *IoK8sAPIAuthenticationV1beta1TokenReviewStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAuthenticationV1beta1TokenReviewStatus) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1beta1TokenReviewStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1beta1TokenReviewStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthenticationV1beta1TokenReviewStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
