package handlers

import "userService/internal/model"

// swagger:route POST /auth/refresh auth RefreshRequest
// Обновление токена.
// security:
//   - Bearer: []
// responses:
//   200: RefreshResponse
//   400: StatusBadRequest

// swagger:parameters RefreshRequest
type RefreshRequest struct {
}

// swagger:response RefreshResponse
type RefreshResponse struct {
	// in:body
	AccessToken model.AuthInfo
}

// swagger:route POST /auth/register auth RegisterRequest
// Регистрация пользователя.
//
// responses:
//   200: RegisterResponse
//   400: StatusBadRequest

// swagger:parameters RegisterRequest
type RegisterRequest struct {
	// in:body
	Body model.UserWithoutBalance `json:"users"`
}

// swagger:response RegisterResponse
type RegisterResponse struct {
	// in:body
	Body model.ResponseMessage
}

// swagger:route POST /auth/login auth LoginRequest
// Авторизация пользователя.
//
// responses:
//   200: LoginResponse
//   400: StatusBadRequest

// swagger:parameters LoginRequest
type LoginRequest struct {
	// in:body
	Body model.UserAuth `json:"usersAuth"`
}

// swagger:response LoginResponse
type LoginResponse struct {
	// in:body
	AccessToken model.AuthInfo
}
