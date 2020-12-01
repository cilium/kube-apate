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

// IoK8sAPICoreV1EphemeralContainer An EphemeralContainer is a container that may be added temporarily to an existing pod for user-initiated activities such as debugging. Ephemeral containers have no resource or scheduling guarantees, and they will not be restarted when they exit or when a pod is removed or restarted. If an ephemeral container causes a pod to exceed its resource allocation, the pod may be evicted. Ephemeral containers may not be added by directly updating the pod spec. They must be added via the pod's ephemeralcontainers subresource, and they will appear in the pod spec once added. This is an alpha feature enabled by the EphemeralContainers feature flag.
//
// swagger:model io.k8s.api.core.v1.EphemeralContainer
type IoK8sAPICoreV1EphemeralContainer struct {

	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args []string `json:"args"`

	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Command []string `json:"command"`

	// List of environment variables to set in the container. Cannot be updated.
	Env []*IoK8sAPICoreV1EnvVar `json:"env"`

	// List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	EnvFrom []*IoK8sAPICoreV1EnvFromSource `json:"envFrom"`

	// Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images
	Image string `json:"image,omitempty"`

	// Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	// Lifecycle is not allowed for ephemeral containers.
	Lifecycle *IoK8sAPICoreV1Lifecycle `json:"lifecycle,omitempty"`

	// Probes are not allowed for ephemeral containers.
	LivenessProbe *IoK8sAPICoreV1Probe `json:"livenessProbe,omitempty"`

	// Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.
	// Required: true
	Name *string `json:"name"`

	// Ports are not allowed for ephemeral containers.
	Ports []*IoK8sAPICoreV1ContainerPort `json:"ports"`

	// Probes are not allowed for ephemeral containers.
	ReadinessProbe *IoK8sAPICoreV1Probe `json:"readinessProbe,omitempty"`

	// Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.
	Resources *IoK8sAPICoreV1ResourceRequirements `json:"resources,omitempty"`

	// SecurityContext is not allowed for ephemeral containers.
	SecurityContext *IoK8sAPICoreV1SecurityContext `json:"securityContext,omitempty"`

	// Probes are not allowed for ephemeral containers.
	StartupProbe *IoK8sAPICoreV1Probe `json:"startupProbe,omitempty"`

	// Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Stdin bool `json:"stdin,omitempty"`

	// Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	StdinOnce bool `json:"stdinOnce,omitempty"`

	// If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container is run in whatever namespaces are shared for the pod. Note that the container runtime must support this feature.
	TargetContainerName string `json:"targetContainerName,omitempty"`

	// Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	TerminationMessagePath string `json:"terminationMessagePath,omitempty"`

	// Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	TerminationMessagePolicy string `json:"terminationMessagePolicy,omitempty"`

	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Tty bool `json:"tty,omitempty"`

	// volumeDevices is the list of block devices to be used by the container.
	VolumeDevices []*IoK8sAPICoreV1VolumeDevice `json:"volumeDevices"`

	// Pod volumes to mount into the container's filesystem. Cannot be updated.
	VolumeMounts []*IoK8sAPICoreV1VolumeMount `json:"volumeMounts"`

	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	WorkingDir string `json:"workingDir,omitempty"`
}

// Validate validates this io k8s api core v1 ephemeral container
func (m *IoK8sAPICoreV1EphemeralContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLifecycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLivenessProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePorts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadinessProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartupProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeMounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateEnv(formats strfmt.Registry) error {

	if swag.IsZero(m.Env) { // not required
		return nil
	}

	for i := 0; i < len(m.Env); i++ {
		if swag.IsZero(m.Env[i]) { // not required
			continue
		}

		if m.Env[i] != nil {
			if err := m.Env[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("env" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateEnvFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.EnvFrom) { // not required
		return nil
	}

	for i := 0; i < len(m.EnvFrom); i++ {
		if swag.IsZero(m.EnvFrom[i]) { // not required
			continue
		}

		if m.EnvFrom[i] != nil {
			if err := m.EnvFrom[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("envFrom" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateLifecycle(formats strfmt.Registry) error {

	if swag.IsZero(m.Lifecycle) { // not required
		return nil
	}

	if m.Lifecycle != nil {
		if err := m.Lifecycle.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lifecycle")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateLivenessProbe(formats strfmt.Registry) error {

	if swag.IsZero(m.LivenessProbe) { // not required
		return nil
	}

	if m.LivenessProbe != nil {
		if err := m.LivenessProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("livenessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validatePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.Ports) { // not required
		return nil
	}

	for i := 0; i < len(m.Ports); i++ {
		if swag.IsZero(m.Ports[i]) { // not required
			continue
		}

		if m.Ports[i] != nil {
			if err := m.Ports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateReadinessProbe(formats strfmt.Registry) error {

	if swag.IsZero(m.ReadinessProbe) { // not required
		return nil
	}

	if m.ReadinessProbe != nil {
		if err := m.ReadinessProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readinessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateSecurityContext(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityContext) { // not required
		return nil
	}

	if m.SecurityContext != nil {
		if err := m.SecurityContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityContext")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateStartupProbe(formats strfmt.Registry) error {

	if swag.IsZero(m.StartupProbe) { // not required
		return nil
	}

	if m.StartupProbe != nil {
		if err := m.StartupProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("startupProbe")
			}
			return err
		}
	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateVolumeDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeDevices) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeDevices); i++ {
		if swag.IsZero(m.VolumeDevices[i]) { // not required
			continue
		}

		if m.VolumeDevices[i] != nil {
			if err := m.VolumeDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IoK8sAPICoreV1EphemeralContainer) validateVolumeMounts(formats strfmt.Registry) error {

	if swag.IsZero(m.VolumeMounts) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeMounts); i++ {
		if swag.IsZero(m.VolumeMounts[i]) { // not required
			continue
		}

		if m.VolumeMounts[i] != nil {
			if err := m.VolumeMounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeMounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sAPICoreV1EphemeralContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sAPICoreV1EphemeralContainer) UnmarshalBinary(b []byte) error {
	var res IoK8sAPICoreV1EphemeralContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
