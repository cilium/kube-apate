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

// IoK8sAPIAuthorizationV1NonResourceAttributes NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface
//
// swagger:model io.k8s.api.authorization.v1.NonResourceAttributes
type IoK8sAPIAuthorizationV1NonResourceAttributes struct {

	// Path is the URL path of the request
	Path string `json:"path,omitempty"`

	// Verb is the standard HTTP verb
	Verb string `json:"verb,omitempty"`
}

// Validate validates this io k8s api authorization v1 non resource attributes
func (m *IoK8sAPIAuthorizationV1NonResourceAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1NonResourceAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAuthorizationV1NonResourceAttributes) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAuthorizationV1NonResourceAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
