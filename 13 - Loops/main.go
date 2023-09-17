package main

import (
	"fmt"
	"time"
)

func main (){
	fmt.Println("Loops")

	// i := 0
	// for i < 10 {
	// 	fmt.Println("Incrementando i",i)
	// 	time.Sleep(time.Second)
	// 	i++
	// }

	// for j := 0; j < 10; j++{
	// 	fmt.Println("Incrementando j",j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João","Leo","André"}

	// O "_" é utilizado para ocultar o índice
	for _,nome := range nomes{
		fmt.Println(nome)
	}

	for i,letra := range "PALAVRA"{
		fmt.Println(i,string(letra)) //Se não utilizar string() ele irá retornar os valores das letras na tabela ASCII
	}

	user := map[string]string{
		"nome": "Leonardo",
		"sobrenome":"Bezerra",
	}

	for chave,valor := range user{
		fmt.Println(chave,valor)
	}

	for {
		fmt.Println("Executando infinito")
		time.Sleep(time.Second/2)
	}
	
}