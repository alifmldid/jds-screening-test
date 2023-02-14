// Code generated by swaggo/swag. DO NOT EDIT
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
        "/aggregate": {
            "get": {
                "description": "Aggregate Data from https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product  ) aggregation by department, product, price IDR and order ascending by price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Aggregate Data"
                ],
                "summary": "fetch data endpoint",
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
                            "$ref": "#/definitions/item.AggregateDataResponse"
                        }
                    }
                }
            }
        },
        "/fetch": {
            "get": {
                "description": "Fetch data from https://60c18de74f7e880017dbfd51.mockapi.io/api/v1/jabar-digital-services/product",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Fetch Data"
                ],
                "summary": "fetch data endpoint",
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
                            "$ref": "#/definitions/item.FetchDataResponse"
                        }
                    }
                }
            }
        },
        "/verify": {
            "get": {
                "description": "Display token private claims data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Verify Token"
                ],
                "summary": "token verify endpoint",
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
                            "$ref": "#/definitions/item.VerifyTokenResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "item.AggregateDataResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/item.ItemAgg"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "item.FetchDataResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/item.Item"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "item.Item": {
            "type": "object",
            "properties": {
                "IDR": {
                    "type": "number",
                    "example": 3317529
                },
                "createdAt": {
                    "type": "string",
                    "example": "2021-06-09T09:37:05.527Z"
                },
                "department": {
                    "type": "string",
                    "example": "Outdoors"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "price": {
                    "type": "number",
                    "example": 218
                },
                "product": {
                    "type": "string",
                    "example": "Salad"
                }
            }
        },
        "item.ItemAgg": {
            "type": "object",
            "properties": {
                "department": {
                    "type": "string",
                    "example": "Games"
                },
                "price": {
                    "type": "number",
                    "example": 60872.094
                },
                "product": {
                    "type": "string",
                    "example": "Computer"
                }
            }
        },
        "item.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 2
                },
                "nik": {
                    "type": "string",
                    "example": "3574526883729838"
                },
                "role": {
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "item.VerifyTokenResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/item.User"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:9000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Fetch App Service",
	Description:      "A fetch app service with 3 endpoint: fetch, aggregate, token verify",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
