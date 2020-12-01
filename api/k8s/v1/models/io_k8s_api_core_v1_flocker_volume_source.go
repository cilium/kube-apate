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

// IoK8sAPICoreV1FlockerVolumeSource Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
//
// swagger:model io.k8s.api.core.v1.FlockerVolumeSource
type IoK8sAPICoreV1FlockerVolumeSource struct {

	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
	DatasetName string `json:"datasetName,omitempty"`

	// UUID of the dataset. This is unique identifier of a Flocker dataset
	DatasetUUID string `json:"datasetUUID,omitempty"`
}

// Validate validates this io k8s api core v1 flocker volume source
func (m *IoK8sAPICoreV1FlockerVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1FlockerVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1FlockerVolumeSource) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1FlockerVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
