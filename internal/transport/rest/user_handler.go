package rest

import (
	"encoding/json"
	"io"
	"net/http"
	"userService/internal/auth"
	"userService/internal/model"

	"github.com/go-chi/chi"
)

func (h Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	if username != u.Username {
		http.Error(w, "invalid resource", http.StatusBadRequest)
		return
	}

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
	username := chi.URLParam(r, "username")

	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	if username != u.Username {
		http.Error(w, "invalid resource", http.StatusBadRequest)
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
