package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", GetTask).Methods("GET")
	r.HandleFunc("/tasks", CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", DeleteTask).Methods("DELETE")

	r.Use(middlewareCORS)

	log.Println("Servidor rodando na porta 8080...")
	http.ListenAndServe(":8080", r)
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS"{
			return
		}

		next.ServeHTTP(w, r)
	})
}