// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"encoding/binary"
	"net"
	"strconv"

	ipPkg "github.com/containernetworking/plugins/pkg/ip"
)

const (
	// octetCap is the capacity of an octet.
	octetCap = 1 << 8 // 256

	// quartetCap is the capacity of a quartet.
	quartetCap = 1 << 16 // 65536

	// maxHosts is the maximum number of hosts accepted on a cluster
	maxHosts    = int64(1 << maxHostsPo2) // 8192
	maxHostsPo2 = 13

	// maxPodsPerHost is the maximum number of pods accepted on a host.
	maxPodsPerHost    = int64(1 << maxPodsPerHostPo2) // 128
	maxPodsPerHostPo2 = 7

	// maxPods is the maximum number of pods accepted on a cluster.
	maxPods    = int64(1 << maxPodsPo2) // 1048576
	maxPodsPo2 = 20

	// maxServices is the maximum number of services accepted on a cluster.
	maxServices    = int64(1 << maxServicesPo2) // 32768
	maxServicesPo2 = 16

	ipV4Bits = 8 * net.IPv4len
	ipV6Bits = 8 * net.IPv6len
)

var (
	// From Kubernetes documentation the maximum supported configurations are:
	// No more than 100 pods per node
	// No more than 5000 nodes
	// No more than 150000 total pods
	// No more than 300000 total containers
	//
	// More info: https://kubernetes.io/docs/setup/best-practices/cluster-large/

	// HostIPs     -> 172.16.0.0/19 and fd00:beef:beef:0000::/115 gives 8192 nodes for the cluster
	hostCIDRMaskv4 = net.CIDRMask(ipV4Bits-maxHostsPo2, ipV4Bits)
	hostCIDRv4     = net.IPNet{
		IP:   net.ParseIP("172.16.0.0"),
		Mask: hostCIDRMaskv4,
	}
	hostCIDRMaskv6 = net.CIDRMask(ipV6Bits-maxHostsPo2, ipV6Bits)
	hostCIDRv6     = net.IPNet{
		IP:   net.ParseIP("fd00:beef:beef::"),
		Mask: hostCIDRMaskv6,
	}

	// PodCIDRs    -> X.X.X.X/25 and ::/121 gives 128 pods for each node
	hostPodCIDRMaskv4 = net.CIDRMask(ipV4Bits-maxPodsPerHostPo2, ipV4Bits)
	hostPodCIDRMaskv6 = net.CIDRMask(ipV6Bits-maxPodsPerHostPo2, ipV6Bits)

	// PodIPs      -> 10.0.0.0/12 and fd01:beef:beef:0000::/108 creates 8192*128 pods for the cluster
	podCIDRMaskv4 = net.CIDRMask(ipV4Bits-maxPodsPo2, ipV4Bits)
	podCIDRv4     = net.IPNet{
		IP:   net.ParseIP("10.0.0.0"),
		Mask: podCIDRMaskv4,
	}
	podCIDRMaskv6 = net.CIDRMask(ipV6Bits-maxPodsPo2, ipV6Bits)
	podCIDRv6     = net.IPNet{
		IP:   net.ParseIP("fd01:beef:beef::"),
		Mask: podCIDRMaskv6,
	}

	// ServicesIPs -> 172.20.0.0/16 and fd02:beef:beef:0000::/112 creates 65536 services for the cluster
	serviceCIDRMaskv4 = net.CIDRMask(ipV4Bits-maxServicesPo2, ipV4Bits)
	serviceCIDRv4     = net.IPNet{
		IP:   net.ParseIP("172.20.0.0"),
		Mask: serviceCIDRMaskv4,
	}
	serviceCIDRMaskv6 = net.CIDRMask(ipV6Bits-maxServicesPo2, ipV6Bits)
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
	hostIPv4 := hostCIDRv4.IP.To4()
	hostIPv4[2] = byte(hostIdx / octetCap)
	hostIPv4[3] = byte(hostIdx % octetCap)

	hostIPv6 := hostCIDRv6.IP.To16()
	binary.BigEndian.PutUint16(hostIPv6[14:], uint16(hostIdx))

	return hostIPv4.String(), hostIPv6.String()
}

