// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package storage_v1

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

// NewWatchStorageV1VolumeAttachmentListParams creates a new WatchStorageV1VolumeAttachmentListParams object
// with the default values initialized.
func NewWatchStorageV1VolumeAttachmentListParams() *WatchStorageV1VolumeAttachmentListParams {
	var ()
	return &WatchStorageV1VolumeAttachmentListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWatchStorageV1VolumeAttachmentListParamsWithTimeout creates a new WatchStorageV1VolumeAttachmentListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWatchStorageV1VolumeAttachmentListParamsWithTimeout(timeout time.Duration) *WatchStorageV1VolumeAttachmentListParams {
	var ()
	return &WatchStorageV1VolumeAttachmentListParams{

		timeout: timeout,
	}
}

// NewWatchStorageV1VolumeAttachmentListParamsWithContext creates a new WatchStorageV1VolumeAttachmentListParams object
// with the default values initialized, and the ability to set a context for a request
func NewWatchStorageV1VolumeAttachmentListParamsWithContext(ctx context.Context) *WatchStorageV1VolumeAttachmentListParams {
	var ()
	return &WatchStorageV1VolumeAttachmentListParams{

		Context: ctx,
	}
}

// NewWatchStorageV1VolumeAttachmentListParamsWithHTTPClient creates a new WatchStorageV1VolumeAttachmentListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWatchStorageV1VolumeAttachmentListParamsWithHTTPClient(client *http.Client) *WatchStorageV1VolumeAttachmentListParams {
	var ()
	return &WatchStorageV1VolumeAttachmentListParams{
		HTTPClient: client,
	}
}

/*WatchStorageV1VolumeAttachmentListParams contains all the parameters to send to the API endpoint
for the watch storage v1 volume attachment list operation typically these are written to a http.Request
*/
type WatchStorageV1VolumeAttachmentListParams struct {

	/*AllowWatchBookmarks
	  allowWatchBookmarks requests watch events with type "BOOKMARK". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored.

	*/
	AllowWatchBookmarks *bool
	/*Continue
	  The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the "next key".

	This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.

	*/
	Continue *string
	/*FieldSelector
	  A selector to restrict the list of returned objects by their fields. Defaults to everything.

	*/
	FieldSelector *string
	/*LabelSelector
	  A selector to restrict the list of returned objects by their labels. Defaults to everything.

	*/
	LabelSelector *string
	/*Limit
	  limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.

	The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.

	*/
	Limit *int64
	/*Pretty
	  If 'true', then the output is pretty printed.

	*/
	Pretty *string
	/*ResourceVersion
	  resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset

	*/
	ResourceVersion *string
	/*ResourceVersionMatch
	  resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset

	*/
	ResourceVersionMatch *string
	/*TimeoutSeconds
	  Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.

	*/
	TimeoutSeconds *int64
	/*Watch
	  Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.

	*/
	Watch *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithTimeout(timeout time.Duration) *WatchStorageV1VolumeAttachmentListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithContext(ctx context.Context) *WatchStorageV1VolumeAttachmentListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithHTTPClient(client *http.Client) *WatchStorageV1VolumeAttachmentListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowWatchBookmarks adds the allowWatchBookmarks to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithAllowWatchBookmarks(allowWatchBookmarks *bool) *WatchStorageV1VolumeAttachmentListParams {
	o.SetAllowWatchBookmarks(allowWatchBookmarks)
	return o
}

// SetAllowWatchBookmarks adds the allowWatchBookmarks to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetAllowWatchBookmarks(allowWatchBookmarks *bool) {
	o.AllowWatchBookmarks = allowWatchBookmarks
}

// WithContinue adds the continueVar to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithContinue(continueVar *string) *WatchStorageV1VolumeAttachmentListParams {
	o.SetContinue(continueVar)
	return o
}

// SetContinue adds the continue to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetContinue(continueVar *string) {
	o.Continue = continueVar
}

// WithFieldSelector adds the fieldSelector to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithFieldSelector(fieldSelector *string) *WatchStorageV1VolumeAttachmentListParams {
	o.SetFieldSelector(fieldSelector)
	return o
}

// SetFieldSelector adds the fieldSelector to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetFieldSelector(fieldSelector *string) {
	o.FieldSelector = fieldSelector
}

// WithLabelSelector adds the labelSelector to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithLabelSelector(labelSelector *string) *WatchStorageV1VolumeAttachmentListParams {
	o.SetLabelSelector(labelSelector)
	return o
}

// SetLabelSelector adds the labelSelector to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetLabelSelector(labelSelector *string) {
	o.LabelSelector = labelSelector
}

// WithLimit adds the limit to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithLimit(limit *int64) *WatchStorageV1VolumeAttachmentListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPretty adds the pretty to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithPretty(pretty *string) *WatchStorageV1VolumeAttachmentListParams {
	o.SetPretty(pretty)
	return o
}

// SetPretty adds the pretty to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetPretty(pretty *string) {
	o.Pretty = pretty
}

// WithResourceVersion adds the resourceVersion to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithResourceVersion(resourceVersion *string) *WatchStorageV1VolumeAttachmentListParams {
	o.SetResourceVersion(resourceVersion)
	return o
}

