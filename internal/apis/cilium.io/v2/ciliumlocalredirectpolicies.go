// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v2

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/cilium"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	ciliumV2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	CLRP = "ciliumlocalredirectpolicies.cilium.io"
)

var CLRPManager = &CLRPMgr{
	staticList:   &ciliumV2.CiliumLocalRedirectPolicyList{},
	elemTemplate: &ciliumV2.CiliumLocalRedirectPolicy{},
}

type CLRPConfig struct {
	Name string
	UUID string
}

type CLRPMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *ciliumV2.CiliumLocalRedirectPolicyList
	elemTemplate *ciliumV2.CiliumLocalRedirectPolicy
}

func (*CLRPMgr) NewListInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumLocalRedirectPolicyList{}
}

func (*CLRPMgr) NewInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumLocalRedirectPolicy{}
}

func (n *CLRPMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CLRPMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumLocalRedirectPolicyList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumLocalRedirectPolicyList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CLRPMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumLocalRedirectPolicy)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumLocalRedirectPolicy", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CLRPMgr) getItem(cfg *CLRPConfig) *ciliumV2.CiliumLocalRedirectPolicy {
	n.RLock()
	defer n.RUnlock()
	ce := n.elemTemplate.DeepCopy()
	ce.Name = cfg.Name
	ce.UID = k8sTypes.UID(cfg.UUID)
	return ce
}

func (n *CLRPMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	clrpList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(clrpList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		clrpList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	clrpList.Continue = utils.Cont(start, totalElems, limit)

	clrpCfg := &CLRPConfig{}
	for i := start; i < maxElemts; i++ {
		podIPv4, _ := generators.GetPodIP(i)
		if generators.IsBlockListedPod(podIPv4) {
			continue
		}
		clrpCfg.Name = generators.CLRPName(i)
		clrpCfg.UUID = generators.CLRPUUID(i)
		ce := n.getItem(clrpCfg)
		clrpList.Items = append(clrpList.Items, *ce)
	}

	return clrpList, nil
}

type postCLRP struct {
	*CLRPMgr
}

func NewPostCLRP() cilium.ListApisCiliumIoV2CiliumLocalRedirectPolicyHandler {
	return &postCLRP{
		CLRPMgr: CLRPManager,
	}
}

func (pcn *postCLRP) read(r io.ReadCloser) (k8sRuntime.Object, error) {
	byts, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var cn ciliumV2.CiliumLocalRedirectPolicy
	_, _, err = codecs.UniversalDecoder(ciliumV2.SchemeGroupVersion).Decode(byts, nil, &cn)
	return &cn, err
}

func (pcn *postCLRP) Handle(params cilium.ListApisCiliumIoV2CiliumLocalRedirectPolicyParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "POST",
		"URL":    "/apis/cilium.io/v2/ciliumlocalredirectpolicies",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	obj, err := pcn.read(params.HTTPRequest.Body)
	if err != nil {
		log.WithError(err).Debug("Unable to handle request")
		return middleware.Error(http.StatusInternalServerError, nil)
	}

	return postResponder(obj)
}
