// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-07-12 13:44:52.532804671 +0300 EEST m=+3.044256576
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://swagger.io/terms/",
        "contact": {
            "name": "Developer",
            "url": "https://github.com/illiafox",
            "email": "illiadimura@gmail.com"
        },
        "license": {
            "name": "Boost Software License 1.0",
            "url": "https://opensource.org/licenses/BSL-1.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/block": {
            "post": {
                "description": "Block balance by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Block user balance",
                "parameters": [
                    {
                        "description": "User ID and Reason",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BlockIN"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BlockOUT"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    }
                }
            }
        },
        "/admin/unblock": {
            "post": {
                "description": "Unblock balance by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Unblock user balance",
                "parameters": [
                    {
                        "description": "User ID",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UnblockIN"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UnblockOUT"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    }
                }
            }
        },
        "/user/change": {
            "patch": {
                "description": "Change balance by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Change user balance",
                "parameters": [
                    {
                        "description": "User ID, Change amount and Description",
                        "name": "input",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/dto.ChangeBalanceIN"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ChangeBalanceOUT"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    }
                }
            }
        },
        "/user/transfer": {
            "post": {
                "description": "Transfer money from one balance to another",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Transfer money between users",
                "parameters": [
                    {
                        "description": "To and From ID, Amount and Description",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TransferBalanceIN"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TransferBalanceOUT"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get balance by User ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "balance"
                ],
                "summary": "Get user balance",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minLength": 3,
                        "type": "string",
                        "description": "currency abbreviation",
                        "name": "base",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Balance data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.GetBalanceOUT"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "balance": {
                                            "type": "integer"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    }
                }
            }
        },
        "/user/{id}/transactions": {
            "get": {
                "description": "View transactions with sorting and pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "View user transactions",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "maximum": 100,
                        "minimum": 0,
                        "type": "integer",
                        "default": 100,
                        "description": "output limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "output offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "DATE_DESC",
                            "DATE_ASC",
                            "SUM_DESC",
                            "SUM_ASC"
                        ],
                        "type": "string",
                        "description": "sort type",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transactions data",
                        "schema": {
                            "$ref": "#/definitions/dto.ViewTransactionsOUT"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BlockIN": {
            "type": "object",
            "properties": {
                "reason": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.BlockOUT": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "dto.ChangeBalanceIN": {
            "type": "object",
            "required": [
                "description"
            ],
            "properties": {
                "change": {
                    "type": "integer"
                },
                "description": {
                    "description": "|min_len:10",
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.ChangeBalanceOUT": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "dto.GetBalanceOUT": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "string"
                },
                "base": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "dto.TransferBalanceIN": {
            "type": "object",
            "required": [
                "description"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "description": "|min_len:10",
                    "type": "string"
                },
                "from_id": {
                    "type": "integer"
                },
                "to_id": {
                    "type": "integer"
                }
            }
        },
        "dto.TransferBalanceOUT": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "dto.UnblockIN": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.UnblockOUT": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                }
            }
        },
        "dto.ViewTransactionsOUT": {
            "type": "object",
            "properties": {
                "ok": {
                    "type": "boolean"
                },
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Transaction"
                    }
                }
            }
        },
        "entity.Transaction": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "from_id": {
                    "description": "zero -\u003e null -\u003e received from other service",
                    "type": "integer"
                },
                "to_id": {
                    "description": "to_id",
                    "type": "integer"
                },
                "transaction_id": {
                    "type": "integer"
                }
            }
        },
        "httputils.Error": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                },
                "ok": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "0.0.0.0:8080",
	BasePath:         "/api",
	Schemes:          []string{"http", "https"},
	Title:            "Balance API",
	Description:      "Balance API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
