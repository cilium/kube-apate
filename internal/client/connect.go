// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	runtime_client "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	clientapi "github.com/cilium/kube-apate/api/management/v1/client"
)

const (
	ClientTimeout = 90 * time.Second
)

type Client struct {
	clientapi.Management
}

func NewClient(host string) (*Client, error) {
	tmp := strings.SplitN(host, "://", 2)
	if len(tmp) != 2 {
		return nil, fmt.Errorf("invalid host format '%s'", host)
	}

	switch tmp[0] {
	case "tcp":
		if _, err := url.Parse("tcp://" + tmp[1]); err != nil {
			return nil, err
		}
		host = "http://" + tmp[1]
	}

	clientTrans := runtime_client.New(
		tmp[1], clientapi.DefaultBasePath,
		clientapi.DefaultSchemes)
	return &Client{*clientapi.New(clientTrans, strfmt.Default)}, nil
}
