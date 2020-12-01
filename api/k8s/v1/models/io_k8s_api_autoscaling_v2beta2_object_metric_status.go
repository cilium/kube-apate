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

// IoK8sAPIAutoscalingV2beta2ObjectMetricStatus ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
//
// swagger:model io.k8s.api.autoscaling.v2beta2.ObjectMetricStatus
type IoK8sAPIAutoscalingV2beta2ObjectMetricStatus struct {

	// current contains the current value for the given metric
	// Required: true
	Current *IoK8sAPIAutoscalingV2beta2MetricValueStatus `json:"current"`

	// described object
	// Required: true
	DescribedObject *IoK8sAPIAutoscalingV2beta2CrossVersionObjectReference `json:"describedObject"`

	// metric identifies the target metric by name and selector
	// Required: true
	Metric *IoK8sAPIAutoscalingV2beta2MetricIdentifier `json:"metric"`
}

// Validate validates this io k8s api autoscaling v2beta2 object metric status
func (m *IoK8sAPIAutoscalingV2beta2ObjectMetricStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescribedObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ObjectMetricStatus) validateCurrent(formats strfmt.Registry) error {

	if err := validate.Required("current", "body", m.Current); err != nil {
		return err
	}

	if m.Current != nil {
		if err := m.Current.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ObjectMetricStatus) validateDescribedObject(formats strfmt.Registry) error {

	if err := validate.Required("describedObject", "body", m.DescribedObject); err != nil {
		return err
	}

	if m.DescribedObject != nil {
		if err := m.DescribedObject.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("describedObject")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta2ObjectMetricStatus) validateMetric(formats strfmt.Registry) error {

	if err := validate.Required("metric", "body", m.Metric); err != nil {
		return err
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2ObjectMetricStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta2ObjectMetricStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta2ObjectMetricStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
