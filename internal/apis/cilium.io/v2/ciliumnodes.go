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
	"github.com/cilium/kube-apate/internal/encoders"
	"github.com/cilium/kube-apate/internal/generators"
	"github.com/cilium/kube-apate/utils"

	ciliumV2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"github.com/cilium/cilium/pkg/node/addressing"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	k8sTypes "k8s.io/apimachinery/pkg/types"
)

const (
	CN = "ciliumnodes.cilium.io"
)

var CiliumNodeManager = &CiliumNodeMgr{}

type CiliumNodeConfig struct {
	Name           string
	IPv4           string
	IPv6           string
	PodCIDRv4      string
	PodCIDRv6      string
	CiliumHostIP   string
	CiliumHealthIP string
	UUID           string
	PodUUID        string
}

type CiliumNodeMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *ciliumV2.CiliumNodeList
	elemTemplate *ciliumV2.CiliumNode
}

func (*CiliumNodeMgr) NewListInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumNodeList{}
}

func (*CiliumNodeMgr) NewInterface() k8sRuntime.Object {
	return &ciliumV2.CiliumNode{}
}

func (n *CiliumNodeMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CiliumNodeMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumNodeList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumNodeList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CiliumNodeMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*ciliumV2.CiliumNode)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *ciliumV2.CiliumNode", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CiliumNodeMgr) getItem(cfg *CiliumNodeConfig) *ciliumV2.CiliumNode {
	n.RLock()
	defer n.RUnlock()
	cn := n.elemTemplate.DeepCopy()
	cn.Name = cfg.Name
	cn.UID = k8sTypes.UID(cfg.UUID)
	owner := cn.OwnerReferences[0]
	owner.Name = cfg.Name
	owner.UID = k8sTypes.UID(cfg.PodUUID)
	cn.Spec.Addresses = nil
	cn.Spec.IPAM.PodCIDRs = nil
	if cfg.IPv4 != "" {
		cn.Spec.Addresses = append(
			cn.Spec.Addresses,
			ciliumV2.NodeAddress{
				Type: addressing.NodeInternalIP,
				IP:   cfg.IPv4,
			})
	}
	if cn.Name != "" {
		cn.Spec.Addresses = append(
			cn.Spec.Addresses,
			ciliumV2.NodeAddress{
				Type: addressing.NodeHostName,
				IP:   cn.Name,
			})
	}

	if cfg.PodCIDRv4 != "" {
		cn.Spec.IPAM.PodCIDRs = append(
			cn.Spec.IPAM.PodCIDRs,
			cfg.PodCIDRv4,
		)
	}
	if cfg.PodCIDRv6 != "" {
		cn.Spec.IPAM.PodCIDRs = append(
			cn.Spec.IPAM.PodCIDRs,
			cfg.PodCIDRv6,
		)
	}
	return cn
}

func (n *CiliumNodeMgr) List(start, limit int64) (k8sRuntime.Object, error) {
	n.RLock()
	cnList := n.staticList.DeepCopy()
	totalElems := n.totalGenElem
	n.RUnlock()

	if start == 0 {
		totalElems += int64(len(cnList.Items))
		// we are going to send the static elements so no need to generate
		// elements from [0, start]
		start = totalElems
	} else {
		cnList.Items = nil
	}

	maxElemts := utils.Min(totalElems, limit)
	cnList.Continue = utils.Cont(start, totalElems, limit)

	nCfg := &CiliumNodeConfig{}
	for i := start; i < maxElemts; i++ {
		nCfg.Name = generators.CNName(i)
		nCfg.UUID = generators.CNUUID(i)
		nCfg.PodUUID = generators.PodUUID(i)
		nCfg.IPv4, nCfg.IPv6 = generators.GetHostIP(i)
		nCfg.PodCIDRv4, nCfg.PodCIDRv6 = generators.GetPodCIDR(i)
		nCfg.CiliumHostIP = generators.IPFromCIDR(nCfg.PodCIDRv4, 125)
		nCfg.CiliumHealthIP = generators.IPFromCIDR(nCfg.PodCIDRv4, 126)
		ns := n.getItem(nCfg)
		cnList.Items = append(cnList.Items, *ns)
	}

	return cnList, nil
}

func (*CiliumNodeMgr) read(r io.ReadCloser) (k8sRuntime.Object, error) {
	byts, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var cn ciliumV2.CiliumNode
	_, _, err = codecs.UniversalDecoder(ciliumV2.SchemeGroupVersion).Decode(byts, nil, &cn)
	return &cn, err
}

type postCiliumNode struct {
	*CiliumNodeMgr
}

func NewPostCiliumNode() cilium.PostApisCiliumIoV2CiliumNodesHandler {
	return &postCiliumNode{
		CiliumNodeMgr: CiliumNodeManager,
	}
}

func (pcn *postCiliumNode) Handle(params cilium.PostApisCiliumIoV2CiliumNodesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "POST",
		"URL":    "/apis/cilium.io/v2/ciliumnodes",
	}).Debug("request received")

	obj, err := pcn.read(params.HTTPRequest.Body)
	if err != nil {
		log.WithError(err).Debug("Unable to handle request")
		return middleware.Error(http.StatusInternalServerError, nil)
	}

	return postResponder(obj)
}

type readCiliumNode struct {
	*CiliumNodeMgr
}

func NewReadCiliumNode() cilium.ReadApisCiliumIoV2CiliumNodeHandler {
	return &readCiliumNode{
		CiliumNodeMgr: CiliumNodeManager,
	}
}

func (rcn *readCiliumNode) Handle(params cilium.ReadApisCiliumIoV2CiliumNodeParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/cilium.io/v2/ciliumnodes/{name}",
		"Name":   params.Name,
	}).Debug("request received")

	// Send the templated structure since Cilium will do a GET on its own node
	cn := rcn.elemTemplate.DeepCopy()

	return getResponder(cn)
}

type putCiliumNode struct {
	*CiliumNodeMgr
}

func NewPutCiliumNode() cilium.PutApisCiliumIoV2CiliumNodesHandler {
	return &putCiliumNode{
		CiliumNodeMgr: CiliumNodeManager,
	}
}

func (pcn *putCiliumNode) Handle(params cilium.PutApisCiliumIoV2CiliumNodesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "PUT",
		"URL":    "/apis/cilium.io/v2/ciliumnodes/{name}",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	obj, err := pcn.read(params.HTTPRequest.Body)
	if err != nil {
		log.WithError(err).Debug("Unable to handle request")
		return middleware.Error(http.StatusInternalServerError, nil)
	}

	return postResponder(obj)
}

type listCiliumNode struct {
	*CiliumNodeMgr
}

func NewListCiliumNode() cilium.ListApisCiliumIoV2CiliumNodesHandler {
	return &listCiliumNode{
		CiliumNodeMgr: CiliumNodeManager,
	}
}

func (lcn *listCiliumNode) Handle(params cilium.ListApisCiliumIoV2CiliumNodesParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/cilium.io/v2/ciliumnodes",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lcn,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
