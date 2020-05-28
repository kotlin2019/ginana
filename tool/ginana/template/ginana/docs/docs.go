// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-05-21 13:23:31.3715995 +0800 CST m=+0.163732701

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
        "contact": {},
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "登陆接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "parameters": [
                    {
                        "description": "User Login",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    }
                }
            }
        },
        "/login/captcha": {
            "get": {
                "description": "获取验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "responses": {
                    "200": {
                        "description": "data:image/png;base64,U3dhZ2dlciByb2Nrcw==",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    }
                }
            }
        },
        "/login/captcha/check": {
            "post": {
                "description": "提前检查验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "parameters": [
                    {
                        "description": "Check Captcha",
                        "name": "captcha",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Captcha"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "安全登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Login"
                ],
                "parameters": [
                    {
                        "description": "User Register",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.JSON"
                        }
                    }
                }
            }
        },
        "/upload/account/avatar": {
            "post": {
                "description": "上传当前账号头像接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Upload"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "上一个资源的地址",
                        "name": "last_src",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "上传类型，本接口只接收type=2",
                        "name": "type",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"success\\\":1,\\\"message\\\":\\\"上传成功\\\",\\\"url\\\":\\\"xxxx/xxxx.jpg\\\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{\\\"success\\\":0,\\\"message\\\":\\\"error\\\",\\\"url\\\":\\\"\\\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Captcha": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code",
                    "type": "string"
                },
                "image": {
                    "description": "验证码图片",
                    "type": "string"
                }
            }
        },
        "model.JSON": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.UserLoginReq": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "loginIP": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UserRegisterReq": {
            "type": "object",
            "properties": {
                "captcha": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "loginIP": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password1": {
                    "type": "string"
                },
                "password2": {
                    "type": "string"
                },
                "username1": {
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
	Version:     "1.0.0",
	Host:        "127.0.0.1:8000",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "GiNana",
	Description: "基于GiNana的个人网站项目，默认端口：8000",
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