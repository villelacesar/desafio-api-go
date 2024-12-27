package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Pessoa struct {
	ID        string    `json: "id, omitempty"`
	Firstname string    `json: "firstname, omitempty"`
	Lastname  string    `json: "lastname, omitempty"`
	Endereco  *Endereco `json: "address, omitempty"`
}

type Endereco struct {
	Cidade string `json: "city, omitempty"`
	Estado string `json: "state, omitempty"`
}

var pessoas []Pessoa

// mostra todos os contatos da variavel Pessoa
func GetPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pessoa{})
}

// CriarPessoa cria um novo contato
func CriarPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = params["id"]
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoas)
}

// Deleta um contato
func DeletarPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pessoas {
		if item.ID == params["id"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pessoas)
	}
}

const port = ":8081"

func main() {

	router := mux.NewRouter()
	pessoas = append(pessoas, Pessoa{ID: "1", Firstname: "Cesar", Lastname: "Villela", Endereco: &Endereco{Cidade: "Pinhais", Estado: "Parana"}})
	pessoas = append(pessoas, Pessoa{ID: "2", Firstname: "Selma", Lastname: "Villela", Endereco: &Endereco{Cidade: "Pinhais", Estado: "Parana"}})
	router.HandleFunc("/contato", GetPessoas).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPessoa).Methods("GET")
	router.HandleFunc("/contato/{id}", CriarPessoa).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletarPessoa).Methods("DELETE")
	log.Fatal(http.ListenAndServe(port, router))
}
