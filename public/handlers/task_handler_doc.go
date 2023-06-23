package handlers

import "userService/internal/model"

// swagger:route POST /users/{username}/orders orders GetUsersOrdersRequest
// Получение заказов пользователя.
// security:
//   - Bearer: []
// responses:
//   200: GetUsersOrdersResponse
//   400: StatusBadRequest

// swagger:parameters GetUsersOrdersRequest
type GetUsersOrdersRequest struct {
	// in:path
	Username string `json:"username"`

	// страница, которую нужно получить, при пустом значении возвращает первую страницу.
	//
	// in:path
	Page string `json:"page"`
}

// swagger:response GetUsersOrdersResponse
type GetUsersOrdersResponse struct {
	// in:body
	Tasks []model.Task
}