// This file is safe to edit. Once it exists it will not be overwritten

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"crypto/tls"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/yamlpc"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/admissionregistration"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/admissionregistration_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/admissionregistration_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiextensions"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiextensions_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiextensions_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiregistration"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiregistration_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiregistration_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apis"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apps"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apps_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/authentication"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/authentication_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/authentication_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/authorization"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/authorization_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/authorization_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/autoscaling"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/autoscaling_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/autoscaling_v2beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/autoscaling_v2beta2"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/batch"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/batch_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/batch_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/batch_v2alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/certificates"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/certificates_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/certificates_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/cilium"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/coordination"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/coordination_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/coordination_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/core"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/core_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/discovery"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/discovery_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/events"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/events_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/events_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/extensions"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/extensions_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/flowcontrol_apiserver"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/flowcontrol_apiserver_v1alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/logs"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/networking"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/networking_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/networking_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/node"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/node_v1alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/node_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/policy"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/policy_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/rbac_authorization"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/rbac_authorization_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/rbac_authorization_v1alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/rbac_authorization_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/scheduling"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/scheduling_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/scheduling_v1alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/scheduling_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/settings"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/settings_v1alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/storage"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/storage_v1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/storage_v1alpha1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/storage_v1beta1"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/version"
	"github.com/cilium/kube-apate/internal/healthz"
)

//go:generate swagger generate server --target ../../v1 --name Kubernetes --spec ../swagger.json --api-package restapi --server-package server --principal interface{}

