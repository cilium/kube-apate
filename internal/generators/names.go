// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"strconv"
)

func CCNPName(idx int64) string {
	return "ccnp-" + strconv.FormatInt(idx, 10)
}

func CNPName(idx int64) string {
	return "cnp-" + strconv.FormatInt(idx, 10)
}

func CIName(idx int64) string {
	// Avoid using any of the reserved identities
	return strconv.FormatInt(2000+idx, 10)
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

// const uniqueSet = 50
// "k8s-app.name":                        fmt.Sprintf("front-a-%d", idx%uniqueSet),

func NamespaceName(idx int64) string {
	return "ns-" + strconv.FormatInt(idx, 10)
}

func NodeName(idx int64) string {
	return "node-" + strconv.FormatInt(idx, 10)
}

func PodName(idx int64) string {
	return "pod-" + strconv.FormatInt(idx, 10)
}

func ServiceName(idx int64) string {
	return "svc-" + strconv.FormatInt(idx, 10)
}
