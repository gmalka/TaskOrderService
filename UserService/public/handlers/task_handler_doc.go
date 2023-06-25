package handlers

import "userService/internal/model"

// swagger:route GET /tasks/{page} orders GetAllWithoutAnswersRequest
// Получение всех задач без ответов.
//
// responses:
//   200: GetAllWithoutAnswersResponse
//   400: StatusBadRequest

// swagger:parameters GetAllWithoutAnswersRequest
type GetAllWithoutAnswersRequest struct {
	// id заказываемого задания
	//
	// in:path
	Pahe string `json:"page"`
}

// swagger:response GetAllWithoutAnswersResponse
type GetAllWithoutAnswersResponse struct {
	// in:body
	Tasks []model.TaskWithoutAnswer `json:"tasks"`
}

// swagger:route POST /users/{username}/orders orders OrderTaskRequest
// Заказ решения для задачи.
// security:
//   - Bearer: []
// responses:
//   200: OrderTaskResponse
//   400: StatusBadRequest

// swagger:parameters OrderTaskRequest
type OrderTaskRequest struct {
	// in:path
	Username string `json:"username"`

	// id заказываемого задания
	//
	// in:header
	TaskId string `json:"taskId"`
}

// swagger:response OrderTaskResponse
type OrderTaskResponse struct {
	// in:body
	Answer model.TaskAnswer `json:"answer"`
}

// swagger:route GET /users/{username}/orders/purchased/{page} orders GetUsersOrdersRequest
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
