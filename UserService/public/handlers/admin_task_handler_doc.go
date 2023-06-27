package handlers

import (
	"userService/pkg/model"
	"userService/transport/rest"
)

// swagger:route POST /users/{username}/admin admin CreateTaskRequest
// Admin: Создание новой задачи.
// security:
//   - Bearer: []
// responses:
//   200: CreateTaskResponse
//   400: StatusBadRequest

// swagger:parameters CreateTaskRequest
type CreateTaskRequest struct {
	// in:path
	Username string `json:"username"`
	// in:body
	Body model.TaskWithoutAnswer `json:"task"`
}

// swagger:response CreateTaskResponse
type CreateTaskResponse struct {
	// in:body
	Message model.ResponseMessage
}

// swagger:route PATCH /users/{username}/admin admin ChangeBalanceRequest
// Admin: Изменение баланса пользователя.
// security:
//   - Bearer: []
// responses:
//   200: ChangeBalanceResponse
//   400: StatusBadRequest

// swagger:parameters ChangeBalanceRequest
type ChangeBalanceRequest struct {
	// in:path
	Username string `json:"username"`
	// in:body
	Body model.BalanceChange
}

// swagger:response ChangeBalanceResponse
type ChangeBalanceResponse struct {
	// in:body
	Message model.ResponseMessage
}

// swagger:route DELETE /users/{username}/admin/{taskId} admin DeleteTaskRequest
// Admin: Удаление задачи.
// security:
//   - Bearer: []
// responses:
//   200: DeleteTaskResponse
//   400: StatusBadRequest

// swagger:parameters DeleteTaskRequest
type DeleteTaskRequest struct {
	// in:path
	Username string `json:"username"`
	// in:path
	TaskId string `json:"taskId"`
}

// swagger:response DeleteTaskResponse
type DeleteTaskResponse struct {
	// in:body
	Message model.ResponseMessage
}

// swagger:route PUT /users/{username}/admin admin UpdateBalanceRequest
// Admin: Изменение цены для задачи.
// security:
//   - Bearer: []
// responses:
//   200: UpdateBalanceResponse
//   400: StatusBadRequest

// swagger:parameters UpdateBalanceRequest
type UpdateBalanceRequest struct {
	// in:path
	Username string `json:"username"`
	// in:body
	Body rest.UpdateReuqest
}

// swagger:response UpdateBalanceResponse
type UpdateBalanceResponse struct {
	// in:body
	Message model.ResponseMessage
}

// swagger:route GET /users/{username}/admin admin GetAllTasksRequest
// Admin: Получение всех доступных задач.
// security:
//   - Bearer: []
// responses:
//   200: GetAllTasksResponse
//   400: StatusBadRequest

// swagger:parameters GetAllTasksRequest
type GetAllTasksRequest struct {
	// in:path
	Username string `json:"username"`
}

// swagger:response GetAllTasksResponse
type GetAllTasksResponse struct {
	// in:body
	Answer []model.Task `json:"answer"`
}
