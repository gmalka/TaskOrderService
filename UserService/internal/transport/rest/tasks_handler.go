package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"userService/internal/auth/tokenManager"
	"userService/internal/model"
	usercontroller "userService/internal/user_controller"

	"github.com/go-chi/chi/v5"
)

type UpdateReuqest struct {
	TaskId     int `json:"taskId"`
	NewBalance int `json:"balance"`
}

func (h Handler) getUsersTasksWithoutAnswer(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		h.logger.Err.Printf("can't parse number %s: %v", page, err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if pageNum <= 0 {
		pageNum = 1
	}

	tasks, err := h.grpcCli.GetAllTasksWithoutAnswers(pageNum)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) getTasksWithoutAnswer(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		h.logger.Err.Printf("can't parse number %s: %v", page, err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if pageNum == 0 {
		pageNum = 1
	}

	orders, err := h.grpcCli.GetAllTasksWithoutAnswers(pageNum)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(orders)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("page №%d:\n", pageNum)))
	w.Write(b)
}

func (h Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	tasks, err := h.grpcCli.GetAllTasks()
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) updateUserBalance(w http.ResponseWriter, r *http.Request) {
	var change model.BalanceChange

	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some parse error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &change)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.controller.UpdateBalance(change.Username, change.Money)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(model.ResponseMessage{Message: "success balance change"})
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	var update UpdateReuqest

	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some parse error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &update)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.grpcCli.UpdatePriceOfTask(update.TaskId, update.NewBalance)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(model.ResponseMessage{Message: "success update"})
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "taskId")

	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(taskId)
	if err != nil {
		h.logger.Err.Printf("can't parse tasks id %s: %v", taskId, err)
		http.Error(w, "message: invalid taskId", http.StatusInternalServerError)
		return
	}

	err = h.grpcCli.DeleteTask(id)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(model.ResponseMessage{Message: "success delete"})
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("can't get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &task)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.grpcCli.CreateNewTask(task)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(model.ResponseMessage{Message: "success create"})
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
