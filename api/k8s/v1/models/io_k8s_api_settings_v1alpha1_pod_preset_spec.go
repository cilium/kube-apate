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

// IoK8sAPISettingsV1alpha1PodPresetSpec PodPresetSpec is a description of a pod preset.
//
// swagger:model io.k8s.api.settings.v1alpha1.PodPresetSpec
type IoK8sAPISettingsV1alpha1PodPresetSpec struct {

	// Env defines the collection of EnvVar to inject into containers.
	Env []*IoK8sAPICoreV1EnvVar `json:"env"`

	// EnvFrom defines the collection of EnvFromSource to inject into containers.
	EnvFrom []*IoK8sAPICoreV1EnvFromSource `json:"envFrom"`

	// Selector is a label query over a set of resources, in this case pods. Required.
	Selector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"selector,omitempty"`

	// VolumeMounts defines the collection of VolumeMount to inject into containers.
	VolumeMounts []*IoK8sAPICoreV1VolumeMount `json:"volumeMounts"`

	// Volumes defines the collection of Volume to inject into the pod.
	Volumes []*IoK8sAPICoreV1Volume `json:"volumes"`
}

// Validate validates this io k8s api settings v1alpha1 pod preset spec
func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeMounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) validateEnv(formats strfmt.Registry) error {

	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) validateEnvFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvFrom) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvFrom); i++ {
		if swag.IsZero(m.EnvFrom[i]) { // not required
			continue
		}

		if m.EnvFrom[i] != nil {
			if err := m.EnvFrom[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envFrom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) validateSelector(formats strfmt.Registry) error {

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

func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) validateVolumeMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeMounts) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeMounts); i++ {
		if swag.IsZero(m.VolumeMounts[i]) { // not required
			continue
		}

		if m.VolumeMounts[i] != nil {
			if err := m.VolumeMounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeMounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) validateVolumes(formats strfmt.Registry) error {

	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPISettingsV1alpha1PodPresetSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sAPISettingsV1alpha1PodPresetSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
