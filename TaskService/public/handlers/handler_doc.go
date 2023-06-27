package handlers

import "userService/pkg/model"

// swagger:route GET / GetTasksRequest
// Admin: Получение всех задач.
// responses:
//   200: GetTasksResponse

// swagger:parameters GetTasksRequest
type GetTasksRequest struct {
}

// swagger:response GetTasksResponse
type GetTasksResponse struct {
	// in:body
	Body []model.Task
}

// swagger:route GET /{taskId} GetTaskRequest
// Admin: Получение задачи.
// responses:
//   200: GetTaskResponse

// swagger:parameters GetTaskRequest
type GetTaskRequest struct {
	// in:path
	TaskId int `json:"taskId"`
}

// swagger:response GetTaskResponse
type GetTaskResponse struct {
	// in:body
	Body model.Task
}

// swagger:route POST / CreateTaskRequest
// Admin: Создание новой задачи.
// responses:
//   200: CreateTaskResponse

// swagger:parameters CreateTaskRequest
type CreateTaskRequest struct {
	// in:body
	Body model.Task
}

// swagger:response CreateTaskResponse
type CreateTaskResponse struct {
}

// swagger:route PATCH / ChangeTaskRequest
// Admin: Изменение цены на задачу.
// responses:
//   200: ChangeTaskResponse

// swagger:parameters ChangeTaskRequest
type ChangeTaskRequest struct {
	// in:body
	Body model.TaskNewPrice
}

// swagger:response ChangeTaskResponse
type ChangeTaskResponse struct {
}

// swagger:route DELETE /{taskId} DeleteTaskRequest
// Admin: Удаление задачи.
// responses:
//   200: DeleteTaskResponse

// swagger:parameters DeleteTaskRequest
type DeleteTaskRequest struct {
	// in:path
	TaskId int `json:"taskId"`
}

// swagger:response DeleteTaskResponse
type DeleteTaskResponse struct {
}

// swagger:route DELETE /users/{username} DeleteTaskForUserRequest
// Admin: Удаление задач для пользователя.
// responses:
//   200: DeleteTaskForUserResponse

// swagger:parameters DeleteTaskForUserRequest
type DeleteTaskForUserRequest struct {
	// in:path
	Username string `json:"username"`
}

// swagger:response DeleteTaskForUserResponse
type DeleteTaskForUserResponse struct {
}
