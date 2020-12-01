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
	Pod       = "pod"
	uniqueSet = 50
)

func PodLabels(idx int64) map[string]string {
	return map[string]string{
		"k8s-app.name":    fmt.Sprintf("front-a-%d", idx%uniqueSet),
		"random-label-1":  "thequickbrownfoxjumpsoverthelazydog-1",
		"random-label-2":  "thequickbrownfoxjumpsoverthelazydog-2",
		"random-label-3":  "thequickbrownfoxjumpsoverthelazydog-3",
		"random-label-4":  "thequickbrownfoxjumpsoverthelazydog-4",
		"random-label-5":  "thequickbrownfoxjumpsoverthelazydog-5",
		"random-label-6":  "thequickbrownfoxjumpsoverthelazydog-6",
		"random-label-7":  "thequickbrownfoxjumpsoverthelazydog-7",
		"random-label-8":  "thequickbrownfoxjumpsoverthelazydog-8",
		"random-label-9":  "thequickbrownfoxjumpsoverthelazydog-9",
		"random-label-10": "thequickbrownfoxjumpsoverthelazydog-10",
	}
}

var PodManager = &PodMgr{}

type PodConfig struct {
	Name     string
	HostIP   string
	PodIPs   []k8sCoreV1.PodIP
	PodIP    string
	NodeName string
	Idx      int64
	UUID     string
}

type PodMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sCoreV1.PodList
	elemTemplate *k8sCoreV1.Pod
}

func (*PodMgr) NewListInterface() k8sRuntime.Object {
	return &k8sCoreV1.PodList{}
}

func (*PodMgr) NewInterface() k8sRuntime.Object {
	return &k8sCoreV1.Pod{}
}

func (p *PodMgr) AddElements(elem int64) int64 {
	p.Lock()
	defer p.Unlock()
	oldElem := p.totalGenElem
	p.totalGenElem = elem
	return oldElem
}

func (p *PodMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.PodList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.PodList", obj)
	}
	p.Lock()
	defer p.Unlock()
	p.staticList = nss
	return nil
}

func (p *PodMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sCoreV1.Pod)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sCoreV1.Pod", obj)
	}
	p.Lock()
	defer p.Unlock()
	p.elemTemplate = nss
	return nil
}

func (p *PodMgr) getItem(cfg *PodConfig) *k8sCoreV1.Pod {
	p.RLock()
	defer p.RUnlock()
	pod := p.elemTemplate.DeepCopy()
	pod.Name = cfg.Name
	pod.UID = k8sTypes.UID(cfg.UUID)
	pod.Status.PodIPs = cfg.PodIPs
	pod.Status.PodIP = cfg.PodIP
	pod.Status.HostIP = cfg.HostIP
	pod.Spec.NodeName = cfg.NodeName
	pod.Labels = PodLabels(cfg.Idx)
	return pod
}

func (p *PodMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	p.RLock()
	podList := p.staticList.DeepCopy()
	totalElems := p.totalGenElem
	p.RUnlock()

	if start == 0 {
		totalElems += int64(len(podList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		podList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	podList.Continue = utils.Cont(start, totalElems, limit)

	pCfg := &PodConfig{}
	for i := start; i < maxElemts; i++ {
		podIPv4, podIPv6 := generators.GetPodIP(i)
		if generators.IsBlockListedPod(podIPv4) {
			continue
		}
		pCfg.Name = generators.PodName(i)
		pCfg.PodIP = podIPv4
		pCfg.PodIPs = []k8sCoreV1.PodIP{
			{
				IP: podIPv4,
			},
			{
				IP: podIPv6,
			},
		}
		pCfg.HostIP = generators.GetHostOfPodIPv4(podIPv4)
		pCfg.NodeName = generators.NodeName(i)
		pCfg.UUID = generators.PodUUID(i)
		pCfg.Idx = i
		pCpy := p.getItem(pCfg)
		podList.Items = append(podList.Items, *pCpy)
	}

	return podList, nil
}

type listAllPods struct {
	*PodMgr
}

func NewListCoreV1Pod() core_v1.ListCoreV1PodForAllNamespacesHandler {
	return &listAllPods{
		PodMgr: PodManager,
	}
}

func (lap *listAllPods) Handle(params core_v1.ListCoreV1PodForAllNamespacesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/api/v1/pods",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lap,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
