// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/index": {
            "get": {
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "tags": [
                    "创建用户"
                ],
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "tags": [
                    "删除用户"
                ],
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserBasic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "tags": [
                    "获取用户列表"
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.UserBasic": {
            "type": "object"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
