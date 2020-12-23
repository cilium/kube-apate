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

// IoK8sAPIAppsV1StatefulSetUpdateStrategy StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
//
// swagger:model io.k8s.api.apps.v1.StatefulSetUpdateStrategy
type IoK8sAPIAppsV1StatefulSetUpdateStrategy struct {

	// RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
	RollingUpdate *IoK8sAPIAppsV1RollingUpdateStatefulSetStrategy `json:"rollingUpdate,omitempty"`

	// Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
	Type string `json:"type,omitempty"`
}

// Validate validates this io k8s api apps v1 stateful set update strategy
func (m *IoK8sAPIAppsV1StatefulSetUpdateStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRollingUpdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAppsV1StatefulSetUpdateStrategy) validateRollingUpdate(formats strfmt.Registry) error {

	if swag.IsZero(m.RollingUpdate) { // not required
		return nil
	}

	if m.RollingUpdate != nil {
		if err := m.RollingUpdate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rollingUpdate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAppsV1StatefulSetUpdateStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAppsV1StatefulSetUpdateStrategy) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAppsV1StatefulSetUpdateStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}