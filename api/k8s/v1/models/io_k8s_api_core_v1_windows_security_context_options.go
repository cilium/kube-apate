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

// IoK8sAPICoreV1WindowsSecurityContextOptions WindowsSecurityContextOptions contain Windows-specific options and credentials.
//
// swagger:model io.k8s.api.core.v1.WindowsSecurityContextOptions
type IoK8sAPICoreV1WindowsSecurityContextOptions struct {

	// GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
	GmsaCredentialSpec string `json:"gmsaCredentialSpec,omitempty"`

	// GMSACredentialSpecName is the name of the GMSA credential spec to use.
	GmsaCredentialSpecName string `json:"gmsaCredentialSpecName,omitempty"`

	// The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsUserName string `json:"runAsUserName,omitempty"`
}

// Validate validates this io k8s api core v1 windows security context options
func (m *IoK8sAPICoreV1WindowsSecurityContextOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1WindowsSecurityContextOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1WindowsSecurityContextOptions) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1WindowsSecurityContextOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
