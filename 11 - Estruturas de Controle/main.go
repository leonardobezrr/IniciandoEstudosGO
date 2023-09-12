package main

import "fmt"

func main (){
	fmt.Println("")
	fmt.Println("Estruturas de Controle")
	fmt.Println("")

	numero := 1
	if numero > 15 {
		fmt.Println("Número maior que 15")
	}else if numero >5{
		fmt.Println("Número maior que 5")
	}else{
		fmt.Println("Número menor que 5")
	}
	
	// Criando a variável limitada a esse if
	if outroNumero := numero; outroNumero > 0{
		fmt.Println("outroNúmero é maior que 0")
	}else {
		fmt.Println("outroNumero é menor que 0")
	}

	//fmt.Println(outroNumero) - dará erro pois a variável só existe la dentro do if
}