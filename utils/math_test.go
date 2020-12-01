// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"testing"
)

func TestCont(t *testing.T) {
	type args struct {
		start int64
		end   int64
		limit int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{
				start: 0,
				end:   1,
				limit: 500,
			},
			want: "",
		},
		{
			name: "test-2",
			args: args{
				start: 0,
				end:   499,
				limit: 500,
			},
			want: "",
		},
		{
			name: "test-3",
			args: args{
				start: 0,
				end:   500,
				limit: 500,
			},
			want: "500",
		},
		{
			name: "test-4",
			args: args{
				start: 250,
				end:   750,
				limit: 500,
			},
			want: "750",
		},
		{
			name: "test-5",
			args: args{
				start: 250,
				end:   751,
				limit: 500,
			},
			want: "750",
		},
		{
			name: "test-6",
			args: args{
				start: 0,
				end:   501,
				limit: 500,
			},
			want: "500",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cont(tt.args.start, tt.args.end, tt.args.limit); got != tt.want {
				t.Errorf("Cont() = %v, want %v", got, tt.want)
			}
		})
	}
}
