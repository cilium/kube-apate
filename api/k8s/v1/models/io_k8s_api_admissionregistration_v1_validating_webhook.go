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

// IoK8sAPIAdmissionregistrationV1ValidatingWebhook ValidatingWebhook describes an admission webhook and the resources and operations it applies to.
//
// swagger:model io.k8s.api.admissionregistration.v1.ValidatingWebhook
type IoK8sAPIAdmissionregistrationV1ValidatingWebhook struct {

	// AdmissionReviewVersions is an ordered list of preferred `AdmissionReview` versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.
	// Required: true
	AdmissionReviewVersions []string `json:"admissionReviewVersions"`

	// ClientConfig defines how to communicate with the hook. Required
	// Required: true
	ClientConfig *IoK8sAPIAdmissionregistrationV1WebhookClientConfig `json:"clientConfig"`

	// FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.
	FailurePolicy string `json:"failurePolicy,omitempty"`

	// matchPolicy defines how the "rules" list is used to match incoming requests. Allowed values are "Exact" or "Equivalent".
	//
	// - Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.
	//
	// - Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and "rules" only included `apiGroups:["apps"], apiVersions:["v1"], resources: ["deployments"]`, a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.
	//
	// Defaults to "Equivalent"
	MatchPolicy string `json:"matchPolicy,omitempty"`

	// The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where "imagepolicy" is the name of the webhook, and kubernetes.io is the name of the organization. Required.
	// Required: true
	Name *string `json:"name"`

	// NamespaceSelector decides whether to run the webhook on an object based on whether the namespace for that object matches the selector. If the object itself is a namespace, the matching is performed on object.metadata.labels. If the object is another cluster scoped resource, it never skips the webhook.
	//
	// For example, to run the webhook on any objects whose namespace is not associated with "runlevel" of "0" or "1";  you will set the selector as follows: "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "runlevel",
	//       "operator": "NotIn",
	//       "values": [
	//         "0",
	//         "1"
	//       ]
	//     }
	//   ]
	// }
	//
	// If instead you want to only run the webhook on any objects whose namespace is associated with the "environment" of "prod" or "staging"; you will set the selector as follows: "namespaceSelector": {
	//   "matchExpressions": [
	//     {
	//       "key": "environment",
	//       "operator": "In",
	//       "values": [
	//         "prod",
	//         "staging"
	//       ]
	//     }
	//   ]
	// }
	//
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels for more examples of label selectors.
	//
	// Default to the empty LabelSelector, which matches everything.
	NamespaceSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"namespaceSelector,omitempty"`

	// ObjectSelector decides whether to run the webhook based on if the object has matching labels. objectSelector is evaluated against both the oldObject and newObject that would be sent to the webhook, and is considered to match if either object matches the selector. A null object (oldObject in the case of create, or newObject in the case of delete) or an object that cannot have labels (like a DeploymentRollback or a PodProxyOptions object) is not considered to match. Use the object selector only if the webhook is opt-in, because end users may skip the admission webhook by setting the labels. Default to the empty LabelSelector, which matches everything.
	ObjectSelector *IoK8sApimachineryPkgApisMetaV1LabelSelector `json:"objectSelector,omitempty"`

	// Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.
	Rules []*IoK8sAPIAdmissionregistrationV1RuleWithOperations `json:"rules"`

	// SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission change and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.
	// Required: true
	SideEffects *string `json:"sideEffects"`

	// TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.
	TimeoutSeconds int32 `json:"timeoutSeconds,omitempty"`
}

// Validate validates this io k8s api admissionregistration v1 validating webhook
func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionReviewVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectSelector(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSideEffects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateAdmissionReviewVersions(formats strfmt.Registry) error {

	if err := validate.Required("admissionReviewVersions", "body", m.AdmissionReviewVersions); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateClientConfig(formats strfmt.Registry) error {

	if err := validate.Required("clientConfig", "body", m.ClientConfig); err != nil {
		return err
	}

	if m.ClientConfig != nil {
		if err := m.ClientConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clientConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateNamespaceSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.NamespaceSelector) { // not required
		return nil
	}

	if m.NamespaceSelector != nil {
		if err := m.NamespaceSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("namespaceSelector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateObjectSelector(formats strfmt.Registry) error {

	if swag.IsZero(m.ObjectSelector) { // not required
		return nil
	}

	if m.ObjectSelector != nil {
		if err := m.ObjectSelector.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("objectSelector")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateRules(formats strfmt.Registry) error {

	if swag.IsZero(m.Rules) { // not required
		return nil
	}

	for i := 0; i < len(m.Rules); i++ {
		if swag.IsZero(m.Rules[i]) { // not required
			continue
		}

		if m.Rules[i] != nil {
			if err := m.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) validateSideEffects(formats strfmt.Registry) error {

	if err := validate.Required("sideEffects", "body", m.SideEffects); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPIAdmissionregistrationV1ValidatingWebhook) UnmarshalBinary(b []byte) error {
	var res IoK8sAPIAdmissionregistrationV1ValidatingWebhook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
