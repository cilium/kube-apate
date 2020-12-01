// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"fmt"
)

const (
	maxCEPsPerNode = 128
)

func CCNPUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0001-0000%08x", idx)
}

func CNPUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0002-0000%08x", idx)
}
func CIUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0003-0000%08x", idx)
}

func CLRPUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0004-0000%08x", idx)
}

func CEUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0005-0000%08x", idx)
}

func CELocalID(idx int64) int64 {
	return idx % maxCEPsPerNode
}

func CNUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0006-0000%08x", idx)
}

func ESUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0007-0000%08x", idx)
}

func NetPolUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0008-0000%08x", idx)
}

func NamespaceUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-0009-0000%08x", idx)
}

func NodeUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-000a-0000%08x", idx)
}

func PodUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-000b-0000%08x", idx)
}

func ServiceUUID(idx int64) string {
	return fmt.Sprintf("0000000-0000-0000-000c-0000%08x", idx)
}

func ContainerID(idx int64) string {
	return fmt.Sprintf("1ab5b89c09169b1fb66b2f344be373ace62ce47865f422f021a80310%08x", idx)
}
