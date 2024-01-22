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
        "/comment/addComment": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "貼文留言"
                ],
                "summary": "貼文留言",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶ID",
                        "name": "userId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "貼文ID",
                        "name": "postId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "留言內容",
                        "name": "userComment",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/comment/getComment": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "貼文留言"
                ],
                "summary": "獲取貼文留言",
                "parameters": [
                    {
                        "type": "string",
                        "description": "貼文ID",
                        "name": "postId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/favorite/addFavoriteMountainRoad": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "最愛的山名及路名"
                ],
                "summary": "最愛的山名及路名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶ID",
                        "name": "userId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "山名ID",
                        "name": "mountainRoadID",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/favorite/delFavoriteMountainRoad": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "最愛的山名及路名"
                ],
                "summary": "移除用戶最愛的山名及路名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶ID",
                        "name": "userId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "山名ID",
                        "name": "mountainRoadID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/favorite/getFavoriteMountainRoad": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "最愛的山名及路名"
                ],
                "summary": "用戶最愛的山名及路名",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶ID",
                        "name": "userId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/like/addLike": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "貼文按讚"
                ],
                "summary": "貼文按讚",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶ID",
                        "name": "userId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "貼文ID",
                        "name": "postId",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mountainRoad/getMountainRoad": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "取得山名路名"
                ],
                "summary": "取得山名路名",
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "tags": [
                    "首頁"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/createPost": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "貼文資料"
                ],
                "summary": "創建貼文",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用戶名",
                        "name": "author",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "authorId",
                        "name": "authorId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "貼文類型",
                        "name": "postType",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "標題",
                        "name": "title",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "內容",
                        "name": "content",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "照片",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/deletePost": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "貼文資料"
                ],
                "summary": "刪除貼文",
                "parameters": [
                    {
                        "type": "string",
                        "description": "postId",
                        "name": "postId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/post/getPostList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "貼文資料"
                ],
                "summary": "貼文列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "postType",
                        "name": "postType",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/RefreshToken": {
            "post": {
                "tags": [
                    "更新token"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用戶名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/appleSignIn": {
            "post": {
                "tags": [
                    "用戶資料"
                ],
                "summary": "AppleSignIn",
                "parameters": [
                    {
                        "type": "string",
                        "description": "idToken",
                        "name": "idToken",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userName",
                        "name": "userName",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/createUser": {
            "post": {
                "tags": [
                    "用戶資料"
                ],
                "summary": "新增用戶",
                "parameters": [
                    {
                        "type": "string",
                        "description": "電子郵件",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用戶名",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密碼",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "確認密碼",
                        "name": "repassword",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用戶資料"
                ],
                "summary": "刪除用戶",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/findUserByNameAndPwd": {
            "post": {
                "tags": [
                    "用戶資料"
                ],
                "summary": "用戶登入",
                "parameters": [
                    {
                        "description": "用戶名和密碼",
                        "name": "UserInput",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "tags": [
                    "用戶資料"
                ],
                "summary": "用戶列表",
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/googleSignIn": {
            "post": {
                "tags": [
                    "用戶資料"
                ],
                "summary": "google登入",
                "parameters": [
                    {
                        "type": "string",
                        "description": "idToken",
                        "name": "idToken",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "用戶資料"
                ],
                "summary": "修改用戶",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "舊用戶名",
                        "name": "oldname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "新用戶名",
                        "name": "newname",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "舊密碼",
                        "name": "oldpassword",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "新密碼",
                        "name": "newpassword",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "phone",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "照片",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version/getVersion": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "獲取版本號"
                ],
                "summary": "獲取版本號",
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "service.UserInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
