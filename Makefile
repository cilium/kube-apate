# Copyright 2017-2020 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

CONTAINER_ENGINE?=docker
SWAGGER_VERSION := v0.25.0
SWAGGER := $(CONTAINER_ENGINE) run -u $(shell id -u):$(shell id -g) --rm -v $(CURDIR):$(CURDIR) -w $(CURDIR) --entrypoint swagger quay.io/goswagger/swagger:$(SWAGGER_VERSION)

kube-apate: main.go
	go build -o $@ $^


generate-api: api/k8s/v1/swagger.json
	-$(QUIET)$(SWAGGER) generate server -s server -a restapi \
		-t api/k8s/v1 \
		-f $^ \
		--default-scheme=http \
		-C api/k8s/v1/cilium-server.yml \
		-r tools/spdx-copyright-header.txt
	-$(QUIET)$(SWAGGER) generate client -a restapi \
		-t api/k8s/v1 \
		-f $^ \
		-r tools/spdx-copyright-header.txt

generate-management-api: api/management/v1/management-swagger.yaml
	-$(QUIET)$(SWAGGER) generate server -s server -a restapi \
		-t api/management/v1 \
		-f $^ \
		--default-scheme=http \
		-C api/management/v1/management-server.yml \
		-r tools/spdx-copyright-header.txt
	-$(QUIET)$(SWAGGER) generate client -a restapi \
		-t api/management/v1 \
		-f $^ \
		-r tools/spdx-copyright-header.txt

sort-imports:
	@# sort goimports automatically
	-$(QUIET) find api/k8s/v1/client/ -type f -name "*.go" -print | PATH="$(PWD)/tools:$(PATH)" xargs goimports -w
	-$(QUIET) find api/k8s/v1/models/ -type f -name "*.go" -print | PATH="$(PWD)/tools:$(PATH)" xargs goimports -w
	-$(QUIET) find api/k8s/v1/server/ -type f -name "*.go" -print | PATH="$(PWD)/tools:$(PATH)" xargs goimports -w
	-$(QUIET) find api/management/v1/models/ -type f -name "*.go" -print | PATH="$(PWD)/tools:$(PATH)" xargs goimports -w
	-$(QUIET) find api/management/v1/server/ -type f -name "*.go" -print | PATH="$(PWD)/tools:$(PATH)" xargs goimports -w

api/v1/swagger.json: api/v1/k8s-swagger.json api/v1/cilium-swagger.json
	-$(QUIET)$(SWAGGER) mixin api/v1/k8s-swagger.json api/v1/cilium-swagger.json \
       -o api/v1/swagger.json

.PHONY: kube-apate
