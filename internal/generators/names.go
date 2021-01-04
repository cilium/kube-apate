// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"strconv"
)

const (
	// maxNamespaces defines the maximum number of namespaces in a cluster.
	// The bigger the number of namespaces, the bigger is the cardinality of
	// labels used in the security identities. The number of namespaces in a
	// cluster only affects Cilium scalability with the cardinality of labels.
	maxNamespaces = int64(50)
)

func CCNPName(idx int64) string {
	return "ccnp-" + strconv.FormatInt(idx, 10)
}

func CNPName(idx int64) string {
	return "cnp-" + strconv.FormatInt(idx, 10)
}

func CIID(idx int64) int64 {
	// Since the cardinality of the labels depend on the namespace, we will
	// create an index depending on the max number of namespaces.
	ciIdx := idx / maxNamespaces
	// Avoid using any of the reserved identities
	ciIdx += 2000
	return ciIdx
}

func CIName(idx int64) string {
	ciIdx := CIID(idx)
	return strconv.FormatInt(ciIdx, 10)
}

func CLRPName(idx int64) string {
	return "clrp-" + strconv.FormatInt(idx, 10)
}

func CEName(idx int64) string {
	return PodName(idx)
}

func CNName(idx int64) string {
	return NodeName(idx)
}

func ESName(idx int64) string {
	return "es-" + strconv.FormatInt(idx, 10)
}

func NetPolName(idx int64) string {
	return "netpol-" + strconv.FormatInt(idx, 10)
}

func NamespaceName(idx int64) string {
	nsIdx := idx / maxNamespaces
	return "ns-" + strconv.FormatInt(nsIdx, 10)
}

func NodeName(idx int64) string {
	return "node-" + strconv.FormatInt(idx, 10)
}

// NodeNameOfPodIdx returns the Node Name that is hosting the Pod with the
// given index.
func NodeNameOfPodIdx(idx int64) string {
	hostIdx := idx / maxPodsPerHost
	return NodeName(hostIdx)
}

func PodName(idx int64) string {
	return "pod-" + strconv.FormatInt(idx, 10)
}

func ServiceName(idx int64) string {
	return "svc-" + strconv.FormatInt(idx, 10)
}
