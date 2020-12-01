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

// IoK8sAPIStorageV1beta1VolumeAttachmentSource VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
//
// swagger:model io.k8s.api.storage.v1beta1.VolumeAttachmentSource
type IoK8sAPIStorageV1beta1VolumeAttachmentSource struct {

	// inlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is alpha-level and is only honored by servers that enabled the CSIMigration feature.
	InlineVolumeSpec *IoK8sAPICoreV1PersistentVolumeSpec `json:"inlineVolumeSpec,omitempty"`

	// Name of the persistent volume to attach.
	PersistentVolumeName string `json:"persistentVolumeName,omitempty"`
}

// Validate validates this io k8s api storage v1beta1 volume attachment source
func (m *IoK8sAPIStorageV1beta1VolumeAttachmentSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInlineVolumeSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIStorageV1beta1VolumeAttachmentSource) validateInlineVolumeSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.InlineVolumeSpec) { // not required
		return nil
	}

	if m.InlineVolumeSpec != nil {
		if err := m.InlineVolumeSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("inlineVolumeSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeAttachmentSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1beta1VolumeAttachmentSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1beta1VolumeAttachmentSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
