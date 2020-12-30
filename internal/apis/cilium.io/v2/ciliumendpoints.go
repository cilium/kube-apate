// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"sync"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/cilium"
	management "github.com/cilium/kube-apate/api/management/v1/server/restapi/cilium"
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	"github.com/cilium/cilium/api/v1/models"
	ciliumV2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	CE = "ciliumendpoints.cilium.io"
)

var CEManager = &CEMgr{
	staticList:   &ciliumV2.CiliumEndpointList{},
	elemTemplate: &ciliumV2.CiliumEndpoint{},
	streamCh:     make(chan k8sRuntime.Object),
}

type CEConfig struct {
	Name        string
	Namespace   string
	ContainerID string
	OwnerUID    string
	UID         string
	LocalID     int64
	Identity    *ciliumV2.EndpointIdentity
	Networking  *ciliumV2.EndpointNetworking
}

type CEMgr struct {
	logger
	sync.RWMutex
	streamCh     chan k8sRuntime.Object
	totalGenElem int64
	staticList   *ciliumV2.CiliumEndpointList
	elemTemplate *ciliumV2.CiliumEndpoint
}

func (*CEMgr) NewListInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumEndpointList{}
}

func (*CEMgr) NewInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumEndpoint{}
}

func (n *CEMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CEMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumEndpointList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumEndpointList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CEMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumEndpoint)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumEndpoint", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CEMgr) getItem(cfg *CEConfig) *ciliumV2.CiliumEndpoint {
	n.RLock()
	defer n.RUnlock()
	ce := n.elemTemplate.DeepCopy()
	ce.Name = cfg.Name
	ce.Namespace = cfg.Namespace
	ce.UID = k8sTypes.UID(cfg.UID)
	ce.OwnerReferences = []v1.OwnerReference{
		{
			APIVersion:         "v1",
			Kind:               "Pod",
			Name:               cfg.Name,
			UID:                k8sTypes.UID(cfg.OwnerUID),
			BlockOwnerDeletion: func() *bool { a := true; return &a }(),
		},
	}
	ce.Status.ExternalIdentifiers = &models.EndpointIdentifiers{
		ContainerID:  cfg.ContainerID,
		K8sNamespace: cfg.Namespace,
		K8sPodName:   cfg.Name,
		PodName:      fmt.Sprintf("%s/%s", cfg.Namespace, cfg.Name),
	}
	ce.Status.ID = cfg.LocalID
	ce.Status.Identity = cfg.Identity
	ce.Status.Networking = cfg.Networking
	return ce
}

func (n *CEMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	ceList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(ceList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		ceList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	ceList.Continue = utils.Cont(start, totalElems, limit)

	genCE := n.GenObjs(start, maxElemts)
	for obj := range genCE {
		ce := obj.(*ciliumV2.CiliumEndpoint)
		ceList.Items = append(ceList.Items, *ce)
	}

	return ceList, nil
}

func (n *CEMgr) GenObjs(start, maxElemts int64) <-chan k8sRuntime.Object {
	ch := make(chan k8sRuntime.Object)
	go func() {
		ceCfg := &CEConfig{}
		for i := start; i < maxElemts; i++ {
			podIPv4, podIPv6 := generators.GetPodIP(i)
			if generators.IsBlockListedPod(podIPv4) {
				continue
			}
			ceCfg.Name = generators.CEName(i)
			ceCfg.Namespace = generators.NamespaceName(i)
			ceCfg.LocalID = generators.CELocalID(i)
			ceCfg.ContainerID = generators.ContainerID(i)
			ceCfg.OwnerUID = generators.PodUUID(i)
			ceCfg.Identity = &ciliumV2.EndpointIdentity{
				ID:     generators.CELocalID(i),
				Labels: generators.CiliumEndpointLabels(i),
			}
			ceCfg.Networking = &ciliumV2.EndpointNetworking{
				Addressing: ciliumV2.AddressPairList{
					{
						IPV4: podIPv4,
						IPV6: podIPv6,
					},
				},
				NodeIP: generators.GetHostOfPodIPv4(podIPv4),
			}
			ch <- n.getItem(ceCfg)
		}
		close(ch)
	}()
	return ch
}

func (n *CEMgr) Stream() chan k8sRuntime.Object {
	n.RWMutex.RLock()
	streamer := n.streamCh
	n.RWMutex.RUnlock()
	return streamer
}

type listCiliumEndpoint struct {
	*CEMgr
}

func NewReadCiliumEndpoint() cilium.ListApisCiliumIoV2CiliumEndpointHandler {
	return &listCiliumEndpoint{
		CEMgr: CEManager,
	}
}

func (lce *listCiliumEndpoint) Handle(params cilium.ListApisCiliumIoV2CiliumEndpointParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/cilium.io/v2/ciliumendpoints",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponderWithEvents(
		lce,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
		ciliumEncoder,
	)
}

type deleteCiliumEndpoint struct {
	*CEMgr
}

func NewDeleteCiliumEndpoint() cilium.DeleteApisCiliumIoV2CiliumEndpointHandler {
	return &deleteCiliumEndpoint{
		CEMgr: CEManager,
	}
}

func (*deleteCiliumEndpoint) Handle(params cilium.DeleteApisCiliumIoV2CiliumEndpointParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "DELETE",
		"URL":    "/apis/cilium.io/v2/namespaces/{namespace}/ciliumendpoints/{name}",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return cilium.NewDeleteApisCiliumIoV2CiliumEndpointOK()
}

type manageCiliumEndpoints struct {
	*CEMgr
}

func NewCiliumEndpointsMgr() management.PostManagementCiliumIoV2CiliumEndpointsHandler {
	return &manageCiliumEndpoints{
		CEMgr: CEManager,
	}
}

func (lce *manageCiliumEndpoints) Handle(params management.PostManagementCiliumIoV2CiliumEndpointsParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "POST",
		"URL":    "/management/cilium.io/v2/ciliumendpoints",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	totalCEPs := int64(params.Options.Add) - int64(params.Options.Del)

	encoders.GenerateK8sEvents(lce, totalCEPs)

	return management.NewPostManagementCiliumIoV2CiliumEndpointsAccepted()
}
