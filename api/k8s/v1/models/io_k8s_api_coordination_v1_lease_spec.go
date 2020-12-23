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

// IoK8sAPICoordinationV1LeaseSpec LeaseSpec is a specification of a Lease.
//
// swagger:model io.k8s.api.coordination.v1.LeaseSpec
type IoK8sAPICoordinationV1LeaseSpec struct {

	// acquireTime is a time when the current lease was acquired.
	// Format: date-time
	AcquireTime IoK8sApimachineryPkgApisMetaV1MicroTime `json:"acquireTime,omitempty"`

	// holderIdentity contains the identity of the holder of a current lease.
	HolderIdentity string `json:"holderIdentity,omitempty"`

	// leaseDurationSeconds is a duration that candidates for a lease need to wait to force acquire it. This is measure against time of last observed RenewTime.
	LeaseDurationSeconds int32 `json:"leaseDurationSeconds,omitempty"`

	// leaseTransitions is the number of transitions of a lease between holders.
	LeaseTransitions int32 `json:"leaseTransitions,omitempty"`

	// renewTime is a time when the current holder of a lease has last updated the lease.
	// Format: date-time
	RenewTime IoK8sApimachineryPkgApisMetaV1MicroTime `json:"renewTime,omitempty"`
}

// Validate validates this io k8s api coordination v1 lease spec
func (m *IoK8sAPICoordinationV1LeaseSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcquireTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoordinationV1LeaseSpec) validateAcquireTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AcquireTime) { // not required
		return nil
	}

	if err := m.AcquireTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("acquireTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICoordinationV1LeaseSpec) validateRenewTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RenewTime) { // not required
		return nil
	}

	if err := m.RenewTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("renewTime")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoordinationV1LeaseSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoordinationV1LeaseSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoordinationV1LeaseSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}