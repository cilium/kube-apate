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

// IoK8sAPIAutoscalingV2beta1ObjectMetricSource ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
//
// swagger:model io.k8s.api.autoscaling.v2beta1.ObjectMetricSource
type IoK8sAPIAutoscalingV2beta1ObjectMetricSource struct {

	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
	AverageValue IoK8sApimachineryPkgAPIResourceQuantity `json:"averageValue,omitempty"`

	// metricName is the name of the metric in question.
	// Required: true
	MetricName *string `json:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
	Selector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"selector,omitempty"`

	// target is the described Kubernetes object.
	// Required: true
	Target *IoK8sAPIAutoscalingV2beta1CrossVersionObjectReference `json:"target"`

	// targetValue is the target value of the metric (as a quantity).
	// Required: true
	TargetValue IoK8sApimachineryPkgAPIResourceQuantity `json:"targetValue"`
}

// Validate validates this io k8s api autoscaling v2beta1 object metric source
func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetricName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) validateAverageValue(formats strfmt.Registry) error {

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

func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) validateMetricName(formats strfmt.Registry) error {

	if err := validate.Required("metricName", "body", m.MetricName); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) validateSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.Selector) { // not required
		return nil
	}

	if m.Selector != nil {
		if err := m.Selector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) validateTargetValue(formats strfmt.Registry) error {

	if err := m.TargetValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("targetValue")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ObjectMetricSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta1ObjectMetricSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
