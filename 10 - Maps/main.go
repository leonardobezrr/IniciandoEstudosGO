package main

import "fmt"

func main (){
	fmt.Println("Maps")

	//[Tipo das chaves] Tipo dos valores
	usuario := map[string]string{
		"nome":"Pedro",
		"sobrenome":"Alves",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome":{
			"primeiro":"João",
			"segundo":"Pedro",
		},
		"Curso":{
			"nome":"Sistemas",
			"campus":"Caico",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2,"nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome":"Gêmeos",
		"mês":"Janeiro",
	}
	fmt.Println(usuario2)
}