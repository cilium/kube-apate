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

// IoK8sAPICoreV1ProjectedVolumeSource Represents a projected volume source
//
// swagger:model io.k8s.api.core.v1.ProjectedVolumeSource
type IoK8sAPICoreV1ProjectedVolumeSource struct {

	// Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode int32 `json:"defaultMode,omitempty"`

	// list of volume projections
	// Required: true
	Sources []*IoK8sAPICoreV1VolumeProjection `json:"sources"`
}

// Validate validates this io k8s api core v1 projected volume source
func (m *IoK8sAPICoreV1ProjectedVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ProjectedVolumeSource) validateSources(formats strfmt.Registry) error {

	if err := validate.Required("sources", "body", m.Sources); err != nil {
		return err
	}

	for i := 0; i < len(m.Sources); i++ {
		if swag.IsZero(m.Sources[i]) { // not required
			continue
		}

		if m.Sources[i] != nil {
			if err := m.Sources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ProjectedVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ProjectedVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ProjectedVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
