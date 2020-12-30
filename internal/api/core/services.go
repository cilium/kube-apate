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
	Service = "service"
)

var ServiceManager = &ServiceMgr{}

type ServiceConfig struct {
	Name        string
	Namespace   string
	ClusterIPv4 string
	ClusterIPv6 string
	UID         string
}

type ServiceMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sCoreV1.ServiceList
	elemTemplate *k8sCoreV1.Service
}

func (*ServiceMgr) NewListInterface() k8sRuntime.Object {
	return &k8sCoreV1.ServiceList{}
}

func (*ServiceMgr) NewInterface() k8sRuntime.Object {
	return &k8sCoreV1.Service{}
}

func (svcMgr *ServiceMgr) AddElements(elem int64) int64 {
	svcMgr.Lock()
	defer svcMgr.Unlock()
	oldElem := svcMgr.totalGenElem
	svcMgr.totalGenElem += elem
	return oldElem
}

func (svcMgr *ServiceMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.ServiceList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.ServiceList", obj)
	}
	svcMgr.Lock()
	defer svcMgr.Unlock()
	svcMgr.staticList = nss
	return nil
}

func (svcMgr *ServiceMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.Service)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.Service", obj)
	}
	svcMgr.Lock()
	defer svcMgr.Unlock()
	svcMgr.elemTemplate = nss
	return nil
}

func (svcMgr *ServiceMgr) getItem(cfg *ServiceConfig) *k8sCoreV1.Service {
	svcMgr.RLock()
	defer svcMgr.RUnlock()
	svc := svcMgr.elemTemplate.DeepCopy()
	svc.Name = cfg.Name
	svc.UID = k8sTypes.UID(cfg.UID)
	svc.Spec.ClusterIP = cfg.ClusterIPv4
	svc.Spec.ClusterIPs = nil
	if cfg.ClusterIPv4 != "" {
		svc.Spec.ClusterIPs = append(svc.Spec.ClusterIPs, cfg.ClusterIPv4)
	}
	if cfg.ClusterIPv6 != "" {
		svc.Spec.ClusterIPs = append(svc.Spec.ClusterIPs, cfg.ClusterIPv6)
	}
	svc.Labels["k8s-app.svc-name"] = cfg.Name
	svc.Spec.Selector = map[string]string{
		"k8s-app.guestbook": "web",
	}
	return svc
}

func (svcMgr *ServiceMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	svcMgr.RLock()
	nsList := svcMgr.staticList.DeepCopy()
	totalElems := svcMgr.totalGenElem
	svcMgr.RUnlock()

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

	svcCfg := &ServiceConfig{}
	for i := start; i < maxElemts; i++ {
		svcCfg.Name = generators.ServiceName(i)
		svcCfg.Namespace = generators.NamespaceName(i)
		svcCfg.UID = generators.ServiceUUID(i)
		svcCfg.ClusterIPv4, svcCfg.ClusterIPv6 = generators.GetServiceIP(i)
		svcCpy := svcMgr.getItem(svcCfg)
		nsList.Items = append(nsList.Items, *svcCpy)
	}

	return nsList, nil
}

type listAllServices struct {
	*ServiceMgr
}

func NewListCoreV1Service() core_v1.ListCoreV1ServiceForAllNamespacesHandler {
	return &listAllServices{
		ServiceMgr: ServiceManager,
	}
}

func (las *listAllServices) Handle(params core_v1.ListCoreV1ServiceForAllNamespacesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/api/v1/services",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		las,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
