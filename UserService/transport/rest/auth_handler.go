package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"userService/auth/tokenManager"
	"userService/pkg/model"

	"github.com/go-chi/chi/v5"
)

func (h Handler) loginIn(w http.ResponseWriter, r *http.Request) {
	var userAuth model.UserAuth

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &userAuth)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	user, err := h.controller.GetUser(userAuth.Username)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err := h.passManager.CheckPassword(userAuth.Password, user.User.Password); err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, "message: passwords mismatch", http.StatusUnauthorized)
		return
	}

	tokens, err := generateTokens(h.tokenManager, tokenManager.UserInfo{
		Username:  user.User.Username,
		Role:      user.Role,
		Firstname: user.User.Info.Firstname,
		Lastname:  user.User.Info.Lastname,
	})
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(tokens)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) refresh(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(tokenManager.UserClaims)
	if !ok {
		h.logger.Err.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	tokens, err := generateTokens(h.tokenManager, tokenManager.UserInfo{
		Username:  u.Username,
		Role:      u.Role,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
	})
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(tokens)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

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

	user.Password, err = h.passManager.HashPassword(user.Password)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	h.logger.Inf.Println(user.Password)

	err = h.controller.CreateUser(user)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(model.ResponseMessage{Message: "success register"})
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) checkAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenRaw := r.Header.Get("Authorization")

		tokenParts := strings.Split(tokenRaw, " ")
		if len(tokenParts) < 2 && tokenParts[0] != "Bearer" {
			h.logger.Err.Printf("wrong authorization: %v\n", tokenParts)
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		u, err := h.tokenManager.ParseToken(tokenParts[1], tokenManager.AccessToken)
		if err != nil {
			h.logger.Err.Println(err.Error())
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		username := chi.URLParam(r, "username")
		if username != u.Username {
			h.logger.Err.Printf("username in token and path are different: %s-%s", username, u.Username)
			http.Error(w, "message: invalid resource", http.StatusNotFound)
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
			h.logger.Err.Printf("wrong authorization: %v\n", tokenParts)
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		u, err := h.tokenManager.ParseToken(tokenParts[1], tokenManager.RefreshToken)
		if err != nil {
			h.logger.Err.Println(err.Error())
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), UserRequest{}, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
