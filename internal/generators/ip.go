// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"fmt"
	"net"

	ipPkg "github.com/containernetworking/plugins/pkg/ip"
)

var (
	// TODO: try making the prefix of the CIDRs a little bit more configurable.

	// From Kubernetes documentation the maximum supported configurations are:
	// No more than 100 pods per node
	// No more than 5000 nodes
	// No more than 150000 total pods
	// No more than 300000 total containers
	//
	// More info: https://kubernetes.io/docs/setup/best-practices/cluster-large/

	// HostIPs     -> 172.16.0.0/19 and fd00:beef:beef:0000::/115 gives 8192 nodes for the cluster
	maxHosts       = int64(1 << (32 - 19)) // 8192
	hostCIDRMaskv4 = net.CIDRMask(19, 32)
	hostCIDRMaskv6 = net.CIDRMask(115, 128)

	// PodCIDRs    -> X.X.X.X/25 and ::/121 gives 128 pods for each node

	// PodIPs      -> 10.0.0.0/12 and fd01:beef:beef:0000::/108 creates 8192*128 pods for the cluster
	maxPods       = int64(1 << (32 - 12)) // 1048576
	podCIDRMaskv4 = net.CIDRMask(12, 32)
	podCIDRv4     = net.IPNet{
		IP:   net.ParseIP("10.0.0.0"),
		Mask: podCIDRMaskv4,
	}
	podCIDRMaskv6 = net.CIDRMask(108, 128)
	podCIDRv6     = net.IPNet{
		IP:   net.ParseIP("fd01:beef:beef::"),
		Mask: podCIDRMaskv6,
	}

	// ServicesIPs -> 172.20.0.0/16 and fd02:beef:beef:0000::/112 creates 65536 services for the cluster
	maxServices       = int64(1 << (32 - 16)) // 32768
	serviceCIDRMaskv4 = net.CIDRMask(16, 32)
	serviceCIDRv4     = net.IPNet{
		IP:   net.ParseIP("172.20.0.0"),
		Mask: serviceCIDRMaskv4,
	}
	serviceCIDRMaskv6 = net.CIDRMask(112, 128)
	serviceCIDRv6     = net.IPNet{
		IP:   net.ParseIP("fd02:beef:beef:0000::"),
		Mask: serviceCIDRMaskv6,
	}
)

// GetHostIP returns the v4 and v6 IP address for the given hostIdx
func GetHostIP(hostIdx int64) (string, string) {
	if hostIdx >= maxHosts {
		return "", ""
	}
	v4C := hostIdx / (1 << 8)
	v4D := hostIdx % (1 << 8)
	v4 := fmt.Sprintf("172.16.%d.%d", v4C, v4D)
	v6 := fmt.Sprintf("fd00:beef:beef::%x", hostIdx)
	return v4, v6
}

// GetPodIP returns the v4 and v6 IP addresses for the given podIdx
func GetPodIP(idx int64) (string, string) {
	if idx >= maxPods {
		return "", ""
	}
	// This will assign IPs from 10.0.0.0/14 -> 10.0.0.0 - 10.15.255.255
	v4B := (idx / 256) / 256
	v4C := (idx / 256) % 256
	v4D := idx % 256
	v4 := fmt.Sprintf("10.%d.%d.%d", v4B, v4C, v4D)

	// This will assign IPs from fd01:beef:beef::/110 -> fd01:beef:beef:: - fd01:beef:beef::f:ffff
	v6G := idx / (1 << 16)
	v6H := idx % (1 << 16)
	v6 := fmt.Sprintf("fd01:beef:beef::%x:%x", v6G, v6H)

	return v4, v6
}

