package main

import "fmt"

func main() {
	fmt.Println("Olá mundo!")
	var nomeOne string = "Major"
	var nomeTwo = "Paola"
	var nomeTree string
	fmt.Println(nomeOne, nomeTwo, nomeTree)
	nomeTree = "Jhonny"
	nomeOne = "Daniel"
	fmt.Println(nomeTree, nomeOne)
	nomeFour := "Caruzo"
	fmt.Println(nomeFour)
	var scoreOne float32 = 0.0
	scoreTwo := 10.0
	fmt.Println(scoreOne, scoreTwo)
	const pi float32 = 3.1465
	const nome string = "Ana"
	fmt.Println(pi, nome)
}
