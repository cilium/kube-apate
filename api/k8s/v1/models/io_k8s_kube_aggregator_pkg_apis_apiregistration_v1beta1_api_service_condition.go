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

// IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition APIServiceCondition describes the state of an APIService at a particular point
//
// swagger:model io.k8s.kube-aggregator.pkg.apis.apiregistration.v1beta1.APIServiceCondition
type IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition struct {

	// Last time the condition transitioned from one status to another.
	// Format: date-time
	LastTransitionTime IoK8sApimachineryPkgApisMetaV1Time `json:"lastTransitionTime,omitempty"`

	// Human-readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// Unique, one-word, CamelCase reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status is the status of the condition. Can be True, False, Unknown.
	// Required: true
	Status *string `json:"status"`

	// Type is the type of the condition.
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this io k8s kube aggregator pkg apis apiregistration v1beta1 API service condition
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastTransitionTime(formats); err != nil {
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

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition) validateLastTransitionTime(formats strfmt.Registry) error {

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

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition) UnmarshalBinary(b []byte) error {
	var res IoK8sKubeAggregatorPkgApisApiregistrationV1beta1APIServiceCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
