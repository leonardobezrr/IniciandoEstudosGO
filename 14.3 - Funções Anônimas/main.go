package main

import "fmt"

func main (){
	//Funções Anônimas 

	//Sem parâmetros
	func (){
		fmt.Println("Olá Mundo")
	}()

	//Com parâmetros
	func (texto string){
		fmt.Println(texto)
	}("Olá mundo")

	//Com retorno
	retorno := func (texto string)string{
		return fmt.Sprintf("Recebido -> %s",texto)
	}("Retornando com parâmetro")
	fmt.Println(retorno)
}