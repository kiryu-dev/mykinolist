// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/auth/signin": {
            "post": {
                "description": "Log in to an existing account, update access and refresh tokens",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in to account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Tokens"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/signout": {
            "post": {
                "description": "Sign out of account by deleting refresh token from cookies",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign out of account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Create an account and an empty movie list, which is linked to the account ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up an account",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignUpUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            }
        },
        "/list": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Get all movies from list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Get movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.ListUnit"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Use a third-party API to search for movie information by title and, if successful, add the movie to the list. You can add the movie to your favorites, rate it, and specify movie status (watching, plan to watch, etc.)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Add movie to list",
                "parameters": [
                    {
                        "description": "movie info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ListUnit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            }
        },
        "/list/{id}": {
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Delete movie from list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Delete movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ListUnit"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Update part of the information about the added movie. For example, you can add the movie to your favorites or change the rating of the movie.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "list"
                ],
                "summary": "Update movie info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
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
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Get user info: id, username, email, last login time, etc.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Delete user account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/controller.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.errorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.ListInfo": {
            "type": "object",
            "properties": {
                "list_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.ListUnit": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "is_favorite": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "model.SignInUserDTO": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.SignUpUserDTO": {
            "type": "object",
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
        },
        "model.Tokens": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "created_on": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_login": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "MyKinoList API",
	Description:      "MyKinoList is an API that provides the ability to create accounts to keep a list of movies you are watching, have watched, or may want to watch sometime later, as well as to rate them and add them to your «favorites».\\n\\nThis API uses JSONWebTokens to give access to authorized users. A third-party, unofficial Kinopoisk API is used to retrieve information about movies.\\n\\nThe documentation uses an access token for authorization, which only lasts 30 seconds.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
