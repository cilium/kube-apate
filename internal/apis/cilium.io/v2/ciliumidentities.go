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
	CI = "ciliumidentities.cilium.io"
)

var CIManager = &CIMgr{
	staticList:   &ciliumV2.CiliumIdentityList{},
	elemTemplate: &ciliumV2.CiliumIdentity{},
	streamCh:     make(chan k8sRuntime.Object),
}

type CIConfig struct {
	Name           string
	Namespace      string
	UUID           string
	Labels         map[string]string
	SecurityLabels map[string]string
}

type CIMgr struct {
	sync.RWMutex
	streamCh     chan k8sRuntime.Object
	totalGenElem int64
	staticList   *ciliumV2.CiliumIdentityList
	elemTemplate *ciliumV2.CiliumIdentity
}

func (*CIMgr) NewListInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumIdentityList{}
}

func (*CIMgr) NewInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumIdentity{}
}

func (n *CIMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CIMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumIdentityList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumIdentityList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CIMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumIdentity)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumIdentity", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CIMgr) getItem(cfg *CIConfig) *ciliumV2.CiliumIdentity {
	n.RLock()
	defer n.RUnlock()
	ci := n.elemTemplate.DeepCopy()
	ci.Name = cfg.Name
	ci.Namespace = cfg.Namespace
	ci.UID = k8sTypes.UID(cfg.UUID)
	ci.Labels = cfg.Labels
	ci.SecurityLabels = cfg.SecurityLabels
	return ci
}

func (n *CIMgr) List(start, limit int64) (k8sRuntime.Object, error) {
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

	genCI := n.genCI(start, maxElemts)
	for ce := range genCI {
		ceList.Items = append(ceList.Items, *ce)
	}

	return ceList, nil
}

func (n *CIMgr) genCI(start, maxElemts int64) <-chan *ciliumV2.CiliumIdentity {
	ch := make(chan *ciliumV2.CiliumIdentity)
	go func() {
		ciCfg := &CIConfig{}
		for i := start; i < maxElemts; i++ {
			ciCfg.Name = generators.CIName(i)
			ciCfg.UUID = generators.CIUUID(i)
			ciCfg.Namespace = generators.NamespaceName(i)
			ciCfg.Labels = generators.CiliumIdentityLabels(i)
			ciCfg.SecurityLabels = generators.CiliumSecurityIdentityLabels(i)
			ch <- n.getItem(ciCfg)
		}
		close(ch)
	}()
	return ch
}

func (n *CIMgr) Stream() chan k8sRuntime.Object {
	n.RWMutex.RLock()
	streamer := n.streamCh
	n.RWMutex.RUnlock()
	return streamer
}

type listCiliumIdentity struct {
	*CIMgr
}

func NewReadCiliumIdentity() cilium.ListApisCiliumIoV2CiliumIdentityHandler {
	return &listCiliumIdentity{
		CIMgr: CIManager,
	}
}

func (lci *listCiliumIdentity) Handle(params cilium.ListApisCiliumIoV2CiliumIdentityParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/cilium.io/v2/ciliumidentities",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponderWithEvents(
		lci,
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
