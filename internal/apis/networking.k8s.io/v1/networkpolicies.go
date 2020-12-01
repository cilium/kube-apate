// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"fmt"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/networking_v1"
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sNetworkingV1 "k8s.io/api/networking/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	NetworkPolicies = "netpol"
)

var NetPolManager = &NetPolMgr{
	staticList:   &k8sNetworkingV1.NetworkPolicyList{},
	elemTemplate: &k8sNetworkingV1.NetworkPolicy{},
}

type NetPolConfig struct {
	Name string
	UUID string
}

type NetPolMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sNetworkingV1.NetworkPolicyList
	elemTemplate *k8sNetworkingV1.NetworkPolicy
}

func (*NetPolMgr) NewListInterface() k8sRuntime.Object {
	return &k8sNetworkingV1.NetworkPolicyList{}
}

func (*NetPolMgr) NewInterface() k8sRuntime.Object {
	return &k8sNetworkingV1.NetworkPolicy{}
}

func (n *NetPolMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *NetPolMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sNetworkingV1.NetworkPolicyList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sNetworkingV1.NetworkPolicyList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *NetPolMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sNetworkingV1.NetworkPolicy)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sNetworkingV1.NetworkPolicy", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *NetPolMgr) getItem(cfg *NetPolConfig) *k8sNetworkingV1.NetworkPolicy {
	n.RLock()
	defer n.RUnlock()
	cNetPol := n.elemTemplate.DeepCopy()
	cNetPol.Name = cfg.Name
	cNetPol.UID = k8sTypes.UID(cfg.UUID)
	return cNetPol
}

func (n *NetPolMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	cNetPolList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(cNetPolList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		cNetPolList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	cNetPolList.Continue = utils.Cont(start, totalElems, limit)

	nCfg := &NetPolConfig{}
	for i := start; i < maxElemts; i++ {
		nCfg.Name = generators.NetPolName(i)
		nCfg.UUID = generators.NetPolUUID(i)
		ns := n.getItem(nCfg)
		cNetPolList.Items = append(cNetPolList.Items, *ns)
	}

	return cNetPolList, nil
}

type listAllNSNetPol struct {
	*NetPolMgr
}

func NewListNetworkingV1NetworkPolicyForAllNamespacesHandler() networking_v1.ListNetworkingV1NetworkPolicyForAllNamespacesHandler {
	return &listAllNSNetPol{
		NetPolMgr: NetPolManager,
	}
}

func (lNetPol *listAllNSNetPol) Handle(params networking_v1.ListNetworkingV1NetworkPolicyForAllNamespacesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/networking.k8s.io/v1/networkpolicies",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lNetPol,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
