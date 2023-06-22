package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"userService/internal/auth"
	"userService/internal/model"
	usercontroller "userService/internal/user_controller"
)

type UpdateReuqest struct {
	TaskId   int `json:"id"`
	NewPrice int `json:"price"`
}

func (h Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.grpcCli.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "token parse error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	var update UpdateReuqest

	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		http.Error(w, "permission denied", http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "some parse error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &update)
	if err != nil {
		http.Error(w, "some parse error", http.StatusBadRequest)
		return
	}

	err = h.grpcCli.UpdatePriceOfTask(update.TaskId, update.NewPrice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success update"))
}

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		http.Error(w, "permission denied", http.StatusBadRequest)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "some parse error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &task)
	if err != nil {
		http.Error(w, "some parse error", http.StatusBadRequest)
		return
	}

	err = h.grpcCli.CreateNewTask(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success create"))
}