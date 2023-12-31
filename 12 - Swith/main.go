package main

import "fmt"

func main(){
	fmt.Println("Switch")
	dia := diaDaSemana(6)
	fmt.Println(dia)

	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)
}

func diaDaSemana(numero int)string{
	switch numero {
	case 1: 
		return "Domingo"
		//fallthrough - Utilizado para empurrar para a prómixa condição 
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4: 
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Informe um valor válido"
	}
}

func diaDaSemana2(numero int)string{
	var diaDaSemana string

	switch{
	case numero == 1: //Mudou somente essa parte
		diaDaSemana =  "domingo"
	case numero == 2:
		diaDaSemana =  "Segunda"
	case numero == 3:
		diaDaSemana =  "Terça"
	case numero == 4: 
		diaDaSemana =  "Quarta"
	case numero == 5:
		diaDaSemana =  "Quinta"
	case numero == 6:
		diaDaSemana =  "Sexta"
	case numero == 7:
		diaDaSemana =  "Sábado"
	default:
		diaDaSemana =  "Igual o resto"
	}
	return diaDaSemana
}