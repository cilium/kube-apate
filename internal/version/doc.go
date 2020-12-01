// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package version

import (
	logging "github.com/cilium/kube-apate/internal/log"
	"github.com/cilium/kube-apate/internal/log/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "version")
