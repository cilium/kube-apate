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

// IoK8sAPIAutoscalingV2beta1ResourceMetricStatus ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
//
// swagger:model io.k8s.api.autoscaling.v2beta1.ResourceMetricStatus
type IoK8sAPIAutoscalingV2beta1ResourceMetricStatus struct {

	// currentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.  It will only be present if `targetAverageValue` was set in the corresponding metric specification.
	CurrentAverageUtilization int32 `json:"currentAverageUtilization,omitempty"`

	// currentAverageValue is the current value of the average of the resource metric across all relevant pods, as a raw value (instead of as a percentage of the request), similar to the "pods" metric source type. It will always be set, regardless of the corresponding metric specification.
	// Required: true
	CurrentAverageValue IoK8sApimachineryPkgAPIResourceQuantity `json:"currentAverageValue"`

	// name is the name of the resource in question.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this io k8s api autoscaling v2beta1 resource metric status
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentAverageValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricStatus) validateCurrentAverageValue(formats strfmt.Registry) error {

	if err := m.CurrentAverageValue.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("currentAverageValue")
		}
		return err
	}

	return nil
}

func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricStatus) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAutoscalingV2beta1ResourceMetricStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAutoscalingV2beta1ResourceMetricStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
