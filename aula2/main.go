package main

import "fmt"

func main() {
    var num1, num2 float64

    fmt.Print("Digite o primeiro número: ")
    fmt.Scan(&num1)
    fmt.Print("Digite o segundo número: ")
    fmt.Scan(&num2)

    fmt.Println("Soma:", num1+num2)
    fmt.Println("Subtração:", num1-num2)
    fmt.Println("Multiplicação:", num1*num2)

    if num2 != 0 {
        fmt.Println("Divisão:", num1/num2)
    } else {
        fmt.Println("Não é possível dividir por zero.")
    }
}
