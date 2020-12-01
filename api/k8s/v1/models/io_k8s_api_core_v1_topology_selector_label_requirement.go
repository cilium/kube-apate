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

// IoK8sAPICoreV1TopologySelectorLabelRequirement A topology selector requirement is a selector that matches given label. This is an alpha feature and may change in the future.
//
// swagger:model io.k8s.api.core.v1.TopologySelectorLabelRequirement
type IoK8sAPICoreV1TopologySelectorLabelRequirement struct {

	// The label key that the selector applies to.
	// Required: true
	Key *string `json:"key"`

	// An array of string values. One value must match the label to be selected. Each entry in Values is ORed.
	// Required: true
	Values []string `json:"values"`
}

// Validate validates this io k8s api core v1 topology selector label requirement
func (m *IoK8sAPICoreV1TopologySelectorLabelRequirement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1TopologySelectorLabelRequirement) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1TopologySelectorLabelRequirement) validateValues(formats strfmt.Registry) error {

	if err := validate.Required("values", "body", m.Values); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1TopologySelectorLabelRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1TopologySelectorLabelRequirement) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1TopologySelectorLabelRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
