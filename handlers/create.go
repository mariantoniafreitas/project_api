package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project_api/models"
)

//Funcao do tipo http Handler
func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) //metodo que recebe 3 parametros: o primeiro é o response writer, onde ele vai escrever a mensagem de  erro; o segundo é a mensagem de erro em si; o terceiro é o status code
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
