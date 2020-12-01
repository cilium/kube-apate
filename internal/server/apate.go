// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"github.com/cilium/kube-apate/api/k8s/v1/server"
	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi"
	internalK8sCoreV1 "github.com/cilium/kube-apate/internal/api/core"
	internalCiliumV2 "github.com/cilium/kube-apate/internal/apis/cilium.io/v2"
	internalK8sDiscoveryV1beta1 "github.com/cilium/kube-apate/internal/apis/discovery.k8s.io/v1beta1"
	internalK8sExtensionsV1 "github.com/cilium/kube-apate/internal/apis/extensions.k8s.io/v1"
	internalK8sNetworkingV1 "github.com/cilium/kube-apate/internal/apis/networking.k8s.io/v1"
	logging "github.com/cilium/kube-apate/internal/log"
	"github.com/cilium/kube-apate/internal/log/logfields"
	"github.com/cilium/kube-apate/internal/version"

	"github.com/go-openapi/loads"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "server")

func initializeAPI() *restapi.KubernetesAPI {
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.WithError(err).Fatal("Cannot load swagger spec")
	}

	log.Info("Initializing kube-apate API")
	restAPI := restapi.NewKubernetesAPI(swaggerSpec)

	restAPI.Logger = log.Infof

	// K8s resources

	// GET /apis/networking.k8s.io/v1/networkpolicies
	restAPI.NetworkingV1ListNetworkingV1NetworkPolicyForAllNamespacesHandler =
		internalK8sNetworkingV1.NewListNetworkingV1NetworkPolicyForAllNamespacesHandler()

	// GET /api/v1/namespaces/{namespace}
	restAPI.CoreV1ReadCoreV1NamespaceHandler =
		internalK8sCoreV1.NewReadCoreV1Namespace()

	// GET /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{endpointslice}
	restAPI.DiscoveryV1beta1ReadDiscoveryV1beta1NamespacedEndpointSliceHandler =
		internalK8sDiscoveryV1beta1.NewReadDiscoveryV1beta1NamespacedEndpointSliceHandler()

	// GET /api/v1/nodes/{nodes}
	restAPI.CoreV1ReadCoreV1NodeHandler =
		internalK8sCoreV1.NewReadCoreV1Node()

	// GET /api/v1/nodes
	restAPI.CoreV1ListCoreV1NodeHandler =
		internalK8sCoreV1.NewListCoreV1Node()

	// PATCH /api/v1/nodes/{nodes}
	restAPI.CoreV1PatchCoreV1NodeHandler =
		internalK8sCoreV1.NewPatchCoreV1Node()

	// PATCH /api/v1/nodes/{nodes}/status
	restAPI.CoreV1PatchCoreV1NodeStatusHandler =
		internalK8sCoreV1.NewPatchCoreV1NodeStatus()

	// GET /apis/apiextensions.k8s.io/v1/customresourcedefinitions
	restAPI.ApiextensionsV1ListApiextensionsV1CustomResourceDefinitionHandler =
		internalK8sExtensionsV1.NewListApiextensionsV1CustomResourceDefinition()

	// GET /apis/discovery.k8s.io/v1beta1/endpointslices
	restAPI.DiscoveryV1beta1ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler =
		internalK8sDiscoveryV1beta1.NewListDiscoveryV1beta1NamespacedEndpointSliceHandler()

	// GET /api/v1/namespaces
	restAPI.CoreV1ListCoreV1NamespaceHandler =
		internalK8sCoreV1.NewListCoreV1Namespace()

	// GET /api/v1/pods
	restAPI.CoreV1ListCoreV1PodForAllNamespacesHandler =
		internalK8sCoreV1.NewListCoreV1Pod()

	// GET /api/v1/services
	restAPI.CoreV1ListCoreV1ServiceForAllNamespacesHandler =
		internalK8sCoreV1.NewListCoreV1Service()

	// Cilium resources

	// GET /apis/cilium.io/v2/ciliumclusterwidenetworkpolicies
	restAPI.CiliumListApisCiliumIoV2CiliumClusterwideNetworkPolicyHandler =
		internalCiliumV2.NewReadCiliumClusterwideNetworkPolicy()

	// GET /apis/cilium.io/v2/ciliumnetworkpolicies
	restAPI.CiliumListApisCiliumIoV2CiliumNetworkPolicyHandler =
		internalCiliumV2.NewReadCiliumNetworkPolicy()

	// GET /apis/cilium.io/v2/ciliumnodes
	restAPI.CiliumListApisCiliumIoV2CiliumNodesHandler =
		internalCiliumV2.NewListCiliumNode()

	// GET /apis/cilium.io/v2/ciliumendpoints
	restAPI.CiliumListApisCiliumIoV2CiliumEndpointHandler =
		internalCiliumV2.NewReadCiliumEndpoint()

	// GET /apis/cilium.io/v2/ciliumidentities
	restAPI.CiliumListApisCiliumIoV2CiliumIdentityHandler =
		internalCiliumV2.NewReadCiliumIdentity()

	// DELETE /apis/cilium.io/v2/namespaces/{namespace}/ciliumendpoints/{name}
	restAPI.CiliumDeleteApisCiliumIoV2CiliumEndpointHandler =
		internalCiliumV2.NewDeleteCiliumEndpoint()

	// POST /apis/cilium.io/v2/ciliumnodes
	restAPI.CiliumPostApisCiliumIoV2CiliumNodesHandler =
		internalCiliumV2.NewPostCiliumNode()

	// GET /apis/cilium.io/v2/ciliumnodes/{name}
	restAPI.CiliumReadApisCiliumIoV2CiliumNodeHandler =
		internalCiliumV2.NewReadCiliumNode()

	// PUT /apis/cilium.io/v2/ciliumnodes/{name}
	restAPI.CiliumPutApisCiliumIoV2CiliumNodesHandler =
		internalCiliumV2.NewPutCiliumNode()

	// GET /version
	restAPI.VersionGetCodeVersionHandler =
		version.NewVersionGetCodeVersionHandler()

	return restAPI
}

func StartApate(portNumber int) {
	fakeAPI := initializeAPI()
	s := server.NewServer(fakeAPI)
	s.Port = portNumber
	s.SetAPI(fakeAPI)
	err := s.Serve()
	if err != nil {
		panic(err)
	}
}
