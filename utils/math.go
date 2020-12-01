// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"strconv"
)

func Min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Cont(start, end, limit int64) string {
	total := end - start
	if total >= limit {
		return strconv.FormatInt(start+limit, 10)
	}
	return ""
}
