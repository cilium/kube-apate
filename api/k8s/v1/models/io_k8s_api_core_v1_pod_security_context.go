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
)

// IoK8sAPICoreV1PodSecurityContext PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.
//
// swagger:model io.k8s.api.core.v1.PodSecurityContext
type IoK8sAPICoreV1PodSecurityContext struct {

	// A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:
	//
	// 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----
	//
	// If unset, the Kubelet will not modify the ownership and permissions of any volume.
	FsGroup int64 `json:"fsGroup,omitempty"`

	// fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are "OnRootMismatch" and "Always". If not specified defaults to "Always".
	FsGroupChangePolicy string `json:"fsGroupChangePolicy,omitempty"`

	// The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	RunAsGroup int64 `json:"runAsGroup,omitempty"`

	// Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsNonRoot bool `json:"runAsNonRoot,omitempty"`

	// The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	RunAsUser int64 `json:"runAsUser,omitempty"`

	// The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.
	SeLinuxOptions *IoK8sAPICoreV1SELinuxOptions `json:"seLinuxOptions,omitempty"`

	// The seccomp options to use by the containers in this pod.
	SeccompProfile *IoK8sAPICoreV1SeccompProfile `json:"seccompProfile,omitempty"`

	// A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.
	SupplementalGroups []int64 `json:"supplementalGroups"`

	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.
	Sysctls []*IoK8sAPICoreV1Sysctl `json:"sysctls"`

	// The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	WindowsOptions *IoK8sAPICoreV1WindowsSecurityContextOptions `json:"windowsOptions,omitempty"`
}

// Validate validates this io k8s api core v1 pod security context
func (m *IoK8sAPICoreV1PodSecurityContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeLinuxOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeccompProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSysctls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWindowsOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1PodSecurityContext) validateSeLinuxOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.SeLinuxOptions) { // not required
		return nil
	}

	if m.SeLinuxOptions != nil {
		if err := m.SeLinuxOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seLinuxOptions")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSecurityContext) validateSeccompProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.SeccompProfile) { // not required
		return nil
	}

	if m.SeccompProfile != nil {
		if err := m.SeccompProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("seccompProfile")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1PodSecurityContext) validateSysctls(formats strfmt.Registry) error {

	if swag.IsZero(m.Sysctls) { // not required
		return nil
	}

	for i := 0; i < len(m.Sysctls); i++ {
		if swag.IsZero(m.Sysctls[i]) { // not required
			continue
		}

		if m.Sysctls[i] != nil {
			if err := m.Sysctls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sysctls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1PodSecurityContext) validateWindowsOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.WindowsOptions) { // not required
		return nil
	}

	if m.WindowsOptions != nil {
		if err := m.WindowsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("windowsOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodSecurityContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1PodSecurityContext) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1PodSecurityContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}