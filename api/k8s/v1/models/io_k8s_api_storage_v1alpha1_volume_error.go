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

// IoK8sAPIStorageV1alpha1VolumeError VolumeError captures an error encountered during a volume operation.
//
// swagger:model io.k8s.api.storage.v1alpha1.VolumeError
type IoK8sAPIStorageV1alpha1VolumeError struct {

	// String detailing the error encountered during Attach or Detach operation. This string maybe logged, so it should not contain sensitive information.
	Message string `json:"message,omitempty"`

	// Time the error was encountered.
	// Format: date-time
	Time IoK8sApimachineryPkgApisMetaV1Time `json:"time,omitempty"`
}

// Validate validates this io k8s api storage v1alpha1 volume error
func (m *IoK8sAPIStorageV1alpha1VolumeError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1alpha1VolumeError) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1alpha1VolumeError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1alpha1VolumeError) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1alpha1VolumeError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}