package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade byte
	altura float32 
}

type estudante struct {
	pessoa // é como se tivesse colocado os itens de pessoa aqui Ex.: nome, sobrenome ... 
	curso string
	periodo string
}

func main (){
	fmt.Println("\n")
	pessoa1 := pessoa{nome: "leo"}
	pessoa2 := pessoa{"leo","alves",19,1.7}
	fmt.Println(pessoa1,"Printando apenas o nome")
	fmt.Println(pessoa2)

	//Utilizando a "herança"
	estudante1 := estudante{pessoa2,"Sistemas","segundo"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)

}