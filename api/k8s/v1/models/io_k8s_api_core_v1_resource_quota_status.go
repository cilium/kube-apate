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

// IoK8sAPICoreV1ResourceQuotaStatus ResourceQuotaStatus defines the enforced hard limits and observed use.
//
// swagger:model io.k8s.api.core.v1.ResourceQuotaStatus
type IoK8sAPICoreV1ResourceQuotaStatus struct {

	// Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
	Hard map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"hard,omitempty"`

	// Used is the current observed total usage of the resource in the namespace.
	Used map[string]IoK8sApimachineryPkgAPIResourceQuantity `json:"used,omitempty"`
}

// Validate validates this io k8s api core v1 resource quota status
func (m *IoK8sAPICoreV1ResourceQuotaStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHard(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1ResourceQuotaStatus) validateHard(formats strfmt.Registry) error {

	if swag.IsZero(m.Hard) { // not required
		return nil
	}

	for k := range m.Hard {

		if val, ok := m.Hard[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1ResourceQuotaStatus) validateUsed(formats strfmt.Registry) error {

	if swag.IsZero(m.Used) { // not required
		return nil
	}

	for k := range m.Used {

		if val, ok := m.Used[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ResourceQuotaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ResourceQuotaStatus) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ResourceQuotaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
