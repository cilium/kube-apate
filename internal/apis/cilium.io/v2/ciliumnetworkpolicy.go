// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/cilium"
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	ciliumV2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	CNP = "ciliumnetworkpolicies.cilium.io"
)

var CNPManager = &CNPMgr{
	staticList:   &ciliumV2.CiliumNetworkPolicyList{},
	elemTemplate: &ciliumV2.CiliumNetworkPolicy{},
}

type CNPConfig struct {
	Name string
	UUID string
}

type CNPMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *ciliumV2.CiliumNetworkPolicyList
	elemTemplate *ciliumV2.CiliumNetworkPolicy
}

func (*CNPMgr) NewListInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumNetworkPolicyList{}
}

func (*CNPMgr) NewInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumNetworkPolicy{}
}

func (n *CNPMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CNPMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumNetworkPolicyList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumNetworkPolicyList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CNPMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumNetworkPolicy)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumNetworkPolicy", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CNPMgr) getItem(cfg *CNPConfig) *ciliumV2.CiliumNetworkPolicy {
	n.RLock()
	defer n.RUnlock()
	ccnp := n.elemTemplate.DeepCopy()
	ccnp.Name = cfg.Name
	ccnp.UID = k8sTypes.UID(cfg.UUID)
	return ccnp
}

func (n *CNPMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	ccnpList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(ccnpList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		ccnpList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	ccnpList.Continue = utils.Cont(start, totalElems, limit)

	nCfg := &CNPConfig{}
	for i := start; i < maxElemts; i++ {
		nCfg.Name = generators.CNPName(i)
		nCfg.UUID = generators.CNPUUID(i)
		ns := n.getItem(nCfg)
		ccnpList.Items = append(ccnpList.Items, *ns)
	}

	return ccnpList, nil
}

type listCiliumNetworkPolicy struct {
	*CNPMgr
}

func NewReadCiliumNetworkPolicy() cilium.ListApisCiliumIoV2CiliumNetworkPolicyHandler {
	return &listCiliumNetworkPolicy{
		CNPMgr: CNPManager,
	}
}

func (lcnp *listCiliumNetworkPolicy) Handle(params cilium.ListApisCiliumIoV2CiliumNetworkPolicyParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/cilium.io/v2/ciliumnetworkpolicies",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lcnp,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
