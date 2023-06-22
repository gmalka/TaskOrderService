package rest

import (
	"encoding/json"
	"net/http"
	"userService/internal/auth"
	usercontroller "userService/internal/user_controller"
)

type ChangeTaskRequest struct {
	Id      int `json:"id"`
}

func (h Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	orders, err := h.grpcCli.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "token parse error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		http.Error(w, "permission denied", http.StatusBadRequest)
		return
	}

	h.grpcCli.UpdatePriceOfTask()

	w.WriteHeader(http.StatusOK)
}

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		http.Error(w, "permission denied", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}