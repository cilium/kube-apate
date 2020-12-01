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

// IoK8sAPICoreV1AttachedVolume AttachedVolume describes a volume attached to a node
//
// swagger:model io.k8s.api.core.v1.AttachedVolume
type IoK8sAPICoreV1AttachedVolume struct {

	// DevicePath represents the device path where the volume should be available
	// Required: true
	DevicePath *string `json:"devicePath"`

	// Name of the attached volume
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this io k8s api core v1 attached volume
func (m *IoK8sAPICoreV1AttachedVolume) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevicePath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1AttachedVolume) validateDevicePath(formats strfmt.Registry) error {

	if err := validate.Required("devicePath", "body", m.DevicePath); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1AttachedVolume) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1AttachedVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1AttachedVolume) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1AttachedVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
