package main

type Tasks struct {
	ID			string `json:"id"`
	Titulo		string `json:"titulo"`
	Descricao	string `json:"descricao"`
	Status		string `json:"status"`
}

var tasks = make(map[string]Tasks)

var validaStatus = map[string]bool{
	"A Fazer": true,
	"Em Progresso": true,
	"Concluídas": true,
}