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

// IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec APIServiceSpec contains information for locating and communicating with a server. Only https is supported, though you are able to disable certificate verification.
//
// swagger:model io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIServiceSpec
type IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec struct {

	// CABundle is a PEM encoded CA bundle which will be used to validate an API server's serving certificate. If unspecified, system trust roots on the apiserver are used.
	// Format: byte
	CaBundle strfmt.Base64 `json:"caBundle,omitempty"`

	// Group is the API group name this server hosts
	Group string `json:"group,omitempty"`

	// GroupPriorityMininum is the priority this group should have at least. Higher priority means that the group is preferred by clients over lower priority ones. Note that other versions of this group might specify even higher GroupPriorityMininum values such that the whole group gets a higher priority. The primary sort is based on GroupPriorityMinimum, ordered highest number to lowest (20 before 10). The secondary sort is based on the alphabetical comparison of the name of the object.  (v1.bar before v1.foo) We'd recommend something like: *.k8s.io (except extensions) at 18000 and PaaSes (OpenShift, Deis) are recommended to be in the 2000s
	// Required: true
	GroupPriorityMinimum *int32 `json:"groupPriorityMinimum"`

	// InsecureSkipTLSVerify disables TLS certificate verification when communicating with this server. This is strongly discouraged.  You should use the CABundle instead.
	InsecureSkipTLSVerify bool `json:"insecureSkipTLSVerify,omitempty"`

	// Service is a reference to the service for this API server.  It must communicate on port 443. If the Service is nil, that means the handling for the API groupversion is handled locally on this server. The call will simply delegate to the normal handler chain to be fulfilled.
	Service *IoK8sKubeAggregatorPkgApisApiregistrationV1ServiceReference `json:"service,omitempty"`

	// Version is the API version this server hosts.  For example, "v1"
	Version string `json:"version,omitempty"`

	// VersionPriority controls the ordering of this API version inside of its group.  Must be greater than zero. The primary sort is based on VersionPriority, ordered highest to lowest (20 before 10). Since it's inside of a group, the number can be small, probably in the 10s. In case of equal version priorities, the version string will be used to compute the order inside a group. If the version string is "kube-like", it will sort above non "kube-like" version strings, which are ordered lexicographically. "Kube-like" versions start with a "v", then are followed by a number (the major version), then optionally the string "alpha" or "beta" and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10.
	// Required: true
	VersionPriority *int32 `json:"versionPriority"`
}

// Validate validates this io k8s kube aggregator pkg apis apiregistration v1 API service spec
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupPriorityMinimum(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionPriority(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) validateGroupPriorityMinimum(formats strfmt.Registry) error {

	if err := validate.Required("groupPriorityMinimum", "body", m.GroupPriorityMinimum); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) validateService(formats strfmt.Registry) error {

	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) validateVersionPriority(formats strfmt.Registry) error {

	if err := validate.Required("versionPriority", "body", m.VersionPriority); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec) UnmarshalBinary(b []byte) error {
	var res IoK8sKubeAggregatorPkgApisApiregistrationV1APIServiceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
