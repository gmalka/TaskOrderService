package rest

import "net/http"

func (h Handler) getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) updateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) createTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}