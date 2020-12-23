// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewConnectCoreV1PostNamespacedPodExecParams creates a new ConnectCoreV1PostNamespacedPodExecParams object
// with the default values initialized.
func NewConnectCoreV1PostNamespacedPodExecParams() *ConnectCoreV1PostNamespacedPodExecParams {
	var ()
	return &ConnectCoreV1PostNamespacedPodExecParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectCoreV1PostNamespacedPodExecParamsWithTimeout creates a new ConnectCoreV1PostNamespacedPodExecParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectCoreV1PostNamespacedPodExecParamsWithTimeout(timeout time.Duration) *ConnectCoreV1PostNamespacedPodExecParams {
	var ()
	return &ConnectCoreV1PostNamespacedPodExecParams{

		timeout: timeout,
	}
}

// NewConnectCoreV1PostNamespacedPodExecParamsWithContext creates a new ConnectCoreV1PostNamespacedPodExecParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectCoreV1PostNamespacedPodExecParamsWithContext(ctx context.Context) *ConnectCoreV1PostNamespacedPodExecParams {
	var ()
	return &ConnectCoreV1PostNamespacedPodExecParams{

		Context: ctx,
	}
}

// NewConnectCoreV1PostNamespacedPodExecParamsWithHTTPClient creates a new ConnectCoreV1PostNamespacedPodExecParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectCoreV1PostNamespacedPodExecParamsWithHTTPClient(client *http.Client) *ConnectCoreV1PostNamespacedPodExecParams {
	var ()
	return &ConnectCoreV1PostNamespacedPodExecParams{
		HTTPClient: client,
	}
}

/*ConnectCoreV1PostNamespacedPodExecParams contains all the parameters to send to the API endpoint
for the connect core v1 post namespaced pod exec operation typically these are written to a http.Request
*/
type ConnectCoreV1PostNamespacedPodExecParams struct {

	/*Command
	  Command is the remote command to execute. argv array. Not executed within a shell.

	*/
	Command *string
	/*Container
	  Container in which to execute the command. Defaults to only container if there is only one container in the pod.

	*/
	Container *string
	/*Name
	  name of the PodExecOptions

	*/
	Name string
	/*Namespace
	  object name and auth scope, such as for teams and projects

	*/
	Namespace string
	/*Stderr
	  Redirect the standard error stream of the pod for this call. Defaults to true.

	*/
	Stderr *bool
	/*Stdin
	  Redirect the standard input stream of the pod for this call. Defaults to false.

	*/
	Stdin *bool
	/*Stdout
	  Redirect the standard output stream of the pod for this call. Defaults to true.

	*/
	Stdout *bool
	/*Tty
	  TTY if true indicates that a tty will be allocated for the exec call. Defaults to false.

	*/
	Tty *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithTimeout(timeout time.Duration) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithContext(ctx context.Context) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithHTTPClient(client *http.Client) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommand adds the command to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithCommand(command *string) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetCommand(command)
	return o
}

// SetCommand adds the command to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetCommand(command *string) {
	o.Command = command
}

// WithContainer adds the container to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithContainer(container *string) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetContainer(container)
	return o
}

// SetContainer adds the container to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetContainer(container *string) {
	o.Container = container
}

// WithName adds the name to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithName(name string) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetName(name string) {
	o.Name = name
}

// WithNamespace adds the namespace to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithNamespace(namespace string) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithStderr adds the stderr to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithStderr(stderr *bool) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetStderr(stderr)
	return o
}

// SetStderr adds the stderr to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetStderr(stderr *bool) {
	o.Stderr = stderr
}

// WithStdin adds the stdin to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithStdin(stdin *bool) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetStdin(stdin)
	return o
}

// SetStdin adds the stdin to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetStdin(stdin *bool) {
	o.Stdin = stdin
}

// WithStdout adds the stdout to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithStdout(stdout *bool) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetStdout(stdout)
	return o
}

// SetStdout adds the stdout to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetStdout(stdout *bool) {
	o.Stdout = stdout
}

// WithTty adds the tty to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) WithTty(tty *bool) *ConnectCoreV1PostNamespacedPodExecParams {
	o.SetTty(tty)
	return o
}

// SetTty adds the tty to the connect core v1 post namespaced pod exec params
func (o *ConnectCoreV1PostNamespacedPodExecParams) SetTty(tty *bool) {
	o.Tty = tty
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectCoreV1PostNamespacedPodExecParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Command != nil {

		// query param command
		var qrCommand string
		if o.Command != nil {
			qrCommand = *o.Command
		}
		qCommand := qrCommand
		if qCommand != "" {
			if err := r.SetQueryParam("command", qCommand); err != nil {
				return err
			}
		}

	}

	if o.Container != nil {

		// query param container
		var qrContainer string
		if o.Container != nil {
			qrContainer = *o.Container
		}
		qContainer := qrContainer
		if qContainer != "" {
			if err := r.SetQueryParam("container", qContainer); err != nil {
				return err
			}
		}

	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.Stderr != nil {

		// query param stderr
		var qrStderr bool
		if o.Stderr != nil {
			qrStderr = *o.Stderr
		}
		qStderr := swag.FormatBool(qrStderr)
		if qStderr != "" {
			if err := r.SetQueryParam("stderr", qStderr); err != nil {
				return err
			}
		}

	}

	if o.Stdin != nil {

		// query param stdin
		var qrStdin bool
		if o.Stdin != nil {
			qrStdin = *o.Stdin
		}
		qStdin := swag.FormatBool(qrStdin)
		if qStdin != "" {
			if err := r.SetQueryParam("stdin", qStdin); err != nil {
				return err
			}
		}

	}

	if o.Stdout != nil {

		// query param stdout
		var qrStdout bool
		if o.Stdout != nil {
			qrStdout = *o.Stdout
		}
		qStdout := swag.FormatBool(qrStdout)
		if qStdout != "" {
			if err := r.SetQueryParam("stdout", qStdout); err != nil {
				return err
			}
		}

	}

	if o.Tty != nil {

		// query param tty
		var qrTty bool
		if o.Tty != nil {
			qrTty = *o.Tty
		}
		qTty := swag.FormatBool(qrTty)
		if qTty != "" {
			if err := r.SetQueryParam("tty", qTty); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}