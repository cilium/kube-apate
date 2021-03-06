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

// IoK8sApimachineryPkgApisMetaV1MicroTime MicroTime is version of Time with microsecond level precision.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime
type IoK8sApimachineryPkgApisMetaV1MicroTime strfmt.DateTime

// UnmarshalJSON sets a IoK8sApimachineryPkgApisMetaV1MicroTime value from JSON input
func (m *IoK8sApimachineryPkgApisMetaV1MicroTime) UnmarshalJSON(b []byte) error {
	return ((*strfmt.DateTime)(m)).UnmarshalJSON(b)
}

// MarshalJSON retrieves a IoK8sApimachineryPkgApisMetaV1MicroTime value as JSON output
func (m IoK8sApimachineryPkgApisMetaV1MicroTime) MarshalJSON() ([]byte, error) {
	return (strfmt.DateTime(m)).MarshalJSON()
}

// Validate validates this io k8s apimachinery pkg apis meta v1 micro time
func (m IoK8sApimachineryPkgApisMetaV1MicroTime) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.FormatOf("", "body", "date-time", strfmt.DateTime(m).String(), formats); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1MicroTime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1MicroTime) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1MicroTime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
