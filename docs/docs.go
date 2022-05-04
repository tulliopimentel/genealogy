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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "url": "LinkedIn    \thttps://www.linkedin.com/in/tulliopimentelbarbosa/"
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
        "/v1/person": {
            "get": {
                "description": "Get all persons of the database, even if is not related to each other",
                "tags": [
                    "Show all persons"
                ],
                "summary": "Get all persons of the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Person"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Insert a new person into the database, and you can create a new family tree with this person",
                "tags": [
                    "Create person"
                ],
                "summary": "CreatePerson() Create a new person",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Person"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/person/family": {
            "post": {
                "description": "In this endpoint you can create a family, that is, put the ID of the reference person, another person and the relationship between the two in the database. All by ID",
                "tags": [
                    "Create a new family"
                ],
                "summary": "Creates the relationship between two people to create a family and family tree",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Family"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/v1/person/genealogy/{id}": {
            "get": {
                "description": "Returns the closest family members to the reference person, without delving into the family tree of the other members",
                "tags": [
                    "Show Family Tree"
                ],
                "summary": "ShowFamily() Returns the closest family members to the reference person",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Genealogy"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "This table has just two columns in database, that means you will just put the ID of the referenced person in the table family.",
                "tags": [
                    "Create a new Genealogy"
                ],
                "summary": "Create a new reference (person ID) for a new Genealogy, which is relate to the table 'family'",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Genealogy"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Family": {
            "type": "object",
            "properties": {
                "family": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Family"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "person": {
                    "$ref": "#/definitions/models.Person"
                },
                "personFamily": {
                    "$ref": "#/definitions/models.Person"
                },
                "relationship": {
                    "$ref": "#/definitions/models.Relationship"
                }
            }
        },
        "models.Genealogy": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "person": {
                    "$ref": "#/definitions/models.Person"
                },
                "personFamily": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Family"
                    }
                },
                "personID": {
                    "type": "integer"
                }
            }
        },
        "models.Person": {
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
        "models.Relationship": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5050",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Genealogy API",
	Description:      "This is the documentation of the Genealogy API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
