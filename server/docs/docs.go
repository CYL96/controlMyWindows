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
        "/api/AddControlClass": {
            "post": {
                "description": "添加一个控制模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "添加一个控制模块",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.AddControlClassReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.AddControlClassResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/AddControlDetail": {
            "post": {
                "description": "控制模块-键",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "添加一个控制模块-键",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.AddControlDetailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.AddControlDetailResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/DeleteControlClass": {
            "post": {
                "description": "删除控制模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "删除控制模块",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.DeleteControlClassReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.DeleteControlClassResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/DeleteControlDetail": {
            "post": {
                "description": "删除控制模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "删除控制模块",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.DeleteControlDetailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.DeleteControlDetailResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/ExecControlDetail": {
            "post": {
                "description": "执行key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "执行key",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.ExecControlDetailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.ExecControlDetailResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/GetControlClassInfo": {
            "post": {
                "description": "获取控制模块详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "获取控制模块详情",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.GetControlClassInfoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.GetControlClassInfoResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/GetControlClassList": {
            "post": {
                "description": "获取控制模块列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "获取控制模块列表",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.GetControlClassListReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.GetControlClassListResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/GetControlDetailList": {
            "post": {
                "description": "获取键列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "获取键列表",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.GetControlDetailListReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.GetControlDetailListResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/StopControlDetail": {
            "post": {
                "description": "执行key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "停止执行key",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.StopControlDetailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.StopControlDetailResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/UpdateControlClass": {
            "post": {
                "description": "更新控制模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "更新控制模块",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.UpdateControlClassReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.UpdateControlClassResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/UpdateControlClassOrder": {
            "post": {
                "description": "更新控制顺序",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "更新控制顺序",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.UpdateControlClassOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.UpdateControlClassOrderResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/UpdateControlDetail": {
            "post": {
                "description": "更新控制模块",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "更新控制模块",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.UpdateControlDetailReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.UpdateControlDetailResp"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/UpdateControlDetailOrder": {
            "post": {
                "description": "更新控制顺序",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "控制模块-键"
                ],
                "summary": "更新控制顺序",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hd.UpdateControlDetailOrderReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "desc",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/hd.GinResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/hd.UpdateControlDetailOrderResp"
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
        "config.ControlDetailExt": {
            "type": "object",
            "properties": {
                "control_type": {
                    "description": "1:快捷键，2：脚本 3：打开文件夹目录  4:打开网页",
                    "default": 0,
                    "allOf": [
                        {
                            "$ref": "#/definitions/control.ControlType"
                        }
                    ],
                    "example": 0
                },
                "detail_color": {
                    "type": "string",
                    "example": ""
                },
                "detail_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail_key": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/control.ControlKeyList"
                    }
                },
                "detail_name": {
                    "type": "string",
                    "example": ""
                },
                "run_state": {
                    "description": "1:空闲 2：运行中",
                    "default": 0,
                    "allOf": [
                        {
                            "$ref": "#/definitions/config.RunState"
                        }
                    ],
                    "example": 0
                }
            }
        },
        "config.ControlDetailIdExt": {
            "type": "object",
            "properties": {
                "detail_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "config.ControlListExt": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "control_name": {
                    "type": "string",
                    "example": ""
                },
                "control_order": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ControlDetailExt"
                    }
                },
                "key_height": {
                    "type": "string"
                },
                "key_width": {
                    "type": "string"
                }
            }
        },
        "config.ControlListIdT": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "config.RunState": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-comments": {
                "RunStateFree": "空闲",
                "RunStateRunning": "运行中"
            },
            "x-enum-varnames": [
                "RunStateFree",
                "RunStateRunning"
            ]
        },
        "control.ControlKeyList": {
            "type": "object",
            "properties": {
                "delay": {
                    "description": "当KeyType == 99时 使用 ms",
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "input": {
                    "description": "当KeyType == 2时 输入文本",
                    "type": "string",
                    "example": ""
                },
                "key": {
                    "type": "string",
                    "example": ""
                },
                "key_list": {
                    "description": "当KeyType == 3时",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/control.KeyListT"
                    }
                },
                "key_press": {
                    "description": "当KeyType == 1时 1：单击 2：双击 3：按下 4：抬起",
                    "default": 0,
                    "allOf": [
                        {
                            "$ref": "#/definitions/control.PressType"
                        }
                    ],
                    "example": 0
                },
                "key_type": {
                    "description": "1 ：单键 2 ：文本 3 ：快捷键 99：延迟",
                    "default": 0,
                    "allOf": [
                        {
                            "$ref": "#/definitions/control.KeyType"
                        }
                    ],
                    "example": 0
                }
            }
        },
        "control.ControlType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-comments": {
                "ControlTypeExplorer": "3：打开文件夹目录",
                "ControlTypeNormal": "1:快捷键",
                "ControlTypeScript": "2：脚本",
                "ControlTypeWebsite": "4:打开网页"
            },
            "x-enum-varnames": [
                "ControlTypeNormal",
                "ControlTypeScript",
                "ControlTypeExplorer",
                "ControlTypeWebsite"
            ]
        },
        "control.KeyListT": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "key": {
                    "type": "string",
                    "example": ""
                }
            }
        },
        "control.KeyType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                99
            ],
            "x-enum-comments": {
                "KeyTypeDefault": "单键",
                "KeyTypeDelay": "延迟",
                "KeyTypeShortcutKey": "快捷键",
                "KeyTypeText": "文本"
            },
            "x-enum-varnames": [
                "KeyTypeDefault",
                "KeyTypeText",
                "KeyTypeShortcutKey",
                "KeyTypeDelay"
            ]
        },
        "control.PressType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-comments": {
                "PressTypeClick": "单击",
                "PressTypeDoubleClick": "双击",
                "PressTypePressDown": "按下",
                "PressTypePressUp": "抬起"
            },
            "x-enum-varnames": [
                "PressTypeClick",
                "PressTypeDoubleClick",
                "PressTypePressDown",
                "PressTypePressUp"
            ]
        },
        "hd.AddControlClassReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "control_name": {
                    "type": "string",
                    "example": ""
                },
                "control_order": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "key_height": {
                    "type": "string"
                },
                "key_width": {
                    "type": "string"
                }
            }
        },
        "hd.AddControlClassResp": {
            "type": "object"
        },
        "hd.AddControlDetailReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail": {
                    "$ref": "#/definitions/config.ControlDetailExt"
                }
            }
        },
        "hd.AddControlDetailResp": {
            "type": "object"
        },
        "hd.DeleteControlClassReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "hd.DeleteControlClassResp": {
            "type": "object"
        },
        "hd.DeleteControlDetailReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "hd.DeleteControlDetailResp": {
            "type": "object"
        },
        "hd.ExecControlDetailReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "hd.ExecControlDetailResp": {
            "type": "object"
        },
        "hd.GetControlClassInfoReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "hd.GetControlClassInfoResp": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "control_name": {
                    "type": "string",
                    "example": ""
                },
                "control_order": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "key_height": {
                    "type": "string"
                },
                "key_width": {
                    "type": "string"
                }
            }
        },
        "hd.GetControlClassListReq": {
            "type": "object"
        },
        "hd.GetControlClassListResp": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ControlListExt"
                    }
                }
            }
        },
        "hd.GetControlDetailListReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "hd.GetControlDetailListResp": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ControlDetailExt"
                    }
                }
            }
        },
        "hd.GinResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                },
                "state": {
                    "$ref": "#/definitions/hd.STATE_CODE"
                }
            }
        },
        "hd.STATE_CODE": {
            "type": "integer",
            "enum": [
                9999,
                0,
                1
            ],
            "x-enum-comments": {
                "StateFailed": "失败",
                "StateOk": "成功",
                "StatePanic": "服务器崩溃"
            },
            "x-enum-varnames": [
                "StatePanic",
                "StateOk",
                "StateFailed"
            ]
        },
        "hd.StopControlDetailReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                }
            }
        },
        "hd.StopControlDetailResp": {
            "type": "object"
        },
        "hd.UpdateControlClassOrderReq": {
            "type": "object",
            "properties": {
                "order_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ControlListIdT"
                    }
                }
            }
        },
        "hd.UpdateControlClassOrderResp": {
            "type": "object"
        },
        "hd.UpdateControlClassReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "control_name": {
                    "type": "string",
                    "example": ""
                },
                "control_order": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "key_height": {
                    "type": "string"
                },
                "key_width": {
                    "type": "string"
                }
            }
        },
        "hd.UpdateControlClassResp": {
            "type": "object"
        },
        "hd.UpdateControlDetailOrderReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/config.ControlDetailIdExt"
                    }
                }
            }
        },
        "hd.UpdateControlDetailOrderResp": {
            "type": "object"
        },
        "hd.UpdateControlDetailReq": {
            "type": "object",
            "properties": {
                "control_id": {
                    "type": "integer",
                    "default": 0,
                    "example": 0
                },
                "detail": {
                    "$ref": "#/definitions/config.ControlDetailExt"
                }
            }
        },
        "hd.UpdateControlDetailResp": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
