package handlers

import "userService/internal/model"

// swagger:route GET /users user EmptyRequest
// Получение ников всех пользователей.
//
// responses:
//   200: UsersGetAllResponse
//   400: StatusBadRequest

// swagger:parameters EmptyRequest
type EmptyRequest struct {
}

// swagger:response UsersGetAllResponse
type UsersGetAllResponse struct {
	// массив ников
	//
	// in:body
	Body []string `json:"users"`
}

// Bad Request replys an error of API calling.
//
// swagger:response StatusBadRequest
type StatusBadRequest struct {
	// in:body
	Message string `json:"message"`
}

// swagger:route PUT /users/{username} user UpdateuserRequest
// Обновление информации о пользователе.
// security:
//   - Bearer: []
// responses:
//   200: updateuserResponse
//   400: StatusBadRequest

// swagger:parameters UpdateuserRequest
type UpdateuserRequest struct {
	// in:path
	Username string `json:"username"`
	// in:body
	Body model.UserForUpdate `json:"user"`
}

// swagger:response updateuserResponse
type updateuserResponse struct {
	// информация о пользователе
	//
	// in:body
	Body model.UserInfo `json:"userInfo"`
}

// swagger:route GET /users/{username} user GetUserInfoRequest
// Получение информации о пользователе.
// security:
//   - Bearer: []
// responses:
//   200: UserInfoResponse
//   400: StatusBadRequest

// swagger:parameters GetUserInfoRequest
type GetUserInfoRequest struct {
	// in:path
	Username string `json:"username"`
}

// swagger:response UserInfoResponse
type UserInfoResponse struct {
	// информация о пользователе
	//
	// in:body
	Body model.UserInfo `json:"userInfo"`
}
