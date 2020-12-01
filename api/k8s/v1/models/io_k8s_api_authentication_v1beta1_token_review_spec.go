// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIAuthenticationV1beta1TokenReviewSpec TokenReviewSpec is a description of the token authentication request.
//
// swagger:model io.k8s.api.authentication.v1beta1.TokenReviewSpec
type IoK8sAPIAuthenticationV1beta1TokenReviewSpec struct {

	// Audiences is a list of the identifiers that the resource server presented with the token identifies as. Audience-aware token authenticators will verify that the token was intended for at least one of the audiences in this list. If no audiences are provided, the audience will default to the audience of the Kubernetes apiserver.
	Audiences []string `json:"audiences"`

	// Token is the opaque bearer token.
	Token string `json:"token,omitempty"`
}

// Validate validates this io k8s api authentication v1beta1 token review spec
func (m *IoK8sAPIAuthenticationV1beta1TokenReviewSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1beta1TokenReviewSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthenticationV1beta1TokenReviewSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthenticationV1beta1TokenReviewSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
