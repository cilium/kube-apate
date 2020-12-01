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
	"github.com/go-openapi/validate"
)

// IoK8sAPINetworkingV1HTTPIngressRuleValue HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
//
// swagger:model io.k8s.api.networking.v1.HTTPIngressRuleValue
type IoK8sAPINetworkingV1HTTPIngressRuleValue struct {

	// A collection of paths that map requests to backends.
	// Required: true
	Paths []*IoK8sAPINetworkingV1HTTPIngressPath `json:"paths"`
}

// Validate validates this io k8s api networking v1 HTTP ingress rule value
func (m *IoK8sAPINetworkingV1HTTPIngressRuleValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaths(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPINetworkingV1HTTPIngressRuleValue) validatePaths(formats strfmt.Registry) error {

	if err := validate.Required("paths", "body", m.Paths); err != nil {
		return err
	}

	for i := 0; i < len(m.Paths); i++ {
		if swag.IsZero(m.Paths[i]) { // not required
			continue
		}

		if m.Paths[i] != nil {
			if err := m.Paths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1HTTPIngressRuleValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPINetworkingV1HTTPIngressRuleValue) UnmarshalBinary(b []byte) error {
	var res IoK8sAPINetworkingV1HTTPIngressRuleValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
