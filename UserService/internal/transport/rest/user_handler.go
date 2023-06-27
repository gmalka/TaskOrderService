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

func (h Handler) getUsersNicknames(w http.ResponseWriter, r *http.Request) {
	s, err := h.controller.GetAllUsernames()
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(s)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) tryToOrderTask(w http.ResponseWriter, r *http.Request) {
	var answer model.TaskAnswer

	username := chi.URLParam(r, "username")
	taskId := r.Header.Get("taskId")

	id, err := strconv.Atoi(taskId)
	if err != nil {
		h.logger.Err.Printf("can't parse task id %s: %v", taskId, err.Error())
		http.Error(w, "message: some sserver error", http.StatusInternalServerError)
		return
	}

	order, err := h.grpcCli.CheckAndGetTask(username, id)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.controller.TryToBuyTask(username, order.Price)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.grpcCli.BuyTaskAnswer(username, order.Id)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	answer.Answer = order.Answer
	b, err := json.Marshal(answer)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) getUsersTasks(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
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

	orders, err := h.grpcCli.GetOrdersForUser(username, pageNum)
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
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("page №%d:\n", pageNum)))
	w.Write(b)
}

func (h Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role == usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("cant update admin %s\n", u.Username)
		http.Error(w, "message: permission denied, can't update admin", http.StatusForbidden)
		return
	}

	err := h.grpcCli.DeleteOrdersForUser(username)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err = h.controller.DeleteUser(username)
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

func (h Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var user model.UserForUpdate

	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	if u.Role == usercontroller.ADMIN_ROLE {
		h.logger.Err.Printf("cant update admin %s\n", u.Username)
		http.Error(w, "message: permission denied, can't update admin", http.StatusForbidden)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.controller.UpdateUser(user)
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

func (h Handler) getInfo(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	user, err := h.controller.GetUser(username)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user.User.Info)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
