package main

import "fmt"

func main (){
	//ARITMÉTRICOS
	soma := 1+2
	sub := 1-2
	div := 10/2
	mult := 5*2
	restoDaDiv := 10 %2

	fmt.Println(soma,sub,div,mult,restoDaDiv)

	soma2 := 1+5
	fmt.Println(soma2)
	
	//ATRIBUIÇÃO
	var variavel1 string = "valor1"
	variavel2 := "valor 2"
	fmt.Println(variavel1,variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	verdadeiro,falso := true,false
	fmt.Println("\n----------------")
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//OPERADORES UNÁRIOS
	//Similar para a operação de subtração
	numero := 10
	numero ++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)
	
}