func configureFlags(api *restapi.KubernetesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *restapi.KubernetesAPI) http.Handler {
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

	api.EmptyConsumer = runtime.ConsumerFunc(func(r io.Reader, target interface{}) error {
		return errors.NotImplemented(" consumer has not yet been implemented")
	})
	api.JSONConsumer = runtime.JSONConsumer()
	api.ProtobufConsumer = runtime.ConsumerFunc(func(r io.Reader, target interface{}) error {
		return errors.NotImplemented("protobuf consumer has not yet been implemented")
	})
	api.YamlConsumer = yamlpc.YAMLConsumer()

	api.EmptyProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented(" producer has not yet been implemented")
	})
	api.JSONProducer = runtime.JSONProducer()
	api.ProtobufProducer = runtime.ProducerFunc(func(w io.Writer, data interface{}) error {
		return errors.NotImplemented("protobuf producer has not yet been implemented")
	})
	api.TxtProducer = runtime.TextProducer()
	api.YamlProducer = yamlpc.YAMLProducer()

	if api.CiliumDeleteApisCiliumIoV2CiliumEndpointHandler == nil {
		api.CiliumDeleteApisCiliumIoV2CiliumEndpointHandler = cilium.DeleteApisCiliumIoV2CiliumEndpointHandlerFunc(func(params cilium.DeleteApisCiliumIoV2CiliumEndpointParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.DeleteApisCiliumIoV2CiliumEndpoint has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandler = cilium.ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumClusterwideLocalRedirectPolicy has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumClusterwideNetworkPolicyHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumClusterwideNetworkPolicyHandler = cilium.ListApisCiliumIoV2CiliumClusterwideNetworkPolicyHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumClusterwideNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumClusterwideNetworkPolicy has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumEndpointHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumEndpointHandler = cilium.ListApisCiliumIoV2CiliumEndpointHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumEndpointParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumEndpoint has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumIdentityHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumIdentityHandler = cilium.ListApisCiliumIoV2CiliumIdentityHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumIdentityParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumIdentity has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumLocalRedirectPolicyHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumLocalRedirectPolicyHandler = cilium.ListApisCiliumIoV2CiliumLocalRedirectPolicyHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumLocalRedirectPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumLocalRedirectPolicy has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumNetworkPolicyHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumNetworkPolicyHandler = cilium.ListApisCiliumIoV2CiliumNetworkPolicyHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumNetworkPolicy has not yet been implemented")
		})
	}
	if api.CiliumListApisCiliumIoV2CiliumNodesHandler == nil {
		api.CiliumListApisCiliumIoV2CiliumNodesHandler = cilium.ListApisCiliumIoV2CiliumNodesHandlerFunc(func(params cilium.ListApisCiliumIoV2CiliumNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ListApisCiliumIoV2CiliumNodes has not yet been implemented")
		})
	}
	if api.CiliumPatchApisCiliumIoV2CiliumEndpointStatusHandler == nil {
		api.CiliumPatchApisCiliumIoV2CiliumEndpointStatusHandler = cilium.PatchApisCiliumIoV2CiliumEndpointStatusHandlerFunc(func(params cilium.PatchApisCiliumIoV2CiliumEndpointStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PatchApisCiliumIoV2CiliumEndpointStatus has not yet been implemented")
		})
	}
	if api.CiliumPatchApisCiliumIoV2CiliumIdentityStatusHandler == nil {
		api.CiliumPatchApisCiliumIoV2CiliumIdentityStatusHandler = cilium.PatchApisCiliumIoV2CiliumIdentityStatusHandlerFunc(func(params cilium.PatchApisCiliumIoV2CiliumIdentityStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PatchApisCiliumIoV2CiliumIdentityStatus has not yet been implemented")
		})
	}
	if api.CiliumPostApisCiliumIoV2CiliumEndpointHandler == nil {
		api.CiliumPostApisCiliumIoV2CiliumEndpointHandler = cilium.PostApisCiliumIoV2CiliumEndpointHandlerFunc(func(params cilium.PostApisCiliumIoV2CiliumEndpointParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PostApisCiliumIoV2CiliumEndpoint has not yet been implemented")
		})
	}
	if api.CiliumPostApisCiliumIoV2CiliumNodesHandler == nil {
		api.CiliumPostApisCiliumIoV2CiliumNodesHandler = cilium.PostApisCiliumIoV2CiliumNodesHandlerFunc(func(params cilium.PostApisCiliumIoV2CiliumNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PostApisCiliumIoV2CiliumNodes has not yet been implemented")
		})
	}
	if api.CiliumPutApisCiliumIoV2CiliumNodesHandler == nil {
		api.CiliumPutApisCiliumIoV2CiliumNodesHandler = cilium.PutApisCiliumIoV2CiliumNodesHandlerFunc(func(params cilium.PutApisCiliumIoV2CiliumNodesParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.PutApisCiliumIoV2CiliumNodes has not yet been implemented")
		})
	}
	if api.CiliumReadApisCiliumIoV2CiliumNodeHandler == nil {
		api.CiliumReadApisCiliumIoV2CiliumNodeHandler = cilium.ReadApisCiliumIoV2CiliumNodeHandlerFunc(func(params cilium.ReadApisCiliumIoV2CiliumNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation cilium.ReadApisCiliumIoV2CiliumNode has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1DeleteNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1DeleteNamespacedPodProxyHandler = core_v1.ConnectCoreV1DeleteNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1DeleteNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1DeleteNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1DeleteNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1DeleteNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1DeleteNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1DeleteNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1DeleteNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1DeleteNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1DeleteNamespacedServiceProxyHandler = core_v1.ConnectCoreV1DeleteNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1DeleteNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1DeleteNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1DeleteNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1DeleteNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1DeleteNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1DeleteNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1DeleteNodeProxyHandler = core_v1.ConnectCoreV1DeleteNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1DeleteNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1DeleteNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1DeleteNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1DeleteNodeProxyWithPathHandler = core_v1.ConnectCoreV1DeleteNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1DeleteNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1DeleteNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedPodAttachHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedPodAttachHandler = core_v1.ConnectCoreV1GetNamespacedPodAttachHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedPodAttachParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedPodAttach has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedPodExecHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedPodExecHandler = core_v1.ConnectCoreV1GetNamespacedPodExecHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedPodExecParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedPodExec has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedPodPortforwardHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedPodPortforwardHandler = core_v1.ConnectCoreV1GetNamespacedPodPortforwardHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedPodPortforwardParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedPodPortforward has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedPodProxyHandler = core_v1.ConnectCoreV1GetNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1GetNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedServiceProxyHandler = core_v1.ConnectCoreV1GetNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1GetNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1GetNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1GetNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1GetNodeProxyHandler = core_v1.ConnectCoreV1GetNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1GetNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1GetNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1GetNodeProxyWithPathHandler = core_v1.ConnectCoreV1GetNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1GetNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1GetNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1HeadNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1HeadNamespacedPodProxyHandler = core_v1.ConnectCoreV1HeadNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1HeadNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1HeadNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1HeadNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1HeadNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1HeadNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1HeadNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1HeadNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1HeadNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1HeadNamespacedServiceProxyHandler = core_v1.ConnectCoreV1HeadNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1HeadNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1HeadNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1HeadNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1HeadNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1HeadNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1HeadNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1HeadNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1HeadNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1HeadNodeProxyHandler = core_v1.ConnectCoreV1HeadNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1HeadNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1HeadNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1HeadNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1HeadNodeProxyWithPathHandler = core_v1.ConnectCoreV1HeadNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1HeadNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1HeadNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1OptionsNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1OptionsNamespacedPodProxyHandler = core_v1.ConnectCoreV1OptionsNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1OptionsNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1OptionsNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1OptionsNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1OptionsNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1OptionsNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1OptionsNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1OptionsNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1OptionsNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1OptionsNamespacedServiceProxyHandler = core_v1.ConnectCoreV1OptionsNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1OptionsNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1OptionsNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1OptionsNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1OptionsNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1OptionsNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1OptionsNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1OptionsNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1OptionsNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1OptionsNodeProxyHandler = core_v1.ConnectCoreV1OptionsNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1OptionsNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1OptionsNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1OptionsNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1OptionsNodeProxyWithPathHandler = core_v1.ConnectCoreV1OptionsNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1OptionsNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1OptionsNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PatchNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1PatchNamespacedPodProxyHandler = core_v1.ConnectCoreV1PatchNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1PatchNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PatchNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PatchNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PatchNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1PatchNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PatchNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PatchNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PatchNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1PatchNamespacedServiceProxyHandler = core_v1.ConnectCoreV1PatchNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1PatchNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PatchNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PatchNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PatchNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1PatchNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PatchNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PatchNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PatchNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1PatchNodeProxyHandler = core_v1.ConnectCoreV1PatchNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1PatchNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PatchNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PatchNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PatchNodeProxyWithPathHandler = core_v1.ConnectCoreV1PatchNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PatchNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PatchNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedPodAttachHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedPodAttachHandler = core_v1.ConnectCoreV1PostNamespacedPodAttachHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedPodAttachParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedPodAttach has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedPodExecHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedPodExecHandler = core_v1.ConnectCoreV1PostNamespacedPodExecHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedPodExecParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedPodExec has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedPodPortforwardHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedPodPortforwardHandler = core_v1.ConnectCoreV1PostNamespacedPodPortforwardHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedPodPortforwardParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedPodPortforward has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedPodProxyHandler = core_v1.ConnectCoreV1PostNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1PostNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedServiceProxyHandler = core_v1.ConnectCoreV1PostNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PostNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1PostNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PostNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1PostNodeProxyHandler = core_v1.ConnectCoreV1PostNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1PostNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PostNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PostNodeProxyWithPathHandler = core_v1.ConnectCoreV1PostNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PostNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PostNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PutNamespacedPodProxyHandler == nil {
		api.CoreV1ConnectCoreV1PutNamespacedPodProxyHandler = core_v1.ConnectCoreV1PutNamespacedPodProxyHandlerFunc(func(params core_v1.ConnectCoreV1PutNamespacedPodProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PutNamespacedPodProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PutNamespacedPodProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PutNamespacedPodProxyWithPathHandler = core_v1.ConnectCoreV1PutNamespacedPodProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PutNamespacedPodProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PutNamespacedPodProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PutNamespacedServiceProxyHandler == nil {
		api.CoreV1ConnectCoreV1PutNamespacedServiceProxyHandler = core_v1.ConnectCoreV1PutNamespacedServiceProxyHandlerFunc(func(params core_v1.ConnectCoreV1PutNamespacedServiceProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PutNamespacedServiceProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PutNamespacedServiceProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PutNamespacedServiceProxyWithPathHandler = core_v1.ConnectCoreV1PutNamespacedServiceProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PutNamespacedServiceProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PutNamespacedServiceProxyWithPath has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PutNodeProxyHandler == nil {
		api.CoreV1ConnectCoreV1PutNodeProxyHandler = core_v1.ConnectCoreV1PutNodeProxyHandlerFunc(func(params core_v1.ConnectCoreV1PutNodeProxyParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PutNodeProxy has not yet been implemented")
		})
	}
	if api.CoreV1ConnectCoreV1PutNodeProxyWithPathHandler == nil {
		api.CoreV1ConnectCoreV1PutNodeProxyWithPathHandler = core_v1.ConnectCoreV1PutNodeProxyWithPathHandlerFunc(func(params core_v1.ConnectCoreV1PutNodeProxyWithPathParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ConnectCoreV1PutNodeProxyWithPath has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1CreateAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1CreateAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.CreateAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.CreateAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.CreateAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1CreateAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1CreateAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.CreateAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.CreateAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.CreateAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.CreateAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.CreateAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.CreateAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.CreateAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.ApiextensionsV1CreateApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1CreateApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.CreateApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.CreateApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.CreateApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1CreateApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1CreateApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.CreateApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.CreateApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.CreateApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiregistrationV1CreateApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1CreateApiregistrationV1APIServiceHandler = apiregistration_v1.CreateApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.CreateApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.CreateApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1CreateApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1CreateApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.CreateApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.CreateApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.CreateApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.AppsV1CreateAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1CreateAppsV1NamespacedControllerRevisionHandler = apps_v1.CreateAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.CreateAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.CreateAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1CreateAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1CreateAppsV1NamespacedDaemonSetHandler = apps_v1.CreateAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.CreateAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.CreateAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1CreateAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1CreateAppsV1NamespacedDeploymentHandler = apps_v1.CreateAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.CreateAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.CreateAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1CreateAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1CreateAppsV1NamespacedReplicaSetHandler = apps_v1.CreateAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.CreateAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.CreateAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1CreateAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1CreateAppsV1NamespacedStatefulSetHandler = apps_v1.CreateAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.CreateAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.CreateAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AuthenticationV1CreateAuthenticationV1TokenReviewHandler == nil {
		api.AuthenticationV1CreateAuthenticationV1TokenReviewHandler = authentication_v1.CreateAuthenticationV1TokenReviewHandlerFunc(func(params authentication_v1.CreateAuthenticationV1TokenReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication_v1.CreateAuthenticationV1TokenReview has not yet been implemented")
		})
	}
	if api.AuthenticationV1beta1CreateAuthenticationV1beta1TokenReviewHandler == nil {
		api.AuthenticationV1beta1CreateAuthenticationV1beta1TokenReviewHandler = authentication_v1beta1.CreateAuthenticationV1beta1TokenReviewHandlerFunc(func(params authentication_v1beta1.CreateAuthenticationV1beta1TokenReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication_v1beta1.CreateAuthenticationV1beta1TokenReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1CreateAuthorizationV1NamespacedLocalSubjectAccessReviewHandler == nil {
		api.AuthorizationV1CreateAuthorizationV1NamespacedLocalSubjectAccessReviewHandler = authorization_v1.CreateAuthorizationV1NamespacedLocalSubjectAccessReviewHandlerFunc(func(params authorization_v1.CreateAuthorizationV1NamespacedLocalSubjectAccessReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1.CreateAuthorizationV1NamespacedLocalSubjectAccessReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1CreateAuthorizationV1SelfSubjectAccessReviewHandler == nil {
		api.AuthorizationV1CreateAuthorizationV1SelfSubjectAccessReviewHandler = authorization_v1.CreateAuthorizationV1SelfSubjectAccessReviewHandlerFunc(func(params authorization_v1.CreateAuthorizationV1SelfSubjectAccessReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1.CreateAuthorizationV1SelfSubjectAccessReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1CreateAuthorizationV1SelfSubjectRulesReviewHandler == nil {
		api.AuthorizationV1CreateAuthorizationV1SelfSubjectRulesReviewHandler = authorization_v1.CreateAuthorizationV1SelfSubjectRulesReviewHandlerFunc(func(params authorization_v1.CreateAuthorizationV1SelfSubjectRulesReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1.CreateAuthorizationV1SelfSubjectRulesReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1CreateAuthorizationV1SubjectAccessReviewHandler == nil {
		api.AuthorizationV1CreateAuthorizationV1SubjectAccessReviewHandler = authorization_v1.CreateAuthorizationV1SubjectAccessReviewHandlerFunc(func(params authorization_v1.CreateAuthorizationV1SubjectAccessReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1.CreateAuthorizationV1SubjectAccessReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1beta1CreateAuthorizationV1beta1NamespacedLocalSubjectAccessReviewHandler == nil {
		api.AuthorizationV1beta1CreateAuthorizationV1beta1NamespacedLocalSubjectAccessReviewHandler = authorization_v1beta1.CreateAuthorizationV1beta1NamespacedLocalSubjectAccessReviewHandlerFunc(func(params authorization_v1beta1.CreateAuthorizationV1beta1NamespacedLocalSubjectAccessReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1beta1.CreateAuthorizationV1beta1NamespacedLocalSubjectAccessReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1beta1CreateAuthorizationV1beta1SelfSubjectAccessReviewHandler == nil {
		api.AuthorizationV1beta1CreateAuthorizationV1beta1SelfSubjectAccessReviewHandler = authorization_v1beta1.CreateAuthorizationV1beta1SelfSubjectAccessReviewHandlerFunc(func(params authorization_v1beta1.CreateAuthorizationV1beta1SelfSubjectAccessReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1beta1.CreateAuthorizationV1beta1SelfSubjectAccessReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1beta1CreateAuthorizationV1beta1SelfSubjectRulesReviewHandler == nil {
		api.AuthorizationV1beta1CreateAuthorizationV1beta1SelfSubjectRulesReviewHandler = authorization_v1beta1.CreateAuthorizationV1beta1SelfSubjectRulesReviewHandlerFunc(func(params authorization_v1beta1.CreateAuthorizationV1beta1SelfSubjectRulesReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1beta1.CreateAuthorizationV1beta1SelfSubjectRulesReview has not yet been implemented")
		})
	}
	if api.AuthorizationV1beta1CreateAuthorizationV1beta1SubjectAccessReviewHandler == nil {
		api.AuthorizationV1beta1CreateAuthorizationV1beta1SubjectAccessReviewHandler = authorization_v1beta1.CreateAuthorizationV1beta1SubjectAccessReviewHandlerFunc(func(params authorization_v1beta1.CreateAuthorizationV1beta1SubjectAccessReviewParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1beta1.CreateAuthorizationV1beta1SubjectAccessReview has not yet been implemented")
		})
	}
	if api.AutoscalingV1CreateAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1CreateAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.CreateAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.CreateAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.CreateAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.CreateAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.CreateAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.BatchV1CreateBatchV1NamespacedJobHandler == nil {
		api.BatchV1CreateBatchV1NamespacedJobHandler = batch_v1.CreateBatchV1NamespacedJobHandlerFunc(func(params batch_v1.CreateBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.CreateBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1CreateBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1CreateBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.CreateBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.CreateBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.CreateBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1CreateBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1CreateBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.CreateBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.CreateBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.CreateBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.CertificatesV1CreateCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1CreateCertificatesV1CertificateSigningRequestHandler = certificates_v1.CreateCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.CreateCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.CreateCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1CreateCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1CreateCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.CreateCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.CreateCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.CreateCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CoordinationV1CreateCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1CreateCoordinationV1NamespacedLeaseHandler = coordination_v1.CreateCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.CreateCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.CreateCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1CreateCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1CreateCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.CreateCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.CreateCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.CreateCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespaceHandler == nil {
		api.CoreV1CreateCoreV1NamespaceHandler = core_v1.CreateCoreV1NamespaceHandlerFunc(func(params core_v1.CreateCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedBindingHandler == nil {
		api.CoreV1CreateCoreV1NamespacedBindingHandler = core_v1.CreateCoreV1NamespacedBindingHandlerFunc(func(params core_v1.CreateCoreV1NamespacedBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedBinding has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1CreateCoreV1NamespacedConfigMapHandler = core_v1.CreateCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.CreateCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1CreateCoreV1NamespacedEndpointsHandler = core_v1.CreateCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.CreateCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedEventHandler == nil {
		api.CoreV1CreateCoreV1NamespacedEventHandler = core_v1.CreateCoreV1NamespacedEventHandlerFunc(func(params core_v1.CreateCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1CreateCoreV1NamespacedLimitRangeHandler = core_v1.CreateCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.CreateCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1CreateCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.CreateCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.CreateCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedPodHandler == nil {
		api.CoreV1CreateCoreV1NamespacedPodHandler = core_v1.CreateCoreV1NamespacedPodHandlerFunc(func(params core_v1.CreateCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedPodBindingHandler == nil {
		api.CoreV1CreateCoreV1NamespacedPodBindingHandler = core_v1.CreateCoreV1NamespacedPodBindingHandlerFunc(func(params core_v1.CreateCoreV1NamespacedPodBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedPodBinding has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedPodEvictionHandler == nil {
		api.CoreV1CreateCoreV1NamespacedPodEvictionHandler = core_v1.CreateCoreV1NamespacedPodEvictionHandlerFunc(func(params core_v1.CreateCoreV1NamespacedPodEvictionParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedPodEviction has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1CreateCoreV1NamespacedPodTemplateHandler = core_v1.CreateCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.CreateCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1CreateCoreV1NamespacedReplicationControllerHandler = core_v1.CreateCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.CreateCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1CreateCoreV1NamespacedResourceQuotaHandler = core_v1.CreateCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.CreateCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedSecretHandler == nil {
		api.CoreV1CreateCoreV1NamespacedSecretHandler = core_v1.CreateCoreV1NamespacedSecretHandlerFunc(func(params core_v1.CreateCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedServiceHandler == nil {
		api.CoreV1CreateCoreV1NamespacedServiceHandler = core_v1.CreateCoreV1NamespacedServiceHandlerFunc(func(params core_v1.CreateCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1CreateCoreV1NamespacedServiceAccountHandler = core_v1.CreateCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.CreateCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NamespacedServiceAccountTokenHandler == nil {
		api.CoreV1CreateCoreV1NamespacedServiceAccountTokenHandler = core_v1.CreateCoreV1NamespacedServiceAccountTokenHandlerFunc(func(params core_v1.CreateCoreV1NamespacedServiceAccountTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1NamespacedServiceAccountToken has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1NodeHandler == nil {
		api.CoreV1CreateCoreV1NodeHandler = core_v1.CreateCoreV1NodeHandlerFunc(func(params core_v1.CreateCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1CreateCoreV1PersistentVolumeHandler == nil {
		api.CoreV1CreateCoreV1PersistentVolumeHandler = core_v1.CreateCoreV1PersistentVolumeHandlerFunc(func(params core_v1.CreateCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.CreateCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1CreateDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1CreateDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.CreateDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.CreateDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.CreateDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.EventsV1CreateEventsV1NamespacedEventHandler == nil {
		api.EventsV1CreateEventsV1NamespacedEventHandler = events_v1.CreateEventsV1NamespacedEventHandlerFunc(func(params events_v1.CreateEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.CreateEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1CreateEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1CreateEventsV1beta1NamespacedEventHandler = events_v1beta1.CreateEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.CreateEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.CreateEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1CreateExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1CreateExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.CreateExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.CreateExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.CreateExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1CreateFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1CreateFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.CreateFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.CreateFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.CreateFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1CreateFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1CreateFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.CreateFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.CreateFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.CreateFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.NetworkingV1CreateNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1CreateNetworkingV1IngressClassHandler = networking_v1.CreateNetworkingV1IngressClassHandlerFunc(func(params networking_v1.CreateNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.CreateNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1CreateNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1CreateNetworkingV1NamespacedIngressHandler = networking_v1.CreateNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.CreateNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.CreateNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1CreateNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1CreateNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.CreateNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.CreateNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.CreateNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1CreateNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1CreateNetworkingV1beta1IngressClassHandler = networking_v1beta1.CreateNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.CreateNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.CreateNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1CreateNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1CreateNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.CreateNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.CreateNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.CreateNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NodeV1alpha1CreateNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1CreateNodeV1alpha1RuntimeClassHandler = node_v1alpha1.CreateNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.CreateNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.CreateNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1CreateNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1CreateNodeV1beta1RuntimeClassHandler = node_v1beta1.CreateNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.CreateNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.CreateNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.PolicyV1beta1CreatePolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1CreatePolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.CreatePolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.CreatePolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.CreatePolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1CreatePolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1CreatePolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.CreatePolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.CreatePolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.CreatePolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1CreateRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1CreateRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.CreateRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.CreateRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.CreateRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1CreateRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1CreateRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.CreateRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.CreateRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.CreateRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1CreateRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1CreateRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.CreateRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.CreateRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.CreateRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1CreateRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1CreateRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.CreateRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.CreateRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.CreateRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.CreateRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1CreateRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.CreateRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.SchedulingV1CreateSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1CreateSchedulingV1PriorityClassHandler = scheduling_v1.CreateSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.CreateSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.CreateSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1CreateSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1CreateSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.CreateSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.CreateSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.CreateSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1CreateSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1CreateSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.CreateSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.CreateSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.CreateSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1CreateSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1CreateSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.CreateSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.CreateSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.CreateSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.StorageV1CreateStorageV1CSIDriverHandler == nil {
		api.StorageV1CreateStorageV1CSIDriverHandler = storage_v1.CreateStorageV1CSIDriverHandlerFunc(func(params storage_v1.CreateStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.CreateStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1CreateStorageV1CSINodeHandler == nil {
		api.StorageV1CreateStorageV1CSINodeHandler = storage_v1.CreateStorageV1CSINodeHandlerFunc(func(params storage_v1.CreateStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.CreateStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1CreateStorageV1StorageClassHandler == nil {
		api.StorageV1CreateStorageV1StorageClassHandler = storage_v1.CreateStorageV1StorageClassHandlerFunc(func(params storage_v1.CreateStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.CreateStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1CreateStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1CreateStorageV1VolumeAttachmentHandler = storage_v1.CreateStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.CreateStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.CreateStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1alpha1CreateStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1CreateStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.CreateStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.CreateStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.CreateStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1CreateStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1CreateStorageV1beta1CSIDriverHandler = storage_v1beta1.CreateStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.CreateStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.CreateStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1CreateStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1CreateStorageV1beta1CSINodeHandler = storage_v1beta1.CreateStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.CreateStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.CreateStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1CreateStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1CreateStorageV1beta1StorageClassHandler = storage_v1beta1.CreateStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.CreateStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.CreateStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1CreateStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1CreateStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.CreateStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.CreateStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.CreateStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1DeleteAdmissionregistrationV1CollectionMutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1DeleteAdmissionregistrationV1CollectionMutatingWebhookConfigurationHandler = admissionregistration_v1.DeleteAdmissionregistrationV1CollectionMutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.DeleteAdmissionregistrationV1CollectionMutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.DeleteAdmissionregistrationV1CollectionMutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationHandler = admissionregistration_v1.DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.DeleteAdmissionregistrationV1CollectionValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.DeleteAdmissionregistrationV1CollectionValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.DeleteAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.DeleteAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.DeleteAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.DeleteAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.DeleteAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.DeleteAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1CollectionMutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1CollectionMutatingWebhookConfigurationHandler = admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1CollectionMutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1CollectionMutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1CollectionMutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1CollectionValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1DeleteAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.DeleteAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.ApiextensionsV1DeleteApiextensionsV1CollectionCustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1DeleteApiextensionsV1CollectionCustomResourceDefinitionHandler = apiextensions_v1.DeleteApiextensionsV1CollectionCustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.DeleteApiextensionsV1CollectionCustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.DeleteApiextensionsV1CollectionCustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1DeleteApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1DeleteApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.DeleteApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.DeleteApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.DeleteApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionHandler = apiextensions_v1beta1.DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.DeleteApiextensionsV1beta1CollectionCustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.DeleteApiextensionsV1beta1CollectionCustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1DeleteApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1DeleteApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.DeleteApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.DeleteApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.DeleteApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiregistrationV1DeleteApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1DeleteApiregistrationV1APIServiceHandler = apiregistration_v1.DeleteApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.DeleteApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.DeleteApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1DeleteApiregistrationV1CollectionAPIServiceHandler == nil {
		api.ApiregistrationV1DeleteApiregistrationV1CollectionAPIServiceHandler = apiregistration_v1.DeleteApiregistrationV1CollectionAPIServiceHandlerFunc(func(params apiregistration_v1.DeleteApiregistrationV1CollectionAPIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.DeleteApiregistrationV1CollectionAPIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1DeleteApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1DeleteApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.DeleteApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.DeleteApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.DeleteApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1DeleteApiregistrationV1beta1CollectionAPIServiceHandler == nil {
		api.ApiregistrationV1beta1DeleteApiregistrationV1beta1CollectionAPIServiceHandler = apiregistration_v1beta1.DeleteApiregistrationV1beta1CollectionAPIServiceHandlerFunc(func(params apiregistration_v1beta1.DeleteApiregistrationV1beta1CollectionAPIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.DeleteApiregistrationV1beta1CollectionAPIService has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1CollectionNamespacedControllerRevisionHandler == nil {
		api.AppsV1DeleteAppsV1CollectionNamespacedControllerRevisionHandler = apps_v1.DeleteAppsV1CollectionNamespacedControllerRevisionHandlerFunc(func(params apps_v1.DeleteAppsV1CollectionNamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1CollectionNamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1CollectionNamespacedDaemonSetHandler == nil {
		api.AppsV1DeleteAppsV1CollectionNamespacedDaemonSetHandler = apps_v1.DeleteAppsV1CollectionNamespacedDaemonSetHandlerFunc(func(params apps_v1.DeleteAppsV1CollectionNamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1CollectionNamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1CollectionNamespacedDeploymentHandler == nil {
		api.AppsV1DeleteAppsV1CollectionNamespacedDeploymentHandler = apps_v1.DeleteAppsV1CollectionNamespacedDeploymentHandlerFunc(func(params apps_v1.DeleteAppsV1CollectionNamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1CollectionNamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1CollectionNamespacedReplicaSetHandler == nil {
		api.AppsV1DeleteAppsV1CollectionNamespacedReplicaSetHandler = apps_v1.DeleteAppsV1CollectionNamespacedReplicaSetHandlerFunc(func(params apps_v1.DeleteAppsV1CollectionNamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1CollectionNamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1CollectionNamespacedStatefulSetHandler == nil {
		api.AppsV1DeleteAppsV1CollectionNamespacedStatefulSetHandler = apps_v1.DeleteAppsV1CollectionNamespacedStatefulSetHandlerFunc(func(params apps_v1.DeleteAppsV1CollectionNamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1CollectionNamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1DeleteAppsV1NamespacedControllerRevisionHandler = apps_v1.DeleteAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.DeleteAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1DeleteAppsV1NamespacedDaemonSetHandler = apps_v1.DeleteAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.DeleteAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1DeleteAppsV1NamespacedDeploymentHandler = apps_v1.DeleteAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.DeleteAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1DeleteAppsV1NamespacedReplicaSetHandler = apps_v1.DeleteAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.DeleteAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1DeleteAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1DeleteAppsV1NamespacedStatefulSetHandler = apps_v1.DeleteAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.DeleteAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.DeleteAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AutoscalingV1DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.DeleteAutoscalingV1CollectionNamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV1DeleteAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1DeleteAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.DeleteAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.DeleteAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.DeleteAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.DeleteAutoscalingV2beta1CollectionNamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.DeleteAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.DeleteAutoscalingV2beta2CollectionNamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.DeleteAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.BatchV1DeleteBatchV1CollectionNamespacedJobHandler == nil {
		api.BatchV1DeleteBatchV1CollectionNamespacedJobHandler = batch_v1.DeleteBatchV1CollectionNamespacedJobHandlerFunc(func(params batch_v1.DeleteBatchV1CollectionNamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.DeleteBatchV1CollectionNamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1DeleteBatchV1NamespacedJobHandler == nil {
		api.BatchV1DeleteBatchV1NamespacedJobHandler = batch_v1.DeleteBatchV1NamespacedJobHandlerFunc(func(params batch_v1.DeleteBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.DeleteBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1DeleteBatchV1beta1CollectionNamespacedCronJobHandler == nil {
		api.BatchV1beta1DeleteBatchV1beta1CollectionNamespacedCronJobHandler = batch_v1beta1.DeleteBatchV1beta1CollectionNamespacedCronJobHandlerFunc(func(params batch_v1beta1.DeleteBatchV1beta1CollectionNamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.DeleteBatchV1beta1CollectionNamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1DeleteBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1DeleteBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.DeleteBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.DeleteBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.DeleteBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1DeleteBatchV2alpha1CollectionNamespacedCronJobHandler == nil {
		api.BatchV2alpha1DeleteBatchV2alpha1CollectionNamespacedCronJobHandler = batch_v2alpha1.DeleteBatchV2alpha1CollectionNamespacedCronJobHandlerFunc(func(params batch_v2alpha1.DeleteBatchV2alpha1CollectionNamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.DeleteBatchV2alpha1CollectionNamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1DeleteBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1DeleteBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.DeleteBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.DeleteBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.DeleteBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.CertificatesV1DeleteCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1DeleteCertificatesV1CertificateSigningRequestHandler = certificates_v1.DeleteCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.DeleteCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.DeleteCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1DeleteCertificatesV1CollectionCertificateSigningRequestHandler == nil {
		api.CertificatesV1DeleteCertificatesV1CollectionCertificateSigningRequestHandler = certificates_v1.DeleteCertificatesV1CollectionCertificateSigningRequestHandlerFunc(func(params certificates_v1.DeleteCertificatesV1CollectionCertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.DeleteCertificatesV1CollectionCertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1DeleteCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1DeleteCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.DeleteCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.DeleteCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.DeleteCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1DeleteCertificatesV1beta1CollectionCertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1DeleteCertificatesV1beta1CollectionCertificateSigningRequestHandler = certificates_v1beta1.DeleteCertificatesV1beta1CollectionCertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.DeleteCertificatesV1beta1CollectionCertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.DeleteCertificatesV1beta1CollectionCertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CoordinationV1DeleteCoordinationV1CollectionNamespacedLeaseHandler == nil {
		api.CoordinationV1DeleteCoordinationV1CollectionNamespacedLeaseHandler = coordination_v1.DeleteCoordinationV1CollectionNamespacedLeaseHandlerFunc(func(params coordination_v1.DeleteCoordinationV1CollectionNamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.DeleteCoordinationV1CollectionNamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1DeleteCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1DeleteCoordinationV1NamespacedLeaseHandler = coordination_v1.DeleteCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.DeleteCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.DeleteCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1DeleteCoordinationV1beta1CollectionNamespacedLeaseHandler == nil {
		api.CoordinationV1beta1DeleteCoordinationV1beta1CollectionNamespacedLeaseHandler = coordination_v1beta1.DeleteCoordinationV1beta1CollectionNamespacedLeaseHandlerFunc(func(params coordination_v1beta1.DeleteCoordinationV1beta1CollectionNamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.DeleteCoordinationV1beta1CollectionNamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1DeleteCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1DeleteCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.DeleteCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.DeleteCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.DeleteCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedConfigMapHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedConfigMapHandler = core_v1.DeleteCoreV1CollectionNamespacedConfigMapHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedEndpointsHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedEndpointsHandler = core_v1.DeleteCoreV1CollectionNamespacedEndpointsHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedEventHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedEventHandler = core_v1.DeleteCoreV1CollectionNamespacedEventHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedLimitRangeHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedLimitRangeHandler = core_v1.DeleteCoreV1CollectionNamespacedLimitRangeHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedPersistentVolumeClaimHandler = core_v1.DeleteCoreV1CollectionNamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedPodHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedPodHandler = core_v1.DeleteCoreV1CollectionNamespacedPodHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedPodTemplateHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedPodTemplateHandler = core_v1.DeleteCoreV1CollectionNamespacedPodTemplateHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedReplicationControllerHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedReplicationControllerHandler = core_v1.DeleteCoreV1CollectionNamespacedReplicationControllerHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedResourceQuotaHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedResourceQuotaHandler = core_v1.DeleteCoreV1CollectionNamespacedResourceQuotaHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedSecretHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedSecretHandler = core_v1.DeleteCoreV1CollectionNamespacedSecretHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNamespacedServiceAccountHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNamespacedServiceAccountHandler = core_v1.DeleteCoreV1CollectionNamespacedServiceAccountHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionNodeHandler == nil {
		api.CoreV1DeleteCoreV1CollectionNodeHandler = core_v1.DeleteCoreV1CollectionNodeHandlerFunc(func(params core_v1.DeleteCoreV1CollectionNodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionNode has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1CollectionPersistentVolumeHandler == nil {
		api.CoreV1DeleteCoreV1CollectionPersistentVolumeHandler = core_v1.DeleteCoreV1CollectionPersistentVolumeHandlerFunc(func(params core_v1.DeleteCoreV1CollectionPersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1CollectionPersistentVolume has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespaceHandler == nil {
		api.CoreV1DeleteCoreV1NamespaceHandler = core_v1.DeleteCoreV1NamespaceHandlerFunc(func(params core_v1.DeleteCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedConfigMapHandler = core_v1.DeleteCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedEndpointsHandler = core_v1.DeleteCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedEventHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedEventHandler = core_v1.DeleteCoreV1NamespacedEventHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedLimitRangeHandler = core_v1.DeleteCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.DeleteCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedPodHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedPodHandler = core_v1.DeleteCoreV1NamespacedPodHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedPodTemplateHandler = core_v1.DeleteCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedReplicationControllerHandler = core_v1.DeleteCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedResourceQuotaHandler = core_v1.DeleteCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedSecretHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedSecretHandler = core_v1.DeleteCoreV1NamespacedSecretHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedServiceHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedServiceHandler = core_v1.DeleteCoreV1NamespacedServiceHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1DeleteCoreV1NamespacedServiceAccountHandler = core_v1.DeleteCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.DeleteCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1NodeHandler == nil {
		api.CoreV1DeleteCoreV1NodeHandler = core_v1.DeleteCoreV1NodeHandlerFunc(func(params core_v1.DeleteCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1DeleteCoreV1PersistentVolumeHandler == nil {
		api.CoreV1DeleteCoreV1PersistentVolumeHandler = core_v1.DeleteCoreV1PersistentVolumeHandlerFunc(func(params core_v1.DeleteCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.DeleteCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceHandler = discovery_v1beta1.DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.DeleteDiscoveryV1beta1CollectionNamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.DeleteDiscoveryV1beta1CollectionNamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1DeleteDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1DeleteDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.DeleteDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.DeleteDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.DeleteDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.EventsV1DeleteEventsV1CollectionNamespacedEventHandler == nil {
		api.EventsV1DeleteEventsV1CollectionNamespacedEventHandler = events_v1.DeleteEventsV1CollectionNamespacedEventHandlerFunc(func(params events_v1.DeleteEventsV1CollectionNamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.DeleteEventsV1CollectionNamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1DeleteEventsV1NamespacedEventHandler == nil {
		api.EventsV1DeleteEventsV1NamespacedEventHandler = events_v1.DeleteEventsV1NamespacedEventHandlerFunc(func(params events_v1.DeleteEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.DeleteEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1DeleteEventsV1beta1CollectionNamespacedEventHandler == nil {
		api.EventsV1beta1DeleteEventsV1beta1CollectionNamespacedEventHandler = events_v1beta1.DeleteEventsV1beta1CollectionNamespacedEventHandlerFunc(func(params events_v1beta1.DeleteEventsV1beta1CollectionNamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.DeleteEventsV1beta1CollectionNamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1DeleteEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1DeleteEventsV1beta1NamespacedEventHandler = events_v1beta1.DeleteEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.DeleteEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.DeleteEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1DeleteExtensionsV1beta1CollectionNamespacedIngressHandler == nil {
		api.ExtensionsV1beta1DeleteExtensionsV1beta1CollectionNamespacedIngressHandler = extensions_v1beta1.DeleteExtensionsV1beta1CollectionNamespacedIngressHandlerFunc(func(params extensions_v1beta1.DeleteExtensionsV1beta1CollectionNamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.DeleteExtensionsV1beta1CollectionNamespacedIngress has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1DeleteExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1DeleteExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.DeleteExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.DeleteExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.DeleteExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandler = flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1CollectionFlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1CollectionPriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1CollectionPriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1CollectionPriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1CollectionPriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1CollectionPriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.DeleteFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.NetworkingV1DeleteNetworkingV1CollectionIngressClassHandler == nil {
		api.NetworkingV1DeleteNetworkingV1CollectionIngressClassHandler = networking_v1.DeleteNetworkingV1CollectionIngressClassHandlerFunc(func(params networking_v1.DeleteNetworkingV1CollectionIngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.DeleteNetworkingV1CollectionIngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1DeleteNetworkingV1CollectionNamespacedIngressHandler == nil {
		api.NetworkingV1DeleteNetworkingV1CollectionNamespacedIngressHandler = networking_v1.DeleteNetworkingV1CollectionNamespacedIngressHandlerFunc(func(params networking_v1.DeleteNetworkingV1CollectionNamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.DeleteNetworkingV1CollectionNamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1DeleteNetworkingV1CollectionNamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1DeleteNetworkingV1CollectionNamespacedNetworkPolicyHandler = networking_v1.DeleteNetworkingV1CollectionNamespacedNetworkPolicyHandlerFunc(func(params networking_v1.DeleteNetworkingV1CollectionNamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.DeleteNetworkingV1CollectionNamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1DeleteNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1DeleteNetworkingV1IngressClassHandler = networking_v1.DeleteNetworkingV1IngressClassHandlerFunc(func(params networking_v1.DeleteNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.DeleteNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1DeleteNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1DeleteNetworkingV1NamespacedIngressHandler = networking_v1.DeleteNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.DeleteNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.DeleteNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1DeleteNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1DeleteNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.DeleteNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.DeleteNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.DeleteNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1DeleteNetworkingV1beta1CollectionIngressClassHandler == nil {
		api.NetworkingV1beta1DeleteNetworkingV1beta1CollectionIngressClassHandler = networking_v1beta1.DeleteNetworkingV1beta1CollectionIngressClassHandlerFunc(func(params networking_v1beta1.DeleteNetworkingV1beta1CollectionIngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.DeleteNetworkingV1beta1CollectionIngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1DeleteNetworkingV1beta1CollectionNamespacedIngressHandler == nil {
		api.NetworkingV1beta1DeleteNetworkingV1beta1CollectionNamespacedIngressHandler = networking_v1beta1.DeleteNetworkingV1beta1CollectionNamespacedIngressHandlerFunc(func(params networking_v1beta1.DeleteNetworkingV1beta1CollectionNamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.DeleteNetworkingV1beta1CollectionNamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1DeleteNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1DeleteNetworkingV1beta1IngressClassHandler = networking_v1beta1.DeleteNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.DeleteNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.DeleteNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1DeleteNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1DeleteNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.DeleteNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.DeleteNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.DeleteNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NodeV1alpha1DeleteNodeV1alpha1CollectionRuntimeClassHandler == nil {
		api.NodeV1alpha1DeleteNodeV1alpha1CollectionRuntimeClassHandler = node_v1alpha1.DeleteNodeV1alpha1CollectionRuntimeClassHandlerFunc(func(params node_v1alpha1.DeleteNodeV1alpha1CollectionRuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.DeleteNodeV1alpha1CollectionRuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1alpha1DeleteNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1DeleteNodeV1alpha1RuntimeClassHandler = node_v1alpha1.DeleteNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.DeleteNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.DeleteNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1DeleteNodeV1beta1CollectionRuntimeClassHandler == nil {
		api.NodeV1beta1DeleteNodeV1beta1CollectionRuntimeClassHandler = node_v1beta1.DeleteNodeV1beta1CollectionRuntimeClassHandlerFunc(func(params node_v1beta1.DeleteNodeV1beta1CollectionRuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.DeleteNodeV1beta1CollectionRuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1DeleteNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1DeleteNodeV1beta1RuntimeClassHandler = node_v1beta1.DeleteNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.DeleteNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.DeleteNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.PolicyV1beta1DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetHandler = policy_v1beta1.DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.DeletePolicyV1beta1CollectionNamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1DeletePolicyV1beta1CollectionPodSecurityPolicyHandler == nil {
		api.PolicyV1beta1DeletePolicyV1beta1CollectionPodSecurityPolicyHandler = policy_v1beta1.DeletePolicyV1beta1CollectionPodSecurityPolicyHandlerFunc(func(params policy_v1beta1.DeletePolicyV1beta1CollectionPodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.DeletePolicyV1beta1CollectionPodSecurityPolicy has not yet been implemented")
		})
	}
	if api.PolicyV1beta1DeletePolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1DeletePolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.DeletePolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.DeletePolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.DeletePolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1DeletePolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1DeletePolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.DeletePolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.DeletePolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.DeletePolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionClusterRoleHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionClusterRoleHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionClusterRoleHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionClusterRoleBindingHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionNamespacedRoleHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionNamespacedRoleHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionNamespacedRoleHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionNamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionNamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionNamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1CollectionNamespacedRoleBindingHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionNamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionNamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1CollectionNamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1DeleteRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1DeleteRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.DeleteRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.DeleteRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.DeleteRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionClusterRoleHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionNamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1CollectionNamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1DeleteRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.DeleteRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionClusterRoleHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionClusterRoleBindingHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionNamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleBindingHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1CollectionNamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1DeleteRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.DeleteRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.SchedulingV1DeleteSchedulingV1CollectionPriorityClassHandler == nil {
		api.SchedulingV1DeleteSchedulingV1CollectionPriorityClassHandler = scheduling_v1.DeleteSchedulingV1CollectionPriorityClassHandlerFunc(func(params scheduling_v1.DeleteSchedulingV1CollectionPriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.DeleteSchedulingV1CollectionPriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1DeleteSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1DeleteSchedulingV1PriorityClassHandler = scheduling_v1.DeleteSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.DeleteSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.DeleteSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1DeleteSchedulingV1alpha1CollectionPriorityClassHandler == nil {
		api.SchedulingV1alpha1DeleteSchedulingV1alpha1CollectionPriorityClassHandler = scheduling_v1alpha1.DeleteSchedulingV1alpha1CollectionPriorityClassHandlerFunc(func(params scheduling_v1alpha1.DeleteSchedulingV1alpha1CollectionPriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.DeleteSchedulingV1alpha1CollectionPriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1DeleteSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1DeleteSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.DeleteSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.DeleteSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.DeleteSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1DeleteSchedulingV1beta1CollectionPriorityClassHandler == nil {
		api.SchedulingV1beta1DeleteSchedulingV1beta1CollectionPriorityClassHandler = scheduling_v1beta1.DeleteSchedulingV1beta1CollectionPriorityClassHandlerFunc(func(params scheduling_v1beta1.DeleteSchedulingV1beta1CollectionPriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.DeleteSchedulingV1beta1CollectionPriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1DeleteSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1DeleteSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.DeleteSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.DeleteSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.DeleteSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandler = settings_v1alpha1.DeleteSettingsV1alpha1CollectionNamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.DeleteSettingsV1alpha1CollectionNamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.DeleteSettingsV1alpha1CollectionNamespacedPodPreset has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1DeleteSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1DeleteSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.DeleteSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.DeleteSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.DeleteSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1CSIDriverHandler == nil {
		api.StorageV1DeleteStorageV1CSIDriverHandler = storage_v1.DeleteStorageV1CSIDriverHandlerFunc(func(params storage_v1.DeleteStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1CSINodeHandler == nil {
		api.StorageV1DeleteStorageV1CSINodeHandler = storage_v1.DeleteStorageV1CSINodeHandlerFunc(func(params storage_v1.DeleteStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1CollectionCSIDriverHandler == nil {
		api.StorageV1DeleteStorageV1CollectionCSIDriverHandler = storage_v1.DeleteStorageV1CollectionCSIDriverHandlerFunc(func(params storage_v1.DeleteStorageV1CollectionCSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1CollectionCSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1CollectionCSINodeHandler == nil {
		api.StorageV1DeleteStorageV1CollectionCSINodeHandler = storage_v1.DeleteStorageV1CollectionCSINodeHandlerFunc(func(params storage_v1.DeleteStorageV1CollectionCSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1CollectionCSINode has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1CollectionStorageClassHandler == nil {
		api.StorageV1DeleteStorageV1CollectionStorageClassHandler = storage_v1.DeleteStorageV1CollectionStorageClassHandlerFunc(func(params storage_v1.DeleteStorageV1CollectionStorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1CollectionStorageClass has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1CollectionVolumeAttachmentHandler == nil {
		api.StorageV1DeleteStorageV1CollectionVolumeAttachmentHandler = storage_v1.DeleteStorageV1CollectionVolumeAttachmentHandlerFunc(func(params storage_v1.DeleteStorageV1CollectionVolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1CollectionVolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1StorageClassHandler == nil {
		api.StorageV1DeleteStorageV1StorageClassHandler = storage_v1.DeleteStorageV1StorageClassHandlerFunc(func(params storage_v1.DeleteStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1DeleteStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1DeleteStorageV1VolumeAttachmentHandler = storage_v1.DeleteStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.DeleteStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.DeleteStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1alpha1DeleteStorageV1alpha1CollectionVolumeAttachmentHandler == nil {
		api.StorageV1alpha1DeleteStorageV1alpha1CollectionVolumeAttachmentHandler = storage_v1alpha1.DeleteStorageV1alpha1CollectionVolumeAttachmentHandlerFunc(func(params storage_v1alpha1.DeleteStorageV1alpha1CollectionVolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.DeleteStorageV1alpha1CollectionVolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1alpha1DeleteStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1DeleteStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.DeleteStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.DeleteStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.DeleteStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1CSIDriverHandler = storage_v1beta1.DeleteStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1CSINodeHandler = storage_v1beta1.DeleteStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1CollectionCSIDriverHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1CollectionCSIDriverHandler = storage_v1beta1.DeleteStorageV1beta1CollectionCSIDriverHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1CollectionCSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1CollectionCSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1CollectionCSINodeHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1CollectionCSINodeHandler = storage_v1beta1.DeleteStorageV1beta1CollectionCSINodeHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1CollectionCSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1CollectionCSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1CollectionStorageClassHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1CollectionStorageClassHandler = storage_v1beta1.DeleteStorageV1beta1CollectionStorageClassHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1CollectionStorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1CollectionStorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1CollectionVolumeAttachmentHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1CollectionVolumeAttachmentHandler = storage_v1beta1.DeleteStorageV1beta1CollectionVolumeAttachmentHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1CollectionVolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1CollectionVolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1StorageClassHandler = storage_v1beta1.DeleteStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1DeleteStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1DeleteStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.DeleteStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.DeleteStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.DeleteStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.ApisGetAPIVersionsHandler == nil {
		api.ApisGetAPIVersionsHandler = apis.GetAPIVersionsHandlerFunc(func(params apis.GetAPIVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation apis.GetAPIVersions has not yet been implemented")
		})
	}
	if api.AdmissionregistrationGetAdmissionregistrationAPIGroupHandler == nil {
		api.AdmissionregistrationGetAdmissionregistrationAPIGroupHandler = admissionregistration.GetAdmissionregistrationAPIGroupHandlerFunc(func(params admissionregistration.GetAdmissionregistrationAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration.GetAdmissionregistrationAPIGroup has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1GetAdmissionregistrationV1APIResourcesHandler == nil {
		api.AdmissionregistrationV1GetAdmissionregistrationV1APIResourcesHandler = admissionregistration_v1.GetAdmissionregistrationV1APIResourcesHandlerFunc(func(params admissionregistration_v1.GetAdmissionregistrationV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.GetAdmissionregistrationV1APIResources has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1GetAdmissionregistrationV1beta1APIResourcesHandler == nil {
		api.AdmissionregistrationV1beta1GetAdmissionregistrationV1beta1APIResourcesHandler = admissionregistration_v1beta1.GetAdmissionregistrationV1beta1APIResourcesHandlerFunc(func(params admissionregistration_v1beta1.GetAdmissionregistrationV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.GetAdmissionregistrationV1beta1APIResources has not yet been implemented")
		})
	}
	if api.ApiextensionsGetApiextensionsAPIGroupHandler == nil {
		api.ApiextensionsGetApiextensionsAPIGroupHandler = apiextensions.GetApiextensionsAPIGroupHandlerFunc(func(params apiextensions.GetApiextensionsAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions.GetApiextensionsAPIGroup has not yet been implemented")
		})
	}
	if api.ApiextensionsV1GetApiextensionsV1APIResourcesHandler == nil {
		api.ApiextensionsV1GetApiextensionsV1APIResourcesHandler = apiextensions_v1.GetApiextensionsV1APIResourcesHandlerFunc(func(params apiextensions_v1.GetApiextensionsV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.GetApiextensionsV1APIResources has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1GetApiextensionsV1beta1APIResourcesHandler == nil {
		api.ApiextensionsV1beta1GetApiextensionsV1beta1APIResourcesHandler = apiextensions_v1beta1.GetApiextensionsV1beta1APIResourcesHandlerFunc(func(params apiextensions_v1beta1.GetApiextensionsV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.GetApiextensionsV1beta1APIResources has not yet been implemented")
		})
	}
	if api.ApiregistrationGetApiregistrationAPIGroupHandler == nil {
		api.ApiregistrationGetApiregistrationAPIGroupHandler = apiregistration.GetApiregistrationAPIGroupHandlerFunc(func(params apiregistration.GetApiregistrationAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration.GetApiregistrationAPIGroup has not yet been implemented")
		})
	}
	if api.ApiregistrationV1GetApiregistrationV1APIResourcesHandler == nil {
		api.ApiregistrationV1GetApiregistrationV1APIResourcesHandler = apiregistration_v1.GetApiregistrationV1APIResourcesHandlerFunc(func(params apiregistration_v1.GetApiregistrationV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.GetApiregistrationV1APIResources has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1GetApiregistrationV1beta1APIResourcesHandler == nil {
		api.ApiregistrationV1beta1GetApiregistrationV1beta1APIResourcesHandler = apiregistration_v1beta1.GetApiregistrationV1beta1APIResourcesHandlerFunc(func(params apiregistration_v1beta1.GetApiregistrationV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.GetApiregistrationV1beta1APIResources has not yet been implemented")
		})
	}
	if api.AppsGetAppsAPIGroupHandler == nil {
		api.AppsGetAppsAPIGroupHandler = apps.GetAppsAPIGroupHandlerFunc(func(params apps.GetAppsAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation apps.GetAppsAPIGroup has not yet been implemented")
		})
	}
	if api.AppsV1GetAppsV1APIResourcesHandler == nil {
		api.AppsV1GetAppsV1APIResourcesHandler = apps_v1.GetAppsV1APIResourcesHandlerFunc(func(params apps_v1.GetAppsV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.GetAppsV1APIResources has not yet been implemented")
		})
	}
	if api.AuthenticationGetAuthenticationAPIGroupHandler == nil {
		api.AuthenticationGetAuthenticationAPIGroupHandler = authentication.GetAuthenticationAPIGroupHandlerFunc(func(params authentication.GetAuthenticationAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication.GetAuthenticationAPIGroup has not yet been implemented")
		})
	}
	if api.AuthenticationV1GetAuthenticationV1APIResourcesHandler == nil {
		api.AuthenticationV1GetAuthenticationV1APIResourcesHandler = authentication_v1.GetAuthenticationV1APIResourcesHandlerFunc(func(params authentication_v1.GetAuthenticationV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication_v1.GetAuthenticationV1APIResources has not yet been implemented")
		})
	}
	if api.AuthenticationV1beta1GetAuthenticationV1beta1APIResourcesHandler == nil {
		api.AuthenticationV1beta1GetAuthenticationV1beta1APIResourcesHandler = authentication_v1beta1.GetAuthenticationV1beta1APIResourcesHandlerFunc(func(params authentication_v1beta1.GetAuthenticationV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation authentication_v1beta1.GetAuthenticationV1beta1APIResources has not yet been implemented")
		})
	}
	if api.AuthorizationGetAuthorizationAPIGroupHandler == nil {
		api.AuthorizationGetAuthorizationAPIGroupHandler = authorization.GetAuthorizationAPIGroupHandlerFunc(func(params authorization.GetAuthorizationAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization.GetAuthorizationAPIGroup has not yet been implemented")
		})
	}
	if api.AuthorizationV1GetAuthorizationV1APIResourcesHandler == nil {
		api.AuthorizationV1GetAuthorizationV1APIResourcesHandler = authorization_v1.GetAuthorizationV1APIResourcesHandlerFunc(func(params authorization_v1.GetAuthorizationV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1.GetAuthorizationV1APIResources has not yet been implemented")
		})
	}
	if api.AuthorizationV1beta1GetAuthorizationV1beta1APIResourcesHandler == nil {
		api.AuthorizationV1beta1GetAuthorizationV1beta1APIResourcesHandler = authorization_v1beta1.GetAuthorizationV1beta1APIResourcesHandlerFunc(func(params authorization_v1beta1.GetAuthorizationV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation authorization_v1beta1.GetAuthorizationV1beta1APIResources has not yet been implemented")
		})
	}
	if api.AutoscalingGetAutoscalingAPIGroupHandler == nil {
		api.AutoscalingGetAutoscalingAPIGroupHandler = autoscaling.GetAutoscalingAPIGroupHandlerFunc(func(params autoscaling.GetAutoscalingAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling.GetAutoscalingAPIGroup has not yet been implemented")
		})
	}
	if api.AutoscalingV1GetAutoscalingV1APIResourcesHandler == nil {
		api.AutoscalingV1GetAutoscalingV1APIResourcesHandler = autoscaling_v1.GetAutoscalingV1APIResourcesHandlerFunc(func(params autoscaling_v1.GetAutoscalingV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.GetAutoscalingV1APIResources has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1GetAutoscalingV2beta1APIResourcesHandler == nil {
		api.AutoscalingV2beta1GetAutoscalingV2beta1APIResourcesHandler = autoscaling_v2beta1.GetAutoscalingV2beta1APIResourcesHandlerFunc(func(params autoscaling_v2beta1.GetAutoscalingV2beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.GetAutoscalingV2beta1APIResources has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2GetAutoscalingV2beta2APIResourcesHandler == nil {
		api.AutoscalingV2beta2GetAutoscalingV2beta2APIResourcesHandler = autoscaling_v2beta2.GetAutoscalingV2beta2APIResourcesHandlerFunc(func(params autoscaling_v2beta2.GetAutoscalingV2beta2APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.GetAutoscalingV2beta2APIResources has not yet been implemented")
		})
	}
	if api.BatchGetBatchAPIGroupHandler == nil {
		api.BatchGetBatchAPIGroupHandler = batch.GetBatchAPIGroupHandlerFunc(func(params batch.GetBatchAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation batch.GetBatchAPIGroup has not yet been implemented")
		})
	}
	if api.BatchV1GetBatchV1APIResourcesHandler == nil {
		api.BatchV1GetBatchV1APIResourcesHandler = batch_v1.GetBatchV1APIResourcesHandlerFunc(func(params batch_v1.GetBatchV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.GetBatchV1APIResources has not yet been implemented")
		})
	}
	if api.BatchV1beta1GetBatchV1beta1APIResourcesHandler == nil {
		api.BatchV1beta1GetBatchV1beta1APIResourcesHandler = batch_v1beta1.GetBatchV1beta1APIResourcesHandlerFunc(func(params batch_v1beta1.GetBatchV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.GetBatchV1beta1APIResources has not yet been implemented")
		})
	}
	if api.BatchV2alpha1GetBatchV2alpha1APIResourcesHandler == nil {
		api.BatchV2alpha1GetBatchV2alpha1APIResourcesHandler = batch_v2alpha1.GetBatchV2alpha1APIResourcesHandlerFunc(func(params batch_v2alpha1.GetBatchV2alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.GetBatchV2alpha1APIResources has not yet been implemented")
		})
	}
	if api.CertificatesGetCertificatesAPIGroupHandler == nil {
		api.CertificatesGetCertificatesAPIGroupHandler = certificates.GetCertificatesAPIGroupHandlerFunc(func(params certificates.GetCertificatesAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates.GetCertificatesAPIGroup has not yet been implemented")
		})
	}
	if api.CertificatesV1GetCertificatesV1APIResourcesHandler == nil {
		api.CertificatesV1GetCertificatesV1APIResourcesHandler = certificates_v1.GetCertificatesV1APIResourcesHandlerFunc(func(params certificates_v1.GetCertificatesV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.GetCertificatesV1APIResources has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1GetCertificatesV1beta1APIResourcesHandler == nil {
		api.CertificatesV1beta1GetCertificatesV1beta1APIResourcesHandler = certificates_v1beta1.GetCertificatesV1beta1APIResourcesHandlerFunc(func(params certificates_v1beta1.GetCertificatesV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.GetCertificatesV1beta1APIResources has not yet been implemented")
		})
	}
	if api.VersionGetCodeVersionHandler == nil {
		api.VersionGetCodeVersionHandler = version.GetCodeVersionHandlerFunc(func(params version.GetCodeVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation version.GetCodeVersion has not yet been implemented")
		})
	}
	if api.CoordinationGetCoordinationAPIGroupHandler == nil {
		api.CoordinationGetCoordinationAPIGroupHandler = coordination.GetCoordinationAPIGroupHandlerFunc(func(params coordination.GetCoordinationAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination.GetCoordinationAPIGroup has not yet been implemented")
		})
	}
	if api.CoordinationV1GetCoordinationV1APIResourcesHandler == nil {
		api.CoordinationV1GetCoordinationV1APIResourcesHandler = coordination_v1.GetCoordinationV1APIResourcesHandlerFunc(func(params coordination_v1.GetCoordinationV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.GetCoordinationV1APIResources has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1GetCoordinationV1beta1APIResourcesHandler == nil {
		api.CoordinationV1beta1GetCoordinationV1beta1APIResourcesHandler = coordination_v1beta1.GetCoordinationV1beta1APIResourcesHandlerFunc(func(params coordination_v1beta1.GetCoordinationV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.GetCoordinationV1beta1APIResources has not yet been implemented")
		})
	}
	if api.CoreGetCoreAPIVersionsHandler == nil {
		api.CoreGetCoreAPIVersionsHandler = core.GetCoreAPIVersionsHandlerFunc(func(params core.GetCoreAPIVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation core.GetCoreAPIVersions has not yet been implemented")
		})
	}
	if api.CoreV1GetCoreV1APIResourcesHandler == nil {
		api.CoreV1GetCoreV1APIResourcesHandler = core_v1.GetCoreV1APIResourcesHandlerFunc(func(params core_v1.GetCoreV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.GetCoreV1APIResources has not yet been implemented")
		})
	}
	if api.DiscoveryGetDiscoveryAPIGroupHandler == nil {
		api.DiscoveryGetDiscoveryAPIGroupHandler = discovery.GetDiscoveryAPIGroupHandlerFunc(func(params discovery.GetDiscoveryAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery.GetDiscoveryAPIGroup has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1GetDiscoveryV1beta1APIResourcesHandler == nil {
		api.DiscoveryV1beta1GetDiscoveryV1beta1APIResourcesHandler = discovery_v1beta1.GetDiscoveryV1beta1APIResourcesHandlerFunc(func(params discovery_v1beta1.GetDiscoveryV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.GetDiscoveryV1beta1APIResources has not yet been implemented")
		})
	}
	if api.EventsGetEventsAPIGroupHandler == nil {
		api.EventsGetEventsAPIGroupHandler = events.GetEventsAPIGroupHandlerFunc(func(params events.GetEventsAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation events.GetEventsAPIGroup has not yet been implemented")
		})
	}
	if api.EventsV1GetEventsV1APIResourcesHandler == nil {
		api.EventsV1GetEventsV1APIResourcesHandler = events_v1.GetEventsV1APIResourcesHandlerFunc(func(params events_v1.GetEventsV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.GetEventsV1APIResources has not yet been implemented")
		})
	}
	if api.EventsV1beta1GetEventsV1beta1APIResourcesHandler == nil {
		api.EventsV1beta1GetEventsV1beta1APIResourcesHandler = events_v1beta1.GetEventsV1beta1APIResourcesHandlerFunc(func(params events_v1beta1.GetEventsV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.GetEventsV1beta1APIResources has not yet been implemented")
		})
	}
	if api.ExtensionsGetExtensionsAPIGroupHandler == nil {
		api.ExtensionsGetExtensionsAPIGroupHandler = extensions.GetExtensionsAPIGroupHandlerFunc(func(params extensions.GetExtensionsAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions.GetExtensionsAPIGroup has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1GetExtensionsV1beta1APIResourcesHandler == nil {
		api.ExtensionsV1beta1GetExtensionsV1beta1APIResourcesHandler = extensions_v1beta1.GetExtensionsV1beta1APIResourcesHandlerFunc(func(params extensions_v1beta1.GetExtensionsV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.GetExtensionsV1beta1APIResources has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverGetFlowcontrolApiserverAPIGroupHandler == nil {
		api.FlowcontrolApiserverGetFlowcontrolApiserverAPIGroupHandler = flowcontrol_apiserver.GetFlowcontrolApiserverAPIGroupHandlerFunc(func(params flowcontrol_apiserver.GetFlowcontrolApiserverAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver.GetFlowcontrolApiserverAPIGroup has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1GetFlowcontrolApiserverV1alpha1APIResourcesHandler == nil {
		api.FlowcontrolApiserverV1alpha1GetFlowcontrolApiserverV1alpha1APIResourcesHandler = flowcontrol_apiserver_v1alpha1.GetFlowcontrolApiserverV1alpha1APIResourcesHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.GetFlowcontrolApiserverV1alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.GetFlowcontrolApiserverV1alpha1APIResources has not yet been implemented")
		})
	}
	if api.NetworkingGetNetworkingAPIGroupHandler == nil {
		api.NetworkingGetNetworkingAPIGroupHandler = networking.GetNetworkingAPIGroupHandlerFunc(func(params networking.GetNetworkingAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation networking.GetNetworkingAPIGroup has not yet been implemented")
		})
	}
	if api.NetworkingV1GetNetworkingV1APIResourcesHandler == nil {
		api.NetworkingV1GetNetworkingV1APIResourcesHandler = networking_v1.GetNetworkingV1APIResourcesHandlerFunc(func(params networking_v1.GetNetworkingV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.GetNetworkingV1APIResources has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1GetNetworkingV1beta1APIResourcesHandler == nil {
		api.NetworkingV1beta1GetNetworkingV1beta1APIResourcesHandler = networking_v1beta1.GetNetworkingV1beta1APIResourcesHandlerFunc(func(params networking_v1beta1.GetNetworkingV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.GetNetworkingV1beta1APIResources has not yet been implemented")
		})
	}
	if api.NodeGetNodeAPIGroupHandler == nil {
		api.NodeGetNodeAPIGroupHandler = node.GetNodeAPIGroupHandlerFunc(func(params node.GetNodeAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation node.GetNodeAPIGroup has not yet been implemented")
		})
	}
	if api.NodeV1alpha1GetNodeV1alpha1APIResourcesHandler == nil {
		api.NodeV1alpha1GetNodeV1alpha1APIResourcesHandler = node_v1alpha1.GetNodeV1alpha1APIResourcesHandlerFunc(func(params node_v1alpha1.GetNodeV1alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.GetNodeV1alpha1APIResources has not yet been implemented")
		})
	}
	if api.NodeV1beta1GetNodeV1beta1APIResourcesHandler == nil {
		api.NodeV1beta1GetNodeV1beta1APIResourcesHandler = node_v1beta1.GetNodeV1beta1APIResourcesHandlerFunc(func(params node_v1beta1.GetNodeV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.GetNodeV1beta1APIResources has not yet been implemented")
		})
	}
	if api.PolicyGetPolicyAPIGroupHandler == nil {
		api.PolicyGetPolicyAPIGroupHandler = policy.GetPolicyAPIGroupHandlerFunc(func(params policy.GetPolicyAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation policy.GetPolicyAPIGroup has not yet been implemented")
		})
	}
	if api.PolicyV1beta1GetPolicyV1beta1APIResourcesHandler == nil {
		api.PolicyV1beta1GetPolicyV1beta1APIResourcesHandler = policy_v1beta1.GetPolicyV1beta1APIResourcesHandlerFunc(func(params policy_v1beta1.GetPolicyV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.GetPolicyV1beta1APIResources has not yet been implemented")
		})
	}
	if api.RbacAuthorizationGetRbacAuthorizationAPIGroupHandler == nil {
		api.RbacAuthorizationGetRbacAuthorizationAPIGroupHandler = rbac_authorization.GetRbacAuthorizationAPIGroupHandlerFunc(func(params rbac_authorization.GetRbacAuthorizationAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization.GetRbacAuthorizationAPIGroup has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1GetRbacAuthorizationV1APIResourcesHandler == nil {
		api.RbacAuthorizationV1GetRbacAuthorizationV1APIResourcesHandler = rbac_authorization_v1.GetRbacAuthorizationV1APIResourcesHandlerFunc(func(params rbac_authorization_v1.GetRbacAuthorizationV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.GetRbacAuthorizationV1APIResources has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1GetRbacAuthorizationV1alpha1APIResourcesHandler == nil {
		api.RbacAuthorizationV1alpha1GetRbacAuthorizationV1alpha1APIResourcesHandler = rbac_authorization_v1alpha1.GetRbacAuthorizationV1alpha1APIResourcesHandlerFunc(func(params rbac_authorization_v1alpha1.GetRbacAuthorizationV1alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.GetRbacAuthorizationV1alpha1APIResources has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1GetRbacAuthorizationV1beta1APIResourcesHandler == nil {
		api.RbacAuthorizationV1beta1GetRbacAuthorizationV1beta1APIResourcesHandler = rbac_authorization_v1beta1.GetRbacAuthorizationV1beta1APIResourcesHandlerFunc(func(params rbac_authorization_v1beta1.GetRbacAuthorizationV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.GetRbacAuthorizationV1beta1APIResources has not yet been implemented")
		})
	}
	if api.SchedulingGetSchedulingAPIGroupHandler == nil {
		api.SchedulingGetSchedulingAPIGroupHandler = scheduling.GetSchedulingAPIGroupHandlerFunc(func(params scheduling.GetSchedulingAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling.GetSchedulingAPIGroup has not yet been implemented")
		})
	}
	if api.SchedulingV1GetSchedulingV1APIResourcesHandler == nil {
		api.SchedulingV1GetSchedulingV1APIResourcesHandler = scheduling_v1.GetSchedulingV1APIResourcesHandlerFunc(func(params scheduling_v1.GetSchedulingV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.GetSchedulingV1APIResources has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1GetSchedulingV1alpha1APIResourcesHandler == nil {
		api.SchedulingV1alpha1GetSchedulingV1alpha1APIResourcesHandler = scheduling_v1alpha1.GetSchedulingV1alpha1APIResourcesHandlerFunc(func(params scheduling_v1alpha1.GetSchedulingV1alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.GetSchedulingV1alpha1APIResources has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1GetSchedulingV1beta1APIResourcesHandler == nil {
		api.SchedulingV1beta1GetSchedulingV1beta1APIResourcesHandler = scheduling_v1beta1.GetSchedulingV1beta1APIResourcesHandlerFunc(func(params scheduling_v1beta1.GetSchedulingV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.GetSchedulingV1beta1APIResources has not yet been implemented")
		})
	}
	if api.SettingsGetSettingsAPIGroupHandler == nil {
		api.SettingsGetSettingsAPIGroupHandler = settings.GetSettingsAPIGroupHandlerFunc(func(params settings.GetSettingsAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation settings.GetSettingsAPIGroup has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1GetSettingsV1alpha1APIResourcesHandler == nil {
		api.SettingsV1alpha1GetSettingsV1alpha1APIResourcesHandler = settings_v1alpha1.GetSettingsV1alpha1APIResourcesHandlerFunc(func(params settings_v1alpha1.GetSettingsV1alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.GetSettingsV1alpha1APIResources has not yet been implemented")
		})
	}
	if api.StorageGetStorageAPIGroupHandler == nil {
		api.StorageGetStorageAPIGroupHandler = storage.GetStorageAPIGroupHandlerFunc(func(params storage.GetStorageAPIGroupParams) middleware.Responder {
			return middleware.NotImplemented("operation storage.GetStorageAPIGroup has not yet been implemented")
		})
	}
	if api.StorageV1GetStorageV1APIResourcesHandler == nil {
		api.StorageV1GetStorageV1APIResourcesHandler = storage_v1.GetStorageV1APIResourcesHandlerFunc(func(params storage_v1.GetStorageV1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.GetStorageV1APIResources has not yet been implemented")
		})
	}
	if api.StorageV1alpha1GetStorageV1alpha1APIResourcesHandler == nil {
		api.StorageV1alpha1GetStorageV1alpha1APIResourcesHandler = storage_v1alpha1.GetStorageV1alpha1APIResourcesHandlerFunc(func(params storage_v1alpha1.GetStorageV1alpha1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.GetStorageV1alpha1APIResources has not yet been implemented")
		})
	}
	if api.StorageV1beta1GetStorageV1beta1APIResourcesHandler == nil {
		api.StorageV1beta1GetStorageV1beta1APIResourcesHandler = storage_v1beta1.GetStorageV1beta1APIResourcesHandlerFunc(func(params storage_v1beta1.GetStorageV1beta1APIResourcesParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.GetStorageV1beta1APIResources has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1ListAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1ListAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.ListAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.ListAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.ListAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1ListAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1ListAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.ListAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.ListAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.ListAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1ListAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1ListAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.ListAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.ListAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.ListAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.ListAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.ListAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.ApiextensionsV1ListApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1ListApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.ListApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.ListApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.ListApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1ListApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1ListApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.ListApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.ListApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.ListApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiregistrationV1ListApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1ListApiregistrationV1APIServiceHandler = apiregistration_v1.ListApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.ListApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.ListApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1ListApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1ListApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.ListApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.ListApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.ListApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1ControllerRevisionForAllNamespacesHandler == nil {
		api.AppsV1ListAppsV1ControllerRevisionForAllNamespacesHandler = apps_v1.ListAppsV1ControllerRevisionForAllNamespacesHandlerFunc(func(params apps_v1.ListAppsV1ControllerRevisionForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1ControllerRevisionForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1DaemonSetForAllNamespacesHandler == nil {
		api.AppsV1ListAppsV1DaemonSetForAllNamespacesHandler = apps_v1.ListAppsV1DaemonSetForAllNamespacesHandlerFunc(func(params apps_v1.ListAppsV1DaemonSetForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1DaemonSetForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1DeploymentForAllNamespacesHandler == nil {
		api.AppsV1ListAppsV1DeploymentForAllNamespacesHandler = apps_v1.ListAppsV1DeploymentForAllNamespacesHandlerFunc(func(params apps_v1.ListAppsV1DeploymentForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1DeploymentForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1ListAppsV1NamespacedControllerRevisionHandler = apps_v1.ListAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.ListAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1ListAppsV1NamespacedDaemonSetHandler = apps_v1.ListAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.ListAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1ListAppsV1NamespacedDeploymentHandler = apps_v1.ListAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.ListAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1ListAppsV1NamespacedReplicaSetHandler = apps_v1.ListAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.ListAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1ListAppsV1NamespacedStatefulSetHandler = apps_v1.ListAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.ListAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1ReplicaSetForAllNamespacesHandler == nil {
		api.AppsV1ListAppsV1ReplicaSetForAllNamespacesHandler = apps_v1.ListAppsV1ReplicaSetForAllNamespacesHandlerFunc(func(params apps_v1.ListAppsV1ReplicaSetForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1ReplicaSetForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1ListAppsV1StatefulSetForAllNamespacesHandler == nil {
		api.AppsV1ListAppsV1StatefulSetForAllNamespacesHandler = apps_v1.ListAppsV1StatefulSetForAllNamespacesHandlerFunc(func(params apps_v1.ListAppsV1StatefulSetForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ListAppsV1StatefulSetForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV1ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesHandler == nil {
		api.AutoscalingV1ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesHandler = autoscaling_v1.ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesHandlerFunc(func(params autoscaling_v1.ListAutoscalingV1HorizontalPodAutoscalerForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.ListAutoscalingV1HorizontalPodAutoscalerForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV1ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.ListAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.ListAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.ListAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1ListAutoscalingV2beta1HorizontalPodAutoscalerForAllNamespacesHandler == nil {
		api.AutoscalingV2beta1ListAutoscalingV2beta1HorizontalPodAutoscalerForAllNamespacesHandler = autoscaling_v2beta1.ListAutoscalingV2beta1HorizontalPodAutoscalerForAllNamespacesHandlerFunc(func(params autoscaling_v2beta1.ListAutoscalingV2beta1HorizontalPodAutoscalerForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.ListAutoscalingV2beta1HorizontalPodAutoscalerForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1ListAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1ListAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.ListAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.ListAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.ListAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandler == nil {
		api.AutoscalingV2beta2ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandler = autoscaling_v2beta2.ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesHandlerFunc(func(params autoscaling_v2beta2.ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.ListAutoscalingV2beta2HorizontalPodAutoscalerForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2ListAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2ListAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.ListAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.ListAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.ListAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.BatchV1ListBatchV1JobForAllNamespacesHandler == nil {
		api.BatchV1ListBatchV1JobForAllNamespacesHandler = batch_v1.ListBatchV1JobForAllNamespacesHandlerFunc(func(params batch_v1.ListBatchV1JobForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.ListBatchV1JobForAllNamespaces has not yet been implemented")
		})
	}
	if api.BatchV1ListBatchV1NamespacedJobHandler == nil {
		api.BatchV1ListBatchV1NamespacedJobHandler = batch_v1.ListBatchV1NamespacedJobHandlerFunc(func(params batch_v1.ListBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.ListBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1ListBatchV1beta1CronJobForAllNamespacesHandler == nil {
		api.BatchV1beta1ListBatchV1beta1CronJobForAllNamespacesHandler = batch_v1beta1.ListBatchV1beta1CronJobForAllNamespacesHandlerFunc(func(params batch_v1beta1.ListBatchV1beta1CronJobForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.ListBatchV1beta1CronJobForAllNamespaces has not yet been implemented")
		})
	}
	if api.BatchV1beta1ListBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1ListBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.ListBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.ListBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.ListBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1ListBatchV2alpha1CronJobForAllNamespacesHandler == nil {
		api.BatchV2alpha1ListBatchV2alpha1CronJobForAllNamespacesHandler = batch_v2alpha1.ListBatchV2alpha1CronJobForAllNamespacesHandlerFunc(func(params batch_v2alpha1.ListBatchV2alpha1CronJobForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.ListBatchV2alpha1CronJobForAllNamespaces has not yet been implemented")
		})
	}
	if api.BatchV2alpha1ListBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1ListBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.ListBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.ListBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.ListBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.CertificatesV1ListCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1ListCertificatesV1CertificateSigningRequestHandler = certificates_v1.ListCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.ListCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ListCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ListCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1ListCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.ListCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.ListCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ListCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CoordinationV1ListCoordinationV1LeaseForAllNamespacesHandler == nil {
		api.CoordinationV1ListCoordinationV1LeaseForAllNamespacesHandler = coordination_v1.ListCoordinationV1LeaseForAllNamespacesHandlerFunc(func(params coordination_v1.ListCoordinationV1LeaseForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.ListCoordinationV1LeaseForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoordinationV1ListCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1ListCoordinationV1NamespacedLeaseHandler = coordination_v1.ListCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.ListCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.ListCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1ListCoordinationV1beta1LeaseForAllNamespacesHandler == nil {
		api.CoordinationV1beta1ListCoordinationV1beta1LeaseForAllNamespacesHandler = coordination_v1beta1.ListCoordinationV1beta1LeaseForAllNamespacesHandlerFunc(func(params coordination_v1beta1.ListCoordinationV1beta1LeaseForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.ListCoordinationV1beta1LeaseForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1ListCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1ListCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.ListCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.ListCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.ListCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1ComponentStatusHandler == nil {
		api.CoreV1ListCoreV1ComponentStatusHandler = core_v1.ListCoreV1ComponentStatusHandlerFunc(func(params core_v1.ListCoreV1ComponentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1ComponentStatus has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1ConfigMapForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1ConfigMapForAllNamespacesHandler = core_v1.ListCoreV1ConfigMapForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1ConfigMapForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1ConfigMapForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1EndpointsForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1EndpointsForAllNamespacesHandler = core_v1.ListCoreV1EndpointsForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1EndpointsForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1EndpointsForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1EventForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1EventForAllNamespacesHandler = core_v1.ListCoreV1EventForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1EventForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1EventForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1LimitRangeForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1LimitRangeForAllNamespacesHandler = core_v1.ListCoreV1LimitRangeForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1LimitRangeForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1LimitRangeForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespaceHandler == nil {
		api.CoreV1ListCoreV1NamespaceHandler = core_v1.ListCoreV1NamespaceHandlerFunc(func(params core_v1.ListCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1ListCoreV1NamespacedConfigMapHandler = core_v1.ListCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.ListCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1ListCoreV1NamespacedEndpointsHandler = core_v1.ListCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.ListCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedEventHandler == nil {
		api.CoreV1ListCoreV1NamespacedEventHandler = core_v1.ListCoreV1NamespacedEventHandlerFunc(func(params core_v1.ListCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1ListCoreV1NamespacedLimitRangeHandler = core_v1.ListCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.ListCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1ListCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.ListCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.ListCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedPodHandler == nil {
		api.CoreV1ListCoreV1NamespacedPodHandler = core_v1.ListCoreV1NamespacedPodHandlerFunc(func(params core_v1.ListCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1ListCoreV1NamespacedPodTemplateHandler = core_v1.ListCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.ListCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1ListCoreV1NamespacedReplicationControllerHandler = core_v1.ListCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.ListCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1ListCoreV1NamespacedResourceQuotaHandler = core_v1.ListCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.ListCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedSecretHandler == nil {
		api.CoreV1ListCoreV1NamespacedSecretHandler = core_v1.ListCoreV1NamespacedSecretHandlerFunc(func(params core_v1.ListCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedServiceHandler == nil {
		api.CoreV1ListCoreV1NamespacedServiceHandler = core_v1.ListCoreV1NamespacedServiceHandlerFunc(func(params core_v1.ListCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1ListCoreV1NamespacedServiceAccountHandler = core_v1.ListCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.ListCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1NodeHandler == nil {
		api.CoreV1ListCoreV1NodeHandler = core_v1.ListCoreV1NodeHandlerFunc(func(params core_v1.ListCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1PersistentVolumeHandler == nil {
		api.CoreV1ListCoreV1PersistentVolumeHandler = core_v1.ListCoreV1PersistentVolumeHandlerFunc(func(params core_v1.ListCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1PersistentVolumeClaimForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1PersistentVolumeClaimForAllNamespacesHandler = core_v1.ListCoreV1PersistentVolumeClaimForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1PersistentVolumeClaimForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1PersistentVolumeClaimForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1PodForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1PodForAllNamespacesHandler = core_v1.ListCoreV1PodForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1PodForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1PodForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1PodTemplateForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1PodTemplateForAllNamespacesHandler = core_v1.ListCoreV1PodTemplateForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1PodTemplateForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1PodTemplateForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1ReplicationControllerForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1ReplicationControllerForAllNamespacesHandler = core_v1.ListCoreV1ReplicationControllerForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1ReplicationControllerForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1ReplicationControllerForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1ResourceQuotaForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1ResourceQuotaForAllNamespacesHandler = core_v1.ListCoreV1ResourceQuotaForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1ResourceQuotaForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1ResourceQuotaForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1SecretForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1SecretForAllNamespacesHandler = core_v1.ListCoreV1SecretForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1SecretForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1SecretForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1ServiceAccountForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1ServiceAccountForAllNamespacesHandler = core_v1.ListCoreV1ServiceAccountForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1ServiceAccountForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1ServiceAccountForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1ListCoreV1ServiceForAllNamespacesHandler == nil {
		api.CoreV1ListCoreV1ServiceForAllNamespacesHandler = core_v1.ListCoreV1ServiceForAllNamespacesHandlerFunc(func(params core_v1.ListCoreV1ServiceForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ListCoreV1ServiceForAllNamespaces has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler == nil {
		api.DiscoveryV1beta1ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler = discovery_v1beta1.ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandlerFunc(func(params discovery_v1beta1.ListDiscoveryV1beta1EndpointSliceForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.ListDiscoveryV1beta1EndpointSliceForAllNamespaces has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1ListDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1ListDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.ListDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.ListDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.ListDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.EventsV1ListEventsV1EventForAllNamespacesHandler == nil {
		api.EventsV1ListEventsV1EventForAllNamespacesHandler = events_v1.ListEventsV1EventForAllNamespacesHandlerFunc(func(params events_v1.ListEventsV1EventForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.ListEventsV1EventForAllNamespaces has not yet been implemented")
		})
	}
	if api.EventsV1ListEventsV1NamespacedEventHandler == nil {
		api.EventsV1ListEventsV1NamespacedEventHandler = events_v1.ListEventsV1NamespacedEventHandlerFunc(func(params events_v1.ListEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.ListEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1ListEventsV1beta1EventForAllNamespacesHandler == nil {
		api.EventsV1beta1ListEventsV1beta1EventForAllNamespacesHandler = events_v1beta1.ListEventsV1beta1EventForAllNamespacesHandlerFunc(func(params events_v1beta1.ListEventsV1beta1EventForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.ListEventsV1beta1EventForAllNamespaces has not yet been implemented")
		})
	}
	if api.EventsV1beta1ListEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1ListEventsV1beta1NamespacedEventHandler = events_v1beta1.ListEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.ListEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.ListEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1ListExtensionsV1beta1IngressForAllNamespacesHandler == nil {
		api.ExtensionsV1beta1ListExtensionsV1beta1IngressForAllNamespacesHandler = extensions_v1beta1.ListExtensionsV1beta1IngressForAllNamespacesHandlerFunc(func(params extensions_v1beta1.ListExtensionsV1beta1IngressForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.ListExtensionsV1beta1IngressForAllNamespaces has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1ListExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1ListExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.ListExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.ListExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.ListExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ListFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1ListFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.ListFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ListFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ListFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ListFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ListFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.NetworkingV1ListNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1ListNetworkingV1IngressClassHandler = networking_v1.ListNetworkingV1IngressClassHandlerFunc(func(params networking_v1.ListNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ListNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1ListNetworkingV1IngressForAllNamespacesHandler == nil {
		api.NetworkingV1ListNetworkingV1IngressForAllNamespacesHandler = networking_v1.ListNetworkingV1IngressForAllNamespacesHandlerFunc(func(params networking_v1.ListNetworkingV1IngressForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ListNetworkingV1IngressForAllNamespaces has not yet been implemented")
		})
	}
	if api.NetworkingV1ListNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1ListNetworkingV1NamespacedIngressHandler = networking_v1.ListNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.ListNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ListNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1ListNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1ListNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.ListNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.ListNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ListNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1ListNetworkingV1NetworkPolicyForAllNamespacesHandler == nil {
		api.NetworkingV1ListNetworkingV1NetworkPolicyForAllNamespacesHandler = networking_v1.ListNetworkingV1NetworkPolicyForAllNamespacesHandlerFunc(func(params networking_v1.ListNetworkingV1NetworkPolicyForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ListNetworkingV1NetworkPolicyForAllNamespaces has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ListNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1ListNetworkingV1beta1IngressClassHandler = networking_v1beta1.ListNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.ListNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ListNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ListNetworkingV1beta1IngressForAllNamespacesHandler == nil {
		api.NetworkingV1beta1ListNetworkingV1beta1IngressForAllNamespacesHandler = networking_v1beta1.ListNetworkingV1beta1IngressForAllNamespacesHandlerFunc(func(params networking_v1beta1.ListNetworkingV1beta1IngressForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ListNetworkingV1beta1IngressForAllNamespaces has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ListNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1ListNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.ListNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.ListNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ListNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NodeV1alpha1ListNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1ListNodeV1alpha1RuntimeClassHandler = node_v1alpha1.ListNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.ListNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.ListNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1ListNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1ListNodeV1beta1RuntimeClassHandler = node_v1beta1.ListNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.ListNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.ListNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ListPolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1ListPolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.ListPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.ListPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ListPolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandler == nil {
		api.PolicyV1beta1ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandler = policy_v1beta1.ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesHandlerFunc(func(params policy_v1beta1.ListPolicyV1beta1PodDisruptionBudgetForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ListPolicyV1beta1PodDisruptionBudgetForAllNamespaces has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ListPolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1ListPolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.ListPolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.ListPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ListPolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ListRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1ListRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.ListRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.ListRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ListRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ListRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1ListRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.ListRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.ListRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ListRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ListRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1ListRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.ListRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.ListRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ListRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ListRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1ListRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.ListRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.ListRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ListRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ListRbacAuthorizationV1RoleBindingForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1ListRbacAuthorizationV1RoleBindingForAllNamespacesHandler = rbac_authorization_v1.ListRbacAuthorizationV1RoleBindingForAllNamespacesHandlerFunc(func(params rbac_authorization_v1.ListRbacAuthorizationV1RoleBindingForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ListRbacAuthorizationV1RoleBindingForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ListRbacAuthorizationV1RoleForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1ListRbacAuthorizationV1RoleForAllNamespacesHandler = rbac_authorization_v1.ListRbacAuthorizationV1RoleForAllNamespacesHandlerFunc(func(params rbac_authorization_v1.ListRbacAuthorizationV1RoleForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ListRbacAuthorizationV1RoleForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandler = rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesHandlerFunc(func(params rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1RoleBindingForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1RoleBindingForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1RoleForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1alpha1ListRbacAuthorizationV1alpha1RoleForAllNamespacesHandler = rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1RoleForAllNamespacesHandlerFunc(func(params rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1RoleForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ListRbacAuthorizationV1alpha1RoleForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1RoleBindingForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1RoleBindingForAllNamespacesHandler = rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1RoleBindingForAllNamespacesHandlerFunc(func(params rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1RoleBindingForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1RoleBindingForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1RoleForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1beta1ListRbacAuthorizationV1beta1RoleForAllNamespacesHandler = rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1RoleForAllNamespacesHandlerFunc(func(params rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1RoleForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ListRbacAuthorizationV1beta1RoleForAllNamespaces has not yet been implemented")
		})
	}
	if api.SchedulingV1ListSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1ListSchedulingV1PriorityClassHandler = scheduling_v1.ListSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.ListSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.ListSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1ListSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1ListSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.ListSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.ListSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.ListSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1ListSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1ListSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.ListSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.ListSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.ListSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1ListSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1ListSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.ListSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.ListSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.ListSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1ListSettingsV1alpha1PodPresetForAllNamespacesHandler == nil {
		api.SettingsV1alpha1ListSettingsV1alpha1PodPresetForAllNamespacesHandler = settings_v1alpha1.ListSettingsV1alpha1PodPresetForAllNamespacesHandlerFunc(func(params settings_v1alpha1.ListSettingsV1alpha1PodPresetForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.ListSettingsV1alpha1PodPresetForAllNamespaces has not yet been implemented")
		})
	}
	if api.StorageV1ListStorageV1CSIDriverHandler == nil {
		api.StorageV1ListStorageV1CSIDriverHandler = storage_v1.ListStorageV1CSIDriverHandlerFunc(func(params storage_v1.ListStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ListStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1ListStorageV1CSINodeHandler == nil {
		api.StorageV1ListStorageV1CSINodeHandler = storage_v1.ListStorageV1CSINodeHandlerFunc(func(params storage_v1.ListStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ListStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1ListStorageV1StorageClassHandler == nil {
		api.StorageV1ListStorageV1StorageClassHandler = storage_v1.ListStorageV1StorageClassHandlerFunc(func(params storage_v1.ListStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ListStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1ListStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1ListStorageV1VolumeAttachmentHandler = storage_v1.ListStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.ListStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ListStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1alpha1ListStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1ListStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.ListStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.ListStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.ListStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1ListStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1ListStorageV1beta1CSIDriverHandler = storage_v1beta1.ListStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.ListStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ListStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1ListStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1ListStorageV1beta1CSINodeHandler = storage_v1beta1.ListStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.ListStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ListStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1ListStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1ListStorageV1beta1StorageClassHandler = storage_v1beta1.ListStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.ListStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ListStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1ListStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1ListStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.ListStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.ListStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ListStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.LogsLogFileHandlerHandler == nil {
		api.LogsLogFileHandlerHandler = logs.LogFileHandlerHandlerFunc(func(params logs.LogFileHandlerParams) middleware.Responder {
			return middleware.NotImplemented("operation logs.LogFileHandler has not yet been implemented")
		})
	}
	if api.LogsLogFileListHandlerHandler == nil {
		api.LogsLogFileListHandlerHandler = logs.LogFileListHandlerHandlerFunc(func(params logs.LogFileListHandlerParams) middleware.Responder {
			return middleware.NotImplemented("operation logs.LogFileListHandler has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1PatchAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1PatchAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.PatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.PatchAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.PatchAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.PatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.PatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.PatchAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1PatchAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1PatchAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.PatchAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.PatchAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.PatchAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.PatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.PatchAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.ApiextensionsV1PatchApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1PatchApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.PatchApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.PatchApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.PatchApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1PatchApiextensionsV1CustomResourceDefinitionStatusHandler == nil {
		api.ApiextensionsV1PatchApiextensionsV1CustomResourceDefinitionStatusHandler = apiextensions_v1.PatchApiextensionsV1CustomResourceDefinitionStatusHandlerFunc(func(params apiextensions_v1.PatchApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.PatchApiextensionsV1CustomResourceDefinitionStatus has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1PatchApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1PatchApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.PatchApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.PatchApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.PatchApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1PatchApiextensionsV1beta1CustomResourceDefinitionStatusHandler == nil {
		api.ApiextensionsV1beta1PatchApiextensionsV1beta1CustomResourceDefinitionStatusHandler = apiextensions_v1beta1.PatchApiextensionsV1beta1CustomResourceDefinitionStatusHandlerFunc(func(params apiextensions_v1beta1.PatchApiextensionsV1beta1CustomResourceDefinitionStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.PatchApiextensionsV1beta1CustomResourceDefinitionStatus has not yet been implemented")
		})
	}
	if api.ApiregistrationV1PatchApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1PatchApiregistrationV1APIServiceHandler = apiregistration_v1.PatchApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.PatchApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.PatchApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1PatchApiregistrationV1APIServiceStatusHandler == nil {
		api.ApiregistrationV1PatchApiregistrationV1APIServiceStatusHandler = apiregistration_v1.PatchApiregistrationV1APIServiceStatusHandlerFunc(func(params apiregistration_v1.PatchApiregistrationV1APIServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.PatchApiregistrationV1APIServiceStatus has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1PatchApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1PatchApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.PatchApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.PatchApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.PatchApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1PatchApiregistrationV1beta1APIServiceStatusHandler == nil {
		api.ApiregistrationV1beta1PatchApiregistrationV1beta1APIServiceStatusHandler = apiregistration_v1beta1.PatchApiregistrationV1beta1APIServiceStatusHandlerFunc(func(params apiregistration_v1beta1.PatchApiregistrationV1beta1APIServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.PatchApiregistrationV1beta1APIServiceStatus has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1PatchAppsV1NamespacedControllerRevisionHandler = apps_v1.PatchAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1PatchAppsV1NamespacedDaemonSetHandler = apps_v1.PatchAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedDaemonSetStatusHandler == nil {
		api.AppsV1PatchAppsV1NamespacedDaemonSetStatusHandler = apps_v1.PatchAppsV1NamespacedDaemonSetStatusHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedDaemonSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedDaemonSetStatus has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1PatchAppsV1NamespacedDeploymentHandler = apps_v1.PatchAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedDeploymentScaleHandler == nil {
		api.AppsV1PatchAppsV1NamespacedDeploymentScaleHandler = apps_v1.PatchAppsV1NamespacedDeploymentScaleHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedDeploymentScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedDeploymentScale has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedDeploymentStatusHandler == nil {
		api.AppsV1PatchAppsV1NamespacedDeploymentStatusHandler = apps_v1.PatchAppsV1NamespacedDeploymentStatusHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedDeploymentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedDeploymentStatus has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1PatchAppsV1NamespacedReplicaSetHandler = apps_v1.PatchAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedReplicaSetScaleHandler == nil {
		api.AppsV1PatchAppsV1NamespacedReplicaSetScaleHandler = apps_v1.PatchAppsV1NamespacedReplicaSetScaleHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedReplicaSetScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedReplicaSetScale has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedReplicaSetStatusHandler == nil {
		api.AppsV1PatchAppsV1NamespacedReplicaSetStatusHandler = apps_v1.PatchAppsV1NamespacedReplicaSetStatusHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedReplicaSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedReplicaSetStatus has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1PatchAppsV1NamespacedStatefulSetHandler = apps_v1.PatchAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedStatefulSetScaleHandler == nil {
		api.AppsV1PatchAppsV1NamespacedStatefulSetScaleHandler = apps_v1.PatchAppsV1NamespacedStatefulSetScaleHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedStatefulSetScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedStatefulSetScale has not yet been implemented")
		})
	}
	if api.AppsV1PatchAppsV1NamespacedStatefulSetStatusHandler == nil {
		api.AppsV1PatchAppsV1NamespacedStatefulSetStatusHandler = apps_v1.PatchAppsV1NamespacedStatefulSetStatusHandlerFunc(func(params apps_v1.PatchAppsV1NamespacedStatefulSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.PatchAppsV1NamespacedStatefulSetStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV1PatchAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1PatchAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.PatchAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.PatchAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.PatchAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV1PatchAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV1PatchAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v1.PatchAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v1.PatchAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.PatchAutoscalingV1NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV2beta1PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v2beta1.PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v2beta1.PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.PatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV2beta2PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v2beta2.PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v2beta2.PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.PatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.BatchV1PatchBatchV1NamespacedJobHandler == nil {
		api.BatchV1PatchBatchV1NamespacedJobHandler = batch_v1.PatchBatchV1NamespacedJobHandlerFunc(func(params batch_v1.PatchBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.PatchBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1PatchBatchV1NamespacedJobStatusHandler == nil {
		api.BatchV1PatchBatchV1NamespacedJobStatusHandler = batch_v1.PatchBatchV1NamespacedJobStatusHandlerFunc(func(params batch_v1.PatchBatchV1NamespacedJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.PatchBatchV1NamespacedJobStatus has not yet been implemented")
		})
	}
	if api.BatchV1beta1PatchBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1PatchBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.PatchBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.PatchBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.PatchBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1PatchBatchV1beta1NamespacedCronJobStatusHandler == nil {
		api.BatchV1beta1PatchBatchV1beta1NamespacedCronJobStatusHandler = batch_v1beta1.PatchBatchV1beta1NamespacedCronJobStatusHandlerFunc(func(params batch_v1beta1.PatchBatchV1beta1NamespacedCronJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.PatchBatchV1beta1NamespacedCronJobStatus has not yet been implemented")
		})
	}
	if api.BatchV2alpha1PatchBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1PatchBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.PatchBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.PatchBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.PatchBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1PatchBatchV2alpha1NamespacedCronJobStatusHandler == nil {
		api.BatchV2alpha1PatchBatchV2alpha1NamespacedCronJobStatusHandler = batch_v2alpha1.PatchBatchV2alpha1NamespacedCronJobStatusHandlerFunc(func(params batch_v2alpha1.PatchBatchV2alpha1NamespacedCronJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.PatchBatchV2alpha1NamespacedCronJobStatus has not yet been implemented")
		})
	}
	if api.CertificatesV1PatchCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1PatchCertificatesV1CertificateSigningRequestHandler = certificates_v1.PatchCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.PatchCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.PatchCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1PatchCertificatesV1CertificateSigningRequestApprovalHandler == nil {
		api.CertificatesV1PatchCertificatesV1CertificateSigningRequestApprovalHandler = certificates_v1.PatchCertificatesV1CertificateSigningRequestApprovalHandlerFunc(func(params certificates_v1.PatchCertificatesV1CertificateSigningRequestApprovalParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.PatchCertificatesV1CertificateSigningRequestApproval has not yet been implemented")
		})
	}
	if api.CertificatesV1PatchCertificatesV1CertificateSigningRequestStatusHandler == nil {
		api.CertificatesV1PatchCertificatesV1CertificateSigningRequestStatusHandler = certificates_v1.PatchCertificatesV1CertificateSigningRequestStatusHandlerFunc(func(params certificates_v1.PatchCertificatesV1CertificateSigningRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.PatchCertificatesV1CertificateSigningRequestStatus has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1PatchCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1PatchCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1PatchCertificatesV1beta1CertificateSigningRequestApprovalHandler == nil {
		api.CertificatesV1beta1PatchCertificatesV1beta1CertificateSigningRequestApprovalHandler = certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestApprovalHandlerFunc(func(params certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestApprovalParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestApproval has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1PatchCertificatesV1beta1CertificateSigningRequestStatusHandler == nil {
		api.CertificatesV1beta1PatchCertificatesV1beta1CertificateSigningRequestStatusHandler = certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestStatusHandlerFunc(func(params certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.PatchCertificatesV1beta1CertificateSigningRequestStatus has not yet been implemented")
		})
	}
	if api.CoordinationV1PatchCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1PatchCoordinationV1NamespacedLeaseHandler = coordination_v1.PatchCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.PatchCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.PatchCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1PatchCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1PatchCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.PatchCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.PatchCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.PatchCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespaceHandler == nil {
		api.CoreV1PatchCoreV1NamespaceHandler = core_v1.PatchCoreV1NamespaceHandlerFunc(func(params core_v1.PatchCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespaceStatusHandler == nil {
		api.CoreV1PatchCoreV1NamespaceStatusHandler = core_v1.PatchCoreV1NamespaceStatusHandlerFunc(func(params core_v1.PatchCoreV1NamespaceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespaceStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1PatchCoreV1NamespacedConfigMapHandler = core_v1.PatchCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.PatchCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1PatchCoreV1NamespacedEndpointsHandler = core_v1.PatchCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.PatchCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedEventHandler == nil {
		api.CoreV1PatchCoreV1NamespacedEventHandler = core_v1.PatchCoreV1NamespacedEventHandlerFunc(func(params core_v1.PatchCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1PatchCoreV1NamespacedLimitRangeHandler = core_v1.PatchCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.PatchCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1PatchCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.PatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.PatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedPersistentVolumeClaimStatusHandler == nil {
		api.CoreV1PatchCoreV1NamespacedPersistentVolumeClaimStatusHandler = core_v1.PatchCoreV1NamespacedPersistentVolumeClaimStatusHandlerFunc(func(params core_v1.PatchCoreV1NamespacedPersistentVolumeClaimStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedPersistentVolumeClaimStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedPodHandler == nil {
		api.CoreV1PatchCoreV1NamespacedPodHandler = core_v1.PatchCoreV1NamespacedPodHandlerFunc(func(params core_v1.PatchCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedPodStatusHandler == nil {
		api.CoreV1PatchCoreV1NamespacedPodStatusHandler = core_v1.PatchCoreV1NamespacedPodStatusHandlerFunc(func(params core_v1.PatchCoreV1NamespacedPodStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedPodStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1PatchCoreV1NamespacedPodTemplateHandler = core_v1.PatchCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.PatchCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1PatchCoreV1NamespacedReplicationControllerHandler = core_v1.PatchCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.PatchCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedReplicationControllerScaleHandler == nil {
		api.CoreV1PatchCoreV1NamespacedReplicationControllerScaleHandler = core_v1.PatchCoreV1NamespacedReplicationControllerScaleHandlerFunc(func(params core_v1.PatchCoreV1NamespacedReplicationControllerScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedReplicationControllerScale has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedReplicationControllerStatusHandler == nil {
		api.CoreV1PatchCoreV1NamespacedReplicationControllerStatusHandler = core_v1.PatchCoreV1NamespacedReplicationControllerStatusHandlerFunc(func(params core_v1.PatchCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedReplicationControllerStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1PatchCoreV1NamespacedResourceQuotaHandler = core_v1.PatchCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.PatchCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedResourceQuotaStatusHandler == nil {
		api.CoreV1PatchCoreV1NamespacedResourceQuotaStatusHandler = core_v1.PatchCoreV1NamespacedResourceQuotaStatusHandlerFunc(func(params core_v1.PatchCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedResourceQuotaStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedSecretHandler == nil {
		api.CoreV1PatchCoreV1NamespacedSecretHandler = core_v1.PatchCoreV1NamespacedSecretHandlerFunc(func(params core_v1.PatchCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedServiceHandler == nil {
		api.CoreV1PatchCoreV1NamespacedServiceHandler = core_v1.PatchCoreV1NamespacedServiceHandlerFunc(func(params core_v1.PatchCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1PatchCoreV1NamespacedServiceAccountHandler = core_v1.PatchCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.PatchCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NamespacedServiceStatusHandler == nil {
		api.CoreV1PatchCoreV1NamespacedServiceStatusHandler = core_v1.PatchCoreV1NamespacedServiceStatusHandlerFunc(func(params core_v1.PatchCoreV1NamespacedServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NamespacedServiceStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NodeHandler == nil {
		api.CoreV1PatchCoreV1NodeHandler = core_v1.PatchCoreV1NodeHandlerFunc(func(params core_v1.PatchCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1NodeStatusHandler == nil {
		api.CoreV1PatchCoreV1NodeStatusHandler = core_v1.PatchCoreV1NodeStatusHandlerFunc(func(params core_v1.PatchCoreV1NodeStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1NodeStatus has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1PersistentVolumeHandler == nil {
		api.CoreV1PatchCoreV1PersistentVolumeHandler = core_v1.PatchCoreV1PersistentVolumeHandlerFunc(func(params core_v1.PatchCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.CoreV1PatchCoreV1PersistentVolumeStatusHandler == nil {
		api.CoreV1PatchCoreV1PersistentVolumeStatusHandler = core_v1.PatchCoreV1PersistentVolumeStatusHandlerFunc(func(params core_v1.PatchCoreV1PersistentVolumeStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.PatchCoreV1PersistentVolumeStatus has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1PatchDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1PatchDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.PatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.PatchDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.PatchDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.EventsV1PatchEventsV1NamespacedEventHandler == nil {
		api.EventsV1PatchEventsV1NamespacedEventHandler = events_v1.PatchEventsV1NamespacedEventHandlerFunc(func(params events_v1.PatchEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.PatchEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1PatchEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1PatchEventsV1beta1NamespacedEventHandler = events_v1beta1.PatchEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.PatchEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.PatchEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1PatchExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1PatchExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.PatchExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.PatchExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.PatchExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1PatchExtensionsV1beta1NamespacedIngressStatusHandler == nil {
		api.ExtensionsV1beta1PatchExtensionsV1beta1NamespacedIngressStatusHandler = extensions_v1beta1.PatchExtensionsV1beta1NamespacedIngressStatusHandlerFunc(func(params extensions_v1beta1.PatchExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.PatchExtensionsV1beta1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler == nil {
		api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler = flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1FlowSchemaStatus has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler == nil {
		api.FlowcontrolApiserverV1alpha1PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler = flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.PatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus has not yet been implemented")
		})
	}
	if api.NetworkingV1PatchNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1PatchNetworkingV1IngressClassHandler = networking_v1.PatchNetworkingV1IngressClassHandlerFunc(func(params networking_v1.PatchNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.PatchNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1PatchNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1PatchNetworkingV1NamespacedIngressHandler = networking_v1.PatchNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.PatchNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.PatchNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1PatchNetworkingV1NamespacedIngressStatusHandler == nil {
		api.NetworkingV1PatchNetworkingV1NamespacedIngressStatusHandler = networking_v1.PatchNetworkingV1NamespacedIngressStatusHandlerFunc(func(params networking_v1.PatchNetworkingV1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.PatchNetworkingV1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.NetworkingV1PatchNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1PatchNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.PatchNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.PatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.PatchNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1PatchNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1PatchNetworkingV1beta1IngressClassHandler = networking_v1beta1.PatchNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.PatchNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.PatchNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1PatchNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1PatchNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.PatchNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.PatchNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.PatchNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1PatchNetworkingV1beta1NamespacedIngressStatusHandler == nil {
		api.NetworkingV1beta1PatchNetworkingV1beta1NamespacedIngressStatusHandler = networking_v1beta1.PatchNetworkingV1beta1NamespacedIngressStatusHandlerFunc(func(params networking_v1beta1.PatchNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.PatchNetworkingV1beta1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.NodeV1alpha1PatchNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1PatchNodeV1alpha1RuntimeClassHandler = node_v1alpha1.PatchNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.PatchNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.PatchNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1PatchNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1PatchNodeV1beta1RuntimeClassHandler = node_v1beta1.PatchNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.PatchNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.PatchNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.PolicyV1beta1PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.PatchPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.PatchPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.PatchPolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler == nil {
		api.PolicyV1beta1PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler = policy_v1beta1.PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc(func(params policy_v1beta1.PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.PatchPolicyV1beta1NamespacedPodDisruptionBudgetStatus has not yet been implemented")
		})
	}
	if api.PolicyV1beta1PatchPolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1PatchPolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.PatchPolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.PatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.PatchPolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1PatchRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1PatchRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.PatchRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.PatchRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.PatchRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1PatchRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1PatchRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.PatchRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.PatchRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.PatchRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1PatchRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1PatchRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.PatchRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.PatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.PatchRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1PatchRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1PatchRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.PatchRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.PatchRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.PatchRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.PatchRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1PatchRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.PatchRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.SchedulingV1PatchSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1PatchSchedulingV1PriorityClassHandler = scheduling_v1.PatchSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.PatchSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.PatchSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1PatchSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1PatchSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.PatchSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.PatchSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.PatchSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1PatchSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1PatchSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.PatchSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.PatchSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.PatchSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1PatchSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1PatchSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.PatchSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.PatchSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.PatchSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.StorageV1PatchStorageV1CSIDriverHandler == nil {
		api.StorageV1PatchStorageV1CSIDriverHandler = storage_v1.PatchStorageV1CSIDriverHandlerFunc(func(params storage_v1.PatchStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.PatchStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1PatchStorageV1CSINodeHandler == nil {
		api.StorageV1PatchStorageV1CSINodeHandler = storage_v1.PatchStorageV1CSINodeHandlerFunc(func(params storage_v1.PatchStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.PatchStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1PatchStorageV1StorageClassHandler == nil {
		api.StorageV1PatchStorageV1StorageClassHandler = storage_v1.PatchStorageV1StorageClassHandlerFunc(func(params storage_v1.PatchStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.PatchStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1PatchStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1PatchStorageV1VolumeAttachmentHandler = storage_v1.PatchStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.PatchStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.PatchStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1PatchStorageV1VolumeAttachmentStatusHandler == nil {
		api.StorageV1PatchStorageV1VolumeAttachmentStatusHandler = storage_v1.PatchStorageV1VolumeAttachmentStatusHandlerFunc(func(params storage_v1.PatchStorageV1VolumeAttachmentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.PatchStorageV1VolumeAttachmentStatus has not yet been implemented")
		})
	}
	if api.StorageV1alpha1PatchStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1PatchStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.PatchStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.PatchStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.PatchStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1PatchStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1PatchStorageV1beta1CSIDriverHandler = storage_v1beta1.PatchStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.PatchStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.PatchStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1PatchStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1PatchStorageV1beta1CSINodeHandler = storage_v1beta1.PatchStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.PatchStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.PatchStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1PatchStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1PatchStorageV1beta1StorageClassHandler = storage_v1beta1.PatchStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.PatchStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.PatchStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1PatchStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1PatchStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.PatchStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.PatchStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.PatchStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1ReadAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1ReadAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.ReadAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.ReadAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.ReadAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1ReadAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1ReadAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.ReadAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.ReadAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.ReadAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1ReadAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1ReadAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.ReadAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.ReadAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.ReadAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.ReadAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.ReadAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.ApiextensionsV1ReadApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1ReadApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.ReadApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.ReadApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.ReadApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1ReadApiextensionsV1CustomResourceDefinitionStatusHandler == nil {
		api.ApiextensionsV1ReadApiextensionsV1CustomResourceDefinitionStatusHandler = apiextensions_v1.ReadApiextensionsV1CustomResourceDefinitionStatusHandlerFunc(func(params apiextensions_v1.ReadApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.ReadApiextensionsV1CustomResourceDefinitionStatus has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1ReadApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1ReadApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.ReadApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.ReadApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.ReadApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1ReadApiextensionsV1beta1CustomResourceDefinitionStatusHandler == nil {
		api.ApiextensionsV1beta1ReadApiextensionsV1beta1CustomResourceDefinitionStatusHandler = apiextensions_v1beta1.ReadApiextensionsV1beta1CustomResourceDefinitionStatusHandlerFunc(func(params apiextensions_v1beta1.ReadApiextensionsV1beta1CustomResourceDefinitionStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.ReadApiextensionsV1beta1CustomResourceDefinitionStatus has not yet been implemented")
		})
	}
	if api.ApiregistrationV1ReadApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1ReadApiregistrationV1APIServiceHandler = apiregistration_v1.ReadApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.ReadApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.ReadApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1ReadApiregistrationV1APIServiceStatusHandler == nil {
		api.ApiregistrationV1ReadApiregistrationV1APIServiceStatusHandler = apiregistration_v1.ReadApiregistrationV1APIServiceStatusHandlerFunc(func(params apiregistration_v1.ReadApiregistrationV1APIServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.ReadApiregistrationV1APIServiceStatus has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1ReadApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1ReadApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.ReadApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.ReadApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.ReadApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1ReadApiregistrationV1beta1APIServiceStatusHandler == nil {
		api.ApiregistrationV1beta1ReadApiregistrationV1beta1APIServiceStatusHandler = apiregistration_v1beta1.ReadApiregistrationV1beta1APIServiceStatusHandlerFunc(func(params apiregistration_v1beta1.ReadApiregistrationV1beta1APIServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.ReadApiregistrationV1beta1APIServiceStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1ReadAppsV1NamespacedControllerRevisionHandler = apps_v1.ReadAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1ReadAppsV1NamespacedDaemonSetHandler = apps_v1.ReadAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedDaemonSetStatusHandler == nil {
		api.AppsV1ReadAppsV1NamespacedDaemonSetStatusHandler = apps_v1.ReadAppsV1NamespacedDaemonSetStatusHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedDaemonSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedDaemonSetStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1ReadAppsV1NamespacedDeploymentHandler = apps_v1.ReadAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedDeploymentScaleHandler == nil {
		api.AppsV1ReadAppsV1NamespacedDeploymentScaleHandler = apps_v1.ReadAppsV1NamespacedDeploymentScaleHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedDeploymentScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedDeploymentScale has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedDeploymentStatusHandler == nil {
		api.AppsV1ReadAppsV1NamespacedDeploymentStatusHandler = apps_v1.ReadAppsV1NamespacedDeploymentStatusHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedDeploymentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedDeploymentStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1ReadAppsV1NamespacedReplicaSetHandler = apps_v1.ReadAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedReplicaSetScaleHandler == nil {
		api.AppsV1ReadAppsV1NamespacedReplicaSetScaleHandler = apps_v1.ReadAppsV1NamespacedReplicaSetScaleHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedReplicaSetScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedReplicaSetScale has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedReplicaSetStatusHandler == nil {
		api.AppsV1ReadAppsV1NamespacedReplicaSetStatusHandler = apps_v1.ReadAppsV1NamespacedReplicaSetStatusHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedReplicaSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedReplicaSetStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1ReadAppsV1NamespacedStatefulSetHandler = apps_v1.ReadAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedStatefulSetScaleHandler == nil {
		api.AppsV1ReadAppsV1NamespacedStatefulSetScaleHandler = apps_v1.ReadAppsV1NamespacedStatefulSetScaleHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedStatefulSetScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedStatefulSetScale has not yet been implemented")
		})
	}
	if api.AppsV1ReadAppsV1NamespacedStatefulSetStatusHandler == nil {
		api.AppsV1ReadAppsV1NamespacedStatefulSetStatusHandler = apps_v1.ReadAppsV1NamespacedStatefulSetStatusHandlerFunc(func(params apps_v1.ReadAppsV1NamespacedStatefulSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReadAppsV1NamespacedStatefulSetStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV1ReadAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1ReadAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.ReadAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.ReadAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.ReadAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV1ReadAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV1ReadAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v1.ReadAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v1.ReadAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.ReadAutoscalingV1NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV2beta1ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v2beta1.ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v2beta1.ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.ReadAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV2beta2ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v2beta2.ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v2beta2.ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.ReadAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.BatchV1ReadBatchV1NamespacedJobHandler == nil {
		api.BatchV1ReadBatchV1NamespacedJobHandler = batch_v1.ReadBatchV1NamespacedJobHandlerFunc(func(params batch_v1.ReadBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.ReadBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1ReadBatchV1NamespacedJobStatusHandler == nil {
		api.BatchV1ReadBatchV1NamespacedJobStatusHandler = batch_v1.ReadBatchV1NamespacedJobStatusHandlerFunc(func(params batch_v1.ReadBatchV1NamespacedJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.ReadBatchV1NamespacedJobStatus has not yet been implemented")
		})
	}
	if api.BatchV1beta1ReadBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1ReadBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.ReadBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.ReadBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.ReadBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1ReadBatchV1beta1NamespacedCronJobStatusHandler == nil {
		api.BatchV1beta1ReadBatchV1beta1NamespacedCronJobStatusHandler = batch_v1beta1.ReadBatchV1beta1NamespacedCronJobStatusHandlerFunc(func(params batch_v1beta1.ReadBatchV1beta1NamespacedCronJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.ReadBatchV1beta1NamespacedCronJobStatus has not yet been implemented")
		})
	}
	if api.BatchV2alpha1ReadBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1ReadBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.ReadBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.ReadBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.ReadBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1ReadBatchV2alpha1NamespacedCronJobStatusHandler == nil {
		api.BatchV2alpha1ReadBatchV2alpha1NamespacedCronJobStatusHandler = batch_v2alpha1.ReadBatchV2alpha1NamespacedCronJobStatusHandlerFunc(func(params batch_v2alpha1.ReadBatchV2alpha1NamespacedCronJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.ReadBatchV2alpha1NamespacedCronJobStatus has not yet been implemented")
		})
	}
	if api.CertificatesV1ReadCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1ReadCertificatesV1CertificateSigningRequestHandler = certificates_v1.ReadCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.ReadCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ReadCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1ReadCertificatesV1CertificateSigningRequestApprovalHandler == nil {
		api.CertificatesV1ReadCertificatesV1CertificateSigningRequestApprovalHandler = certificates_v1.ReadCertificatesV1CertificateSigningRequestApprovalHandlerFunc(func(params certificates_v1.ReadCertificatesV1CertificateSigningRequestApprovalParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ReadCertificatesV1CertificateSigningRequestApproval has not yet been implemented")
		})
	}
	if api.CertificatesV1ReadCertificatesV1CertificateSigningRequestStatusHandler == nil {
		api.CertificatesV1ReadCertificatesV1CertificateSigningRequestStatusHandler = certificates_v1.ReadCertificatesV1CertificateSigningRequestStatusHandlerFunc(func(params certificates_v1.ReadCertificatesV1CertificateSigningRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ReadCertificatesV1CertificateSigningRequestStatus has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ReadCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1ReadCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ReadCertificatesV1beta1CertificateSigningRequestApprovalHandler == nil {
		api.CertificatesV1beta1ReadCertificatesV1beta1CertificateSigningRequestApprovalHandler = certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestApprovalHandlerFunc(func(params certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestApprovalParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestApproval has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ReadCertificatesV1beta1CertificateSigningRequestStatusHandler == nil {
		api.CertificatesV1beta1ReadCertificatesV1beta1CertificateSigningRequestStatusHandler = certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestStatusHandlerFunc(func(params certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ReadCertificatesV1beta1CertificateSigningRequestStatus has not yet been implemented")
		})
	}
	if api.CoordinationV1ReadCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1ReadCoordinationV1NamespacedLeaseHandler = coordination_v1.ReadCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.ReadCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.ReadCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1ReadCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1ReadCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.ReadCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.ReadCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.ReadCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1ComponentStatusHandler == nil {
		api.CoreV1ReadCoreV1ComponentStatusHandler = core_v1.ReadCoreV1ComponentStatusHandlerFunc(func(params core_v1.ReadCoreV1ComponentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1ComponentStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespaceHandler == nil {
		api.CoreV1ReadCoreV1NamespaceHandler = core_v1.ReadCoreV1NamespaceHandlerFunc(func(params core_v1.ReadCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespaceStatusHandler == nil {
		api.CoreV1ReadCoreV1NamespaceStatusHandler = core_v1.ReadCoreV1NamespaceStatusHandlerFunc(func(params core_v1.ReadCoreV1NamespaceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespaceStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1ReadCoreV1NamespacedConfigMapHandler = core_v1.ReadCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.ReadCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1ReadCoreV1NamespacedEndpointsHandler = core_v1.ReadCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.ReadCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedEventHandler == nil {
		api.CoreV1ReadCoreV1NamespacedEventHandler = core_v1.ReadCoreV1NamespacedEventHandlerFunc(func(params core_v1.ReadCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1ReadCoreV1NamespacedLimitRangeHandler = core_v1.ReadCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.ReadCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1ReadCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.ReadCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.ReadCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedPersistentVolumeClaimStatusHandler == nil {
		api.CoreV1ReadCoreV1NamespacedPersistentVolumeClaimStatusHandler = core_v1.ReadCoreV1NamespacedPersistentVolumeClaimStatusHandlerFunc(func(params core_v1.ReadCoreV1NamespacedPersistentVolumeClaimStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedPersistentVolumeClaimStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedPodHandler == nil {
		api.CoreV1ReadCoreV1NamespacedPodHandler = core_v1.ReadCoreV1NamespacedPodHandlerFunc(func(params core_v1.ReadCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedPodLogHandler == nil {
		api.CoreV1ReadCoreV1NamespacedPodLogHandler = core_v1.ReadCoreV1NamespacedPodLogHandlerFunc(func(params core_v1.ReadCoreV1NamespacedPodLogParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedPodLog has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedPodStatusHandler == nil {
		api.CoreV1ReadCoreV1NamespacedPodStatusHandler = core_v1.ReadCoreV1NamespacedPodStatusHandlerFunc(func(params core_v1.ReadCoreV1NamespacedPodStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedPodStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1ReadCoreV1NamespacedPodTemplateHandler = core_v1.ReadCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.ReadCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1ReadCoreV1NamespacedReplicationControllerHandler = core_v1.ReadCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.ReadCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedReplicationControllerScaleHandler == nil {
		api.CoreV1ReadCoreV1NamespacedReplicationControllerScaleHandler = core_v1.ReadCoreV1NamespacedReplicationControllerScaleHandlerFunc(func(params core_v1.ReadCoreV1NamespacedReplicationControllerScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedReplicationControllerScale has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedReplicationControllerStatusHandler == nil {
		api.CoreV1ReadCoreV1NamespacedReplicationControllerStatusHandler = core_v1.ReadCoreV1NamespacedReplicationControllerStatusHandlerFunc(func(params core_v1.ReadCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedReplicationControllerStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1ReadCoreV1NamespacedResourceQuotaHandler = core_v1.ReadCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.ReadCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedResourceQuotaStatusHandler == nil {
		api.CoreV1ReadCoreV1NamespacedResourceQuotaStatusHandler = core_v1.ReadCoreV1NamespacedResourceQuotaStatusHandlerFunc(func(params core_v1.ReadCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedResourceQuotaStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedSecretHandler == nil {
		api.CoreV1ReadCoreV1NamespacedSecretHandler = core_v1.ReadCoreV1NamespacedSecretHandlerFunc(func(params core_v1.ReadCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedServiceHandler == nil {
		api.CoreV1ReadCoreV1NamespacedServiceHandler = core_v1.ReadCoreV1NamespacedServiceHandlerFunc(func(params core_v1.ReadCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1ReadCoreV1NamespacedServiceAccountHandler = core_v1.ReadCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.ReadCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NamespacedServiceStatusHandler == nil {
		api.CoreV1ReadCoreV1NamespacedServiceStatusHandler = core_v1.ReadCoreV1NamespacedServiceStatusHandlerFunc(func(params core_v1.ReadCoreV1NamespacedServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NamespacedServiceStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NodeHandler == nil {
		api.CoreV1ReadCoreV1NodeHandler = core_v1.ReadCoreV1NodeHandlerFunc(func(params core_v1.ReadCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1NodeStatusHandler == nil {
		api.CoreV1ReadCoreV1NodeStatusHandler = core_v1.ReadCoreV1NodeStatusHandlerFunc(func(params core_v1.ReadCoreV1NodeStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1NodeStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1PersistentVolumeHandler == nil {
		api.CoreV1ReadCoreV1PersistentVolumeHandler = core_v1.ReadCoreV1PersistentVolumeHandlerFunc(func(params core_v1.ReadCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.CoreV1ReadCoreV1PersistentVolumeStatusHandler == nil {
		api.CoreV1ReadCoreV1PersistentVolumeStatusHandler = core_v1.ReadCoreV1PersistentVolumeStatusHandlerFunc(func(params core_v1.ReadCoreV1PersistentVolumeStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReadCoreV1PersistentVolumeStatus has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1ReadDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1ReadDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.ReadDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.ReadDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.ReadDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.EventsV1ReadEventsV1NamespacedEventHandler == nil {
		api.EventsV1ReadEventsV1NamespacedEventHandler = events_v1.ReadEventsV1NamespacedEventHandlerFunc(func(params events_v1.ReadEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.ReadEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1ReadEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1ReadEventsV1beta1NamespacedEventHandler = events_v1beta1.ReadEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.ReadEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.ReadEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1ReadExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1ReadExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.ReadExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.ReadExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.ReadExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1ReadExtensionsV1beta1NamespacedIngressStatusHandler == nil {
		api.ExtensionsV1beta1ReadExtensionsV1beta1NamespacedIngressStatusHandler = extensions_v1beta1.ReadExtensionsV1beta1NamespacedIngressStatusHandlerFunc(func(params extensions_v1beta1.ReadExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.ReadExtensionsV1beta1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler = flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1FlowSchemaStatus has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler = flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReadFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus has not yet been implemented")
		})
	}
	if api.NetworkingV1ReadNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1ReadNetworkingV1IngressClassHandler = networking_v1.ReadNetworkingV1IngressClassHandlerFunc(func(params networking_v1.ReadNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReadNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1ReadNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1ReadNetworkingV1NamespacedIngressHandler = networking_v1.ReadNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.ReadNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReadNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1ReadNetworkingV1NamespacedIngressStatusHandler == nil {
		api.NetworkingV1ReadNetworkingV1NamespacedIngressStatusHandler = networking_v1.ReadNetworkingV1NamespacedIngressStatusHandlerFunc(func(params networking_v1.ReadNetworkingV1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReadNetworkingV1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.NetworkingV1ReadNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1ReadNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.ReadNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.ReadNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReadNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ReadNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1ReadNetworkingV1beta1IngressClassHandler = networking_v1beta1.ReadNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.ReadNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ReadNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ReadNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1ReadNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.ReadNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.ReadNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ReadNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ReadNetworkingV1beta1NamespacedIngressStatusHandler == nil {
		api.NetworkingV1beta1ReadNetworkingV1beta1NamespacedIngressStatusHandler = networking_v1beta1.ReadNetworkingV1beta1NamespacedIngressStatusHandlerFunc(func(params networking_v1beta1.ReadNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ReadNetworkingV1beta1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.NodeV1alpha1ReadNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1ReadNodeV1alpha1RuntimeClassHandler = node_v1alpha1.ReadNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.ReadNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.ReadNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1ReadNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1ReadNodeV1beta1RuntimeClassHandler = node_v1beta1.ReadNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.ReadNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.ReadNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ReadPolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1ReadPolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.ReadPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.ReadPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ReadPolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler == nil {
		api.PolicyV1beta1ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler = policy_v1beta1.ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc(func(params policy_v1beta1.ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ReadPolicyV1beta1NamespacedPodDisruptionBudgetStatus has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ReadPolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1ReadPolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.ReadPolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.ReadPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ReadPolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReadRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1ReadRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.ReadRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.ReadRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReadRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReadRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1ReadRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.ReadRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.ReadRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReadRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReadRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1ReadRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.ReadRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.ReadRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReadRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReadRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1ReadRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.ReadRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.ReadRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReadRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1ReadRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReadRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1ReadRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReadRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.SchedulingV1ReadSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1ReadSchedulingV1PriorityClassHandler = scheduling_v1.ReadSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.ReadSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.ReadSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1ReadSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1ReadSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.ReadSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.ReadSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.ReadSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1ReadSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1ReadSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.ReadSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.ReadSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.ReadSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1ReadSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1ReadSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.ReadSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.ReadSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.ReadSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.StorageV1ReadStorageV1CSIDriverHandler == nil {
		api.StorageV1ReadStorageV1CSIDriverHandler = storage_v1.ReadStorageV1CSIDriverHandlerFunc(func(params storage_v1.ReadStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReadStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1ReadStorageV1CSINodeHandler == nil {
		api.StorageV1ReadStorageV1CSINodeHandler = storage_v1.ReadStorageV1CSINodeHandlerFunc(func(params storage_v1.ReadStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReadStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1ReadStorageV1StorageClassHandler == nil {
		api.StorageV1ReadStorageV1StorageClassHandler = storage_v1.ReadStorageV1StorageClassHandlerFunc(func(params storage_v1.ReadStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReadStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1ReadStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1ReadStorageV1VolumeAttachmentHandler = storage_v1.ReadStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.ReadStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReadStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1ReadStorageV1VolumeAttachmentStatusHandler == nil {
		api.StorageV1ReadStorageV1VolumeAttachmentStatusHandler = storage_v1.ReadStorageV1VolumeAttachmentStatusHandlerFunc(func(params storage_v1.ReadStorageV1VolumeAttachmentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReadStorageV1VolumeAttachmentStatus has not yet been implemented")
		})
	}
	if api.StorageV1alpha1ReadStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1ReadStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.ReadStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.ReadStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.ReadStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReadStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1ReadStorageV1beta1CSIDriverHandler = storage_v1beta1.ReadStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.ReadStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReadStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReadStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1ReadStorageV1beta1CSINodeHandler = storage_v1beta1.ReadStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.ReadStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReadStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReadStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1ReadStorageV1beta1StorageClassHandler = storage_v1beta1.ReadStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.ReadStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReadStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReadStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1ReadStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.ReadStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.ReadStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReadStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.ReplaceAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.ReplaceAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.ReplaceAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1ReplaceAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1ReplaceAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.ReplaceAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.ReplaceAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.ReplaceAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.ReplaceAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.ReplaceAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.ReplaceAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.ApiextensionsV1ReplaceApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1ReplaceApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.ReplaceApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.ReplaceApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.ReplaceApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1ReplaceApiextensionsV1CustomResourceDefinitionStatusHandler == nil {
		api.ApiextensionsV1ReplaceApiextensionsV1CustomResourceDefinitionStatusHandler = apiextensions_v1.ReplaceApiextensionsV1CustomResourceDefinitionStatusHandlerFunc(func(params apiextensions_v1.ReplaceApiextensionsV1CustomResourceDefinitionStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.ReplaceApiextensionsV1CustomResourceDefinitionStatus has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1ReplaceApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1ReplaceApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.ReplaceApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.ReplaceApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.ReplaceApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusHandler == nil {
		api.ApiextensionsV1beta1ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusHandler = apiextensions_v1beta1.ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusHandlerFunc(func(params apiextensions_v1beta1.ReplaceApiextensionsV1beta1CustomResourceDefinitionStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.ReplaceApiextensionsV1beta1CustomResourceDefinitionStatus has not yet been implemented")
		})
	}
	if api.ApiregistrationV1ReplaceApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1ReplaceApiregistrationV1APIServiceHandler = apiregistration_v1.ReplaceApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.ReplaceApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.ReplaceApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1ReplaceApiregistrationV1APIServiceStatusHandler == nil {
		api.ApiregistrationV1ReplaceApiregistrationV1APIServiceStatusHandler = apiregistration_v1.ReplaceApiregistrationV1APIServiceStatusHandlerFunc(func(params apiregistration_v1.ReplaceApiregistrationV1APIServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.ReplaceApiregistrationV1APIServiceStatus has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1ReplaceApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1ReplaceApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.ReplaceApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.ReplaceApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.ReplaceApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1ReplaceApiregistrationV1beta1APIServiceStatusHandler == nil {
		api.ApiregistrationV1beta1ReplaceApiregistrationV1beta1APIServiceStatusHandler = apiregistration_v1beta1.ReplaceApiregistrationV1beta1APIServiceStatusHandlerFunc(func(params apiregistration_v1beta1.ReplaceApiregistrationV1beta1APIServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.ReplaceApiregistrationV1beta1APIServiceStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedControllerRevisionHandler = apps_v1.ReplaceAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedDaemonSetHandler = apps_v1.ReplaceAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedDaemonSetStatusHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedDaemonSetStatusHandler = apps_v1.ReplaceAppsV1NamespacedDaemonSetStatusHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedDaemonSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedDaemonSetStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedDeploymentHandler = apps_v1.ReplaceAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedDeploymentScaleHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedDeploymentScaleHandler = apps_v1.ReplaceAppsV1NamespacedDeploymentScaleHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedDeploymentScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedDeploymentScale has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedDeploymentStatusHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedDeploymentStatusHandler = apps_v1.ReplaceAppsV1NamespacedDeploymentStatusHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedDeploymentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedDeploymentStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedReplicaSetHandler = apps_v1.ReplaceAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedReplicaSetScaleHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedReplicaSetScaleHandler = apps_v1.ReplaceAppsV1NamespacedReplicaSetScaleHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedReplicaSetScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedReplicaSetScale has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedReplicaSetStatusHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedReplicaSetStatusHandler = apps_v1.ReplaceAppsV1NamespacedReplicaSetStatusHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedReplicaSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedReplicaSetStatus has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedStatefulSetHandler = apps_v1.ReplaceAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedStatefulSetScaleHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedStatefulSetScaleHandler = apps_v1.ReplaceAppsV1NamespacedStatefulSetScaleHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedStatefulSetScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedStatefulSetScale has not yet been implemented")
		})
	}
	if api.AppsV1ReplaceAppsV1NamespacedStatefulSetStatusHandler == nil {
		api.AppsV1ReplaceAppsV1NamespacedStatefulSetStatusHandler = apps_v1.ReplaceAppsV1NamespacedStatefulSetStatusHandlerFunc(func(params apps_v1.ReplaceAppsV1NamespacedStatefulSetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.ReplaceAppsV1NamespacedStatefulSetStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV1ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.ReplaceAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV1ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV1ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v1.ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v1.ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.ReplaceAutoscalingV1NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV2beta1ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v2beta1.ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v2beta1.ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.ReplaceAutoscalingV2beta1NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler == nil {
		api.AutoscalingV2beta2ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandler = autoscaling_v2beta2.ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusHandlerFunc(func(params autoscaling_v2beta2.ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.ReplaceAutoscalingV2beta2NamespacedHorizontalPodAutoscalerStatus has not yet been implemented")
		})
	}
	if api.BatchV1ReplaceBatchV1NamespacedJobHandler == nil {
		api.BatchV1ReplaceBatchV1NamespacedJobHandler = batch_v1.ReplaceBatchV1NamespacedJobHandlerFunc(func(params batch_v1.ReplaceBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.ReplaceBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1ReplaceBatchV1NamespacedJobStatusHandler == nil {
		api.BatchV1ReplaceBatchV1NamespacedJobStatusHandler = batch_v1.ReplaceBatchV1NamespacedJobStatusHandlerFunc(func(params batch_v1.ReplaceBatchV1NamespacedJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.ReplaceBatchV1NamespacedJobStatus has not yet been implemented")
		})
	}
	if api.BatchV1beta1ReplaceBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1ReplaceBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.ReplaceBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.ReplaceBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.ReplaceBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1ReplaceBatchV1beta1NamespacedCronJobStatusHandler == nil {
		api.BatchV1beta1ReplaceBatchV1beta1NamespacedCronJobStatusHandler = batch_v1beta1.ReplaceBatchV1beta1NamespacedCronJobStatusHandlerFunc(func(params batch_v1beta1.ReplaceBatchV1beta1NamespacedCronJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.ReplaceBatchV1beta1NamespacedCronJobStatus has not yet been implemented")
		})
	}
	if api.BatchV2alpha1ReplaceBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1ReplaceBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.ReplaceBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.ReplaceBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.ReplaceBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1ReplaceBatchV2alpha1NamespacedCronJobStatusHandler == nil {
		api.BatchV2alpha1ReplaceBatchV2alpha1NamespacedCronJobStatusHandler = batch_v2alpha1.ReplaceBatchV2alpha1NamespacedCronJobStatusHandlerFunc(func(params batch_v2alpha1.ReplaceBatchV2alpha1NamespacedCronJobStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.ReplaceBatchV2alpha1NamespacedCronJobStatus has not yet been implemented")
		})
	}
	if api.CertificatesV1ReplaceCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1ReplaceCertificatesV1CertificateSigningRequestHandler = certificates_v1.ReplaceCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.ReplaceCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ReplaceCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1ReplaceCertificatesV1CertificateSigningRequestApprovalHandler == nil {
		api.CertificatesV1ReplaceCertificatesV1CertificateSigningRequestApprovalHandler = certificates_v1.ReplaceCertificatesV1CertificateSigningRequestApprovalHandlerFunc(func(params certificates_v1.ReplaceCertificatesV1CertificateSigningRequestApprovalParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ReplaceCertificatesV1CertificateSigningRequestApproval has not yet been implemented")
		})
	}
	if api.CertificatesV1ReplaceCertificatesV1CertificateSigningRequestStatusHandler == nil {
		api.CertificatesV1ReplaceCertificatesV1CertificateSigningRequestStatusHandler = certificates_v1.ReplaceCertificatesV1CertificateSigningRequestStatusHandlerFunc(func(params certificates_v1.ReplaceCertificatesV1CertificateSigningRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.ReplaceCertificatesV1CertificateSigningRequestStatus has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ReplaceCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1ReplaceCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ReplaceCertificatesV1beta1CertificateSigningRequestApprovalHandler == nil {
		api.CertificatesV1beta1ReplaceCertificatesV1beta1CertificateSigningRequestApprovalHandler = certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestApprovalHandlerFunc(func(params certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestApprovalParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestApproval has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1ReplaceCertificatesV1beta1CertificateSigningRequestStatusHandler == nil {
		api.CertificatesV1beta1ReplaceCertificatesV1beta1CertificateSigningRequestStatusHandler = certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestStatusHandlerFunc(func(params certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.ReplaceCertificatesV1beta1CertificateSigningRequestStatus has not yet been implemented")
		})
	}
	if api.CoordinationV1ReplaceCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1ReplaceCoordinationV1NamespacedLeaseHandler = coordination_v1.ReplaceCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.ReplaceCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.ReplaceCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1ReplaceCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1ReplaceCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.ReplaceCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.ReplaceCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.ReplaceCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespaceHandler == nil {
		api.CoreV1ReplaceCoreV1NamespaceHandler = core_v1.ReplaceCoreV1NamespaceHandlerFunc(func(params core_v1.ReplaceCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespaceFinalizeHandler == nil {
		api.CoreV1ReplaceCoreV1NamespaceFinalizeHandler = core_v1.ReplaceCoreV1NamespaceFinalizeHandlerFunc(func(params core_v1.ReplaceCoreV1NamespaceFinalizeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespaceFinalize has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespaceStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NamespaceStatusHandler = core_v1.ReplaceCoreV1NamespaceStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NamespaceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespaceStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedConfigMapHandler = core_v1.ReplaceCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedEndpointsHandler = core_v1.ReplaceCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedEventHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedEventHandler = core_v1.ReplaceCoreV1NamespacedEventHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedLimitRangeHandler = core_v1.ReplaceCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.ReplaceCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandler = core_v1.ReplaceCoreV1NamespacedPersistentVolumeClaimStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedPersistentVolumeClaimStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedPersistentVolumeClaimStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedPodHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedPodHandler = core_v1.ReplaceCoreV1NamespacedPodHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedPodStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedPodStatusHandler = core_v1.ReplaceCoreV1NamespacedPodStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedPodStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedPodStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedPodTemplateHandler = core_v1.ReplaceCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedReplicationControllerHandler = core_v1.ReplaceCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedReplicationControllerScaleHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedReplicationControllerScaleHandler = core_v1.ReplaceCoreV1NamespacedReplicationControllerScaleHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedReplicationControllerScaleParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedReplicationControllerScale has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedReplicationControllerStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedReplicationControllerStatusHandler = core_v1.ReplaceCoreV1NamespacedReplicationControllerStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedReplicationControllerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedReplicationControllerStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedResourceQuotaHandler = core_v1.ReplaceCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedResourceQuotaStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedResourceQuotaStatusHandler = core_v1.ReplaceCoreV1NamespacedResourceQuotaStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedResourceQuotaStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedResourceQuotaStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedSecretHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedSecretHandler = core_v1.ReplaceCoreV1NamespacedSecretHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedServiceHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedServiceHandler = core_v1.ReplaceCoreV1NamespacedServiceHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedServiceAccountHandler = core_v1.ReplaceCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NamespacedServiceStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NamespacedServiceStatusHandler = core_v1.ReplaceCoreV1NamespacedServiceStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NamespacedServiceStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NamespacedServiceStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NodeHandler == nil {
		api.CoreV1ReplaceCoreV1NodeHandler = core_v1.ReplaceCoreV1NodeHandlerFunc(func(params core_v1.ReplaceCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1NodeStatusHandler == nil {
		api.CoreV1ReplaceCoreV1NodeStatusHandler = core_v1.ReplaceCoreV1NodeStatusHandlerFunc(func(params core_v1.ReplaceCoreV1NodeStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1NodeStatus has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1PersistentVolumeHandler == nil {
		api.CoreV1ReplaceCoreV1PersistentVolumeHandler = core_v1.ReplaceCoreV1PersistentVolumeHandlerFunc(func(params core_v1.ReplaceCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.CoreV1ReplaceCoreV1PersistentVolumeStatusHandler == nil {
		api.CoreV1ReplaceCoreV1PersistentVolumeStatusHandler = core_v1.ReplaceCoreV1PersistentVolumeStatusHandlerFunc(func(params core_v1.ReplaceCoreV1PersistentVolumeStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.ReplaceCoreV1PersistentVolumeStatus has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1ReplaceDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1ReplaceDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.ReplaceDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.ReplaceDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.ReplaceDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.EventsV1ReplaceEventsV1NamespacedEventHandler == nil {
		api.EventsV1ReplaceEventsV1NamespacedEventHandler = events_v1.ReplaceEventsV1NamespacedEventHandlerFunc(func(params events_v1.ReplaceEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.ReplaceEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1ReplaceEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1ReplaceEventsV1beta1NamespacedEventHandler = events_v1beta1.ReplaceEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.ReplaceEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.ReplaceEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1ReplaceExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1ReplaceExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.ReplaceExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.ReplaceExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.ReplaceExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1ReplaceExtensionsV1beta1NamespacedIngressStatusHandler == nil {
		api.ExtensionsV1beta1ReplaceExtensionsV1beta1NamespacedIngressStatusHandler = extensions_v1beta1.ReplaceExtensionsV1beta1NamespacedIngressStatusHandlerFunc(func(params extensions_v1beta1.ReplaceExtensionsV1beta1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.ReplaceExtensionsV1beta1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandler = flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1FlowSchemaStatus has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler == nil {
		api.FlowcontrolApiserverV1alpha1ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandler = flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.ReplaceFlowcontrolApiserverV1alpha1PriorityLevelConfigurationStatus has not yet been implemented")
		})
	}
	if api.NetworkingV1ReplaceNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1ReplaceNetworkingV1IngressClassHandler = networking_v1.ReplaceNetworkingV1IngressClassHandlerFunc(func(params networking_v1.ReplaceNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReplaceNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1ReplaceNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1ReplaceNetworkingV1NamespacedIngressHandler = networking_v1.ReplaceNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.ReplaceNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReplaceNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1ReplaceNetworkingV1NamespacedIngressStatusHandler == nil {
		api.NetworkingV1ReplaceNetworkingV1NamespacedIngressStatusHandler = networking_v1.ReplaceNetworkingV1NamespacedIngressStatusHandlerFunc(func(params networking_v1.ReplaceNetworkingV1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReplaceNetworkingV1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.NetworkingV1ReplaceNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1ReplaceNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.ReplaceNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.ReplaceNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.ReplaceNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ReplaceNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1ReplaceNetworkingV1beta1IngressClassHandler = networking_v1beta1.ReplaceNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.ReplaceNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ReplaceNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ReplaceNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1ReplaceNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.ReplaceNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.ReplaceNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ReplaceNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1ReplaceNetworkingV1beta1NamespacedIngressStatusHandler == nil {
		api.NetworkingV1beta1ReplaceNetworkingV1beta1NamespacedIngressStatusHandler = networking_v1beta1.ReplaceNetworkingV1beta1NamespacedIngressStatusHandlerFunc(func(params networking_v1beta1.ReplaceNetworkingV1beta1NamespacedIngressStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.ReplaceNetworkingV1beta1NamespacedIngressStatus has not yet been implemented")
		})
	}
	if api.NodeV1alpha1ReplaceNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1ReplaceNodeV1alpha1RuntimeClassHandler = node_v1alpha1.ReplaceNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.ReplaceNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.ReplaceNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1ReplaceNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1ReplaceNodeV1beta1RuntimeClassHandler = node_v1beta1.ReplaceNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.ReplaceNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.ReplaceNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ReplacePolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1ReplacePolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.ReplacePolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.ReplacePolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ReplacePolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler == nil {
		api.PolicyV1beta1ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusHandler = policy_v1beta1.ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusHandlerFunc(func(params policy_v1beta1.ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ReplacePolicyV1beta1NamespacedPodDisruptionBudgetStatus has not yet been implemented")
		})
	}
	if api.PolicyV1beta1ReplacePolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1ReplacePolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.ReplacePolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.ReplacePolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.ReplacePolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReplaceRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1ReplaceRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.ReplaceRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.ReplaceRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReplaceRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReplaceRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1ReplaceRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.ReplaceRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.ReplaceRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReplaceRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReplaceRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1ReplaceRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.ReplaceRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.ReplaceRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReplaceRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1ReplaceRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1ReplaceRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.ReplaceRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.ReplaceRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.ReplaceRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1ReplaceRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.ReplaceRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.ReplaceRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.SchedulingV1ReplaceSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1ReplaceSchedulingV1PriorityClassHandler = scheduling_v1.ReplaceSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.ReplaceSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.ReplaceSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1ReplaceSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1ReplaceSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.ReplaceSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.ReplaceSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.ReplaceSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1ReplaceSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1ReplaceSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.ReplaceSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.ReplaceSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.ReplaceSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1ReplaceSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1ReplaceSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.ReplaceSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.ReplaceSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.ReplaceSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.StorageV1ReplaceStorageV1CSIDriverHandler == nil {
		api.StorageV1ReplaceStorageV1CSIDriverHandler = storage_v1.ReplaceStorageV1CSIDriverHandlerFunc(func(params storage_v1.ReplaceStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReplaceStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1ReplaceStorageV1CSINodeHandler == nil {
		api.StorageV1ReplaceStorageV1CSINodeHandler = storage_v1.ReplaceStorageV1CSINodeHandlerFunc(func(params storage_v1.ReplaceStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReplaceStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1ReplaceStorageV1StorageClassHandler == nil {
		api.StorageV1ReplaceStorageV1StorageClassHandler = storage_v1.ReplaceStorageV1StorageClassHandlerFunc(func(params storage_v1.ReplaceStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReplaceStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1ReplaceStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1ReplaceStorageV1VolumeAttachmentHandler = storage_v1.ReplaceStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.ReplaceStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReplaceStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1ReplaceStorageV1VolumeAttachmentStatusHandler == nil {
		api.StorageV1ReplaceStorageV1VolumeAttachmentStatusHandler = storage_v1.ReplaceStorageV1VolumeAttachmentStatusHandlerFunc(func(params storage_v1.ReplaceStorageV1VolumeAttachmentStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.ReplaceStorageV1VolumeAttachmentStatus has not yet been implemented")
		})
	}
	if api.StorageV1alpha1ReplaceStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1ReplaceStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.ReplaceStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.ReplaceStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.ReplaceStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReplaceStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1ReplaceStorageV1beta1CSIDriverHandler = storage_v1beta1.ReplaceStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.ReplaceStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReplaceStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReplaceStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1ReplaceStorageV1beta1CSINodeHandler = storage_v1beta1.ReplaceStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.ReplaceStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReplaceStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReplaceStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1ReplaceStorageV1beta1StorageClassHandler = storage_v1beta1.ReplaceStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.ReplaceStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReplaceStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1ReplaceStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1ReplaceStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.ReplaceStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.ReplaceStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.ReplaceStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1WatchAdmissionregistrationV1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1WatchAdmissionregistrationV1MutatingWebhookConfigurationHandler = admissionregistration_v1.WatchAdmissionregistrationV1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.WatchAdmissionregistrationV1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.WatchAdmissionregistrationV1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandler == nil {
		api.AdmissionregistrationV1WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandler = admissionregistration_v1.WatchAdmissionregistrationV1MutatingWebhookConfigurationListHandlerFunc(func(params admissionregistration_v1.WatchAdmissionregistrationV1MutatingWebhookConfigurationListParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.WatchAdmissionregistrationV1MutatingWebhookConfigurationList has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandler = admissionregistration_v1.WatchAdmissionregistrationV1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1.WatchAdmissionregistrationV1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.WatchAdmissionregistrationV1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1WatchAdmissionregistrationV1ValidatingWebhookConfigurationListHandler == nil {
		api.AdmissionregistrationV1WatchAdmissionregistrationV1ValidatingWebhookConfigurationListHandler = admissionregistration_v1.WatchAdmissionregistrationV1ValidatingWebhookConfigurationListHandlerFunc(func(params admissionregistration_v1.WatchAdmissionregistrationV1ValidatingWebhookConfigurationListParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1.WatchAdmissionregistrationV1ValidatingWebhookConfigurationList has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationHandler = admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1MutatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationListHandler == nil {
		api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationListHandler = admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationListHandlerFunc(func(params admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationListParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1MutatingWebhookConfigurationList has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler == nil {
		api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandler = admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationHandlerFunc(func(params admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1ValidatingWebhookConfiguration has not yet been implemented")
		})
	}
	if api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandler == nil {
		api.AdmissionregistrationV1beta1WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandler = admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListHandlerFunc(func(params admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationListParams) middleware.Responder {
			return middleware.NotImplemented("operation admissionregistration_v1beta1.WatchAdmissionregistrationV1beta1ValidatingWebhookConfigurationList has not yet been implemented")
		})
	}
	if api.ApiextensionsV1WatchApiextensionsV1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1WatchApiextensionsV1CustomResourceDefinitionHandler = apiextensions_v1.WatchApiextensionsV1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1.WatchApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.WatchApiextensionsV1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1WatchApiextensionsV1CustomResourceDefinitionListHandler == nil {
		api.ApiextensionsV1WatchApiextensionsV1CustomResourceDefinitionListHandler = apiextensions_v1.WatchApiextensionsV1CustomResourceDefinitionListHandlerFunc(func(params apiextensions_v1.WatchApiextensionsV1CustomResourceDefinitionListParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1.WatchApiextensionsV1CustomResourceDefinitionList has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1WatchApiextensionsV1beta1CustomResourceDefinitionHandler == nil {
		api.ApiextensionsV1beta1WatchApiextensionsV1beta1CustomResourceDefinitionHandler = apiextensions_v1beta1.WatchApiextensionsV1beta1CustomResourceDefinitionHandlerFunc(func(params apiextensions_v1beta1.WatchApiextensionsV1beta1CustomResourceDefinitionParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.WatchApiextensionsV1beta1CustomResourceDefinition has not yet been implemented")
		})
	}
	if api.ApiextensionsV1beta1WatchApiextensionsV1beta1CustomResourceDefinitionListHandler == nil {
		api.ApiextensionsV1beta1WatchApiextensionsV1beta1CustomResourceDefinitionListHandler = apiextensions_v1beta1.WatchApiextensionsV1beta1CustomResourceDefinitionListHandlerFunc(func(params apiextensions_v1beta1.WatchApiextensionsV1beta1CustomResourceDefinitionListParams) middleware.Responder {
			return middleware.NotImplemented("operation apiextensions_v1beta1.WatchApiextensionsV1beta1CustomResourceDefinitionList has not yet been implemented")
		})
	}
	if api.ApiregistrationV1WatchApiregistrationV1APIServiceHandler == nil {
		api.ApiregistrationV1WatchApiregistrationV1APIServiceHandler = apiregistration_v1.WatchApiregistrationV1APIServiceHandlerFunc(func(params apiregistration_v1.WatchApiregistrationV1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.WatchApiregistrationV1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1WatchApiregistrationV1APIServiceListHandler == nil {
		api.ApiregistrationV1WatchApiregistrationV1APIServiceListHandler = apiregistration_v1.WatchApiregistrationV1APIServiceListHandlerFunc(func(params apiregistration_v1.WatchApiregistrationV1APIServiceListParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1.WatchApiregistrationV1APIServiceList has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1WatchApiregistrationV1beta1APIServiceHandler == nil {
		api.ApiregistrationV1beta1WatchApiregistrationV1beta1APIServiceHandler = apiregistration_v1beta1.WatchApiregistrationV1beta1APIServiceHandlerFunc(func(params apiregistration_v1beta1.WatchApiregistrationV1beta1APIServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.WatchApiregistrationV1beta1APIService has not yet been implemented")
		})
	}
	if api.ApiregistrationV1beta1WatchApiregistrationV1beta1APIServiceListHandler == nil {
		api.ApiregistrationV1beta1WatchApiregistrationV1beta1APIServiceListHandler = apiregistration_v1beta1.WatchApiregistrationV1beta1APIServiceListHandlerFunc(func(params apiregistration_v1beta1.WatchApiregistrationV1beta1APIServiceListParams) middleware.Responder {
			return middleware.NotImplemented("operation apiregistration_v1beta1.WatchApiregistrationV1beta1APIServiceList has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1ControllerRevisionListForAllNamespacesHandler == nil {
		api.AppsV1WatchAppsV1ControllerRevisionListForAllNamespacesHandler = apps_v1.WatchAppsV1ControllerRevisionListForAllNamespacesHandlerFunc(func(params apps_v1.WatchAppsV1ControllerRevisionListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1ControllerRevisionListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1DaemonSetListForAllNamespacesHandler == nil {
		api.AppsV1WatchAppsV1DaemonSetListForAllNamespacesHandler = apps_v1.WatchAppsV1DaemonSetListForAllNamespacesHandlerFunc(func(params apps_v1.WatchAppsV1DaemonSetListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1DaemonSetListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1DeploymentListForAllNamespacesHandler == nil {
		api.AppsV1WatchAppsV1DeploymentListForAllNamespacesHandler = apps_v1.WatchAppsV1DeploymentListForAllNamespacesHandlerFunc(func(params apps_v1.WatchAppsV1DeploymentListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1DeploymentListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedControllerRevisionHandler == nil {
		api.AppsV1WatchAppsV1NamespacedControllerRevisionHandler = apps_v1.WatchAppsV1NamespacedControllerRevisionHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedControllerRevisionParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedControllerRevision has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedControllerRevisionListHandler == nil {
		api.AppsV1WatchAppsV1NamespacedControllerRevisionListHandler = apps_v1.WatchAppsV1NamespacedControllerRevisionListHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedControllerRevisionListParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedControllerRevisionList has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedDaemonSetHandler == nil {
		api.AppsV1WatchAppsV1NamespacedDaemonSetHandler = apps_v1.WatchAppsV1NamespacedDaemonSetHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedDaemonSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedDaemonSet has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedDaemonSetListHandler == nil {
		api.AppsV1WatchAppsV1NamespacedDaemonSetListHandler = apps_v1.WatchAppsV1NamespacedDaemonSetListHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedDaemonSetListParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedDaemonSetList has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedDeploymentHandler == nil {
		api.AppsV1WatchAppsV1NamespacedDeploymentHandler = apps_v1.WatchAppsV1NamespacedDeploymentHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedDeploymentParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedDeployment has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedDeploymentListHandler == nil {
		api.AppsV1WatchAppsV1NamespacedDeploymentListHandler = apps_v1.WatchAppsV1NamespacedDeploymentListHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedDeploymentListParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedDeploymentList has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedReplicaSetHandler == nil {
		api.AppsV1WatchAppsV1NamespacedReplicaSetHandler = apps_v1.WatchAppsV1NamespacedReplicaSetHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedReplicaSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedReplicaSet has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedReplicaSetListHandler == nil {
		api.AppsV1WatchAppsV1NamespacedReplicaSetListHandler = apps_v1.WatchAppsV1NamespacedReplicaSetListHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedReplicaSetListParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedReplicaSetList has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedStatefulSetHandler == nil {
		api.AppsV1WatchAppsV1NamespacedStatefulSetHandler = apps_v1.WatchAppsV1NamespacedStatefulSetHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedStatefulSetParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedStatefulSet has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1NamespacedStatefulSetListHandler == nil {
		api.AppsV1WatchAppsV1NamespacedStatefulSetListHandler = apps_v1.WatchAppsV1NamespacedStatefulSetListHandlerFunc(func(params apps_v1.WatchAppsV1NamespacedStatefulSetListParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1NamespacedStatefulSetList has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1ReplicaSetListForAllNamespacesHandler == nil {
		api.AppsV1WatchAppsV1ReplicaSetListForAllNamespacesHandler = apps_v1.WatchAppsV1ReplicaSetListForAllNamespacesHandlerFunc(func(params apps_v1.WatchAppsV1ReplicaSetListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1ReplicaSetListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AppsV1WatchAppsV1StatefulSetListForAllNamespacesHandler == nil {
		api.AppsV1WatchAppsV1StatefulSetListForAllNamespacesHandler = apps_v1.WatchAppsV1StatefulSetListForAllNamespacesHandlerFunc(func(params apps_v1.WatchAppsV1StatefulSetListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation apps_v1.WatchAppsV1StatefulSetListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV1WatchAutoscalingV1HorizontalPodAutoscalerListForAllNamespacesHandler == nil {
		api.AutoscalingV1WatchAutoscalingV1HorizontalPodAutoscalerListForAllNamespacesHandler = autoscaling_v1.WatchAutoscalingV1HorizontalPodAutoscalerListForAllNamespacesHandlerFunc(func(params autoscaling_v1.WatchAutoscalingV1HorizontalPodAutoscalerListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.WatchAutoscalingV1HorizontalPodAutoscalerListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV1WatchAutoscalingV1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV1WatchAutoscalingV1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v1.WatchAutoscalingV1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v1.WatchAutoscalingV1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.WatchAutoscalingV1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV1WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandler == nil {
		api.AutoscalingV1WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandler = autoscaling_v1.WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListHandlerFunc(func(params autoscaling_v1.WatchAutoscalingV1NamespacedHorizontalPodAutoscalerListParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v1.WatchAutoscalingV1NamespacedHorizontalPodAutoscalerList has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandler == nil {
		api.AutoscalingV2beta1WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandler = autoscaling_v2beta1.WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesHandlerFunc(func(params autoscaling_v2beta1.WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.WatchAutoscalingV2beta1HorizontalPodAutoscalerListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta1WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta1.WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta1.WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta1WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandler == nil {
		api.AutoscalingV2beta1WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandler = autoscaling_v2beta1.WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListHandlerFunc(func(params autoscaling_v2beta1.WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerListParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta1.WatchAutoscalingV2beta1NamespacedHorizontalPodAutoscalerList has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandler == nil {
		api.AutoscalingV2beta2WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandler = autoscaling_v2beta2.WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesHandlerFunc(func(params autoscaling_v2beta2.WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.WatchAutoscalingV2beta2HorizontalPodAutoscalerListForAllNamespaces has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler == nil {
		api.AutoscalingV2beta2WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandler = autoscaling_v2beta2.WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerHandlerFunc(func(params autoscaling_v2beta2.WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscaler has not yet been implemented")
		})
	}
	if api.AutoscalingV2beta2WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerListHandler == nil {
		api.AutoscalingV2beta2WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerListHandler = autoscaling_v2beta2.WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerListHandlerFunc(func(params autoscaling_v2beta2.WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerListParams) middleware.Responder {
			return middleware.NotImplemented("operation autoscaling_v2beta2.WatchAutoscalingV2beta2NamespacedHorizontalPodAutoscalerList has not yet been implemented")
		})
	}
	if api.BatchV1WatchBatchV1JobListForAllNamespacesHandler == nil {
		api.BatchV1WatchBatchV1JobListForAllNamespacesHandler = batch_v1.WatchBatchV1JobListForAllNamespacesHandlerFunc(func(params batch_v1.WatchBatchV1JobListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.WatchBatchV1JobListForAllNamespaces has not yet been implemented")
		})
	}
	if api.BatchV1WatchBatchV1NamespacedJobHandler == nil {
		api.BatchV1WatchBatchV1NamespacedJobHandler = batch_v1.WatchBatchV1NamespacedJobHandlerFunc(func(params batch_v1.WatchBatchV1NamespacedJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.WatchBatchV1NamespacedJob has not yet been implemented")
		})
	}
	if api.BatchV1WatchBatchV1NamespacedJobListHandler == nil {
		api.BatchV1WatchBatchV1NamespacedJobListHandler = batch_v1.WatchBatchV1NamespacedJobListHandlerFunc(func(params batch_v1.WatchBatchV1NamespacedJobListParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1.WatchBatchV1NamespacedJobList has not yet been implemented")
		})
	}
	if api.BatchV1beta1WatchBatchV1beta1CronJobListForAllNamespacesHandler == nil {
		api.BatchV1beta1WatchBatchV1beta1CronJobListForAllNamespacesHandler = batch_v1beta1.WatchBatchV1beta1CronJobListForAllNamespacesHandlerFunc(func(params batch_v1beta1.WatchBatchV1beta1CronJobListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.WatchBatchV1beta1CronJobListForAllNamespaces has not yet been implemented")
		})
	}
	if api.BatchV1beta1WatchBatchV1beta1NamespacedCronJobHandler == nil {
		api.BatchV1beta1WatchBatchV1beta1NamespacedCronJobHandler = batch_v1beta1.WatchBatchV1beta1NamespacedCronJobHandlerFunc(func(params batch_v1beta1.WatchBatchV1beta1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.WatchBatchV1beta1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV1beta1WatchBatchV1beta1NamespacedCronJobListHandler == nil {
		api.BatchV1beta1WatchBatchV1beta1NamespacedCronJobListHandler = batch_v1beta1.WatchBatchV1beta1NamespacedCronJobListHandlerFunc(func(params batch_v1beta1.WatchBatchV1beta1NamespacedCronJobListParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v1beta1.WatchBatchV1beta1NamespacedCronJobList has not yet been implemented")
		})
	}
	if api.BatchV2alpha1WatchBatchV2alpha1CronJobListForAllNamespacesHandler == nil {
		api.BatchV2alpha1WatchBatchV2alpha1CronJobListForAllNamespacesHandler = batch_v2alpha1.WatchBatchV2alpha1CronJobListForAllNamespacesHandlerFunc(func(params batch_v2alpha1.WatchBatchV2alpha1CronJobListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.WatchBatchV2alpha1CronJobListForAllNamespaces has not yet been implemented")
		})
	}
	if api.BatchV2alpha1WatchBatchV2alpha1NamespacedCronJobHandler == nil {
		api.BatchV2alpha1WatchBatchV2alpha1NamespacedCronJobHandler = batch_v2alpha1.WatchBatchV2alpha1NamespacedCronJobHandlerFunc(func(params batch_v2alpha1.WatchBatchV2alpha1NamespacedCronJobParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.WatchBatchV2alpha1NamespacedCronJob has not yet been implemented")
		})
	}
	if api.BatchV2alpha1WatchBatchV2alpha1NamespacedCronJobListHandler == nil {
		api.BatchV2alpha1WatchBatchV2alpha1NamespacedCronJobListHandler = batch_v2alpha1.WatchBatchV2alpha1NamespacedCronJobListHandlerFunc(func(params batch_v2alpha1.WatchBatchV2alpha1NamespacedCronJobListParams) middleware.Responder {
			return middleware.NotImplemented("operation batch_v2alpha1.WatchBatchV2alpha1NamespacedCronJobList has not yet been implemented")
		})
	}
	if api.CertificatesV1WatchCertificatesV1CertificateSigningRequestHandler == nil {
		api.CertificatesV1WatchCertificatesV1CertificateSigningRequestHandler = certificates_v1.WatchCertificatesV1CertificateSigningRequestHandlerFunc(func(params certificates_v1.WatchCertificatesV1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.WatchCertificatesV1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1WatchCertificatesV1CertificateSigningRequestListHandler == nil {
		api.CertificatesV1WatchCertificatesV1CertificateSigningRequestListHandler = certificates_v1.WatchCertificatesV1CertificateSigningRequestListHandlerFunc(func(params certificates_v1.WatchCertificatesV1CertificateSigningRequestListParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1.WatchCertificatesV1CertificateSigningRequestList has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1WatchCertificatesV1beta1CertificateSigningRequestHandler == nil {
		api.CertificatesV1beta1WatchCertificatesV1beta1CertificateSigningRequestHandler = certificates_v1beta1.WatchCertificatesV1beta1CertificateSigningRequestHandlerFunc(func(params certificates_v1beta1.WatchCertificatesV1beta1CertificateSigningRequestParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.WatchCertificatesV1beta1CertificateSigningRequest has not yet been implemented")
		})
	}
	if api.CertificatesV1beta1WatchCertificatesV1beta1CertificateSigningRequestListHandler == nil {
		api.CertificatesV1beta1WatchCertificatesV1beta1CertificateSigningRequestListHandler = certificates_v1beta1.WatchCertificatesV1beta1CertificateSigningRequestListHandlerFunc(func(params certificates_v1beta1.WatchCertificatesV1beta1CertificateSigningRequestListParams) middleware.Responder {
			return middleware.NotImplemented("operation certificates_v1beta1.WatchCertificatesV1beta1CertificateSigningRequestList has not yet been implemented")
		})
	}
	if api.CoordinationV1WatchCoordinationV1LeaseListForAllNamespacesHandler == nil {
		api.CoordinationV1WatchCoordinationV1LeaseListForAllNamespacesHandler = coordination_v1.WatchCoordinationV1LeaseListForAllNamespacesHandlerFunc(func(params coordination_v1.WatchCoordinationV1LeaseListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.WatchCoordinationV1LeaseListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoordinationV1WatchCoordinationV1NamespacedLeaseHandler == nil {
		api.CoordinationV1WatchCoordinationV1NamespacedLeaseHandler = coordination_v1.WatchCoordinationV1NamespacedLeaseHandlerFunc(func(params coordination_v1.WatchCoordinationV1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.WatchCoordinationV1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1WatchCoordinationV1NamespacedLeaseListHandler == nil {
		api.CoordinationV1WatchCoordinationV1NamespacedLeaseListHandler = coordination_v1.WatchCoordinationV1NamespacedLeaseListHandlerFunc(func(params coordination_v1.WatchCoordinationV1NamespacedLeaseListParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1.WatchCoordinationV1NamespacedLeaseList has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1WatchCoordinationV1beta1LeaseListForAllNamespacesHandler == nil {
		api.CoordinationV1beta1WatchCoordinationV1beta1LeaseListForAllNamespacesHandler = coordination_v1beta1.WatchCoordinationV1beta1LeaseListForAllNamespacesHandlerFunc(func(params coordination_v1beta1.WatchCoordinationV1beta1LeaseListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.WatchCoordinationV1beta1LeaseListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1WatchCoordinationV1beta1NamespacedLeaseHandler == nil {
		api.CoordinationV1beta1WatchCoordinationV1beta1NamespacedLeaseHandler = coordination_v1beta1.WatchCoordinationV1beta1NamespacedLeaseHandlerFunc(func(params coordination_v1beta1.WatchCoordinationV1beta1NamespacedLeaseParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.WatchCoordinationV1beta1NamespacedLease has not yet been implemented")
		})
	}
	if api.CoordinationV1beta1WatchCoordinationV1beta1NamespacedLeaseListHandler == nil {
		api.CoordinationV1beta1WatchCoordinationV1beta1NamespacedLeaseListHandler = coordination_v1beta1.WatchCoordinationV1beta1NamespacedLeaseListHandlerFunc(func(params coordination_v1beta1.WatchCoordinationV1beta1NamespacedLeaseListParams) middleware.Responder {
			return middleware.NotImplemented("operation coordination_v1beta1.WatchCoordinationV1beta1NamespacedLeaseList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1ConfigMapListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1ConfigMapListForAllNamespacesHandler = core_v1.WatchCoreV1ConfigMapListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1ConfigMapListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1ConfigMapListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1EndpointsListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1EndpointsListForAllNamespacesHandler = core_v1.WatchCoreV1EndpointsListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1EndpointsListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1EndpointsListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1EventListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1EventListForAllNamespacesHandler = core_v1.WatchCoreV1EventListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1EventListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1EventListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1LimitRangeListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1LimitRangeListForAllNamespacesHandler = core_v1.WatchCoreV1LimitRangeListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1LimitRangeListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1LimitRangeListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespaceHandler == nil {
		api.CoreV1WatchCoreV1NamespaceHandler = core_v1.WatchCoreV1NamespaceHandlerFunc(func(params core_v1.WatchCoreV1NamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1Namespace has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespaceListHandler == nil {
		api.CoreV1WatchCoreV1NamespaceListHandler = core_v1.WatchCoreV1NamespaceListHandlerFunc(func(params core_v1.WatchCoreV1NamespaceListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespaceList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedConfigMapHandler == nil {
		api.CoreV1WatchCoreV1NamespacedConfigMapHandler = core_v1.WatchCoreV1NamespacedConfigMapHandlerFunc(func(params core_v1.WatchCoreV1NamespacedConfigMapParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedConfigMap has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedConfigMapListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedConfigMapListHandler = core_v1.WatchCoreV1NamespacedConfigMapListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedConfigMapListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedConfigMapList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedEndpointsHandler == nil {
		api.CoreV1WatchCoreV1NamespacedEndpointsHandler = core_v1.WatchCoreV1NamespacedEndpointsHandlerFunc(func(params core_v1.WatchCoreV1NamespacedEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedEndpoints has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedEndpointsListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedEndpointsListHandler = core_v1.WatchCoreV1NamespacedEndpointsListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedEndpointsListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedEndpointsList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedEventHandler == nil {
		api.CoreV1WatchCoreV1NamespacedEventHandler = core_v1.WatchCoreV1NamespacedEventHandlerFunc(func(params core_v1.WatchCoreV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedEventListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedEventListHandler = core_v1.WatchCoreV1NamespacedEventListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedEventListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedEventList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedLimitRangeHandler == nil {
		api.CoreV1WatchCoreV1NamespacedLimitRangeHandler = core_v1.WatchCoreV1NamespacedLimitRangeHandlerFunc(func(params core_v1.WatchCoreV1NamespacedLimitRangeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedLimitRange has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedLimitRangeListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedLimitRangeListHandler = core_v1.WatchCoreV1NamespacedLimitRangeListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedLimitRangeListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedLimitRangeList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedPersistentVolumeClaimHandler == nil {
		api.CoreV1WatchCoreV1NamespacedPersistentVolumeClaimHandler = core_v1.WatchCoreV1NamespacedPersistentVolumeClaimHandlerFunc(func(params core_v1.WatchCoreV1NamespacedPersistentVolumeClaimParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedPersistentVolumeClaim has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedPersistentVolumeClaimListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedPersistentVolumeClaimListHandler = core_v1.WatchCoreV1NamespacedPersistentVolumeClaimListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedPersistentVolumeClaimListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedPersistentVolumeClaimList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedPodHandler == nil {
		api.CoreV1WatchCoreV1NamespacedPodHandler = core_v1.WatchCoreV1NamespacedPodHandlerFunc(func(params core_v1.WatchCoreV1NamespacedPodParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedPod has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedPodListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedPodListHandler = core_v1.WatchCoreV1NamespacedPodListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedPodListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedPodList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedPodTemplateHandler == nil {
		api.CoreV1WatchCoreV1NamespacedPodTemplateHandler = core_v1.WatchCoreV1NamespacedPodTemplateHandlerFunc(func(params core_v1.WatchCoreV1NamespacedPodTemplateParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedPodTemplate has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedPodTemplateListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedPodTemplateListHandler = core_v1.WatchCoreV1NamespacedPodTemplateListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedPodTemplateListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedPodTemplateList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedReplicationControllerHandler == nil {
		api.CoreV1WatchCoreV1NamespacedReplicationControllerHandler = core_v1.WatchCoreV1NamespacedReplicationControllerHandlerFunc(func(params core_v1.WatchCoreV1NamespacedReplicationControllerParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedReplicationController has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedReplicationControllerListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedReplicationControllerListHandler = core_v1.WatchCoreV1NamespacedReplicationControllerListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedReplicationControllerListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedReplicationControllerList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedResourceQuotaHandler == nil {
		api.CoreV1WatchCoreV1NamespacedResourceQuotaHandler = core_v1.WatchCoreV1NamespacedResourceQuotaHandlerFunc(func(params core_v1.WatchCoreV1NamespacedResourceQuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedResourceQuota has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedResourceQuotaListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedResourceQuotaListHandler = core_v1.WatchCoreV1NamespacedResourceQuotaListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedResourceQuotaListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedResourceQuotaList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedSecretHandler == nil {
		api.CoreV1WatchCoreV1NamespacedSecretHandler = core_v1.WatchCoreV1NamespacedSecretHandlerFunc(func(params core_v1.WatchCoreV1NamespacedSecretParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedSecret has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedSecretListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedSecretListHandler = core_v1.WatchCoreV1NamespacedSecretListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedSecretListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedSecretList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedServiceHandler == nil {
		api.CoreV1WatchCoreV1NamespacedServiceHandler = core_v1.WatchCoreV1NamespacedServiceHandlerFunc(func(params core_v1.WatchCoreV1NamespacedServiceParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedService has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedServiceAccountHandler == nil {
		api.CoreV1WatchCoreV1NamespacedServiceAccountHandler = core_v1.WatchCoreV1NamespacedServiceAccountHandlerFunc(func(params core_v1.WatchCoreV1NamespacedServiceAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedServiceAccount has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedServiceAccountListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedServiceAccountListHandler = core_v1.WatchCoreV1NamespacedServiceAccountListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedServiceAccountListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedServiceAccountList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NamespacedServiceListHandler == nil {
		api.CoreV1WatchCoreV1NamespacedServiceListHandler = core_v1.WatchCoreV1NamespacedServiceListHandlerFunc(func(params core_v1.WatchCoreV1NamespacedServiceListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NamespacedServiceList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NodeHandler == nil {
		api.CoreV1WatchCoreV1NodeHandler = core_v1.WatchCoreV1NodeHandlerFunc(func(params core_v1.WatchCoreV1NodeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1Node has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1NodeListHandler == nil {
		api.CoreV1WatchCoreV1NodeListHandler = core_v1.WatchCoreV1NodeListHandlerFunc(func(params core_v1.WatchCoreV1NodeListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1NodeList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1PersistentVolumeHandler == nil {
		api.CoreV1WatchCoreV1PersistentVolumeHandler = core_v1.WatchCoreV1PersistentVolumeHandlerFunc(func(params core_v1.WatchCoreV1PersistentVolumeParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1PersistentVolume has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1PersistentVolumeClaimListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1PersistentVolumeClaimListForAllNamespacesHandler = core_v1.WatchCoreV1PersistentVolumeClaimListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1PersistentVolumeClaimListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1PersistentVolumeClaimListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1PersistentVolumeListHandler == nil {
		api.CoreV1WatchCoreV1PersistentVolumeListHandler = core_v1.WatchCoreV1PersistentVolumeListHandlerFunc(func(params core_v1.WatchCoreV1PersistentVolumeListParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1PersistentVolumeList has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1PodListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1PodListForAllNamespacesHandler = core_v1.WatchCoreV1PodListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1PodListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1PodListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1PodTemplateListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1PodTemplateListForAllNamespacesHandler = core_v1.WatchCoreV1PodTemplateListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1PodTemplateListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1PodTemplateListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1ReplicationControllerListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1ReplicationControllerListForAllNamespacesHandler = core_v1.WatchCoreV1ReplicationControllerListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1ReplicationControllerListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1ReplicationControllerListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1ResourceQuotaListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1ResourceQuotaListForAllNamespacesHandler = core_v1.WatchCoreV1ResourceQuotaListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1ResourceQuotaListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1ResourceQuotaListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1SecretListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1SecretListForAllNamespacesHandler = core_v1.WatchCoreV1SecretListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1SecretListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1SecretListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1ServiceAccountListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1ServiceAccountListForAllNamespacesHandler = core_v1.WatchCoreV1ServiceAccountListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1ServiceAccountListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1ServiceAccountListForAllNamespaces has not yet been implemented")
		})
	}
	if api.CoreV1WatchCoreV1ServiceListForAllNamespacesHandler == nil {
		api.CoreV1WatchCoreV1ServiceListForAllNamespacesHandler = core_v1.WatchCoreV1ServiceListForAllNamespacesHandlerFunc(func(params core_v1.WatchCoreV1ServiceListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation core_v1.WatchCoreV1ServiceListForAllNamespaces has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandler == nil {
		api.DiscoveryV1beta1WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandler = discovery_v1beta1.WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesHandlerFunc(func(params discovery_v1beta1.WatchDiscoveryV1beta1EndpointSliceListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.WatchDiscoveryV1beta1EndpointSliceListForAllNamespaces has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1WatchDiscoveryV1beta1NamespacedEndpointSliceHandler == nil {
		api.DiscoveryV1beta1WatchDiscoveryV1beta1NamespacedEndpointSliceHandler = discovery_v1beta1.WatchDiscoveryV1beta1NamespacedEndpointSliceHandlerFunc(func(params discovery_v1beta1.WatchDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.WatchDiscoveryV1beta1NamespacedEndpointSlice has not yet been implemented")
		})
	}
	if api.DiscoveryV1beta1WatchDiscoveryV1beta1NamespacedEndpointSliceListHandler == nil {
		api.DiscoveryV1beta1WatchDiscoveryV1beta1NamespacedEndpointSliceListHandler = discovery_v1beta1.WatchDiscoveryV1beta1NamespacedEndpointSliceListHandlerFunc(func(params discovery_v1beta1.WatchDiscoveryV1beta1NamespacedEndpointSliceListParams) middleware.Responder {
			return middleware.NotImplemented("operation discovery_v1beta1.WatchDiscoveryV1beta1NamespacedEndpointSliceList has not yet been implemented")
		})
	}
	if api.EventsV1WatchEventsV1EventListForAllNamespacesHandler == nil {
		api.EventsV1WatchEventsV1EventListForAllNamespacesHandler = events_v1.WatchEventsV1EventListForAllNamespacesHandlerFunc(func(params events_v1.WatchEventsV1EventListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.WatchEventsV1EventListForAllNamespaces has not yet been implemented")
		})
	}
	if api.EventsV1WatchEventsV1NamespacedEventHandler == nil {
		api.EventsV1WatchEventsV1NamespacedEventHandler = events_v1.WatchEventsV1NamespacedEventHandlerFunc(func(params events_v1.WatchEventsV1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.WatchEventsV1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1WatchEventsV1NamespacedEventListHandler == nil {
		api.EventsV1WatchEventsV1NamespacedEventListHandler = events_v1.WatchEventsV1NamespacedEventListHandlerFunc(func(params events_v1.WatchEventsV1NamespacedEventListParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1.WatchEventsV1NamespacedEventList has not yet been implemented")
		})
	}
	if api.EventsV1beta1WatchEventsV1beta1EventListForAllNamespacesHandler == nil {
		api.EventsV1beta1WatchEventsV1beta1EventListForAllNamespacesHandler = events_v1beta1.WatchEventsV1beta1EventListForAllNamespacesHandlerFunc(func(params events_v1beta1.WatchEventsV1beta1EventListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.WatchEventsV1beta1EventListForAllNamespaces has not yet been implemented")
		})
	}
	if api.EventsV1beta1WatchEventsV1beta1NamespacedEventHandler == nil {
		api.EventsV1beta1WatchEventsV1beta1NamespacedEventHandler = events_v1beta1.WatchEventsV1beta1NamespacedEventHandlerFunc(func(params events_v1beta1.WatchEventsV1beta1NamespacedEventParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.WatchEventsV1beta1NamespacedEvent has not yet been implemented")
		})
	}
	if api.EventsV1beta1WatchEventsV1beta1NamespacedEventListHandler == nil {
		api.EventsV1beta1WatchEventsV1beta1NamespacedEventListHandler = events_v1beta1.WatchEventsV1beta1NamespacedEventListHandlerFunc(func(params events_v1beta1.WatchEventsV1beta1NamespacedEventListParams) middleware.Responder {
			return middleware.NotImplemented("operation events_v1beta1.WatchEventsV1beta1NamespacedEventList has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1WatchExtensionsV1beta1IngressListForAllNamespacesHandler == nil {
		api.ExtensionsV1beta1WatchExtensionsV1beta1IngressListForAllNamespacesHandler = extensions_v1beta1.WatchExtensionsV1beta1IngressListForAllNamespacesHandlerFunc(func(params extensions_v1beta1.WatchExtensionsV1beta1IngressListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.WatchExtensionsV1beta1IngressListForAllNamespaces has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1WatchExtensionsV1beta1NamespacedIngressHandler == nil {
		api.ExtensionsV1beta1WatchExtensionsV1beta1NamespacedIngressHandler = extensions_v1beta1.WatchExtensionsV1beta1NamespacedIngressHandlerFunc(func(params extensions_v1beta1.WatchExtensionsV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.WatchExtensionsV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.ExtensionsV1beta1WatchExtensionsV1beta1NamespacedIngressListHandler == nil {
		api.ExtensionsV1beta1WatchExtensionsV1beta1NamespacedIngressListHandler = extensions_v1beta1.WatchExtensionsV1beta1NamespacedIngressListHandlerFunc(func(params extensions_v1beta1.WatchExtensionsV1beta1NamespacedIngressListParams) middleware.Responder {
			return middleware.NotImplemented("operation extensions_v1beta1.WatchExtensionsV1beta1NamespacedIngressList has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1FlowSchemaHandler == nil {
		api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1FlowSchemaHandler = flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1FlowSchemaHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1FlowSchemaParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1FlowSchema has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandler == nil {
		api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandler = flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1FlowSchemaListHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1FlowSchemaListParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1FlowSchemaList has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler == nil {
		api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandler = flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1PriorityLevelConfiguration has not yet been implemented")
		})
	}
	if api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationListHandler == nil {
		api.FlowcontrolApiserverV1alpha1WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationListHandler = flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationListHandlerFunc(func(params flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationListParams) middleware.Responder {
			return middleware.NotImplemented("operation flowcontrol_apiserver_v1alpha1.WatchFlowcontrolApiserverV1alpha1PriorityLevelConfigurationList has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1IngressClassHandler == nil {
		api.NetworkingV1WatchNetworkingV1IngressClassHandler = networking_v1.WatchNetworkingV1IngressClassHandlerFunc(func(params networking_v1.WatchNetworkingV1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1IngressClassListHandler == nil {
		api.NetworkingV1WatchNetworkingV1IngressClassListHandler = networking_v1.WatchNetworkingV1IngressClassListHandlerFunc(func(params networking_v1.WatchNetworkingV1IngressClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1IngressClassList has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1IngressListForAllNamespacesHandler == nil {
		api.NetworkingV1WatchNetworkingV1IngressListForAllNamespacesHandler = networking_v1.WatchNetworkingV1IngressListForAllNamespacesHandlerFunc(func(params networking_v1.WatchNetworkingV1IngressListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1IngressListForAllNamespaces has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1NamespacedIngressHandler == nil {
		api.NetworkingV1WatchNetworkingV1NamespacedIngressHandler = networking_v1.WatchNetworkingV1NamespacedIngressHandlerFunc(func(params networking_v1.WatchNetworkingV1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1NamespacedIngressListHandler == nil {
		api.NetworkingV1WatchNetworkingV1NamespacedIngressListHandler = networking_v1.WatchNetworkingV1NamespacedIngressListHandlerFunc(func(params networking_v1.WatchNetworkingV1NamespacedIngressListParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1NamespacedIngressList has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1NamespacedNetworkPolicyHandler == nil {
		api.NetworkingV1WatchNetworkingV1NamespacedNetworkPolicyHandler = networking_v1.WatchNetworkingV1NamespacedNetworkPolicyHandlerFunc(func(params networking_v1.WatchNetworkingV1NamespacedNetworkPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1NamespacedNetworkPolicy has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1NamespacedNetworkPolicyListHandler == nil {
		api.NetworkingV1WatchNetworkingV1NamespacedNetworkPolicyListHandler = networking_v1.WatchNetworkingV1NamespacedNetworkPolicyListHandlerFunc(func(params networking_v1.WatchNetworkingV1NamespacedNetworkPolicyListParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1NamespacedNetworkPolicyList has not yet been implemented")
		})
	}
	if api.NetworkingV1WatchNetworkingV1NetworkPolicyListForAllNamespacesHandler == nil {
		api.NetworkingV1WatchNetworkingV1NetworkPolicyListForAllNamespacesHandler = networking_v1.WatchNetworkingV1NetworkPolicyListForAllNamespacesHandlerFunc(func(params networking_v1.WatchNetworkingV1NetworkPolicyListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1.WatchNetworkingV1NetworkPolicyListForAllNamespaces has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1WatchNetworkingV1beta1IngressClassHandler == nil {
		api.NetworkingV1beta1WatchNetworkingV1beta1IngressClassHandler = networking_v1beta1.WatchNetworkingV1beta1IngressClassHandlerFunc(func(params networking_v1beta1.WatchNetworkingV1beta1IngressClassParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.WatchNetworkingV1beta1IngressClass has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1WatchNetworkingV1beta1IngressClassListHandler == nil {
		api.NetworkingV1beta1WatchNetworkingV1beta1IngressClassListHandler = networking_v1beta1.WatchNetworkingV1beta1IngressClassListHandlerFunc(func(params networking_v1beta1.WatchNetworkingV1beta1IngressClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.WatchNetworkingV1beta1IngressClassList has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1WatchNetworkingV1beta1IngressListForAllNamespacesHandler == nil {
		api.NetworkingV1beta1WatchNetworkingV1beta1IngressListForAllNamespacesHandler = networking_v1beta1.WatchNetworkingV1beta1IngressListForAllNamespacesHandlerFunc(func(params networking_v1beta1.WatchNetworkingV1beta1IngressListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.WatchNetworkingV1beta1IngressListForAllNamespaces has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1WatchNetworkingV1beta1NamespacedIngressHandler == nil {
		api.NetworkingV1beta1WatchNetworkingV1beta1NamespacedIngressHandler = networking_v1beta1.WatchNetworkingV1beta1NamespacedIngressHandlerFunc(func(params networking_v1beta1.WatchNetworkingV1beta1NamespacedIngressParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.WatchNetworkingV1beta1NamespacedIngress has not yet been implemented")
		})
	}
	if api.NetworkingV1beta1WatchNetworkingV1beta1NamespacedIngressListHandler == nil {
		api.NetworkingV1beta1WatchNetworkingV1beta1NamespacedIngressListHandler = networking_v1beta1.WatchNetworkingV1beta1NamespacedIngressListHandlerFunc(func(params networking_v1beta1.WatchNetworkingV1beta1NamespacedIngressListParams) middleware.Responder {
			return middleware.NotImplemented("operation networking_v1beta1.WatchNetworkingV1beta1NamespacedIngressList has not yet been implemented")
		})
	}
	if api.NodeV1alpha1WatchNodeV1alpha1RuntimeClassHandler == nil {
		api.NodeV1alpha1WatchNodeV1alpha1RuntimeClassHandler = node_v1alpha1.WatchNodeV1alpha1RuntimeClassHandlerFunc(func(params node_v1alpha1.WatchNodeV1alpha1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.WatchNodeV1alpha1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1alpha1WatchNodeV1alpha1RuntimeClassListHandler == nil {
		api.NodeV1alpha1WatchNodeV1alpha1RuntimeClassListHandler = node_v1alpha1.WatchNodeV1alpha1RuntimeClassListHandlerFunc(func(params node_v1alpha1.WatchNodeV1alpha1RuntimeClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1alpha1.WatchNodeV1alpha1RuntimeClassList has not yet been implemented")
		})
	}
	if api.NodeV1beta1WatchNodeV1beta1RuntimeClassHandler == nil {
		api.NodeV1beta1WatchNodeV1beta1RuntimeClassHandler = node_v1beta1.WatchNodeV1beta1RuntimeClassHandlerFunc(func(params node_v1beta1.WatchNodeV1beta1RuntimeClassParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.WatchNodeV1beta1RuntimeClass has not yet been implemented")
		})
	}
	if api.NodeV1beta1WatchNodeV1beta1RuntimeClassListHandler == nil {
		api.NodeV1beta1WatchNodeV1beta1RuntimeClassListHandler = node_v1beta1.WatchNodeV1beta1RuntimeClassListHandlerFunc(func(params node_v1beta1.WatchNodeV1beta1RuntimeClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation node_v1beta1.WatchNodeV1beta1RuntimeClassList has not yet been implemented")
		})
	}
	if api.PolicyV1beta1WatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler == nil {
		api.PolicyV1beta1WatchPolicyV1beta1NamespacedPodDisruptionBudgetHandler = policy_v1beta1.WatchPolicyV1beta1NamespacedPodDisruptionBudgetHandlerFunc(func(params policy_v1beta1.WatchPolicyV1beta1NamespacedPodDisruptionBudgetParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.WatchPolicyV1beta1NamespacedPodDisruptionBudget has not yet been implemented")
		})
	}
	if api.PolicyV1beta1WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandler == nil {
		api.PolicyV1beta1WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandler = policy_v1beta1.WatchPolicyV1beta1NamespacedPodDisruptionBudgetListHandlerFunc(func(params policy_v1beta1.WatchPolicyV1beta1NamespacedPodDisruptionBudgetListParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.WatchPolicyV1beta1NamespacedPodDisruptionBudgetList has not yet been implemented")
		})
	}
	if api.PolicyV1beta1WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesHandler == nil {
		api.PolicyV1beta1WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesHandler = policy_v1beta1.WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesHandlerFunc(func(params policy_v1beta1.WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.WatchPolicyV1beta1PodDisruptionBudgetListForAllNamespaces has not yet been implemented")
		})
	}
	if api.PolicyV1beta1WatchPolicyV1beta1PodSecurityPolicyHandler == nil {
		api.PolicyV1beta1WatchPolicyV1beta1PodSecurityPolicyHandler = policy_v1beta1.WatchPolicyV1beta1PodSecurityPolicyHandlerFunc(func(params policy_v1beta1.WatchPolicyV1beta1PodSecurityPolicyParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.WatchPolicyV1beta1PodSecurityPolicy has not yet been implemented")
		})
	}
	if api.PolicyV1beta1WatchPolicyV1beta1PodSecurityPolicyListHandler == nil {
		api.PolicyV1beta1WatchPolicyV1beta1PodSecurityPolicyListHandler = policy_v1beta1.WatchPolicyV1beta1PodSecurityPolicyListHandlerFunc(func(params policy_v1beta1.WatchPolicyV1beta1PodSecurityPolicyListParams) middleware.Responder {
			return middleware.NotImplemented("operation policy_v1beta1.WatchPolicyV1beta1PodSecurityPolicyList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleHandler = rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleBindingHandler = rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleBindingListHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleBindingListHandler = rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleBindingListHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleBindingListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleBindingList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleListHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1ClusterRoleListHandler = rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleListHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1ClusterRoleList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleHandler = rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleBindingHandler = rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleBindingListHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleBindingListHandler = rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleBindingListHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleBindingListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleBindingList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleListHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1NamespacedRoleListHandler = rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleListHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1NamespacedRoleList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1RoleBindingListForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1RoleBindingListForAllNamespacesHandler = rbac_authorization_v1.WatchRbacAuthorizationV1RoleBindingListForAllNamespacesHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1RoleBindingListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1RoleBindingListForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1WatchRbacAuthorizationV1RoleListForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1WatchRbacAuthorizationV1RoleListForAllNamespacesHandler = rbac_authorization_v1.WatchRbacAuthorizationV1RoleListForAllNamespacesHandlerFunc(func(params rbac_authorization_v1.WatchRbacAuthorizationV1RoleListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1.WatchRbacAuthorizationV1RoleListForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleBindingHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleBindingListHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleBindingListHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleBindingListHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleBindingListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleBindingList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleListHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1ClusterRoleListHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleListHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1ClusterRoleList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleBindingListHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleBindingListHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleBindingListHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleBindingListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleBindingList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleListHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1NamespacedRoleListHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleListHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1NamespacedRoleList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1RoleBindingListForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1RoleListForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1alpha1WatchRbacAuthorizationV1alpha1RoleListForAllNamespacesHandler = rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1RoleListForAllNamespacesHandlerFunc(func(params rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1RoleListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1alpha1.WatchRbacAuthorizationV1alpha1RoleListForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleBindingHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleBindingListHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleBindingListHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleBindingListHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleBindingListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleBindingList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleListHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1ClusterRoleListHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleListHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1ClusterRoleList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRole has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleBindingHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleBindingHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleBindingHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleBindingParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleBinding has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleBindingListHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleBindingListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleBindingList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleListHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1NamespacedRoleListHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleListHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleListParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1NamespacedRoleList has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1RoleBindingListForAllNamespaces has not yet been implemented")
		})
	}
	if api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandler == nil {
		api.RbacAuthorizationV1beta1WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandler = rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1RoleListForAllNamespacesHandlerFunc(func(params rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1RoleListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation rbac_authorization_v1beta1.WatchRbacAuthorizationV1beta1RoleListForAllNamespaces has not yet been implemented")
		})
	}
	if api.SchedulingV1WatchSchedulingV1PriorityClassHandler == nil {
		api.SchedulingV1WatchSchedulingV1PriorityClassHandler = scheduling_v1.WatchSchedulingV1PriorityClassHandlerFunc(func(params scheduling_v1.WatchSchedulingV1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.WatchSchedulingV1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1WatchSchedulingV1PriorityClassListHandler == nil {
		api.SchedulingV1WatchSchedulingV1PriorityClassListHandler = scheduling_v1.WatchSchedulingV1PriorityClassListHandlerFunc(func(params scheduling_v1.WatchSchedulingV1PriorityClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1.WatchSchedulingV1PriorityClassList has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1WatchSchedulingV1alpha1PriorityClassHandler == nil {
		api.SchedulingV1alpha1WatchSchedulingV1alpha1PriorityClassHandler = scheduling_v1alpha1.WatchSchedulingV1alpha1PriorityClassHandlerFunc(func(params scheduling_v1alpha1.WatchSchedulingV1alpha1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.WatchSchedulingV1alpha1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1alpha1WatchSchedulingV1alpha1PriorityClassListHandler == nil {
		api.SchedulingV1alpha1WatchSchedulingV1alpha1PriorityClassListHandler = scheduling_v1alpha1.WatchSchedulingV1alpha1PriorityClassListHandlerFunc(func(params scheduling_v1alpha1.WatchSchedulingV1alpha1PriorityClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1alpha1.WatchSchedulingV1alpha1PriorityClassList has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1WatchSchedulingV1beta1PriorityClassHandler == nil {
		api.SchedulingV1beta1WatchSchedulingV1beta1PriorityClassHandler = scheduling_v1beta1.WatchSchedulingV1beta1PriorityClassHandlerFunc(func(params scheduling_v1beta1.WatchSchedulingV1beta1PriorityClassParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.WatchSchedulingV1beta1PriorityClass has not yet been implemented")
		})
	}
	if api.SchedulingV1beta1WatchSchedulingV1beta1PriorityClassListHandler == nil {
		api.SchedulingV1beta1WatchSchedulingV1beta1PriorityClassListHandler = scheduling_v1beta1.WatchSchedulingV1beta1PriorityClassListHandlerFunc(func(params scheduling_v1beta1.WatchSchedulingV1beta1PriorityClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation scheduling_v1beta1.WatchSchedulingV1beta1PriorityClassList has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1WatchSettingsV1alpha1NamespacedPodPresetHandler == nil {
		api.SettingsV1alpha1WatchSettingsV1alpha1NamespacedPodPresetHandler = settings_v1alpha1.WatchSettingsV1alpha1NamespacedPodPresetHandlerFunc(func(params settings_v1alpha1.WatchSettingsV1alpha1NamespacedPodPresetParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.WatchSettingsV1alpha1NamespacedPodPreset has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1WatchSettingsV1alpha1NamespacedPodPresetListHandler == nil {
		api.SettingsV1alpha1WatchSettingsV1alpha1NamespacedPodPresetListHandler = settings_v1alpha1.WatchSettingsV1alpha1NamespacedPodPresetListHandlerFunc(func(params settings_v1alpha1.WatchSettingsV1alpha1NamespacedPodPresetListParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.WatchSettingsV1alpha1NamespacedPodPresetList has not yet been implemented")
		})
	}
	if api.SettingsV1alpha1WatchSettingsV1alpha1PodPresetListForAllNamespacesHandler == nil {
		api.SettingsV1alpha1WatchSettingsV1alpha1PodPresetListForAllNamespacesHandler = settings_v1alpha1.WatchSettingsV1alpha1PodPresetListForAllNamespacesHandlerFunc(func(params settings_v1alpha1.WatchSettingsV1alpha1PodPresetListForAllNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation settings_v1alpha1.WatchSettingsV1alpha1PodPresetListForAllNamespaces has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1CSIDriverHandler == nil {
		api.StorageV1WatchStorageV1CSIDriverHandler = storage_v1.WatchStorageV1CSIDriverHandlerFunc(func(params storage_v1.WatchStorageV1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1CSIDriverListHandler == nil {
		api.StorageV1WatchStorageV1CSIDriverListHandler = storage_v1.WatchStorageV1CSIDriverListHandlerFunc(func(params storage_v1.WatchStorageV1CSIDriverListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1CSIDriverList has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1CSINodeHandler == nil {
		api.StorageV1WatchStorageV1CSINodeHandler = storage_v1.WatchStorageV1CSINodeHandlerFunc(func(params storage_v1.WatchStorageV1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1CSINodeListHandler == nil {
		api.StorageV1WatchStorageV1CSINodeListHandler = storage_v1.WatchStorageV1CSINodeListHandlerFunc(func(params storage_v1.WatchStorageV1CSINodeListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1CSINodeList has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1StorageClassHandler == nil {
		api.StorageV1WatchStorageV1StorageClassHandler = storage_v1.WatchStorageV1StorageClassHandlerFunc(func(params storage_v1.WatchStorageV1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1StorageClassListHandler == nil {
		api.StorageV1WatchStorageV1StorageClassListHandler = storage_v1.WatchStorageV1StorageClassListHandlerFunc(func(params storage_v1.WatchStorageV1StorageClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1StorageClassList has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1VolumeAttachmentHandler == nil {
		api.StorageV1WatchStorageV1VolumeAttachmentHandler = storage_v1.WatchStorageV1VolumeAttachmentHandlerFunc(func(params storage_v1.WatchStorageV1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1WatchStorageV1VolumeAttachmentListHandler == nil {
		api.StorageV1WatchStorageV1VolumeAttachmentListHandler = storage_v1.WatchStorageV1VolumeAttachmentListHandlerFunc(func(params storage_v1.WatchStorageV1VolumeAttachmentListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1.WatchStorageV1VolumeAttachmentList has not yet been implemented")
		})
	}
	if api.StorageV1alpha1WatchStorageV1alpha1VolumeAttachmentHandler == nil {
		api.StorageV1alpha1WatchStorageV1alpha1VolumeAttachmentHandler = storage_v1alpha1.WatchStorageV1alpha1VolumeAttachmentHandlerFunc(func(params storage_v1alpha1.WatchStorageV1alpha1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.WatchStorageV1alpha1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1alpha1WatchStorageV1alpha1VolumeAttachmentListHandler == nil {
		api.StorageV1alpha1WatchStorageV1alpha1VolumeAttachmentListHandler = storage_v1alpha1.WatchStorageV1alpha1VolumeAttachmentListHandlerFunc(func(params storage_v1alpha1.WatchStorageV1alpha1VolumeAttachmentListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1alpha1.WatchStorageV1alpha1VolumeAttachmentList has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1CSIDriverHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1CSIDriverHandler = storage_v1beta1.WatchStorageV1beta1CSIDriverHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1CSIDriverParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1CSIDriver has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1CSIDriverListHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1CSIDriverListHandler = storage_v1beta1.WatchStorageV1beta1CSIDriverListHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1CSIDriverListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1CSIDriverList has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1CSINodeHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1CSINodeHandler = storage_v1beta1.WatchStorageV1beta1CSINodeHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1CSINodeParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1CSINode has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1CSINodeListHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1CSINodeListHandler = storage_v1beta1.WatchStorageV1beta1CSINodeListHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1CSINodeListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1CSINodeList has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1StorageClassHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1StorageClassHandler = storage_v1beta1.WatchStorageV1beta1StorageClassHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1StorageClassParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1StorageClass has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1StorageClassListHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1StorageClassListHandler = storage_v1beta1.WatchStorageV1beta1StorageClassListHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1StorageClassListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1StorageClassList has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1VolumeAttachmentHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1VolumeAttachmentHandler = storage_v1beta1.WatchStorageV1beta1VolumeAttachmentHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1VolumeAttachmentParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1VolumeAttachment has not yet been implemented")
		})
	}
	if api.StorageV1beta1WatchStorageV1beta1VolumeAttachmentListHandler == nil {
		api.StorageV1beta1WatchStorageV1beta1VolumeAttachmentListHandler = storage_v1beta1.WatchStorageV1beta1VolumeAttachmentListHandlerFunc(func(params storage_v1beta1.WatchStorageV1beta1VolumeAttachmentListParams) middleware.Responder {
			return middleware.NotImplemented("operation storage_v1beta1.WatchStorageV1beta1VolumeAttachmentList has not yet been implemented")
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
	// add /healthz manually since it's not part of the API
	return healthz.NewGetHealthzHandler(handler)
}
