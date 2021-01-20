// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package swagger

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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/todo": {
            "post": {
                "tags": [
                    "UserTodo"
                ],
                "summary": "添加新的UserTodo",
                "parameters": [
                    {
                        "description": "创建新的todo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.addUserTodoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{Code: 200, Msg: \"请求成功\", Data: \"0\"}",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    },
                    "500": {
                        "description": "{Code: 500, Msg: \"服务器内部错误\", Data: \"服务器内部错误\"}",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/todo/user_todo_id": {
            "delete": {
                "tags": [
                    "UserTodo"
                ],
                "summary": "删除一个UserTodo数据",
                "parameters": [
                    {
                        "description": "删除一个todo",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.findUserTodo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{Code: 200, Msg: \"请求成功\", Data: \"0\"}",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    },
                    "500": {
                        "description": "{Code: 500, Msg: \"服务器内部错误\", Data: \"服务器内部错误\"}",
                        "schema": {
                            "$ref": "#/definitions/response.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.addUserTodoInput": {
            "type": "object",
            "required": [
                "user_todo_title"
            ],
            "properties": {
                "status": {
                    "description": "用户Todo状态",
                    "type": "boolean"
                },
                "user_id": {
                    "description": "用户id(新增)",
                    "type": "integer"
                },
                "user_todo_description": {
                    "description": "用户Todo的描述",
                    "type": "string"
                },
                "user_todo_due_time": {
                    "description": "用户Todo截止时间",
                    "type": "string",
                    "example": "2021-01-22 16:09:00"
                },
                "user_todo_remind_time": {
                    "description": "用户Todo提前多久通知",
                    "type": "string",
                    "example": "2021-01-21 16:09:00"
                },
                "user_todo_title": {
                    "description": "用户Todo的标题",
                    "type": "string"
                }
            }
        },
        "api.findUserTodo": {
            "type": "object",
            "properties": {
                "user_id": {
                    "description": "用户id",
                    "type": "string"
                },
                "user_todo_id": {
                    "description": "用户Todo id",
                    "type": "string"
                }
            }
        },
        "response.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "消息代码",
                    "type": "integer"
                },
                "data": {
                    "description": "消息体",
                    "type": "object"
                },
                "msg": {
                    "description": "消息信息",
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
