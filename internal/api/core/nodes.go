// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"fmt"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/core_v1"
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sCoreV1 "k8s.io/api/core/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	Node = "node"
)

var NodeManager = &NodeMgr{}

type NodeConfig struct {
	Name           string
	IPv4           string
	IPv6           string
	PodCIDRv4      string
	PodCIDRv6      string
	CiliumHostIP   string
	CiliumHealthIP string
	UUID           string
}

type NodeMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sCoreV1.NodeList
	elemTemplate *k8sCoreV1.Node
}

func (*NodeMgr) NewListInterface() k8sRuntime.Object {
	return &k8sCoreV1.NodeList{}
}

func (*NodeMgr) NewInterface() k8sRuntime.Object {
	return &k8sCoreV1.Node{}
}

func (n *NodeMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *NodeMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.NodeList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.NodeList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *NodeMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.Node)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.Node", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *NodeMgr) getItem(cfg *NodeConfig) *k8sCoreV1.Node {
	n.RLock()
	defer n.RUnlock()
	ns := n.elemTemplate.DeepCopy()
	ns.Name = cfg.Name
	ns.UID = k8sTypes.UID(cfg.UUID)
	ns.Status.Addresses = nil
	ns.Spec.PodCIDRs = nil
	if cfg.IPv4 != "" {
		ns.Status.Addresses = append(
			ns.Status.Addresses,
			k8sCoreV1.NodeAddress{
				Type:    k8sCoreV1.NodeInternalIP,
				Address: cfg.IPv4,
			})
	}
	if ns.Name != "" {
		ns.Status.Addresses = append(
			ns.Status.Addresses,
			k8sCoreV1.NodeAddress{
				Type:    k8sCoreV1.NodeHostName,
				Address: ns.Name,
			})
	}
	if cfg.PodCIDRv4 != "" {
		ns.Spec.PodCIDR = cfg.PodCIDRv4
		ns.Spec.PodCIDRs = append(ns.Spec.PodCIDRs, cfg.PodCIDRv4)
		ns.ObjectMeta.Annotations["io.cilium.network.ipv4-pod-cidr"] = cfg.PodCIDRv4
	}
	if cfg.PodCIDRv6 != "" {
		ns.Spec.PodCIDRs = append(ns.Spec.PodCIDRs, cfg.PodCIDRv6)
		ns.ObjectMeta.Annotations["io.cilium.network.ipv6-pod-cidr"] = cfg.PodCIDRv6
	}
	if cfg.CiliumHostIP != "" {
		ns.ObjectMeta.Annotations["io.cilium.network.ipv4-cilium-host"] = cfg.CiliumHostIP
	}
	if cfg.CiliumHealthIP != "" {
		ns.ObjectMeta.Annotations["io.cilium.network.ipv4-health-ip"] = cfg.CiliumHealthIP
	}
	return ns
}

func (n *NodeMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	nodeList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(nodeList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		nodeList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	nodeList.Continue = utils.Cont(start, totalElems, limit)

	nCfg := &NodeConfig{}
	for i := start; i < maxElemts; i++ {
		nCfg.Name = generators.NodeName(i)
		nCfg.IPv4, nCfg.IPv6 = generators.GetHostIP(i)
		nCfg.PodCIDRv4, nCfg.PodCIDRv6 = generators.GetPodCIDR(i)
		nCfg.CiliumHostIP = generators.IPFromCIDR(nCfg.PodCIDRv4, 125)
		nCfg.CiliumHealthIP = generators.IPFromCIDR(nCfg.PodCIDRv4, 126)
		nCfg.UUID = generators.NodeUUID(i)
		ns := n.getItem(nCfg)
		nodeList.Items = append(nodeList.Items, *ns)
	}

	return nodeList, nil
}

type getNode struct {
	*NodeMgr
}

func NewReadCoreV1Node() core_v1.ReadCoreV1NodeHandler {
	return &getNode{
		NodeMgr: NodeManager,
	}
}

func (gn *getNode) Handle(params core_v1.ReadCoreV1NodeParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/api/v1/nodes/{node}",
		"Name":   params.Name,
	}).Debug("request received")

	obj := gn.getItem(&NodeConfig{Name: params.Name})

	return getResponder(obj)
}

type patchNode struct {
	*NodeMgr
}

func NewPatchCoreV1Node() core_v1.PatchCoreV1NodeHandler {
	return &patchNode{
		NodeMgr: NodeManager,
	}
}

func (pn *patchNode) Handle(params core_v1.PatchCoreV1NodeParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "PATCH",
		"URL":    "/api/v1/nodes/{node}",
		"Name":   params.Name,
	}).Debug("request received")

	obj := pn.getItem(&NodeConfig{Name: params.Name})

	return getResponder(obj)
}

type patchNodeStatus struct {
	*NodeMgr
}

func NewPatchCoreV1NodeStatus() core_v1.PatchCoreV1NodeStatusHandler {
	return &patchNodeStatus{
		NodeMgr: NodeManager,
	}
}

func (pn *patchNodeStatus) Handle(params core_v1.PatchCoreV1NodeStatusParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "PATCH",
		"URL":    "/api/v1/nodes/{node}/status",
		"Name":   params.Name,
	}).Debug("request received")

	obj := pn.getItem(&NodeConfig{Name: params.Name})

	return getResponder(obj)
}

type listAllNodes struct {
	*NodeMgr
}

func NewListCoreV1Node() core_v1.ListCoreV1NodeHandler {
	return &listAllNodes{
		NodeMgr: NodeManager,
	}
}

func (lan *listAllNodes) Handle(params core_v1.ListCoreV1NodeParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/api/v1/nodes",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lan,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
