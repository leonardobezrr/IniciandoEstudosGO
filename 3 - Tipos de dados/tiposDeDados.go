package main

import (
	"fmt"
)

func main (){
	//int8,int16,int32,int64 Int's disponíveis 
	//int irá usar a quantidade de bits referente a arquitetura do computador

	//var numero int = 1000000000000000000
	numero := 1000000000000000000
	fmt.Println(numero)

	var numero2 uint = 5000 //não aceita valores negativos
	fmt.Println(numero2)

	// INT32  = RUNE
	var numero3 rune = 1212452155
	fmt.Println(numero3)

	//flaot
	var numeroComVirgula float32 = 126.54551321
	fmt.Println(numeroComVirgula)

	//Devolve o valor do caracter na tabela Ascii
	char := 'L'
	fmt.Println(char)
}