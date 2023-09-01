package main 

import "fmt"

func somar(n1 byte,n2 byte)byte{		//byte equivale ao int8
	return n1 + n2
}

func calculosMat(n1,n2 int8) (int8,int8){  //Declarando o último valor, ele serve para o primeiro também
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main (){
	soma := somar(10,20) //Inferencia de tipo
	fmt.Println(soma)

	var f = func(txt string)string{
		fmt.Println(txt)
		return txt
	}
	
	resultado := f("retornando função f")
	f("Texto da função 1")
	fmt.Println(resultado,"oi")

	//Printar os dois valores 
	resultadoSoma,resultadoSub := calculosMat(10,15)
	fmt.Println(resultadoSoma,resultadoSub)

	//Printar somente um dos valores
	resultadoSoma1, _ := calculosMat(10,15)
	fmt.Println(resultadoSoma1)
}