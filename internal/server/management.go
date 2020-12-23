// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"encoding/json"
	"os"
	"path/filepath"

	managementServer "github.com/cilium/kube-apate/api/management/v1/server"
	managementRestAPI "github.com/cilium/kube-apate/api/management/v1/server/restapi"
	internalK8sCoreV1 "github.com/cilium/kube-apate/internal/api/core"
	internalCiliumV2 "github.com/cilium/kube-apate/internal/apis/cilium.io/v2"
	internalK8sDiscoveryV1beta1 "github.com/cilium/kube-apate/internal/apis/discovery.k8s.io/v1beta1"
	internalK8sExtensionsV1 "github.com/cilium/kube-apate/internal/apis/extensions.k8s.io/v1"
	internalK8sNetworkingV1 "github.com/cilium/kube-apate/internal/apis/networking.k8s.io/v1"

	"github.com/go-openapi/loads"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
)

func initializeManagementAPI() *managementRestAPI.ManagementAPI {
	swaggerSpec, err := loads.Analyzed(managementServer.SwaggerJSON, "")
	if err != nil {
		log.WithError(err).Fatal("Cannot load swagger spec")
	}

	log.Info("Initializing kube-apate management API")
	restAPI := managementRestAPI.NewManagementAPI(swaggerSpec)

	restAPI.Logger = log.Infof

	// // K8s resources
	//
	// // GET /apis/networking.k8s.io/v1/networkpolicies
	// restAPI.NetworkingV1ListNetworkingV1NetworkPolicyForAllNamespacesHandler =
	// 	internalK8sNetworkingV1.NewListNetworkingV1NetworkPolicyForAllNamespacesHandler()
	//
	// // GET /api/v1/namespaces/{namespace}
	// restAPI.CoreV1ReadCoreV1NamespaceHandler =
	// 	internalK8sCoreV1.NewReadCoreV1Namespace()
	//
	// // GET /apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{endpointslice}
	// restAPI.DiscoveryV1beta1ReadDiscoveryV1beta1NamespacedEndpointSliceHandler =
	// 	internalK8sDiscoveryV1beta1.NewReadDiscoveryV1beta1NamespacedEndpointSliceHandler()
	//
	// // GET /api/v1/nodes/{nodes}
	// restAPI.CoreV1ReadCoreV1NodeHandler =
	// 	internalK8sCoreV1.NewReadCoreV1Node()
	//
	// // GET /api/v1/nodes
	// restAPI.CoreV1ListCoreV1NodeHandler =
	// 	internalK8sCoreV1.NewListCoreV1Node()
	//
	// // PATCH /api/v1/nodes/{nodes}
	// restAPI.CoreV1PatchCoreV1NodeHandler =
	// 	internalK8sCoreV1.NewPatchCoreV1Node()
	//
	// // PATCH /api/v1/nodes/{nodes}/status
	// restAPI.CoreV1PatchCoreV1NodeStatusHandler =
	// 	internalK8sCoreV1.NewPatchCoreV1NodeStatus()
	//
	// // GET /apis/apiextensions.k8s.io/v1/customresourcedefinitions
	// restAPI.ApiextensionsV1ListApiextensionsV1CustomResourceDefinitionHandler =
	// 	internalK8sExtensionsV1.NewListApiextensionsV1CustomResourceDefinition()
	//
	// // GET /apis/discovery.k8s.io/v1beta1/endpointslices
	// restAPI.DiscoveryV1beta1ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler =
	// 	internalK8sDiscoveryV1beta1.NewListDiscoveryV1beta1NamespacedEndpointSliceHandler()
	//
	// // GET /api/v1/namespaces
	// restAPI.CoreV1ListCoreV1NamespaceHandler =
	// 	internalK8sCoreV1.NewListCoreV1Namespace()
	//
	// // GET /api/v1/pods
	// restAPI.CoreV1ListCoreV1PodForAllNamespacesHandler =
	// 	internalK8sCoreV1.NewListCoreV1Pod()
	//
	// // GET /api/v1/services
	// restAPI.CoreV1ListCoreV1ServiceForAllNamespacesHandler =
	// 	internalK8sCoreV1.NewListCoreV1Service()

	// Cilium resources

	// // GET /apis/cilium.io/v2/ciliumclusterwidenetworkpolicies
	// restAPI.CiliumListApisCiliumIoV2CiliumClusterwideNetworkPolicyHandler =
	// 	internalCiliumV2.NewReadCiliumClusterwideNetworkPolicy()
	//
	// // GET /apis/cilium.io/v2/ciliumnetworkpolicies
	// restAPI.CiliumListApisCiliumIoV2CiliumNetworkPolicyHandler =
	// 	internalCiliumV2.NewReadCiliumNetworkPolicy()
	//
	// POST /management/cilium.io/v2/ciliumnodes
	restAPI.CiliumPostManagementCiliumIoV2CiliumNodesHandler =
		internalCiliumV2.NewCiliumNodesMgr()

	// POST /management/cilium.io/v2/ciliumendpoints
	restAPI.CiliumPostManagementCiliumIoV2CiliumEndpointsHandler =
		internalCiliumV2.NewCiliumEndpointsMgr()

	// POST /management/cilium.io/v2/ciliumidentities
	restAPI.CiliumPostManagementCiliumIoV2CiliumIdentitiesHandler =
		internalCiliumV2.NewCiliumIdentitiesMgr()
	//
	// // DELETE /apis/cilium.io/v2/namespaces/{namespace}/ciliumendpoints/{name}
	// restAPI.CiliumDeleteApisCiliumIoV2CiliumEndpointHandler =
	// 	internalCiliumV2.NewDeleteCiliumEndpoint()
	//
	// // POST /apis/cilium.io/v2/ciliumnodes
	// restAPI.CiliumPostApisCiliumIoV2CiliumNodesHandler =
	// 	internalCiliumV2.NewPostCiliumNode()
	//
	// // GET /apis/cilium.io/v2/ciliumnodes/{name}
	// restAPI.CiliumReadApisCiliumIoV2CiliumNodeHandler =
	// 	internalCiliumV2.NewReadCiliumNode()
	//
	// // PUT /apis/cilium.io/v2/ciliumnodes/{name}
	// restAPI.CiliumPutApisCiliumIoV2CiliumNodesHandler =
	// 	internalCiliumV2.NewPutCiliumNode()
	//
	// // GET /version
	// restAPI.VersionGetCodeVersionHandler =
	// 	version.NewVersionGetCodeVersionHandler()

	return restAPI
}

