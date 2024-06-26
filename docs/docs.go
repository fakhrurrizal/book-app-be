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
        "/v1/book": {
            "get": {
                "description": "Get Book with Pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get Book with Pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search (string)",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page (int)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit (int)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "sort (id or publication_year or title)",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "order (asc or desc)",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "status (status)",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "category_id (int)",
                        "name": "category_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Create New Book",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create Book",
                "parameters": [
                    {
                        "description": "Create body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqres.BookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/book-category": {
            "get": {
                "description": "Get category Book with Pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Get Category Book  with Pagination",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search (string)",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page (int)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit (int)",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "status (status)",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Create New Book Category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Create Book Category",
                "parameters": [
                    {
                        "description": "Create body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqres.BookCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/book-category/{id}": {
            "put": {
                "description": "Update Single Book Category by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Update Single Book Category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqres.BookCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete Single Book Category by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Delete Single Book Category by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/book/{id}": {
            "get": {
                "description": "Get Single Book",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get Single Book",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "description": "Update Single Book by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Update Single Book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reqres.BookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete Single Book by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Delete Single Book by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/file": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    },
                    {
                        "JwtToken": []
                    }
                ],
                "description": "File Uploader",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "File"
                ],
                "summary": "File Uploader",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "reqres.BookCategoryRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "reqres.BookRequest": {
            "type": "object",
            "required": [
                "category_id",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string"
                },
                "book_code": {
                    "type": "string"
                },
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "number_of_pages": {
                    "type": "integer"
                },
                "publication_year": {
                    "type": "integer"
                },
                "publisher": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                },
                "title": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "BOOK APP",
	Description:      "API documentation by BOOK APP",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
