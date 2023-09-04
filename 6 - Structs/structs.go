package main

import "fmt"

type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	rua string
	numero uint8
}

func main (){
	fmt.Println("Arquivos Structs")

	var u usuario
	u.nome = "Leo"
	u.idade = 21
	fmt.Println(u)

	enderecoEx := endereco {"Rua manel gomes",212}
	u2 := usuario {"Lionel",19,enderecoEx}
	fmt.Println(u2)

	u3 := usuario{nome :"Leandro"}
	fmt.Println(u3)
}