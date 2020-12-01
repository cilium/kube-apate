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

// IoK8sAPICoreV1ContainerState ContainerState holds a possible state of container. Only one of its members may be specified. If none of them is specified, the default one is ContainerStateWaiting.
//
// swagger:model io.k8s.api.core.v1.ContainerState
type IoK8sAPICoreV1ContainerState struct {

	// Details about a running container
	Running *IoK8sAPICoreV1ContainerStateRunning `json:"running,omitempty"`

	// Details about a terminated container
	Terminated *IoK8sAPICoreV1ContainerStateTerminated `json:"terminated,omitempty"`

	// Details about a waiting container
	Waiting *IoK8sAPICoreV1ContainerStateWaiting `json:"waiting,omitempty"`
}

// Validate validates this io k8s api core v1 container state
func (m *IoK8sAPICoreV1ContainerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunning(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTerminated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWaiting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ContainerState) validateRunning(formats strfmt.Registry) error {

	if swag.IsZero(m.Running) { // not required
		return nil
	}

	if m.Running != nil {
		if err := m.Running.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("running")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerState) validateTerminated(formats strfmt.Registry) error {

	if swag.IsZero(m.Terminated) { // not required
		return nil
	}

	if m.Terminated != nil {
		if err := m.Terminated.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("terminated")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1ContainerState) validateWaiting(formats strfmt.Registry) error {

	if swag.IsZero(m.Waiting) { // not required
		return nil
	}

	if m.Waiting != nil {
		if err := m.Waiting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("waiting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ContainerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ContainerState) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ContainerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
