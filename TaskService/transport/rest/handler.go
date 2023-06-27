package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"taskServer/model"
	ordercontroller "taskServer/order_controller"
	"text/template"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	controller ordercontroller.Controller
	logger     Log
}

func NewHandler(controller ordercontroller.Controller, logger Log) Handler {
	return Handler{controller: controller, logger: logger}
}

func (h Handler) InitRouter(logging bool) http.Handler {
	r := chi.NewRouter()

	if logging {
		r.Use(middleware.Logger)
	}

	r.Get("/", h.getTasks)
	r.Get("/{taskId}", h.getTask)
	r.Post("/", h.createTask)
	r.Patch("/", h.changeTaskPrice)
	r.Delete("/{taskId}", h.deleteTask)
	r.Delete("/users/{username}", h.deleteTaskForUser)

	r.Get("/swagger", h.swaggerUI)
	r.Get("/public/*", func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
	})

	return r
}

func (h Handler) deleteTaskForUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	err := h.controller.DeleteAllTasksOfUser(username)
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

func (h Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "taskId")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		h.logger.Err.Printf("can't parse tasks id %s: %v", taskId, err)
		http.Error(w, "message: invalid taskId", http.StatusInternalServerError)
		return
	}

	err = h.controller.DeleteTask(id)
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

func (h Handler) changeTaskPrice(w http.ResponseWriter, r *http.Request) {
	var newPrice model.TaskNewPrice

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &newPrice)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.controller.ChangeTaskPrice(newPrice.Id, newPrice.Price)
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

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Err.Printf("cant read body: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &task)
	if err != nil {
		h.logger.Err.Printf("unmarshal error: %v\n", err.Error())
		http.Error(w, "message: wrong input data", http.StatusBadRequest)
		return
	}

	err = h.controller.CreateTask(task)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err = json.Marshal(model.ResponseMessage{Message: "success create"})
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) getTask(w http.ResponseWriter, r *http.Request) {
	taskId := chi.URLParam(r, "taskId")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		h.logger.Err.Printf("can't parse tasks id %s: %v", taskId, err)
		http.Error(w, "message: invalid taskId", http.StatusInternalServerError)
		return
	}

	task, err := h.controller.CheckAndGetTask("", id)
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(task)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h Handler) getTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.controller.GetAllTasks()
	if err != nil {
		h.logger.Err.Println(err.Error())
		http.Error(w, fmt.Sprintf("message: %s", err.Error()), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(tasks)
	if err != nil {
		h.logger.Err.Printf("marshal error: %v\n", err.Error())
		http.Error(w, "message: some server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
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
