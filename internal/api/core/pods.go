// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package core

import (
	"fmt"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/core_v1"
	management "github.com/cilium/kube-apate/api/management/v1/server/restapi/cilium"
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sCoreV1 "k8s.io/api/core/v1"
	k8sMetaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sFields "k8s.io/apimachinery/pkg/fields"
	k8sLabels "k8s.io/apimachinery/pkg/labels"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	k8sGeneric "k8s.io/apiserver/pkg/registry/generic"
	k8sStorage "k8s.io/apiserver/pkg/storage"
)

const (
	Pod = "pod"
)

var PodManager = &PodMgr{
	staticList:   &k8sCoreV1.PodList{},
	elemTemplate: &k8sCoreV1.Pod{},
	streamCh:     make(chan k8sRuntime.Object),
}

type PodConfig struct {
	Name      string
	HostIP    string
	PodIPs    []k8sCoreV1.PodIP
	PodIP     string
	NodeName  string
	Namespace string
	Idx       int64
	UUID      string
	PodLabels map[string]string
}

type PodMgr struct {
	logger
	sync.RWMutex
	streamCh     chan k8sRuntime.Object
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
	p.totalGenElem += elem
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
	pod.Namespace = cfg.Namespace
	pod.UID = k8sTypes.UID(cfg.UUID)
	pod.Status.PodIPs = cfg.PodIPs
	pod.Status.PodIP = cfg.PodIP
	pod.Status.HostIP = cfg.HostIP
	pod.Spec.NodeName = cfg.NodeName
	pod.Labels = cfg.PodLabels
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

	genPods := p.GenObjs(start, maxElemts)
	for obj := range genPods {
		pod := obj.(*k8sCoreV1.Pod)
		podList.Items = append(podList.Items, *pod)
	}

	return podList, nil
}

func (p *PodMgr) GenObjs(start, maxElemts int64) <-chan k8sRuntime.Object {
	ch := make(chan k8sRuntime.Object)
	go func() {
		for i := start; i < maxElemts; i++ {
			ch <- p.GenObj(i)
		}
		close(ch)
	}()
	return ch
}

func (p *PodMgr) GenObj(idx int64) k8sRuntime.Object {
	pCfg := &PodConfig{}
	podIPv4, podIPv6 := generators.GetPodIP(idx)
	if generators.IsBlockListedPod(podIPv4) {
		return nil
	}
	pCfg.Name = generators.PodName(idx)
	pCfg.Namespace = generators.NamespaceName(idx)
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
	pCfg.NodeName = generators.NodeNameOfPodIdx(idx)
	pCfg.UUID = generators.PodUUID(idx)
	pCfg.Idx = idx
	pCfg.PodLabels = generators.PodLabels(idx)
	return p.getItem(pCfg)
}

func (p *PodMgr) Stream() chan k8sRuntime.Object {
	p.RWMutex.RLock()
	streamer := p.streamCh
	p.RWMutex.RUnlock()
	return streamer
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

	return encoders.WatchOrListResponderWithEvents(
		lap,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
			Matcher: func(label k8sLabels.Selector, field k8sFields.Selector) k8sStorage.SelectionPredicate {
				return k8sStorage.SelectionPredicate{
					Label:       label,
					Field:       field,
					GetAttrs:    GetPodAttrs,
					IndexFields: []string{"spec.nodeName"},
				}
			},
			FieldSelector: params.FieldSelector,
			LabelSelector: params.LabelSelector,
		},
		getResponder,
		coreEncoder,
	)
}

type managePods struct {
	podMgr *PodMgr
	cepMgr encoders.ObjHandler
}

func NewPodsMgr(cepMgr encoders.ObjHandler) management.PostManagementKubernetesIoV1PodsHandler {
	return &managePods{
		podMgr: PodManager,
		cepMgr: cepMgr,
	}
}

func (lPods *managePods) Handle(params management.PostManagementKubernetesIoV1PodsParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "POST",
		"URL":    "/management/kubernetes.io/api/v1/pods",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	totalPods := params.Options.Add - params.Options.Del

	if params.Options.WithDependents {
		encoders.GenerateK8sEventsWithDependent(lPods.podMgr, lPods.cepMgr, totalPods)
	} else {
		encoders.GenerateK8sEvents(lPods.podMgr, totalPods)
	}

	return management.NewPostManagementKubernetesIoV1PodsAccepted()
}

// Code retrieved from Kubernetes to avoid importing the k8s.io/kubernetes

// ToSelectableFields returns a field set that represents the object
// TODO: fields are not labels, and the validation rules for them do not apply.
func ToSelectableFields(pod *k8sCoreV1.Pod) k8sFields.Set {
	// The purpose of allocation with a given number of elements is to reduce
	// amount of allocations needed to create the fields.Set. If you add any
	// field here or the number of object-meta related fields changes, this should
	// be adjusted.
	podSpecificFieldsSet := make(k8sFields.Set, 9)
	podSpecificFieldsSet["spec.nodeName"] = pod.Spec.NodeName
	podSpecificFieldsSet["spec.restartPolicy"] = string(pod.Spec.RestartPolicy)
	podSpecificFieldsSet["spec.schedulerName"] = string(pod.Spec.SchedulerName)
	podSpecificFieldsSet["spec.serviceAccountName"] = string(pod.Spec.ServiceAccountName)
	podSpecificFieldsSet["status.phase"] = string(pod.Status.Phase)
	// TODO: add podIPs as a downward API value(s) with proper format
	podIP := ""
	if len(pod.Status.PodIPs) > 0 {
		podIP = string(pod.Status.PodIPs[0].IP)
	}
	podSpecificFieldsSet["status.podIP"] = podIP
	podSpecificFieldsSet["status.nominatedNodeName"] = string(pod.Status.NominatedNodeName)
	return k8sGeneric.AddObjectMetaFieldsSet(podSpecificFieldsSet, &pod.ObjectMeta, true)
}

// GetPodAttrs returns labels and fields of a given object for filtering purposes.
func GetPodAttrs(obj k8sRuntime.Object) (k8sLabels.Set, k8sFields.Set, error) {
	watchEvent, ok := obj.(*k8sMetaV1.WatchEvent)
	if ok {
		obj = watchEvent.Object.Object
	}
	pod, ok := obj.(*k8sCoreV1.Pod)
	if !ok {
		return nil, nil, fmt.Errorf("not a pod, but a %T", obj)
	}
	return pod.ObjectMeta.Labels, ToSelectableFields(pod), nil
}
