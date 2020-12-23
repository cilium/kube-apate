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

// IoK8sAPICoreV1ExecAction ExecAction describes a "run in container" action.
//
// swagger:model io.k8s.api.core.v1.ExecAction
type IoK8sAPICoreV1ExecAction struct {

	// Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
	Command []string `json:"command"`
}

// Validate validates this io k8s api core v1 exec action
func (m *IoK8sAPICoreV1ExecAction) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1ExecAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1ExecAction) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1ExecAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}