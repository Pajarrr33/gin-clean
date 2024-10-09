package response

import "gin-db/model"

// type ResponseSuccessTodo struct {
// 	Code int    `json:"code"`
// 	Data []model.Todo `json:"data"`
// }

// type ResponseSuccessUser struct {
// 	Code int    `json:"code"`
// 	Data []model.User `json:"data"`
// }


// type ResponseTodoSingle struct{
// 	Code int    `json:"code"`
// 	Data model.Todo `json:"data"`
// }

// type ResponseUserSingle struct{
// 	Code int    `json:"code"`
// 	Data model.User `json:"data"`
// }

type ResponseCredentialSingle struct {
	Code int    `json:"code"`
	Data model.Credential `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseDelete struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}