// GetPodIP returns the v4 and v6 IP addresses for the given podIdx
func GetPodIP(idx int64) (string, string) {
	if idx >= maxPods {
		return "", ""
	}
	// This will assign IPs from 10.0.0.0/14 -> 10.0.0.0 - 10.15.255.255
	podIPv4 := podCIDRv4.IP.To4()
	podIPv4[1] = byte((idx / octetCap) / octetCap)
	podIPv4[2] = byte((idx / octetCap) % octetCap)
	podIPv4[3] = byte(idx % octetCap)

	// This will assign IPs from fd01:beef:beef::/110 -> fd01:beef:beef:: - fd01:beef:beef::f:ffff
	v6G := ((idx / quartetCap) << 16) & 0xFFFF0000
	v6H := idx % quartetCap & 0xFFFF

	podIPv6 := podCIDRv6.IP.To16()
	binary.BigEndian.PutUint32(podIPv6[12:], uint32(v6G|v6H))

	return podIPv4.String(), podIPv6.String()
}

// GetPodCIDR returns the v4 and v6 podCIDR for the given hostIdx
func GetPodCIDR(hostIdx int64) (string, string) {
	if hostIdx >= maxHosts {
		return "", ""
	}

	// Since we have 128 pods per host we can divide the hostIdx by 2 because a
	// single octet can "generate" 256 IPs.
	podIPv4 := podCIDRv4.IP.To4()
	podIPv4[1] = byte((hostIdx / 2) / octetCap)
	podIPv4[2] = byte((hostIdx / 2) % octetCap)
	if hostIdx%2 != 0 {
		podIPv4[3] = byte(maxPodsPerHost)
	} else {
		podIPv4[3] = 0
	}
	podIPv4CIDR := net.IPNet{
		IP:   podIPv4,
		Mask: hostPodCIDRMaskv4,
	}

	// Same here, we are dividing by 2 because we only allow 128 Pods per host.
	v6G := (((hostIdx / 2) % quartetCap) << 8) & 0xFFFFFF00
	var v6H int64
	if hostIdx%2 != 0 {
		v6H = maxPodsPerHost & 0x00FF
	}

	podIPv6 := podCIDRv6.IP.To16()
	binary.BigEndian.PutUint32(podIPv6[12:], uint32(v6G|v6H))
	podIPv6CIDR := net.IPNet{
		IP:   podIPv6,
		Mask: hostPodCIDRMaskv6,
	}

	return podIPv4CIDR.String(), podIPv6CIDR.String()
}

// GetHostOfPodIPv4 returns the Host IPv4 address of the that has the given pod
// IP.
func GetHostOfPodIPv4(podIP string) string {
	ip, _, err := net.ParseCIDR(podIP + "/" + strconv.Itoa(ipV4Bits-maxPodsPerHostPo2))
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
	hostIdx := int64(ip[1])*512 + int64(ip[2])*2 + (int64(ip[3]) / maxPodsPerHost)
	v4, _ := GetHostIP(hostIdx)
	return v4
}

// GetHostOfPodIPv6 returns the Host IPv6 address of the that has the given pod
// IP.
func GetHostOfPodIPv6(podIP string) string {
	ip, _, err := net.ParseCIDR(podIP + "/" + strconv.Itoa(ipV6Bits-maxPodsPerHostPo2))
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
	hostIdx := (int64(ip[13])*512)&0xFFFF + ((int64(ip[14])*octetCap + int64(ip[15])) / maxPodsPerHost)
	_, v6 := GetHostIP(hostIdx)
	return v6
}

// GetServiceIP returns the v4 and v6 IP address for the given serviceIdx.
func GetServiceIP(svcIdx int64) (string, string) {
	if svcIdx >= maxServices {
		return "", ""
	}
	svcIPv4 := serviceCIDRv4.IP.To4()
	svcIPv4[2] = byte(svcIdx / octetCap)
	svcIPv4[3] = byte(svcIdx % octetCap)

	svcIPv6 := serviceCIDRv6.IP.To16()
	binary.BigEndian.PutUint16(svcIPv6[14:], uint16(svcIdx))

	return svcIPv4.String(), svcIPv6.String()
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
