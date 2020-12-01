// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package apiregistration_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListApiregistrationV1APIServiceParams creates a new ListApiregistrationV1APIServiceParams object
// no default values defined in spec.
func NewListApiregistrationV1APIServiceParams() ListApiregistrationV1APIServiceParams {

	return ListApiregistrationV1APIServiceParams{}
}

// ListApiregistrationV1APIServiceParams contains all the bound params for the list apiregistration v1 API service operation
// typically these are obtained from a http.Request
//
// swagger:parameters listApiregistrationV1APIService
type ListApiregistrationV1APIServiceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*allowWatchBookmarks requests watch events with type "BOOKMARK". Servers that do not implement bookmarks may ignore this flag and bookmarks are sent at the server's discretion. Clients should not assume bookmarks are returned at any specific interval, nor may they assume the server will send any BOOKMARK event during a session. If this is not a watch, this field is ignored. If the feature gate WatchBookmarks is not enabled in apiserver, this field is ignored.
	  Unique: true
	  In: query
	*/
	AllowWatchBookmarks *bool
	/*The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server, the server will respond with a 410 ResourceExpired error together with a continue token. If the client needs a consistent list, it must restart their list without the continue field. Otherwise, the client may send another list request with the token received with the 410 error, the server will respond with a list starting from the next key, but from the latest snapshot, which is inconsistent from the previous list results - objects that are created, modified, or deleted after the first list request will be included in the response, as long as their keys are after the "next key".

	This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications.
	  Unique: true
	  In: query
	*/
	Continue *string
	/*A selector to restrict the list of returned objects by their fields. Defaults to everything.
	  Unique: true
	  In: query
	*/
	FieldSelector *string
	/*A selector to restrict the list of returned objects by their labels. Defaults to everything.
	  Unique: true
	  In: query
	*/
	LabelSelector *string
	/*limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.

	The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned.
	  Unique: true
	  In: query
	*/
	Limit *int64
	/*If 'true', then the output is pretty printed.
	  Unique: true
	  In: query
	*/
	Pretty *string
	/*resourceVersion sets a constraint on what resource versions a request may be served from. See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset
	  Unique: true
	  In: query
	*/
	ResourceVersion *string
	/*resourceVersionMatch determines how resourceVersion is applied to list calls. It is highly recommended that resourceVersionMatch be set for list calls where resourceVersion is set See https://kubernetes.io/docs/reference/using-api/api-concepts/#resource-versions for details.

	Defaults to unset
	  Unique: true
	  In: query
	*/
	ResourceVersionMatch *string
	/*Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity.
	  Unique: true
	  In: query
	*/
	TimeoutSeconds *int64
	/*Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.
	  Unique: true
	  In: query
	*/
	Watch *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewListApiregistrationV1APIServiceParams() beforehand.
