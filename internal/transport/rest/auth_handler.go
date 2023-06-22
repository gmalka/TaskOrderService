package rest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"userService/internal/auth"
	"userService/internal/model"
)

func (h Handler) loginIn(w http.ResponseWriter, r *http.Request) {
	var userAuth model.UserAuth

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "input data read error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &userAuth)
	if err != nil {
		http.Error(w, "data parsing error", http.StatusBadRequest)
		return
	}

	user, err := h.controller.GetUser(userAuth.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ok := auth.CheckPassword(userAuth.Password, user.User.Password); !ok {
		http.Error(w, "passwords mismatch", http.StatusBadRequest)
		return
	}

	access, refresh, err := generateTokens(h.tokenManager, auth.UserInfo{
		Username: user.User.Username,
		Role: user.Role,
		Firstname: user.User.Info.Firstname,
		Lastname: user.User.Info.Lastname,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success login"))
	w.Write([]byte(access))
	w.Write([]byte(refresh))
}

func (h Handler) refresh(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		http.Error(w, "some token error", http.StatusBadRequest)
		return
	}

	access, refresh, err := generateTokens(h.tokenManager, auth.UserInfo{
		Username: u.Username,
		Role: u.Role,
		Firstname: u.Firstname,
		Lastname: u.Lastname,
	})
	if err != nil {
		http.Error(w, "token generate error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success login"))
	w.Write([]byte(access))
	w.Write([]byte(refresh))
}

func (h Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "input data read error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, "data parsing error", http.StatusBadRequest)
		return
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "data parsing error", http.StatusBadRequest)
		return
	}

	err = h.controller.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success registartion"))
}

func (h Handler) checkAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenRaw := r.Header.Get("Authorization")

		tokenParts := strings.Split(tokenRaw, " ")
		if len(tokenParts) < 2 && tokenParts[0] != "Bearer" {
			http.Error(w, "wrong input data", http.StatusBadRequest)
			return
		}

		u, err := h.tokenManager.ParseToken(tokenParts[1], auth.AccessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), UserRequest{}, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h Handler) checkRefresh(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenRaw := r.Header.Get("Authorization")

		tokenParts := strings.Split(tokenRaw, " ")
		if len(tokenParts) < 2 && tokenParts[0] != "Bearer" {
			http.Error(w, "wrong input data", http.StatusBadRequest)
			return
		}

		u, err := h.tokenManager.ParseToken(tokenParts[1], auth.RefreshToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), UserRequest{}, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
