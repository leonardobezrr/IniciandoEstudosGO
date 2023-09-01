package main

import "fmt"

func main(){
	var frases string = "String 1"
	frases2 := "String 2" //Utilizando a técnica de inferência de tipos, ele seleciona o tipo a partir do valor
	fmt.Println(frases,frases2)

	frases,frases2 = frases2,frases //Inverter de forma rápida o valor das variáveis

	fmt.Println(frases,frases2)
}