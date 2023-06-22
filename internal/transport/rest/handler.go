package rest

import (
	"encoding/json"
	"net/http"
	"userService/internal/auth"
	"userService/internal/transport/grpc"
	usercontroller "userService/internal/user_controller"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	controller   usercontroller.Controller
	tokenManager auth.TokenManager
	grpcCli      grpc.RemoteOrderService
}

func NewHandler(controller usercontroller.Controller, tokenManager auth.TokenManager, grpcCli grpc.RemoteOrderService) Handler {
	return Handler{controller: controller, tokenManager: tokenManager, grpcCli: grpcCli}
}

func (h Handler) InitRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/users", h.getUsersNicknames)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.registerUser)
		r.Post("/login", h.loginIn)
		r.Route("/refresh", func(r chi.Router) {
			r.Use(h.checkRefresh)
			r.Post("/", h.refresh)
		})
	})

	r.Route("/{username}", func(r chi.Router) {
		r.Use(h.checkAccess)
		r.Get("/", h.getInfo)
		r.Put("/", h.updateUser)

		r.Route("/orders", func(r chi.Router) {
			r.Use(h.checkAccess)
			r.Get("/{page:^(|[1-9][0-9]*)$}", h.getOrdersForUser)
			r.Post("/", h.tryToOrderTask)

			r.Get("/", h.getAllTasks)
			r.Put("/", h.updateTask)
			r.Post("/", h.createTask)
		})

	})

	return r
}

func (h Handler) getUsersNicknames(w http.ResponseWriter, r *http.Request) {
	s, err := h.controller.GetAllUsernames()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(s)
	if err != nil {
		http.Error(w, "data parsing error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func generateTokens(tokenManager auth.TokenManager, user auth.UserInfo) (string, string, error) {
	accessToken, err := tokenManager.CreateToken(user, auth.ACCESS_TOKEN_TTL, auth.AccessToken)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := tokenManager.CreateToken(user, auth.REFRESH_TOKEN_TTL, auth.RefreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

type UserRequest struct{}
