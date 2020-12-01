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
)

// IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus io k8s api certificates v1beta1 certificate signing request status
//
// swagger:model io.k8s.api.certificates.v1beta1.CertificateSigningRequestStatus
type IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus struct {

	// If request was approved, the controller will place the issued certificate here.
	// Format: byte
	Certificate strfmt.Base64 `json:"certificate,omitempty"`

	// Conditions applied to the request, such as approval or denial.
	Conditions []*IoK8sAPICertificatesV1beta1CertificateSigningRequestCondition `json:"conditions"`
}

// Validate validates this io k8s api certificates v1beta1 certificate signing request status
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICertificatesV1beta1CertificateSigningRequestStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
