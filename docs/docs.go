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
        "contact": {
            "name": "Alexandre VINET",
            "email": "contact@alexandrevinet.fr"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars": {
            "get": {
                "description": "Retrieves the collection of Cars resources.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Retrieves the collection of Cars resources.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.ResponseCollection-entities_Car"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "member": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Car"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Cars resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Creates a Cars resource.",
                "parameters": [
                    {
                        "description": "Car payload",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Car"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.Car"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseViolations"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/cars/{id}": {
            "get": {
                "description": "Retrieves a Cars resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Retrieves a Cars resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cars ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Car"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Replaces the Cars resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Replaces the Cars resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cars ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Car payload",
                        "name": "car",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Car"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseViolations"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes the Cars resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cars"
                ],
                "summary": "Removes the Cars resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cars ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/customers": {
            "get": {
                "description": "Retrieves the collection of Customers resources.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Retrieves the collection of Customers resources.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.ResponseCollection-entities_Customer"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "member": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Customer"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Customers resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Creates a Customers resource.",
                "parameters": [
                    {
                        "description": "Customer payload",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Customer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.Customer"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseViolations"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/customers/{id}": {
            "get": {
                "description": "Retrieves a Customers resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Retrieves a Customers resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customers ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Customer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Replaces the Customers resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Replaces the Customers resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customers ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer payload",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Customer"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseViolations"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes the Customers resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Removes the Customers resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customers ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/schedules": {
            "get": {
                "description": "Retrieves the collection of Schedules resources.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "summary": "Retrieves the collection of Schedules resources.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/httputils.ResponseCollection-entities_Schedule"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "member": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Schedule"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a Schedules resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "summary": "Creates a Schedules resource.",
                "parameters": [
                    {
                        "description": "Schedule payload",
                        "name": "schedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Schedule"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.Schedule"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseViolations"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        },
        "/schedules/{id}": {
            "get": {
                "description": "Retrieves a Schedules resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "summary": "Retrieves a Schedules resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedules ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Schedule"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "description": "Replaces the Schedules resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "summary": "Replaces the Schedules resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedules ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Schedule payload",
                        "name": "schedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ScheduleUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.Schedule"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseViolations"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes the Schedules resource.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "summary": "Removes the Schedules resource.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Schedules ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputils.ResponseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Car": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "dto.Customer": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 255
                }
            }
        },
        "dto.Schedule": {
            "type": "object",
            "required": [
                "beginAt",
                "carId",
                "customerId",
                "endAt"
            ],
            "properties": {
                "beginAt": {
                    "type": "string"
                },
                "carId": {
                    "type": "string"
                },
                "customerId": {
                    "type": "string"
                },
                "endAt": {
                    "type": "string"
                }
            }
        },
        "dto.ScheduleUpdate": {
            "type": "object",
            "required": [
                "beginAt",
                "endAt"
            ],
            "properties": {
                "beginAt": {
                    "type": "string"
                },
                "endAt": {
                    "type": "string"
                }
            }
        },
        "entities.Car": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "schedules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Schedule"
                    }
                }
            }
        },
        "entities.Customer": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "schedules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Schedule"
                    }
                }
            }
        },
        "entities.Schedule": {
            "type": "object",
            "properties": {
                "beginAt": {
                    "type": "string"
                },
                "carId": {
                    "type": "string"
                },
                "customerId": {
                    "type": "string"
                },
                "endAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "httputils.ResponseCollection-entities_Car": {
            "type": "object",
            "properties": {
                "member": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Car"
                    }
                },
                "totalItems": {
                    "type": "integer"
                }
            }
        },
        "httputils.ResponseCollection-entities_Customer": {
            "type": "object",
            "properties": {
                "member": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Customer"
                    }
                },
                "totalItems": {
                    "type": "integer"
                }
            }
        },
        "httputils.ResponseCollection-entities_Schedule": {
            "type": "object",
            "properties": {
                "member": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Schedule"
                    }
                },
                "totalItems": {
                    "type": "integer"
                }
            }
        },
        "httputils.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        },
        "httputils.ResponseViolations": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                },
                "violations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validator.Violation"
                    }
                }
            }
        },
        "validator.Violation": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "propertyPath": {
                    "type": "string"
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
	Title:            "Leasing",
	Description:      "An api to manage leasing locations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
