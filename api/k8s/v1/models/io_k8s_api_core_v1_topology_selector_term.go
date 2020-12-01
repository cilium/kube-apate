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

// IoK8sAPICoreV1TopologySelectorTerm A topology selector term represents the result of label queries. A null or empty topology selector term matches no objects. The requirements of them are ANDed. It provides a subset of functionality as NodeSelectorTerm. This is an alpha feature and may change in the future.
//
// swagger:model io.k8s.api.core.v1.TopologySelectorTerm
type IoK8sAPICoreV1TopologySelectorTerm struct {

	// A list of topology selector requirements by labels.
	MatchLabelExpressions []*IoK8sAPICoreV1TopologySelectorLabelRequirement `json:"matchLabelExpressions"`
}

// Validate validates this io k8s api core v1 topology selector term
func (m *IoK8sAPICoreV1TopologySelectorTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatchLabelExpressions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1TopologySelectorTerm) validateMatchLabelExpressions(formats strfmt.Registry) error {

	if swag.IsZero(m.MatchLabelExpressions) { // not required
		return nil
	}

	for i := 0; i < len(m.MatchLabelExpressions); i++ {
		if swag.IsZero(m.MatchLabelExpressions[i]) { // not required
			continue
		}

		if m.MatchLabelExpressions[i] != nil {
			if err := m.MatchLabelExpressions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matchLabelExpressions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1TopologySelectorTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1TopologySelectorTerm) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1TopologySelectorTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
