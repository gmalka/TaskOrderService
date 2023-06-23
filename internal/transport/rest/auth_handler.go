package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"userService/internal/auth"
	"userService/internal/model"

	"github.com/go-chi/chi/v5"
)

func (h Handler) loginIn(w http.ResponseWriter, r *http.Request) {
	var userAuth model.UserAuth

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &userAuth)
	if err != nil {
		h.logger.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	user, err := h.controller.GetUser(userAuth.Username)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	if err := auth.CheckPassword(userAuth.Password, user.User.Password); err != nil {
		h.logger.Println(err.Error())
		http.Error(w, "message: passwords mismatch", http.StatusUnauthorized)
		return
	}

	access, refresh, err := generateTokens(h.tokenManager, auth.UserInfo{
		Username:  user.User.Username,
		Role:      user.Role,
		Firstname: user.User.Info.Firstname,
		Lastname:  user.User.Info.Lastname,
	})
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success login\n"))
	w.Write([]byte("Access: "))
	w.Write([]byte(access))
	w.Write([]byte("\nRefresh: "))
	w.Write([]byte(refresh))
}

func (h Handler) refresh(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(UserRequest{}).(auth.UserClaims)
	if !ok {
		h.logger.Println("cant get data from context")
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	access, refresh, err := generateTokens(h.tokenManager, auth.UserInfo{
		Username:  u.Username,
		Role:      u.Role,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
	})
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success login\n"))
	w.Write([]byte("Access: "))
	w.Write([]byte(access))
	w.Write([]byte("\nRefresh: "))
	w.Write([]byte(refresh))
}

func (h Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		h.logger.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	user.Password, err = auth.HashPassword(user.Password)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = h.controller.CreateUser(user)
	if err != nil {
		h.logger.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
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
			h.logger.Printf("wrong authorization: %v\n", tokenParts)
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		u, err := h.tokenManager.ParseToken(tokenParts[1], auth.AccessToken)
		if err != nil {
			h.logger.Println(err.Error())
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		username := chi.URLParam(r, "username")
		if username != u.Username {
			h.logger.Printf("username in token and path are different: %s-%s", username, u.Username)
			http.Error(w, "message: invalid resource", http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), UserRequest{}, u)
		log.Println("Success")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h Handler) checkRefresh(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenRaw := r.Header.Get("Authorization")

		tokenParts := strings.Split(tokenRaw, " ")
		if len(tokenParts) < 2 && tokenParts[0] != "Bearer" {
			h.logger.Printf("wrong authorization: %v\n", tokenParts)
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		u, err := h.tokenManager.ParseToken(tokenParts[1], auth.RefreshToken)
		if err != nil {
			h.logger.Println(err.Error())
			http.Error(w, "message: wrong authorization token", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), UserRequest{}, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
