package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"userService/internal/auth"
	"userService/internal/model"

	"github.com/go-chi/chi/v5"
)

type OrderRequest struct {
	Id      int `json:"id"`
	Balance int `json:"balance"`
}

func (h Handler) tryToOrderTask(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	h.controller.GetUserForUpdate(username)

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "data parsing error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &order)
	if err != nil {
		http.Error(w, "data parsing error", http.StatusBadRequest)
		return
	}

	balance, ok, err := h.grpcCli.TryOrderTask(username, order.Id, order.Balance)
	if !ok {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success ordering"))
}

func (h Handler) getOrdersForUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	page := chi.URLParam(r, "page")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "page num error", http.StatusBadRequest)
		return
	}

	if pageNum == 0 {
		pageNum = 1
	}

	orders, err := h.grpcCli.GetOrdersForUser(username, pageNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("page №%d:\n", pageNum)))
	w.Write(b)
}

func (h Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	var user model.UserForUpdate

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body read error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, "unmarshal error", http.StatusBadRequest)
		return
	}

	err = h.controller.UpdateUser(user)
	if err != nil {
		http.Error(w, "cant update user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) getInfo(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	user, err := h.controller.GetUser(u.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user.User.Info)
	if err != nil {
		http.Error(w, "parsing error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
