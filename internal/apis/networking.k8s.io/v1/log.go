// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	logging "github.com/cilium/kube-apate/internal/log"
	"github.com/cilium/kube-apate/internal/log/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "v1/networking.k8s.io")
