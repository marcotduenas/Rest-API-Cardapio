package main

import(
    "net/http"
    "log"
    "fmt"
    "io"
    "encoding/json"
)

type alimentoDoCardapio struct{
    ID int `json:"id"`
    Nome string `json:"alimento"`
}

func main(){
    res, err := http.Get("http://localhost:3000/")
    if err != nil{
        log.Fatal(err)
        return
    }

    if res.StatusCode != 200{
        log.Fatal("Erro: ", res.StatusCode)
        return
    }

    content, err := io.ReadAll(res.Body)

    var resposta []alimentoDoCardapio
    err = json.Unmarshal(content, &resposta)
    if err != nil{
        log.Fatal(err)
        return
    }

    fmt.Println(resposta)

}
