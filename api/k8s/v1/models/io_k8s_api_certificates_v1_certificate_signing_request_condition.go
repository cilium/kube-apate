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

// IoK8sAPICertificatesV1CertificateSigningRequestCondition CertificateSigningRequestCondition describes a condition of a CertificateSigningRequest object
//
// swagger:model io.k8s.api.certificates.v1.CertificateSigningRequestCondition
type IoK8sAPICertificatesV1CertificateSigningRequestCondition struct {

	// lastTransitionTime is the time the condition last transitioned from one status to another. If unset, when a new condition type is added or an existing condition's status is changed, the server defaults this to the current time.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// lastUpdateTime is the time of the last update to this condition
	// Format: date-time
	LastUpdateTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastUpdateTime,omitempty"`

	// message contains a human readable message with details about the request state
	Message string `json:"message,omitempty"`

	// reason indicates a brief reason for the request state
	Reason string `json:"reason,omitempty"`

	// status of the condition, one of True, False, Unknown. Approved, Denied, and Failed conditions may not be "False" or "Unknown".
	// Required: true
	Status *string `json:"status"`

	// type of the condition. Known conditions are "Approved", "Denied", and "Failed".
	//
	// An "Approved" condition is added via the /approval subresource, indicating the request was approved and should be issued by the signer.
	//
	// A "Denied" condition is added via the /approval subresource, indicating the request was denied and should not be issued by the signer.
	//
	// A "Failed" condition is added via the /status subresource, indicating the signer failed to issue the certificate.
	//
	// Approved and Denied conditions are mutually exclusive. Approved, Denied, and Failed conditions cannot be removed once added.
	//
	// Only one condition of a given type is allowed.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s api certificates v1 certificate signing request condition
func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) validateLastTransitionTime(formats strfmt.Registry) error {

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

func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) validateLastUpdateTime(formats strfmt.Registry) error {

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

func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1CertificateSigningRequestCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICertificatesV1CertificateSigningRequestCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
