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

// IoK8sAPIAutoscalingV2beta2MetricValueStatus MetricValueStatus holds the current value for a metric
//
// swagger:model io.k8s.api.autoscaling.v2beta2.MetricValueStatus
type IoK8sAPIAutoscalingV2beta2MetricValueStatus struct {

	// currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
	AverageUtilization int32 `json:"averageUtilization,omitempty"`

	// averageValue is the current value of the average of the metric across all relevant pods (as a quantity)
	AverageValue IoK8sApimachineryPkgAPIResourceQuantity `json:"averageValue,omitempty"`

	// value is the current value of the metric (as a quantity).
	Value IoK8sApimachineryPkgAPIResourceQuantity `json:"value,omitempty"`
}

// Validate validates this io k8s api autoscaling v2beta2 metric value status
func (m *IoK8sAPIAutoscalingV2beta2MetricValueStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2MetricValueStatus) validateAverageValue(formats strfmt.Registry) error {

	if swag.IsZero(m.AverageValue) { // not required
		return nil
	}

	if err := m.AverageValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("averageValue")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2MetricValueStatus) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := m.Value.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("value")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2MetricValueStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2MetricValueStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta2MetricValueStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
