// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package generators

import (
	"testing"
)

func TestGetHostOfPodIPv4(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{
				ipStr: "10.0.0.0",
			},
			want: "172.16.0.0",
		},
		{
			name: "test-2",
			args: args{
				ipStr: "10.0.1.0",
			},
			want: "172.16.0.2",
		},
		{
			name: "test-3",
			args: args{
				ipStr: "10.0.2.0",
			},
			want: "172.16.0.4",
		},
		{
			name: "test-4",
			args: args{
				ipStr: "10.1.0.0",
			},
			want: "172.16.2.0",
		},
		{
			name: "test-5",
			args: args{
				ipStr: "10.15.255.255",
			},
			want: "172.16.31.255",
		},
		{
			name: "test-6",
			args: args{
				ipStr: "10.255.255.244",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHostOfPodIPv4(tt.args.ipStr); got != tt.want {
				t.Errorf("TestGetHostOfPodIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetHostOfPodIPv6(t *testing.T) {
	type args struct {
		ipStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{
				ipStr: "fd01:beef:beef::0:0",
			},
			want: "fd00:beef:beef::0",
		},
		{
			name: "test-2",
			args: args{
				ipStr: "fd01:beef:beef::0:100",
			},
			want: "fd00:beef:beef::2",
		},
		{
			name: "test-3",
			args: args{
				ipStr: "fd01:beef:beef::0:200",
			},
			want: "fd00:beef:beef::4",
		},
		{
			name: "test-4",
			args: args{
				ipStr: "fd01:beef:beef::1:0",
			},
			want: "fd00:beef:beef::200",
		},
		{
			name: "test-5",
			args: args{
				ipStr: "fd01:beef:beef::f:ffff",
			},
			want: "fd00:beef:beef::1fff",
		},
		{
			name: "test-6",
			args: args{
				ipStr: "fd01:beef:beef::ff00:0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHostOfPodIPv6(tt.args.ipStr); got != tt.want {
				t.Errorf("GetHostOfPodIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPodIP(t *testing.T) {
	type args struct {
		idx int64
	}
	tests := []struct {
		name   string
		args   args
		wantV4 string
		wantV6 string
	}{
		{
			name: "test-1",
			args: args{
				idx: 0,
			},
			wantV4: "10.0.0.0",
			wantV6: "fd01:beef:beef::0:0",
		},
		{
			name: "test-2",
			args: args{
				idx: 256 - 1,
			},
			wantV4: "10.0.0.255",
			wantV6: "fd01:beef:beef::0:ff",
		},
		{
			name: "test-3",
			args: args{
				idx: 256,
			},
			wantV4: "10.0.1.0",
			wantV6: "fd01:beef:beef::0:100",
		},
		{
			name: "test-4",
			args: args{
				idx: 257,
			},
			wantV4: "10.0.1.1",
			wantV6: "fd01:beef:beef::0:101",
		},
		{
			name: "test-5",
			args: args{
				idx: 258,
			},
			wantV4: "10.0.1.2",
			wantV6: "fd01:beef:beef::0:102",
		},
		{
			name: "test-6",
			args: args{
				idx: 65536 - 1,
			},
			wantV4: "10.0.255.255",
			wantV6: "fd01:beef:beef::0:ffff",
		},
		{
			name: "test-7",
			args: args{
				idx: 65536,
			},
			wantV4: "10.1.0.0",
			wantV6: "fd01:beef:beef::1:0",
		},
		{
			name: "test-8",
			args: args{
				idx: maxPods - 1,
			},
			wantV4: "10.15.255.255",
			wantV6: "fd01:beef:beef::f:ffff",
		},
		{
			name: "test-9",
			args: args{
				idx: maxPods,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v4, v6 := GetPodIP(tt.args.idx)
			if v4 != tt.wantV4 {
				t.Errorf("GetPodIP() got v4 = %v, want %v", v4, tt.wantV4)
			}
			if v6 != tt.wantV6 {
				t.Errorf("GetPodIP() got v6 = %v, want %v", v6, tt.wantV6)
			}
		})
	}
}

func TestGetPodCIDR(t *testing.T) {
	type args struct {
		hostIdx int64
	}
	tests := []struct {
		name   string
		args   args
		wantV4 string
		wantV6 string
	}{
		{
			name: "test-1",
			args: args{
				hostIdx: 0,
			},
			wantV4: "10.0.0.0/25",
			wantV6: "fd01:beef:beef::0:0/121",
		},
		{
			name: "test-2",
			args: args{
				hostIdx: 1,
			},
			wantV4: "10.0.0.128/25",
			wantV6: "fd01:beef:beef::0:80/121",
		},
		{
			name: "test-3",
			args: args{
				hostIdx: 2,
			},
			wantV4: "10.0.1.0/25",
			wantV6: "fd01:beef:beef::0:100/121",
		},
		{
			name: "test-4",
			args: args{
				hostIdx: 3,
			},
			wantV4: "10.0.1.128/25",
			wantV6: "fd01:beef:beef::0:180/121",
		},
		{
			name: "test-5",
			args: args{
				hostIdx: 512 - 1,
			},
			wantV4: "10.0.255.128/25",
			wantV6: "fd01:beef:beef::0:ff80/121",
		},
		{
			name: "test-6",
			args: args{
				hostIdx: 512,
			},
			wantV4: "10.1.0.0/25",
			wantV6: "fd01:beef:beef::1:0/121",
		},
		{
			name: "test-7",
			args: args{
				hostIdx: maxHosts - 1,
			},
			wantV4: "10.15.255.128/25",
			wantV6: "fd01:beef:beef::f:ff80/121",
		},
		{
			name: "test-8",
			args: args{
				hostIdx: maxHosts,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v4, v6 := GetPodCIDR(tt.args.hostIdx)
			if v4 != tt.wantV4 {
				t.Errorf("GetPodCIDR() got v4 = %v, want %v", v4, tt.wantV4)
			}
			if v6 != tt.wantV6 {
				t.Errorf("GetPodCIDR() got v6 = %v, want %v", v6, tt.wantV6)
			}
		})
	}
}

func TestGetHostIP(t *testing.T) {
	type args struct {
		hostIdx int64
	}
	tests := []struct {
		name   string
		args   args
		wantV4 string
		wantV6 string
	}{
		{
			name: "test-1",
			args: args{
				hostIdx: 0,
			},
			wantV4: "172.16.0.0",
			wantV6: "fd00:beef:beef::0",
		},
		{
			name: "test-2",
			args: args{
				hostIdx: 1,
			},
			wantV4: "172.16.0.1",
			wantV6: "fd00:beef:beef::1",
		},
		{
			name: "test-3",
			args: args{
				hostIdx: 2,
			},
			wantV4: "172.16.0.2",
			wantV6: "fd00:beef:beef::2",
		},
		{
			name: "test-4",
			args: args{
				hostIdx: 30,
			},
			wantV4: "172.16.0.30",
			wantV6: "fd00:beef:beef::1e",
		},
		{
			name: "test-5",
			args: args{
				hostIdx: 512 - 1,
			},
			wantV4: "172.16.1.255",
			wantV6: "fd00:beef:beef::1ff",
		},
		{
			name: "test-6",
			args: args{
				hostIdx: 512,
			},
			wantV4: "172.16.2.0",
			wantV6: "fd00:beef:beef::200",
		},
		{
			name: "test-7",
			args: args{
				hostIdx: 2047,
			},
			wantV4: "172.16.7.255",
			wantV6: "fd00:beef:beef::7ff",
		},
		{
			name: "test-8",
			args: args{
				hostIdx: maxHosts - 1,
			},
			wantV4: "172.16.31.255",
			wantV6: "fd00:beef:beef::1fff",
		},
		{
			name: "test-9",
			args: args{
				hostIdx: maxHosts,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v4, v6 := GetHostIP(tt.args.hostIdx)
			if v4 != tt.wantV4 {
				t.Errorf("GetHostIP() got v4 = %v, want %v", v4, tt.wantV4)
			}
			if v6 != tt.wantV6 {
				t.Errorf("GetHostIP() got v6 = %v, want %v", v6, tt.wantV6)
			}
		})
	}
}

func TestGetServiceIP(t *testing.T) {
	type args struct {
		svcIdx int64
	}
	tests := []struct {
		name   string
		args   args
		wantV4 string
		wantV6 string
	}{
		{
			name: "test-1",
			args: args{
				svcIdx: 0,
			},
			wantV4: "172.20.0.0",
			wantV6: "fd02:beef:beef::0",
		},
		{
			name: "test-2",
			args: args{
				svcIdx: 1,
			},
			wantV4: "172.20.0.1",
			wantV6: "fd02:beef:beef::1",
		},
		{
			name: "test-3",
			args: args{
				svcIdx: 2,
			},
			wantV4: "172.20.0.2",
			wantV6: "fd02:beef:beef::2",
		},
		{
			name: "test-4",
			args: args{
				svcIdx: 30,
			},
			wantV4: "172.20.0.30",
			wantV6: "fd02:beef:beef::1e",
		},
		{
			name: "test-5",
			args: args{
				svcIdx: 512 - 1,
			},
			wantV4: "172.20.1.255",
			wantV6: "fd02:beef:beef::1ff",
		},
		{
			name: "test-6",
			args: args{
				svcIdx: 512,
			},
			wantV4: "172.20.2.0",
			wantV6: "fd02:beef:beef::200",
		},
		{
			name: "test-7",
			args: args{
				svcIdx: 2047,
			},
			wantV4: "172.20.7.255",
			wantV6: "fd02:beef:beef::7ff",
		},
		{
			name: "test-8",
			args: args{
				svcIdx: maxServices - 1,
			},
			wantV4: "172.20.255.255",
			wantV6: "fd02:beef:beef::ffff",
		},
		{
			name: "test-9",
			args: args{
				svcIdx: maxServices,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v4, v6 := GetServiceIP(tt.args.svcIdx)
			if v4 != tt.wantV4 {
				t.Errorf("GetServiceIP() got v4 = %v, want %v", v4, tt.wantV4)
			}
			if v6 != tt.wantV6 {
				t.Errorf("GetServiceIP() got v6 = %v, want %v", v6, tt.wantV6)
			}
		})
	}
}
