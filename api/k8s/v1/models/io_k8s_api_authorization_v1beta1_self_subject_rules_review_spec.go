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

// IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReviewSpec io k8s api authorization v1beta1 self subject rules review spec
//
// swagger:model io.k8s.api.authorization.v1beta1.SelfSubjectRulesReviewSpec
type IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReviewSpec struct {

	// Namespace to evaluate rules for. Required.
	Namespace string `json:"namespace,omitempty"`
}

// Validate validates this io k8s api authorization v1beta1 self subject rules review spec
func (m *IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReviewSpec) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReviewSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReviewSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthorizationV1beta1SelfSubjectRulesReviewSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