// SetResourceVersion adds the resourceVersion to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetResourceVersion(resourceVersion *string) {
	o.ResourceVersion = resourceVersion
}

// WithResourceVersionMatch adds the resourceVersionMatch to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithResourceVersionMatch(resourceVersionMatch *string) *WatchStorageV1VolumeAttachmentListParams {
	o.SetResourceVersionMatch(resourceVersionMatch)
	return o
}

// SetResourceVersionMatch adds the resourceVersionMatch to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetResourceVersionMatch(resourceVersionMatch *string) {
	o.ResourceVersionMatch = resourceVersionMatch
}

// WithTimeoutSeconds adds the timeoutSeconds to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithTimeoutSeconds(timeoutSeconds *int64) *WatchStorageV1VolumeAttachmentListParams {
	o.SetTimeoutSeconds(timeoutSeconds)
	return o
}

// SetTimeoutSeconds adds the timeoutSeconds to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetTimeoutSeconds(timeoutSeconds *int64) {
	o.TimeoutSeconds = timeoutSeconds
}

// WithWatch adds the watch to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) WithWatch(watch *bool) *WatchStorageV1VolumeAttachmentListParams {
	o.SetWatch(watch)
	return o
}

// SetWatch adds the watch to the watch storage v1 volume attachment list params
func (o *WatchStorageV1VolumeAttachmentListParams) SetWatch(watch *bool) {
	o.Watch = watch
}

// WriteToRequest writes these params to a swagger request
func (o *WatchStorageV1VolumeAttachmentListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowWatchBookmarks != nil {

		// query param allowWatchBookmarks
		var qrAllowWatchBookmarks bool
		if o.AllowWatchBookmarks != nil {
			qrAllowWatchBookmarks = *o.AllowWatchBookmarks
		}
		qAllowWatchBookmarks := swag.FormatBool(qrAllowWatchBookmarks)
		if qAllowWatchBookmarks != "" {
			if err := r.SetQueryParam("allowWatchBookmarks", qAllowWatchBookmarks); err != nil {
				return err
			}
		}

	}

	if o.Continue != nil {

		// query param continue
		var qrContinue string
		if o.Continue != nil {
			qrContinue = *o.Continue
		}
		qContinue := qrContinue
		if qContinue != "" {
			if err := r.SetQueryParam("continue", qContinue); err != nil {
				return err
			}
		}

	}

	if o.FieldSelector != nil {

		// query param fieldSelector
		var qrFieldSelector string
		if o.FieldSelector != nil {
			qrFieldSelector = *o.FieldSelector
		}
		qFieldSelector := qrFieldSelector
		if qFieldSelector != "" {
			if err := r.SetQueryParam("fieldSelector", qFieldSelector); err != nil {
				return err
			}
		}

	}

	if o.LabelSelector != nil {

		// query param labelSelector
		var qrLabelSelector string
		if o.LabelSelector != nil {
			qrLabelSelector = *o.LabelSelector
		}
		qLabelSelector := qrLabelSelector
		if qLabelSelector != "" {
			if err := r.SetQueryParam("labelSelector", qLabelSelector); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Pretty != nil {

		// query param pretty
		var qrPretty string
		if o.Pretty != nil {
			qrPretty = *o.Pretty
		}
		qPretty := qrPretty
		if qPretty != "" {
			if err := r.SetQueryParam("pretty", qPretty); err != nil {
				return err
			}
		}

	}

	if o.ResourceVersion != nil {

		// query param resourceVersion
		var qrResourceVersion string
		if o.ResourceVersion != nil {
			qrResourceVersion = *o.ResourceVersion
		}
		qResourceVersion := qrResourceVersion
		if qResourceVersion != "" {
			if err := r.SetQueryParam("resourceVersion", qResourceVersion); err != nil {
				return err
			}
		}

	}

	if o.ResourceVersionMatch != nil {

		// query param resourceVersionMatch
		var qrResourceVersionMatch string
		if o.ResourceVersionMatch != nil {
			qrResourceVersionMatch = *o.ResourceVersionMatch
		}
		qResourceVersionMatch := qrResourceVersionMatch
		if qResourceVersionMatch != "" {
			if err := r.SetQueryParam("resourceVersionMatch", qResourceVersionMatch); err != nil {
				return err
			}
		}

	}

	if o.TimeoutSeconds != nil {

		// query param timeoutSeconds
		var qrTimeoutSeconds int64
		if o.TimeoutSeconds != nil {
			qrTimeoutSeconds = *o.TimeoutSeconds
		}
		qTimeoutSeconds := swag.FormatInt64(qrTimeoutSeconds)
		if qTimeoutSeconds != "" {
			if err := r.SetQueryParam("timeoutSeconds", qTimeoutSeconds); err != nil {
				return err
			}
		}

	}

	if o.Watch != nil {

		// query param watch
		var qrWatch bool
		if o.Watch != nil {
			qrWatch = *o.Watch
		}
		qWatch := swag.FormatBool(qrWatch)
		if qWatch != "" {
			if err := r.SetQueryParam("watch", qWatch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
