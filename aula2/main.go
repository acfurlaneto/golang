package main

import "fmt"

func main() {
    var usuario, senha string

    fmt.Print("Digite o usuário: ")
    fmt.Scan(&usuario)

    fmt.Print("Digite a senha: ")
    fmt.Scan(&senha)

    if usuario == "admin" && senha == "1234" {
        fmt.Println("Acesso permitido!")
    } else {
        fmt.Println("Acesso negado!")
    }
}
