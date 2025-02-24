// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

// ModelsUser defines model for Models.User.
type ModelsUser struct {
	// EmailAddress メールアドレス
	EmailAddress string `json:"email_address"`

	// Id ユーザーID
	Id int32 `json:"id"`

	// Name ユーザー名
	Name string `json:"name"`
}

// ResponsesApiV1UsersCreateUserResponse ユーザー作成成功のレスポンス
type ResponsesApiV1UsersCreateUserResponse struct {
	// User ユーザー
	User ModelsUser `json:"user"`
}

// UsersCreateUserJSONBody defines parameters for UsersCreateUser.
type UsersCreateUserJSONBody struct {
	// EmailAddress メールアドレス
	EmailAddress string `json:"email_address"`

	// Name ユーザー名
	Name string `json:"name"`
}

// UsersCreateUserJSONRequestBody defines body for UsersCreateUser for application/json ContentType.
type UsersCreateUserJSONRequestBody UsersCreateUserJSONBody
