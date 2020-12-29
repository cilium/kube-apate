// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2020 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Management",
    "version": "unversioned"
  },
  "paths": {
    "/management/cilium.io/v2/ciliumendpoints": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementCiliumIoV2CiliumEndpoints",
        "parameters": [
          {
            "$ref": "#/parameters/options"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    },
    "/management/cilium.io/v2/ciliumidentities": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementCiliumIoV2CiliumIdentities",
        "parameters": [
          {
            "$ref": "#/parameters/options"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    },
    "/management/cilium.io/v2/ciliumnodes": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementCiliumIoV2CiliumNodes",
        "parameters": [
          {
            "$ref": "#/parameters/options"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    },
    "/management/kubernetes.io/v1/pods": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementKubernetesIoV1Pods",
        "parameters": [
          {
            "$ref": "#/parameters/options"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    }
  },
  "definitions": {
    "add": {
      "description": "add N random new Objects",
      "type": "integer"
    },
    "del": {
      "description": "del N random new Objects",
      "type": "integer"
    },
    "options": {
      "type": "object",
      "properties": {
        "add": {
          "$ref": "#/definitions/add"
        },
        "del": {
          "$ref": "#/definitions/del"
        }
      }
    }
  },
  "parameters": {
    "options": {
      "description": "Options to modify structures",
      "name": "options",
      "in": "body",
      "schema": {
        "$ref": "#/definitions/options"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Management",
    "version": "unversioned"
  },
  "paths": {
    "/management/cilium.io/v2/ciliumendpoints": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementCiliumIoV2CiliumEndpoints",
        "parameters": [
          {
            "description": "Options to modify structures",
            "name": "options",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/options"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    },
    "/management/cilium.io/v2/ciliumidentities": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementCiliumIoV2CiliumIdentities",
        "parameters": [
          {
            "description": "Options to modify structures",
            "name": "options",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/options"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    },
    "/management/cilium.io/v2/ciliumnodes": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementCiliumIoV2CiliumNodes",
        "parameters": [
          {
            "description": "Options to modify structures",
            "name": "options",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/options"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    },
    "/management/kubernetes.io/v1/pods": {
      "post": {
        "schemes": [
          "http"
        ],
        "tags": [
          "cilium"
        ],
        "operationId": "PostManagementKubernetesIoV1Pods",
        "parameters": [
          {
            "description": "Options to modify structures",
            "name": "options",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/options"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "202": {
            "description": "Accepted"
          }
        }
      }
    }
  },
  "definitions": {
    "add": {
      "description": "add N random new Objects",
      "type": "integer"
    },
    "del": {
      "description": "del N random new Objects",
      "type": "integer"
    },
    "options": {
      "type": "object",
      "properties": {
        "add": {
          "$ref": "#/definitions/add"
        },
        "del": {
          "$ref": "#/definitions/del"
        }
      }
    }
  },
  "parameters": {
    "options": {
      "description": "Options to modify structures",
      "name": "options",
      "in": "body",
      "schema": {
        "$ref": "#/definitions/options"
      }
    }
  }
}`))
}
