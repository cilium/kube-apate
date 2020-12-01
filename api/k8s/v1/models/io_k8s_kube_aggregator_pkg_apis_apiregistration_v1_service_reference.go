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

// IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference ServiceReference holds a reference to Service.legacy.k8s.io
//
// swagger:model io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.ServiceReference
type IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference struct {

	// Name is the name of the service
	Name string `json:"name,omitempty"`

	// Namespace is the namespace of the service
	Namespace string `json:"namespace,omitempty"`

	// If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
	Port int32 `json:"port,omitempty"`
}

// Validate validates this io k8s kube aggregator pkg apis apiregistration v1 service reference
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference) UnmarshalBinary(b []byte) error {
	var res IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
