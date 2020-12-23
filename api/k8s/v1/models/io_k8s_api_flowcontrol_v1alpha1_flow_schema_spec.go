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
	"github.com/go-openapi/validate"
)

// IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec FlowSchemaSpec describes how the FlowSchema's specification looks like.
//
// swagger:model io.k8s.api.flowcontrol.v1alpha1.FlowSchemaSpec
type IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec struct {

	// `distinguisherMethod` defines how to compute the flow distinguisher for requests that match this schema. `nil` specifies that the distinguisher is disabled and thus will always be the empty string.
	DistinguisherMethod *IoK8sAPIFlowcontrolV1alpha1FlowDistinguisherMethod `json:"distinguisherMethod,omitempty"`

	// `matchingPrecedence` is used to choose among the FlowSchemas that match a given request. The chosen FlowSchema is among those with the numerically lowest (which we take to be logically highest) MatchingPrecedence.  Each MatchingPrecedence value must be ranged in [1,10000]. Note that if the precedence is not specified, it will be set to 1000 as default.
	MatchingPrecedence int32 `json:"matchingPrecedence,omitempty"`

	// `priorityLevelConfiguration` should reference a PriorityLevelConfiguration in the cluster. If the reference cannot be resolved, the FlowSchema will be ignored and marked as invalid in its status. Required.
	// Required: true
	PriorityLevelConfiguration *IoK8sAPIFlowcontrolV1alpha1PriorityLevelConfigurationReference `json:"priorityLevelConfiguration"`

	// `rules` describes which requests will match this flow schema. This FlowSchema matches a request if and only if at least one member of rules matches the request. if it is an empty slice, there will be no requests matching the FlowSchema.
	Rules []*IoK8sAPIFlowcontrolV1alpha1PolicyRulesWithSubjects `json:"rules"`
}

// Validate validates this io k8s api flowcontrol v1alpha1 flow schema spec
func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistinguisherMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriorityLevelConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec) validateDistinguisherMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.DistinguisherMethod) { // not required
		return nil
	}

	if m.DistinguisherMethod != nil {
		if err := m.DistinguisherMethod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("distinguisherMethod")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec) validatePriorityLevelConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("priorityLevelConfiguration", "body", m.PriorityLevelConfiguration); err != nil {
		return err
	}

	if m.PriorityLevelConfiguration != nil {
		if err := m.PriorityLevelConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priorityLevelConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIFlowcontrolV1alpha1FlowSchemaSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}