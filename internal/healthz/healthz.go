// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package healthz

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type getHealthz struct {
	h http.Handler
}

func NewGetHealthzHandler(h http.Handler) http.Handler {
	return &getHealthz{h}
}

func (gnp *getHealthz) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{
		"Method": r.URL.Path,
		"URL":    r.Method,
	}).Debug("request IN")
	if r.URL.Path == "/healthz" {
		log.WithField("", "").Debug("GET /healthz request")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK\n"))
		return
	}
	gnp.h.ServeHTTP(w, r)
}