func (o *ListApiregistrationV1APIServiceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qAllowWatchBookmarks, qhkAllowWatchBookmarks, _ := qs.GetOK("allowWatchBookmarks")
	if err := o.bindAllowWatchBookmarks(qAllowWatchBookmarks, qhkAllowWatchBookmarks, route.Formats); err != nil {
		res = append(res, err)
	}

	qContinue, qhkContinue, _ := qs.GetOK("continue")
	if err := o.bindContinue(qContinue, qhkContinue, route.Formats); err != nil {
		res = append(res, err)
	}

	qFieldSelector, qhkFieldSelector, _ := qs.GetOK("fieldSelector")
	if err := o.bindFieldSelector(qFieldSelector, qhkFieldSelector, route.Formats); err != nil {
		res = append(res, err)
	}

	qLabelSelector, qhkLabelSelector, _ := qs.GetOK("labelSelector")
	if err := o.bindLabelSelector(qLabelSelector, qhkLabelSelector, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qPretty, qhkPretty, _ := qs.GetOK("pretty")
	if err := o.bindPretty(qPretty, qhkPretty, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceVersion, qhkResourceVersion, _ := qs.GetOK("resourceVersion")
	if err := o.bindResourceVersion(qResourceVersion, qhkResourceVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	qResourceVersionMatch, qhkResourceVersionMatch, _ := qs.GetOK("resourceVersionMatch")
	if err := o.bindResourceVersionMatch(qResourceVersionMatch, qhkResourceVersionMatch, route.Formats); err != nil {
		res = append(res, err)
	}

	qTimeoutSeconds, qhkTimeoutSeconds, _ := qs.GetOK("timeoutSeconds")
	if err := o.bindTimeoutSeconds(qTimeoutSeconds, qhkTimeoutSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	qWatch, qhkWatch, _ := qs.GetOK("watch")
	if err := o.bindWatch(qWatch, qhkWatch, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAllowWatchBookmarks binds and validates parameter AllowWatchBookmarks from query.
func (o *ListApiregistrationV1APIServiceParams) bindAllowWatchBookmarks(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("allowWatchBookmarks", "query", "bool", raw)
	}
	o.AllowWatchBookmarks = &value

	if err := o.validateAllowWatchBookmarks(formats); err != nil {
		return err
	}

	return nil
}

// validateAllowWatchBookmarks carries on validations for parameter AllowWatchBookmarks
func (o *ListApiregistrationV1APIServiceParams) validateAllowWatchBookmarks(formats strfmt.Registry) error {

	return nil
}

// bindContinue binds and validates parameter Continue from query.
func (o *ListApiregistrationV1APIServiceParams) bindContinue(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Continue = &raw

	if err := o.validateContinue(formats); err != nil {
		return err
	}

	return nil
}

// validateContinue carries on validations for parameter Continue
func (o *ListApiregistrationV1APIServiceParams) validateContinue(formats strfmt.Registry) error {

	return nil
}

// bindFieldSelector binds and validates parameter FieldSelector from query.
func (o *ListApiregistrationV1APIServiceParams) bindFieldSelector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FieldSelector = &raw

	if err := o.validateFieldSelector(formats); err != nil {
		return err
	}

	return nil
}

// validateFieldSelector carries on validations for parameter FieldSelector
func (o *ListApiregistrationV1APIServiceParams) validateFieldSelector(formats strfmt.Registry) error {

	return nil
}

// bindLabelSelector binds and validates parameter LabelSelector from query.
func (o *ListApiregistrationV1APIServiceParams) bindLabelSelector(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.LabelSelector = &raw

	if err := o.validateLabelSelector(formats); err != nil {
		return err
	}

	return nil
}

// validateLabelSelector carries on validations for parameter LabelSelector
func (o *ListApiregistrationV1APIServiceParams) validateLabelSelector(formats strfmt.Registry) error {

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *ListApiregistrationV1APIServiceParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	if err := o.validateLimit(formats); err != nil {
		return err
	}

	return nil
}

// validateLimit carries on validations for parameter Limit
func (o *ListApiregistrationV1APIServiceParams) validateLimit(formats strfmt.Registry) error {

	return nil
}

// bindPretty binds and validates parameter Pretty from query.
func (o *ListApiregistrationV1APIServiceParams) bindPretty(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Pretty = &raw

	if err := o.validatePretty(formats); err != nil {
		return err
	}

	return nil
}

// validatePretty carries on validations for parameter Pretty
func (o *ListApiregistrationV1APIServiceParams) validatePretty(formats strfmt.Registry) error {

	return nil
}

// bindResourceVersion binds and validates parameter ResourceVersion from query.
func (o *ListApiregistrationV1APIServiceParams) bindResourceVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ResourceVersion = &raw

	if err := o.validateResourceVersion(formats); err != nil {
		return err
	}

	return nil
}

// validateResourceVersion carries on validations for parameter ResourceVersion
func (o *ListApiregistrationV1APIServiceParams) validateResourceVersion(formats strfmt.Registry) error {

	return nil
}

// bindResourceVersionMatch binds and validates parameter ResourceVersionMatch from query.
func (o *ListApiregistrationV1APIServiceParams) bindResourceVersionMatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ResourceVersionMatch = &raw

	if err := o.validateResourceVersionMatch(formats); err != nil {
		return err
	}

	return nil
}

// validateResourceVersionMatch carries on validations for parameter ResourceVersionMatch
func (o *ListApiregistrationV1APIServiceParams) validateResourceVersionMatch(formats strfmt.Registry) error {

	return nil
}

// bindTimeoutSeconds binds and validates parameter TimeoutSeconds from query.
func (o *ListApiregistrationV1APIServiceParams) bindTimeoutSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("timeoutSeconds", "query", "int64", raw)
	}
	o.TimeoutSeconds = &value

	if err := o.validateTimeoutSeconds(formats); err != nil {
		return err
	}

	return nil
}

// validateTimeoutSeconds carries on validations for parameter TimeoutSeconds
func (o *ListApiregistrationV1APIServiceParams) validateTimeoutSeconds(formats strfmt.Registry) error {

	return nil
}

// bindWatch binds and validates parameter Watch from query.
func (o *ListApiregistrationV1APIServiceParams) bindWatch(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("watch", "query", "bool", raw)
	}
	o.Watch = &value

	if err := o.validateWatch(formats); err != nil {
		return err
	}

	return nil
}

// validateWatch carries on validations for parameter Watch
func (o *ListApiregistrationV1APIServiceParams) validateWatch(formats strfmt.Registry) error {

	return nil
}
