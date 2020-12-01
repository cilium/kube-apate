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

// IoK8sAPIDiscoveryV1beta1EndpointConditions EndpointConditions represents the current condition of an endpoint.
//
// swagger:model io.k8s.api.discovery.v1beta1.EndpointConditions
type IoK8sAPIDiscoveryV1beta1EndpointConditions struct {

	// ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
	Ready bool `json:"ready,omitempty"`
}

// Validate validates this io k8s api discovery v1beta1 endpoint conditions
func (m *IoK8sAPIDiscoveryV1beta1EndpointConditions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1beta1EndpointConditions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIDiscoveryV1beta1EndpointConditions) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIDiscoveryV1beta1EndpointConditions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