type Manager interface {
	AddElements(elem int64) int64
	SetList(obj k8sRuntime.Object) error
	SetObj(obj k8sRuntime.Object) error
	NewListInterface() k8sRuntime.Object
	NewInterface() k8sRuntime.Object
}

var managers = map[string]Manager{}

func init() {
	managers[internalK8sCoreV1.Namespace] = internalK8sCoreV1.NamespaceManager
	managers[internalK8sCoreV1.Node] = internalK8sCoreV1.NodeManager
	managers[internalK8sCoreV1.Pod] = internalK8sCoreV1.PodManager
	managers[internalK8sCoreV1.Service] = internalK8sCoreV1.ServiceManager
	managers[internalK8sDiscoveryV1beta1.EndpointSlice] = internalK8sDiscoveryV1beta1.ESManager
	managers[internalK8sExtensionsV1.CustomResourceDefinitions] = internalK8sExtensionsV1.CRDManager
	managers[internalK8sNetworkingV1.NetworkPolicies] = internalK8sNetworkingV1.NetPolManager

	// Cilium
	managers[internalCiliumV2.CCNP] = internalCiliumV2.CCNPManager
	managers[internalCiliumV2.CE] = internalCiliumV2.CEManager
	managers[internalCiliumV2.CI] = internalCiliumV2.CIManager
	managers[internalCiliumV2.CLRP] = internalCiliumV2.CEManager
	managers[internalCiliumV2.CNP] = internalCiliumV2.CNPManager
	managers[internalCiliumV2.CN] = internalCiliumV2.CiliumNodeManager

	for name, mgr := range managers {
		f, err := os.Open(filepath.Join("structs", "static", name+".json"))
		switch {
		case os.IsNotExist(err):
			break
		case err != nil:
			panic(err)
		default:
			i := mgr.NewInterface()
			err = json.NewDecoder(f).Decode(i)
			if err != nil {
				panic(err)
			}
			err = mgr.SetObj(i)
			if err != nil {
				panic(err)
			}
			f.Close()
		}
		//
		f, err = os.Open(filepath.Join("structs", "static", name+"-list.json"))
		switch {
		case os.IsNotExist(err):
			break
		case err != nil:
			panic(err)
		default:
			i := mgr.NewListInterface()
			err = json.NewDecoder(f).Decode(i)
			if err != nil {
				panic(err)
			}
			err = mgr.SetList(i)
			if err != nil {
				panic(err)
			}
			f.Close()
		}
	}
}

func StartManagement(portNumber int) {
	fakeAPI := initializeManagementAPI()
	s := managementServer.NewServer(fakeAPI)
	s.Port = portNumber
	s.SetAPI(fakeAPI)
	err := s.Serve()
	if err != nil {
		panic(err)
	}
}
