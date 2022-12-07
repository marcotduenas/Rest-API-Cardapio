package main

import(
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "log"
    "encoding/json"
)

type alimentoDoCardapio struct{
    ID int `json:"id"`
    Nome string `json:"alimento"`
}

func cardapio(w http.ResponseWriter, r *http.Request){
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Contet-Type", "application/json")
    json.NewEncoder(w).Encode([]alimentoDoCardapio{{
        ID: 1,
        Nome: "Macarrão com carne moída",
    }})
}

func main(){
    fmt.Println("API sendo executada: http://localhost:3000")
    rota := mux.NewRouter()
    rota.HandleFunc("/", cardapio)
    log.Fatal(http.ListenAndServe(":3000", rota))
}
