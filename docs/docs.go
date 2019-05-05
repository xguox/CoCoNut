// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-05-05 17:19:37.965963 +0800 CST m=+0.053235168

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/api/v1/admin/categories": {
            "post": {
                "description": "添加新的商品分类",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "category"
                ],
                "summary": "添加新的商品分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "认证 Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建分类请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.CategoryValidator"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{msg:\"请求处理成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{msg:\"请求参数有误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{msg:\"服务器错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/login": {
            "post": {
                "description": "后台账号登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "后台账号登录",
                "parameters": [
                    {
                        "description": "账号登录请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.LoginValidator"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{msg:\"请求处理成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "{msg:\"账号或密码有误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{msg:\"请求参数有误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{msg:\"服务器错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/admin/users": {
            "post": {
                "description": "添加新的 User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "添加新的 User",
                "parameters": [
                    {
                        "description": "创建User请求参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserValidator"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{msg:\"请求处理成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "{msg:\"请求参数有误\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "{msg:\"服务器错误\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CategoryValidator": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "object",
                    "required": [
                        "name",
                        "slug"
                    ],
                    "properties": {
                        "name": {
                            "type": "string"
                        },
                        "slug": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "model.LoginValidator": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "object",
                    "required": [
                        "email",
                        "password"
                    ],
                    "properties": {
                        "email": {
                            "type": "string"
                        },
                        "password": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "model.UserValidator": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "object",
                    "required": [
                        "username",
                        "email",
                        "password"
                    ],
                    "properties": {
                        "email": {
                            "type": "string"
                        },
                        "password": {
                            "type": "string"
                        },
                        "username": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
