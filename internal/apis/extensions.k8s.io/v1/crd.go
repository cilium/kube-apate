// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"fmt"
	"sync"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/apiextensions_v1"
	"github.com/cilium/kube-apate/internal/encoders"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"
	k8sAPIExtensionsV1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	k8sMetaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
)

const (
	CustomResourceDefinitions = "customresourcedefinitions"
)

var CRDManager = &CRDMgr{
	staticList:   &k8sAPIExtensionsV1.CustomResourceDefinitionList{},
	elemTemplate: &k8sAPIExtensionsV1.CustomResourceDefinition{},
}

type CRDConfig struct {
	Name string
}

type CRDMgr struct {
	sync.RWMutex
	totalGenElem int64
	staticList   *k8sAPIExtensionsV1.CustomResourceDefinitionList
	elemTemplate *k8sAPIExtensionsV1.CustomResourceDefinition
}

func (*CRDMgr) NewListInterface() k8sRuntime.Object {
	return &k8sAPIExtensionsV1.CustomResourceDefinitionList{}
}

func (*CRDMgr) NewInterface() k8sRuntime.Object {
	return &k8sAPIExtensionsV1.CustomResourceDefinition{}
}

func (n *CRDMgr) AddElements(elem int64) int64 {
	n.Lock()
	defer n.Unlock()
	oldElem := n.totalGenElem
	n.totalGenElem += elem
	return oldElem
}

func (n *CRDMgr) SetList(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sAPIExtensionsV1.CustomResourceDefinitionList)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sAPIExtensionsV1.CustomResourceDefinitionList", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.staticList = nss
	return nil
}

func (n *CRDMgr) SetObj(obj k8sRuntime.Object) error {
	nss, ok := obj.(*k8sAPIExtensionsV1.CustomResourceDefinition)
	if !ok {
		return fmt.Errorf("invalid object got %T, expecting *k8sAPIExtensionsV1.CustomResourceDefinition", obj)
	}
	n.Lock()
	defer n.Unlock()
	n.elemTemplate = nss
	return nil
}

func (n *CRDMgr) List(_, _ int64) (k8sRuntime.Object, error) {
	var (
		pom k8sMetaV1.PartialObjectMetadataList
	)

	for _, item := range n.staticList.Items {
		pom.Items = append(pom.Items,
			k8sMetaV1.PartialObjectMetadata{
				ObjectMeta: k8sMetaV1.ObjectMeta{
					Name: item.GetName(),
				},
			},
		)
	}

	return &pom, nil
}

type listCRDs struct {
	*CRDMgr
}

func NewListApiextensionsV1CustomResourceDefinition() apiextensions_v1.ListApiextensionsV1CustomResourceDefinitionHandler {
	return &listCRDs{
		CRDMgr: CRDManager,
	}
}

func (lCRD *listCRDs) Handle(params apiextensions_v1.ListApiextensionsV1CustomResourceDefinitionParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/apis/apiextensions.k8s.io/v1/customresourcedefinitions",
		"Params": params.HTTPRequest.URL.RawQuery,
	}).Debug("request received")

	return encoders.WatchOrListResponder(
		lCRD,
		&encoders.WatchOrListResponderCfg{
			Watch:          params.Watch,
			TimeoutSeconds: params.TimeoutSeconds,
			Limit:          params.Limit,
			Cont:           params.Continue,
		},
		getResponder,
	)
}
