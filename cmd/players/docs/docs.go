// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "erupshis",
            "email": "e.rupshis@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/player": {
            "get": {
                "description": "provides player data by playerID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "player"
                ],
                "summary": "select player by ID",
                "operationId": "player-select",
                "parameters": [
                    {
                        "type": "string",
                        "description": "player search by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserData"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/withdraw": {
            "patch": {
                "description": "withdraw currency from player balance",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "player"
                ],
                "summary": "withdraw currency from player account by ID",
                "operationId": "player-withdraw",
                "parameters": [
                    {
                        "type": "string",
                        "description": "player search by id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "currency amount",
                        "name": "amount",
                        "in": "query",
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
                    "204": {
                        "description": "user not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "incorrect query in request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "405": {
                        "description": "insufficient funds",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "unexpected marshalling error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.UserData": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Players service Swagger API",
	Description:      "Swagger API for players storage.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
