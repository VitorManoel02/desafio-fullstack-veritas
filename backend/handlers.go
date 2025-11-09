package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/google/uuid"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lista []Task

	for _, task := range tasks {
		lista = append(lista, task)
	}

	json.NewEncoder(w).Encode(lista)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var nova Task
	json.NewDecoder(r.Body).Decode(&nova)

	if nova.Titulo == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"erro": "Titulo é obrigatorio}`))
		return
	}

	if !validaStatus[nova.Status] {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"erro": "Status inválido"}`))
		return
	}

	nova.ID = uuid.New().String()
	tasks[nova.ID] = nova

	json.NewEncoder(w).Encode(nova)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	task, existe := tasks[id]
	
	if !existe {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"erro": "Tarefa não encontrada"}`))
		return

	}

	json.NewDecoder(r.Body).Decode(&task)

	if task.Titulo == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"erro:" "Titulo é obrigatorio"}`))
		return
	}

	if !validaStatus[task.Status] {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"erro:" "Status invalido"}`))
		return
	}

	tasks[id] = task
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	_, existe := tasks[id]

	if !existe {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"erro: " "Tarefa não encontrada"}`))
		return
	}

	delete(tasks, id)
	w.Write([]byte(`{"mensagem :" "Tarefa excluida"}`))
}