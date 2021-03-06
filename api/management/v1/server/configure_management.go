// This file is safe to edit. Once it exists it will not be overwritten

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cilium/kube-apate/api/management/v1/server/restapi"
	"github.com/cilium/kube-apate/api/management/v1/server/restapi/cilium"
)

//go:generate swagger generate server --target ../../v1 --name Management --spec ../management-swagger.yaml --api-package restapi --server-package server --principal interface{}

func configureFlags(api *restapi.ManagementAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *restapi.ManagementAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CiliumPostManagementCiliumIoV2CiliumEndpointsHandler == nil {
		api.CiliumPostManagementCiliumIoV2CiliumEndpointsHandler = cilium.PostManagementCiliumIoV2CiliumEndpointsHandlerFunc(func(params cilium.PostManagementCiliumIoV2CiliumEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PostManagementCiliumIoV2CiliumEndpoints has not yet been implemented")
		})
	}
	if api.CiliumPostManagementCiliumIoV2CiliumIdentitiesHandler == nil {
		api.CiliumPostManagementCiliumIoV2CiliumIdentitiesHandler = cilium.PostManagementCiliumIoV2CiliumIdentitiesHandlerFunc(func(params cilium.PostManagementCiliumIoV2CiliumIdentitiesParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PostManagementCiliumIoV2CiliumIdentities has not yet been implemented")
		})
	}
	if api.CiliumPostManagementCiliumIoV2CiliumNodesHandler == nil {
		api.CiliumPostManagementCiliumIoV2CiliumNodesHandler = cilium.PostManagementCiliumIoV2CiliumNodesHandlerFunc(func(params cilium.PostManagementCiliumIoV2CiliumNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PostManagementCiliumIoV2CiliumNodes has not yet been implemented")
		})
	}
	if api.CiliumPostManagementKubernetesIoV1PodsHandler == nil {
		api.CiliumPostManagementKubernetesIoV1PodsHandler = cilium.PostManagementKubernetesIoV1PodsHandlerFunc(func(params cilium.PostManagementKubernetesIoV1PodsParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PostManagementKubernetesIoV1Pods has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
