// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "haupham404"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/gormdb/adduser": {
            "post": {
                "description": "Add new user to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gormDB"
                ],
                "summary": "AddUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorizationc",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "model.RegisterUser",
                        "name": "token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gormdb/login": {
            "post": {
                "description": "login username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gormDB"
                ],
                "summary": "LoginUser",
                "parameters": [
                    {
                        "description": "model.Login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gormdb/user": {
            "get": {
                "description": "get username from token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gormDB"
                ],
                "summary": "GetUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "GetLogout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "MonggoDB"
                ],
                "summary": "GetLogout",
                "parameters": [
                    {
                        "type": "string",
                        "default": "application/json",
                        "description": "application/json",
                        "name": "Content-Type",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1000",
                        "description": "1000",
                        "name": "Content-Length",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "localhost",
                        "description": "localhost",
                        "name": "Host",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Login": {
            "type": "object",
            "properties": {
                "Password": {
                    "type": "string"
                },
                "UserID": {
                    "type": "string"
                }
            }
        },
        "model.RegisterUser": {
            "type": "object",
            "properties": {
                "Password": {
                    "type": "string"
                },
                "UserID": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "model.Token": {
            "type": "object",
            "properties": {
                "Authorization": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:1909",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
