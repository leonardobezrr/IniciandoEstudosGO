package main

import (
	"fmt"
	"reflect"
)

func main (){
	fmt.Println("Array e Slices")

	var array1 [5]int
	array1 [0] = 1
	fmt.Println(array1)

	var array2 [3]string
	array2 = [3]string {"1","2","3"} // Podemos declarar assim
	fmt.Println(array2)

	array3 := [5]string {"ola 1,","2","3","4","5"}// e assim, por inferência de tipo 
	array3[1] = "1"
	fmt.Println(array3,"Array 3")

	array4 := [...]int{1,2,3,4,5,6} //Vai decidir o tamanho dependendo de quantos itens tiver no array 
	fmt.Println(array4,"Array 4")

	lista1 := []int {1,2,3,4,12,23,123,}	// SLICE Possui tamnho dinâmico, possuindo limitação apenas do tipo
	fmt.Println(lista1,"Lista 1")

	lista1 = append(lista1, 21) //Adicionando valores ao SLICE
	fmt.Println(lista1,"Lista 1 utilizando append")

	lista2 := array4[0:3]
	fmt.Println(lista2,"Selecionando valores do Array 4")

	fmt.Println(reflect.TypeOf(lista1))
	fmt.Println(reflect.TypeOf(array4)) // Possuem tipos diferentes 
}