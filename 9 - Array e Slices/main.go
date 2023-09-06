package main

import "fmt"

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
	fmt.Println(array3)

	array4 := [...]int{1,2,3,4,5,6} //Vai decidir o tamanho dependendo de quantos itens tiver no array 
	fmt.Println(array4)
}