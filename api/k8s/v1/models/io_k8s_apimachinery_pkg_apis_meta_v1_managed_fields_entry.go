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

// IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry ManagedFieldsEntry is a workflow-id, a FieldSet and the group version of the resource that the fieldset applies to.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
type IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry struct {

	// APIVersion defines the version of this resource that this field set applies to. The format is "group/version" just like the top-level APIVersion field. It is necessary to track the version of a field set because it cannot be automatically converted.
	APIVersion string `json:"apiVersion,omitempty"`

	// FieldsType is the discriminator for the different fields format and version. There is currently only one possible value: "FieldsV1"
	FieldsType string `json:"fieldsType,omitempty"`

	// FieldsV1 holds the first JSON version format as described in the "FieldsV1" type.
	FieldsV1 IoK8sApimachineryPkgApisMetaV1FieldsV1 `json:"fieldsV1,omitempty"`

	// Manager is an identifier of the workflow managing these fields.
	Manager string `json:"manager,omitempty"`

	// Operation is the type of operation which lead to this ManagedFieldsEntry being created. The only valid values for this field are 'Apply' and 'Update'.
	Operation string `json:"operation,omitempty"`

	// Time is timestamp of when these fields were set. It should always be empty if Operation is 'Apply'
	// Format: date-time
	Time IoK8sApimachineryPkgApisMetaV1Time `json:"time,omitempty"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 managed fields entry
func (m *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := m.Time.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("time")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1ManagedFieldsEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}