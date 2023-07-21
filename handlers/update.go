package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project_api/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) //metodo que recebe 3 parametros: o primeiro é o response writer, onde ele vai escrever a mensagem de  erro; o segundo é a mensagem de erro em si; o terceiro é o status code
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) //metodo que recebe 3 parametros: o primeiro é o response writer, onde ele vai escrever a mensagem de  erro; o segundo é a mensagem de erro em si; o terceiro é o status code
		return
	}

	rows, err := models.Update(int64(id), todo)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) //metodo que recebe 3 parametros: o primeiro é o response writer, onde ele vai escrever a mensagem de  erro; o segundo é a mensagem de erro em si; o terceiro é o status code
		return
	}
	if rows > 1 {
		log.Printf("Erro: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Erro":    false,
		"Message": "dados atualizados com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
