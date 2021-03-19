// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Saidamir Botirov",
            "url": "https://www.linkedin.com/in/saidamir-botirov-a08559192",
            "email": "saidamir.botirov@gmail.com"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/config": {
            "get": {
                "description": "shows config of the project only on the development phase",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "gets project config",
                "operationId": "get-config",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/config.Config"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "this returns \"pong\" messsage to show service is working",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "returns \"pong\" message",
                "operationId": "ping",
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/has-access": {
            "post": {
                "description": "has access",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "has access",
                "operationId": "has-access",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform Id",
                        "name": "platform-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Access Info",
                        "name": "access",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.AccessModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.HasAccessModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/logout": {
            "delete": {
                "description": "logout user by his/her token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "logout user",
                "operationId": "logout",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform Id",
                        "name": "platform-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/refresh": {
            "put": {
                "description": "refresh user token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "refresh user token",
                "operationId": "refresh-token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform Id",
                        "name": "platform-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Token Info",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.RefreshTokenModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.TokenModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/standard/login": {
            "post": {
                "description": "standard login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "standard login",
                "operationId": "standard-login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Platform Id",
                        "name": "platform-id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.StandardLoginModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.TokenModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/user/register": {
            "post": {
                "description": "register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "register user",
                "operationId": "register-user",
                "parameters": [
                    {
                        "description": "body",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.RegisterUserModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.ResponseModel"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Config": {
            "type": "object",
            "properties": {
                "app": {
                    "type": "string"
                },
                "basePath": {
                    "type": "string"
                },
                "cassandraHost": {
                    "type": "string"
                },
                "cassandraKeyspace": {
                    "type": "string"
                },
                "cassandraPassword": {
                    "type": "string"
                },
                "cassandraPort": {
                    "type": "integer"
                },
                "cassandraUser": {
                    "type": "string"
                },
                "defaultLimit": {
                    "type": "string"
                },
                "defaultOffset": {
                    "type": "string"
                },
                "environment": {
                    "description": "development, staging, production",
                    "type": "string"
                },
                "httpport": {
                    "type": "string"
                },
                "logLevel": {
                    "description": "debug, info, warn, error, dpanic, panic, fatal",
                    "type": "string"
                },
                "postgresDatabase": {
                    "type": "string"
                },
                "postgresHost": {
                    "type": "string"
                },
                "postgresPassword": {
                    "type": "string"
                },
                "postgresPort": {
                    "type": "integer"
                },
                "postgresUser": {
                    "type": "string"
                },
                "rabbitURI": {
                    "type": "string"
                },
                "secretKey": {
                    "type": "string"
                },
                "serviceHost": {
                    "type": "string"
                }
            }
        },
        "rest.AccessModel": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "rest.HasAccessModel": {
            "type": "object",
            "properties": {
                "client_platform_id": {
                    "type": "string"
                },
                "client_type_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "role_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "rest.RefreshTokenModel": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "rest.RegisterUserModel": {
            "type": "object",
            "properties": {
                "client_type_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phones": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role_id": {
                    "type": "string"
                }
            }
        },
        "rest.ResponseModel": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "error": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "rest.StandardLoginModel": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "rest.TokenModel": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "refresh_in_seconds": {
                    "type": "integer"
                },
                "refresh_token": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Go Boilerplate API",
	Description: "This is a Go Boilerplate for medium sized projects",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
