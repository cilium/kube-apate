// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sAPIStorageV1VolumeNodeResources VolumeNodeResources is a set of resource limits for scheduling of volumes.
//
// swagger:model io.k8s.api.storage.v1.VolumeNodeResources
type IoK8sAPIStorageV1VolumeNodeResources struct {

	// Maximum number of unique volumes managed by the CSI driver that can be used on a node. A volume that is both attached and mounted on a node is considered to be used once, not twice. The same rule applies for a unique volume that is shared among multiple pods on the same node. If this field is not specified, then the supported number of volumes on this node is unbounded.
	Count int32 `json:"count,omitempty"`
}

// Validate validates this io k8s api storage v1 volume node resources
func (m *IoK8sAPIStorageV1VolumeNodeResources) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIStorageV1VolumeNodeResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIStorageV1VolumeNodeResources) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIStorageV1VolumeNodeResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
