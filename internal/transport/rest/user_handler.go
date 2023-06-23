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

func (h Handler) tryToOrderTask(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	taskId := r.Header.Get("taskId")

	id, err := strconv.Atoi(taskId)
	if err != nil {
		h.logger.Printf("can't parse task id %s: %v", taskId, err.Error())
		http.Error(w, "some sserver error", http.StatusInternalServerError)
		return
	}

	order, err := h.grpcCli.GetTask(id)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.controller.TryToBuyTask(username, order.Price)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.grpcCli.BuyTaskAnswer(username, order.Id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Answer: "))
	w.Write([]byte(strconv.Itoa(order.Answer)))
}

func (h Handler) getUsersOrders(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	page := chi.URLParam(r, "page")

	pageNum, err := strconv.Atoi(page)
	if err != nil {
		h.logger.Printf("can't parse number %s: %v", page, err.Error())
		http.Error(w, "some server error", http.StatusInternalServerError)
		return
	}

	if pageNum == 0 {
		pageNum = 1
	}

	orders, err := h.grpcCli.GetOrdersForUser(username, pageNum)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(orders)
	if err != nil {
		h.logger.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "wrong input data", http.StatusBadRequest)
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
		h.logger.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		h.logger.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "wrong input data", http.StatusBadRequest)
		return
	}

	err = h.controller.UpdateUser(user)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) getInfo(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		h.logger.Println("cant get data from context")
		http.Error(w, "some server error", http.StatusInternalServerError)
		return
	}

	user, err := h.controller.GetUser(u.Username)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user.User.Info)
	if err != nil {
		h.logger.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "wrong input data", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}