// GetPodCIDR returns the v4 and v6 podCIDR for the given hostIdx
func GetPodCIDR(hostIdx int64) (string, string) {
	if hostIdx >= maxHosts {
		return "", ""
	}
	v4B := (hostIdx / 2) / 256
	v4C := (hostIdx / 2) % 256
	var v4D int
	if hostIdx%2 != 0 {
		v4D = 128
	}
	v4 := fmt.Sprintf("10.%d.%d.%d/25", v4B, v4C, v4D)

	var v6H int64
	if hostIdx%2 != 0 {
		v6H = 128
	}
	v6G := (((hostIdx / 2) % 65536) >> 8) & 0xFFFF
	v6H = ((((hostIdx / 2) % 65536) << 8) | v6H) & 0xFFFF
	v6 := fmt.Sprintf("fd01:beef:beef::%x:%x/121", v6G, v6H)

	return v4, v6
}

// GetHostOfPodIPv4 returns the Host IPv4 address of the that has the given pod
// IP.
func GetHostOfPodIPv4(podIP string) string {
	ip, _, err := net.ParseCIDR(podIP + "/25")
	if err != nil {
		return ""
	}
	ip = ip.To4()
	if !podCIDRv4.Contains(ip) {
		return ""
	}
	// 10.0.0.0       ->  0 * 512 + 0*2 + 0/128     = host idx 0
	// 10.0.2.0       ->  0 * 512 + 2*2 + 0/128     = host idx 4
	// 10.0.2.129     ->  0 * 512 + 2*2 + 129/128   = host idx 5
	// 10.0.3.0       ->  0 * 512 + 3*2 + 0/128     = host idx 6
	// 10.1.0.0       ->  1 * 512 + 0*2 + 0/128     = host idx 512
	// 10.1.1.0       ->  1 * 512 + 1*2 + 0/128     = host idx 514
	// 10.3.0.0       ->  3 * 512 + 0*2 + 0/128     = host idx 1536
	// 10.15.255.128  -> 15 * 512 + 255*2 + 0/128   = host idx 8191
	hostIdx := int64(ip[1])*512 + int64(ip[2])*2 + (int64(ip[3]) / 128)
	v4, _ := GetHostIP(hostIdx)
	return v4
}

// GetHostOfPodIPv6 returns the Host IPv6 address of the that has the given pod
// IP.
func GetHostOfPodIPv6(podIP string) string {
	ip, _, err := net.ParseCIDR(podIP + "/121")
	if err != nil {
		return ""
	}
	ip = ip.To16()
	if !podCIDRv6.Contains(ip) {
		return ""
	}
	// fd01:beef:beef::0:0/121    -> 0 * 65536 + 0/128  = host idx 0
	// fd01:beef:beef::0:ff80/121 -> 0 * 65536 + 65408/128 = host idx 511
	// fd01:beef:beef::f:ff80/121 -> 2 * 65536 + 65408/128 = host idx 131583
	hostIdx := (int64(ip[13])*2<<8)&0xFFFF + ((int64(ip[14])*256 + int64(ip[15])) / 128)
	_, v6 := GetHostIP(hostIdx)
	return v6
}

// GetServiceIP returns the v4 and v6 IP address for the given serviceIdx.
func GetServiceIP(svcIdx int64) (string, string) {
	if svcIdx >= maxServices {
		return "", ""
	}
	v4C := svcIdx / (1 << 8)
	v4D := svcIdx % (1 << 8)
	v4 := fmt.Sprintf("172.20.%d.%d", v4C, v4D)
	v6 := fmt.Sprintf("fd02:beef:beef::%x", svcIdx)
	return v4, v6
}

// IPFromCIDR returns the IP address with the index 'idx' from the provided CIDR
func IPFromCIDR(cidr string, idx int) string {
	ip, _, err := net.ParseCIDR(cidr)
	if err != nil {
		return ""
	}
	for i := 0; i < idx; i++ {
		ip = ipPkg.NextIP(ip)
	}
	return ip.String()
}

func IsBlockListedPod(ipStr string) bool {
	_, ipnet, _ := net.ParseCIDR("10.128.0.0/26")
	ip := net.ParseIP(ipStr)
	return ipnet.Contains(ip)
}
