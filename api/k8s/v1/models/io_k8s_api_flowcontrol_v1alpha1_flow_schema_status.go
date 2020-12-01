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

// IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus FlowSchemaStatus represents the current state of a FlowSchema.
//
// swagger:model io.k8s.api.flowcontrol.v1alpha1.FlowSchemaStatus
type IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus struct {

	// `conditions` is a list of the current states of FlowSchema.
	Conditions []*IoK8sAPIFlowcontrolV1alpha1FlowSchemaCondition `json:"conditions"`
}

// Validate validates this io k8s api flowcontrol v1alpha1 flow schema status
func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus) validateConditions(formats strfmt.Registry) error {

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
func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1alpha1FlowSchemaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
