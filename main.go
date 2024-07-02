package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Numero struct {
	Num int `json:"num"`
}
type Total struct {
	Total     int   `json:"total"`
	Historico []int `json:"historico"`
}

var total Total

func adicao(w http.ResponseWriter, r *http.Request) {
	var numero Numero
	err := json.NewDecoder(r.Body).Decode(&numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	total.Historico = append(total.Historico, numero.Num)
	total.Total += numero.Num

	response := struct {
		Numero Numero `json:"numero"`
		Total  int    `json:"total"`
	}{
		Numero: numero,
		Total:  total.Total,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func subtracao(w http.ResponseWriter, r *http.Request) {
	var numero Numero
	err := json.NewDecoder(r.Body).Decode(&numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	total.Historico = append(total.Historico, numero.Num)
	total.Total -= numero.Num

	response := struct {
		Numero Numero `json:"numero"`
		Total  int    `json:"total"`
	}{
		Numero: numero,
		Total:  total.Total,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func multiplicacao(w http.ResponseWriter, r *http.Request) {
	var numero Numero
	err := json.NewDecoder(r.Body).Decode(&numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	total.Historico = append(total.Historico, numero.Num)
	total.Total *= numero.Num

	response := struct {
		Numero Numero `json:"numero"`
		Total  int    `json:"total"`
	}{
		Numero: numero,
		Total:  total.Total,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func divisao(w http.ResponseWriter, r *http.Request) {
	var numero Numero
	err := json.NewDecoder(r.Body).Decode(&numero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	total.Historico = append(total.Historico, numero.Num)
	total.Total /= numero.Num

	response := struct {
		Numero Numero `json:"numero"`
		Total  int    `json:"total"`
	}{
		Numero: numero,
		Total:  total.Total,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func seeTotal(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(total.Total)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func limparTotal(w http.ResponseWriter, r *http.Request) {
	total.Total = 0
	fmt.Fprintf(w, "Reiniciado ...")
	w.WriteHeader(http.StatusNoContent)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API de adição:\n- Adicionar número POST /adicao, subtracao, multplicacao, divisao\n- Ver total: GET /total\n - Limpar números: DELETE /limpar")
}
func main() {
	router := mux.NewRouter()

	log.Println("Executando ...")

	router.HandleFunc("/home", homepage).Methods("GET")
	router.HandleFunc("/adicao", adicao).Methods("POST")
	router.HandleFunc("/subtracao", subtracao).Methods("POST")
	router.HandleFunc("/mult", multiplicacao).Methods("POST")
	router.HandleFunc("/div", divisao).Methods("POST")
	router.HandleFunc("/total", seeTotal).Methods("GET")
	router.HandleFunc("/limpar", limparTotal).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
