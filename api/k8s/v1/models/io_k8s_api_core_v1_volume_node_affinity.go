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

// IoK8sAPICoreV1VolumeNodeAffinity VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
//
// swagger:model io.k8s.api.core.v1.VolumeNodeAffinity
type IoK8sAPICoreV1VolumeNodeAffinity struct {

	// Required specifies hard node constraints that must be met.
	Required *IoK8sAPICoreV1NodeSelector `json:"required,omitempty"`
}

// Validate validates this io k8s api core v1 volume node affinity
func (m *IoK8sAPICoreV1VolumeNodeAffinity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequired(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1VolumeNodeAffinity) validateRequired(formats strfmt.Registry) error {

	if swag.IsZero(m.Required) { // not required
		return nil
	}

	if m.Required != nil {
		if err := m.Required.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("required")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1VolumeNodeAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1VolumeNodeAffinity) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1VolumeNodeAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
