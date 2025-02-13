// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/devices": {
            "get": {
                "description": "Retrieves a list of all devices",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get all devices",
                "responses": {
                    "200": {
                        "description": "List of devices",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Device"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new device to the inventory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Create a new device",
                "parameters": [
                    {
                        "description": "Device details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created device",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/devices/brand/{brand}": {
            "get": {
                "description": "Retrieves a single device by its brand",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get a device by brand",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Device Brand",
                        "name": "brand",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Device details",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/devices/state/{state}": {
            "get": {
                "description": "Retrieves a single device by its state",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get a device by state",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Device State",
                        "name": "brand",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Device details",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/v1/devices/{id}": {
            "get": {
                "description": "Retrieves a single device by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Get a device by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Device details",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "404": {
                        "description": "Device not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Updates device details if allowed. ` + "`" + `PUT` + "`" + ` requires a full update, while ` + "`" + `PATCH` + "`" + ` allows partial updates.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Update an existing device",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Device update details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Update"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated device",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden update",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Device not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes a device from the inventory",
                "tags": [
                    "Device"
                ],
                "summary": "Delete a device",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content"
                    },
                    "403": {
                        "description": "Cannot delete device in use",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Device not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Allows partial updates to a device. Only provided fields are modified.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device"
                ],
                "summary": "Partially update an existing device",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Partial device update details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Patch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated device",
                        "schema": {
                            "$ref": "#/definitions/domain.Device"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden update",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Device not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Device": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "creation_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/domain.State"
                }
            }
        },
        "domain.Patch": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "brand": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/domain.State"
                }
            }
        },
        "domain.State": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "AvailableState",
                "InUseState",
                "InactiveState"
            ]
        },
        "domain.Update": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "brand": {
                    "type": "string"
                },
                "creation_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/domain.State"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
