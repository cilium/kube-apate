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

// IoK8sAPICoreV1CinderPersistentVolumeSource Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.CinderPersistentVolumeSource
type IoK8sAPICoreV1CinderPersistentVolumeSource struct {

	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	FsType string `json:"fsType,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	ReadOnly bool `json:"readOnly,omitempty"`

	// Optional: points to a secret object containing parameters used to connect to OpenStack.
	SecretRef *IoK8sAPICoreV1SecretReference `json:"secretRef,omitempty"`

	// volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	// Required: true
	VolumeID *string `json:"volumeID"`
}

// Validate validates this io k8s api core v1 cinder persistent volume source
func (m *IoK8sAPICoreV1CinderPersistentVolumeSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecretRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1CinderPersistentVolumeSource) validateSecretRef(formats strfmt.Registry) error {

	if swag.IsZero(m.SecretRef) { // not required
		return nil
	}

	if m.SecretRef != nil {
		if err := m.SecretRef.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secretRef")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1CinderPersistentVolumeSource) validateVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("volumeID", "body", m.VolumeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1CinderPersistentVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1CinderPersistentVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1CinderPersistentVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}