// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IoK8sAPICoreV1LimitRangeSpec LimitRangeSpec defines a min/max usage limit for resources that match on kind.
//
// swagger:model io.k8s.api.core.v1.LimitRangeSpec
type IoK8sAPICoreV1LimitRangeSpec struct {

	// Limits is the list of LimitRangeItem objects that are enforced.
	// Required: true
	Limits []*IoK8sAPICoreV1LimitRangeItem `json:"limits"`
}

// Validate validates this io k8s api core v1 limit range spec
func (m *IoK8sAPICoreV1LimitRangeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1LimitRangeSpec) validateLimits(formats strfmt.Registry) error {

	if err := validate.Required("limits", "body", m.Limits); err != nil {
		return err
	}

	for i := 0; i < len(m.Limits); i++ {
		if swag.IsZero(m.Limits[i]) { // not required
			continue
		}

		if m.Limits[i] != nil {
			if err := m.Limits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("limits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1LimitRangeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1LimitRangeSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1LimitRangeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
