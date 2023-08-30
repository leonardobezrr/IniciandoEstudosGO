package auxiliar

import "fmt"

// Inicia com letra maisúcula para indicar que é uma função pública

//Registra uma mensagem na tela   - boa prática do go
func Escrever() {
	fmt.Println("Escrevendo pacote auxiliar")
	escrever2()  //Não precisou ser pública pois se encontra no mesmo pacote
}

