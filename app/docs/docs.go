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
        "/product": {
            "get": {
                "description": "Получение списка продуктов с возможным применением фильтров",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product list",
                "parameters": [
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "description": "category id",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    },
                    "404": {
                        "description": "Unexpected error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Получение детальной информации по товару",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product detail",
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "404": {
                        "description": "Unexpected error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "brand": {
                    "type": "string"
                },
                "category": {
                    "$ref": "#/definitions/models.Category"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "preview": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "property": {
                    "$ref": "#/definitions/models.Property"
                }
            }
        },
        "models.Property": {
            "type": "object",
            "properties": {
                "barcode": {
                    "type": "string"
                },
                "color": {
                    "type": "string"
                },
                "height": {
                    "type": "integer"
                },
                "vat": {
                    "type": "string"
                },
                "weight": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://188.225.77.14",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Product API",
	Description:      "This is a service product",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
