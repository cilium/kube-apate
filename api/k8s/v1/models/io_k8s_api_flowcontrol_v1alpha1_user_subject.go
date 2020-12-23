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
	"github.com/go-openapi/validate"
)

// IoK8sAPIFlowcontrolV1alpha1UserSubject UserSubject holds detailed information for user-kind subject.
//
// swagger:model io.k8s.api.flowcontrol.v1alpha1.UserSubject
type IoK8sAPIFlowcontrolV1alpha1UserSubject struct {

	// `name` is the username that matches, or "*" to match all usernames. Required.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this io k8s api flowcontrol v1alpha1 user subject
func (m *IoK8sAPIFlowcontrolV1alpha1UserSubject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1UserSubject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1UserSubject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1UserSubject) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1alpha1UserSubject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}