package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"project_api/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) //metodo que recebe 3 parametros: o primeiro é o response writer, onde ele vai escrever a mensagem de  erro; o segundo é a mensagem de erro em si; o terceiro é o status code
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) //metodo que recebe 3 parametros: o primeiro é o response writer, onde ele vai escrever a mensagem de  erro; o segundo é a mensagem de erro em si; o terceiro é o status code
		return
	}
	if rows > 1 {
		log.Printf("Erro: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Erro":    false,
		"Message": "registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
