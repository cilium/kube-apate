// Copyright 2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package version

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	"github.com/cilium/kube-apate/api/k8s/v1/server/restapi/version"
)

type getVersion struct {
}

func NewVersionGetCodeVersionHandler() version.GetCodeVersionHandler {
	return &getVersion{}
}

func (gnp *getVersion) Handle(params version.GetCodeVersionParams) middleware.Responder {
	log.WithFields(logrus.Fields{
		"Method": "GET",
		"URL":    "/version",
	}).Debug("request received")

	return middleware.ResponderFunc(func(w http.ResponseWriter, producer runtime.Producer) {
		f, err := os.Open(filepath.Join("structs", "version.json"))
		defer f.Close()
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, f)
	})
}
