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

// IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion CustomResourceConversion describes how to convert different versions of a CR.
//
// swagger:model io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceConversion
type IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion struct {

	// conversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail. Defaults to `["v1beta1"]`.
	ConversionReviewVersions []string `json:"conversionReviewVersions"`

	// strategy specifies how custom resources are converted between versions. Allowed values are: - `None`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `Webhook`: API Server will call to an external webhook to do the conversion. Additional information
	//   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhookClientConfig to be set.
	// Required: true
	Strategy *string `json:"strategy"`

	// webhookClientConfig is the instructions for how to call the webhook if strategy is `Webhook`. Required when `strategy` is set to `Webhook`.
	WebhookClientConfig *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1WebhookClientConfig `json:"webhookClientConfig,omitempty"`
}

// Validate validates this io k8s apiextensions apiserver pkg apis apiextensions v1beta1 custom resource conversion
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookClientConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion) validateStrategy(formats strfmt.Registry) error {

	if err := validate.Required("strategy", "body", m.Strategy); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion) validateWebhookClientConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.WebhookClientConfig) { // not required
		return nil
	}

	if m.WebhookClientConfig != nil {
		if err := m.WebhookClientConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webhookClientConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion) UnmarshalBinary(b []byte) error {
	var res IoK8sApiextensionsApiserverPkgApisApiextensionsV1beta1CustomResourceConversion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}