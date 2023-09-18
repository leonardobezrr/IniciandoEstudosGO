package main

import "fmt"

//Funções variáticas "genéricas"
func soma(numeros ... int)int{
	total := 0
	for _,numero := range numeros{
		total += numero
	}
	return total
}

// Combinando string com int em funções variáticas
func escrever (text string,numeros ...int){
	for _ ,numero := range numeros{
		fmt.Println(text,numero)
	}
}

func main (){
	totalDaSoma := soma(1,2,3)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo",1,2,3,4)
}