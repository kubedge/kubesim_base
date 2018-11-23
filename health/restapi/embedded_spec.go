// Code generated by go-swagger; DO NOT EDIT.

package restapi

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
  "swagger": "2.0",
  "info": {
    "title": "KUBEDGE ConfigAPI Server",
    "version": "1.0.0"
  },
  "paths": {
    "/config": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getConfig",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to all if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a config file",
            "schema": {
              "description": "contains the actual config as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getGreeting",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains the actual greeting as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/liveness": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "probeLiveness",
        "responses": {
          "200": {
            "description": "returns a liveness state of the simulator",
            "schema": {
              "description": "contains the actual liveness state as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/readiness": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "probeReadiness",
        "responses": {
          "200": {
            "description": "returns a readiness state of the simulator",
            "schema": {
              "description": "contains the actual readiness state as plain text",
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "KUBEDGE ConfigAPI Server",
    "version": "1.0.0"
  },
  "paths": {
    "/config": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getConfig",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to all if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a config file",
            "schema": {
              "description": "contains the actual config as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "getGreeting",
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains the actual greeting as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/liveness": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "probeLiveness",
        "responses": {
          "200": {
            "description": "returns a liveness state of the simulator",
            "schema": {
              "description": "contains the actual liveness state as plain text",
              "type": "string"
            }
          }
        }
      }
    },
    "/readiness": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "operationId": "probeReadiness",
        "responses": {
          "200": {
            "description": "returns a readiness state of the simulator",
            "schema": {
              "description": "contains the actual readiness state as plain text",
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
}
