// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"fmt"
)

func CiliumIdentityLabels(idx int64) map[string]string {
	return map[string]string{
		"io.cilium.k8s.policy.cluster":        "default",
		"io.cilium.k8s.policy.serviceaccount": "default",
		"io.kubernetes.pod.namespace":         NamespaceName(idx),
		"random-label-1":                      "thequickbrownfoxjumpsoverthelazydog-1",
		"random-label-2":                      "thequickbrownfoxjumpsoverthelazydog-2",
		"random-label-3":                      "thequickbrownfoxjumpsoverthelazydog-3",
		"random-label-4":                      "thequickbrownfoxjumpsoverthelazydog-4",
		"random-label-5":                      "thequickbrownfoxjumpsoverthelazydog-5",
		"random-label-6":                      "thequickbrownfoxjumpsoverthelazydog-6",
		"random-label-7":                      "thequickbrownfoxjumpsoverthelazydog-7",
		"random-label-8":                      "thequickbrownfoxjumpsoverthelazydog-8",
		"random-label-9":                      "thequickbrownfoxjumpsoverthelazydog-9",
		"random-label-10":                     "thequickbrownfoxjumpsoverthelazydog-10",
		"k8s-app.name":                        "frontend-a",
	}
}

func CiliumSecurityIdentityLabels(idx int64) map[string]string {
	m := CiliumIdentityLabels(idx)
	for k, v := range m {
		m["k8s:"+k] = v
		delete(m, k)
	}
	return m
}

func CiliumEndpointLabels(idx int64) []string {
	m := CiliumIdentityLabels(idx)
	var l []string
	for k, v := range m {
		l = append(l, fmt.Sprintf("k8s:%s=%s", k, v))
	}
	return l
}
