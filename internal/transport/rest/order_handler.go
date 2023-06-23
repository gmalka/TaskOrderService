package rest

import (
	"encoding/json"
	"fmt"
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
		h.logger.Println(err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		h.logger.Printf("marshal error: %v\n", err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	var update UpdateReuqest

	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		h.logger.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some parse error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &update)
	if err != nil {
		h.logger.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.grpcCli.UpdatePriceOfTask(update.TaskId, update.NewPrice)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success update"))
}

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		h.logger.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role != usercontroller.ADMIN_ROLE {
		h.logger.Printf("permission denied for user %s\n", u.Username)
		http.Error(w, "message: permission denied", http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &task)
	if err != nil {
		h.logger.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.grpcCli.CreateNewTask(task)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success create"))
}
