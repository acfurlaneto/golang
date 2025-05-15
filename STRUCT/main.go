package main

import "fmt"
type Jogador struct {
	nome string
	vida int
	nivel int
}

func exibeDados(j Jogador){
	fmt.Println("Nome do Jogador:", j.nome)
	fmt.Println("Vida do Jogador:", j.vida)
	fmt.Println("Nivel do Jogador:", j.vida)
}
func main() {
	j1 := Jogador{"Nine9", 100, 1}
	exibeDados(j1)
}

type Carro struct {
	modelo string
	ano int
	velocidade string
}

func exibeDados(c Carro){
	fmt.Println("Modelo do Carro:", c.modelo)
	fmt.Println("Ano do Carro:", c.ano)
	fmt.Println("Velocidade do Carro:", c.velocidade)
}
func main() {
	c1 := Carro{"Porsche", 2017, "120km"}
	exibeDados(c1)
}