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
	CCNP = "ciliumclusterwidenetworkpolicies.cilium.io"
)

var CCNPManager = &CCNPMgr{
	staticList:   &ciliumV2.CiliumClusterwideNetworkPolicyList{},
	elemTemplate: &ciliumV2.CiliumClusterwideNetworkPolicy{},
}

type CCNPConfig struct {
	Name string
	UID  string
}

type CCNPMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *ciliumV2.CiliumClusterwideNetworkPolicyList
	elemTemplate *ciliumV2.CiliumClusterwideNetworkPolicy
}

func (*CCNPMgr) NewListInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumClusterwideNetworkPolicyList{}
}

func (*CCNPMgr) NewInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumClusterwideNetworkPolicy{}
}

func (n *CCNPMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CCNPMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumClusterwideNetworkPolicyList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumClusterwideNetworkPolicyList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CCNPMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumClusterwideNetworkPolicy)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumClusterwideNetworkPolicy", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CCNPMgr) getItem(cfg *CCNPConfig) *ciliumV2.CiliumClusterwideNetworkPolicy {
	n.RLock()
	defer n.RUnlock()
	ccnp := n.elemTemplate.DeepCopy()
	ccnp.Name = cfg.Name
	ccnp.UID = k8sTypes.UID(cfg.UID)
	return ccnp
}

func (n *CCNPMgr) List(start, limit int64) (k8sRuntime.Object, error) {
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

	nCfg := &CCNPConfig{}
	for i := start; i < maxElemts; i++ {
		nCfg.Name = generators.CCNPName(i)
		nCfg.UID = generators.CCNPUUID(i)
		ns := n.getItem(nCfg)
		ccnpList.Items = append(ccnpList.Items, *ns)
	}

	return ccnpList, nil
}

type listCiliumClusterwideNetworkPolicy struct {
	*CCNPMgr
}

func NewReadCiliumClusterwideNetworkPolicy() cilium.ListApisCiliumIoV2CiliumClusterwideNetworkPolicyHandler {
	return &listCiliumClusterwideNetworkPolicy{
		CCNPMgr: CCNPManager,
	}
}

func (lccnp *listCiliumClusterwideNetworkPolicy) Handle(params cilium.ListApisCiliumIoV2CiliumClusterwideNetworkPolicyParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/cilium.io/v2/ciliumclusterwidenetworkpolicies",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lccnp,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
