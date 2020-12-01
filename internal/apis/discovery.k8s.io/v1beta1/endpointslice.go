// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"fmt"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/discovery_v1beta1"
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sDiscoveryV1beta1 "k8s.io/api/discovery/v1beta1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	EndpointSlice = "endpointslice"
)

var ESManager = &ESMgr{
	staticList:   &k8sDiscoveryV1beta1.EndpointSliceList{},
	elemTemplate: &k8sDiscoveryV1beta1.EndpointSlice{},
}

type ESConfig struct {
	Name string
	UUID string
}

type ESMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sDiscoveryV1beta1.EndpointSliceList
	elemTemplate *k8sDiscoveryV1beta1.EndpointSlice
}

func (*ESMgr) NewListInterface() k8sRuntime.Object {
	return &k8sDiscoveryV1beta1.EndpointSliceList{}
}

func (*ESMgr) NewInterface() k8sRuntime.Object {
	return &k8sDiscoveryV1beta1.EndpointSlice{}
}

func (n *ESMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *ESMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sDiscoveryV1beta1.EndpointSliceList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sDiscoveryV1beta1.EndpointSliceList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *ESMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sDiscoveryV1beta1.EndpointSlice)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sDiscoveryV1beta1.EndpointSlice", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *ESMgr) getItem(cfg *ESConfig) *k8sDiscoveryV1beta1.EndpointSlice {
	n.RLock()
	defer n.RUnlock()
	es := n.elemTemplate.DeepCopy()
	es.Name = cfg.Name
	es.UID = k8sTypes.UID(cfg.UUID)
	return es
}

func (n *ESMgr) List(start, limit int64) (k8sRuntime.Object, error) {
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

	nCfg := &ESConfig{}
	for i := start; i < maxElemts; i++ {
		nCfg.Name = generators.ESName(i)
		nCfg.UUID = generators.ESUUID(i)
		ns := n.getItem(nCfg)
		ccnpList.Items = append(ccnpList.Items, *ns)
	}

	return ccnpList, nil
}

type getES struct {
	*ESMgr
}

func NewReadDiscoveryV1beta1NamespacedEndpointSliceHandler() discovery_v1beta1.ReadDiscoveryV1beta1NamespacedEndpointSliceHandler {
	return &getES{
		ESMgr: ESManager,
	}
}

func (ges *getES) Handle(params discovery_v1beta1.ReadDiscoveryV1beta1NamespacedEndpointSliceParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method":    "GET",
		"URL":       "/apis/discovery.k8s.io/v1beta1/namespaces/{namespace}/endpointslices/{endpointslice}",
		"Namespace": params.Namespace,
		"Name":      params.Name,
	}).Debug("request received")

	obj := ges.getItem(&ESConfig{Name: params.Name})

	return getResponder(obj)
}

type listAllES struct {
	*ESMgr
}

func NewListDiscoveryV1beta1NamespacedEndpointSliceHandler() discovery_v1beta1.ListDiscoveryV1beta1EndpointSliceForAllNamespacesHandler {
	return &listAllES{
		ESMgr: ESManager,
	}
}

func (les *listAllES) Handle(params discovery_v1beta1.ListDiscoveryV1beta1EndpointSliceForAllNamespacesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/discovery.k8s.io/v1beta1/endpointslices",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		les,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
