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
	Namespace = "namespace"
)

var NamespaceManager = &NamespaceMgr{}

type NamespaceConfig struct {
	Name string
	UUID string
}

type NamespaceMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sCoreV1.NamespaceList
	elemTemplate *k8sCoreV1.Namespace
}

func (*NamespaceMgr) NewListInterface() k8sRuntime.Object {
	return &k8sCoreV1.NamespaceList{}
}

func (*NamespaceMgr) NewInterface() k8sRuntime.Object {
	return &k8sCoreV1.Namespace{}
}

func (n *NamespaceMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *NamespaceMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.NamespaceList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.NamespaceList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *NamespaceMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.Namespace)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.Namespace", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *NamespaceMgr) getItem(cfg *NamespaceConfig) *k8sCoreV1.Namespace {
	n.RLock()
	defer n.RUnlock()
	ns := n.elemTemplate.DeepCopy()
	ns.Name = cfg.Name
	ns.UID = k8sTypes.UID(cfg.UUID)
	return ns
}

func (n *NamespaceMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	nsList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(nsList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		nsList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	nsList.Continue = utils.Cont(start, totalElems, limit)

	nsCfg := &NamespaceConfig{}

	for i := start; i < maxElemts; i++ {
		nsCfg.Name = generators.NamespaceName(i)
		nsCfg.UUID = generators.NamespaceUUID(i)
		ns := n.getItem(nsCfg)
		nsList.Items = append(nsList.Items, *ns)
	}

	return nsList, nil
}

type readNamespace struct {
	*NamespaceMgr
}

func NewReadCoreV1Namespace() core_v1.ReadCoreV1NamespaceHandler {
	return &readNamespace{
		NamespaceMgr: NamespaceManager,
	}
}

func (lan *readNamespace) Handle(params core_v1.ReadCoreV1NamespaceParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/api/v1/namespaces/{namespace}",
		"Name":   params.Name,
	}).Debug("request received")

	obj := lan.getItem(&NamespaceConfig{Name: params.Name})

	return getResponder(obj)
}

type listAllNamespaces struct {
	*NamespaceMgr
}

func NewListCoreV1Namespace() core_v1.ListCoreV1NamespaceHandler {
	return &listAllNamespaces{
		NamespaceMgr: NamespaceManager,
	}
}

func (lan *listAllNamespaces) Handle(params core_v1.ListCoreV1NamespaceParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/api/v1/namespaces",
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
