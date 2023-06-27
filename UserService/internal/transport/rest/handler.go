package rest

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"userService/internal/auth/passwordManager"
	"userService/internal/auth/tokenManager"
	"userService/internal/model"
	mygrpc "userService/internal/transport/grpc"
	usercontroller "userService/internal/user_controller"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

const (
	swaggerTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"></script>
    <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-standalone-preset.js"></script> -->
    <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
    <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-bundle.js"></script> -->
    <link rel="stylesheet" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css" />
    <!-- <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui.css" /> -->
	<style>
		body {
			margin: 0;
		}
	</style>
    <title>Swagger</title>
</head>
<body>
    <div id="swagger-ui"></div>
    <script>
        window.onload = function() {
          SwaggerUIBundle({
            url: "/public/swagger.json?{{.Time}}",
            dom_id: '#swagger-ui',
            presets: [
              SwaggerUIBundle.presets.apis,
              SwaggerUIStandalonePreset
            ],
            layout: "StandaloneLayout"
          })
        }
    </script>
</body>
</html>
`
)

type Log struct {
	Err *log.Logger
	Inf *log.Logger
}

type Handler struct {
	controller   usercontroller.Controller
	tokenManager tokenManager.TokenManager
	grpcCli      mygrpc.RemoteOrderClient
	passManager  passwordManager.PasswordManager
	logger       Log
}

func NewHandler(controller usercontroller.Controller, tokenManager tokenManager.TokenManager, grpcCli mygrpc.RemoteOrderClient, passManager passwordManager.PasswordManager, logger Log) Handler {
	return Handler{controller: controller, tokenManager: tokenManager, grpcCli: grpcCli, passManager: passManager, logger: logger}
}

func (h Handler) InitRouter(logging bool) http.Handler {
	r := chi.NewRouter()

	if logging {
		r.Use(middleware.Logger)
	}

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", h.registerUser)
		r.Post("/login", h.loginIn)
		r.Route("/refresh", func(r chi.Router) {
			r.Use(h.checkRefresh)
			r.Post("/", h.refresh)
		})
	})

	r.Get("/tasks/{page:^(|0|[1-9][0-9]*)$}", h.getUsersTasksWithoutAnswer)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.getUsersNicknames)

		r.Route("/{username}", func(r chi.Router) {
			r.Use(h.checkAccess)
			r.Get("/", h.getInfo)
			r.Put("/", h.updateUser)
			r.Delete("/", h.deleteUser)
			r.Post("/", h.tryToOrderTask)

			r.Patch("/", h.updateUserBalance)

			r.Route("/orders", func(r chi.Router) {
				r.Get("/purchased/{page:^(|0|[1-9][0-9]*)$}", h.getUsersTasks)

				r.Get("/", h.getAllTasks)
				r.Put("/", h.updateTask)
				r.Post("/", h.createTask)
				r.Delete("/{taskId}", h.deleteTask)
			})
		})
	})

	r.Get("/swagger", h.swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})

	return r
}

func (h Handler) swaggerUI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.New("swagger").Parse(swaggerTemplate)
	if err != nil {
		return
	}

	err = tmpl.Execute(w, struct {
		Time int64
	}{
		Time: time.Now().Unix(),
	})
	if err != nil {
		return
	}
}

func generateTokens(tm tokenManager.TokenManager, user tokenManager.UserInfo) (model.AuthInfo, error) {
	accessToken, err := tm.CreateToken(user, tokenManager.ACCESS_TOKEN_TTL, tokenManager.AccessToken)
	if err != nil {
		return model.AuthInfo{}, err
	}

	refreshToken, err := tm.CreateToken(user, tokenManager.REFRESH_TOKEN_TTL, tokenManager.RefreshToken)
	if err != nil {
		return model.AuthInfo{}, err
	}

	return model.AuthInfo{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

type UserRequest struct{}
