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

// IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition io k8s api certificates v1beta1 certificate signing request condition
//
// swagger:model io.k8s.api.certificates.v1beta1.CertificateSigningRequestCondition
type IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition struct {

	// lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// timestamp for the last update to this condition
	// Format: date-time
	LastUpdateTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastUpdateTime,omitempty"`

	// human readable message with details about the request state
	Message string `json:"message,omitempty"`

	// brief reason for the request state
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown". Defaults to "True". If unset, should be treated as "True".
	Status string `json:"status,omitempty"`

	// type of the condition. Known conditions include "Approved", "Denied", and "Failed".
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api certificates v1beta1 certificate signing request condition
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) validateLastTransitionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := m.LastTransitionTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastTransitionTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) validateLastUpdateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdateTime) { // not required
		return nil
	}

	if err := m.LastUpdateTime.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lastUpdateTime")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
