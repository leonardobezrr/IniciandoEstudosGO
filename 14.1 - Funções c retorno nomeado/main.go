package main

import "fmt"

func calcMat(n1,n2 int)(soma int,sub int){
	soma = n1 + n2
	sub = n1-n2
	return
}

func main(){
	n1 := 10
	n2 := 30
	soma, sub := calcMat(n1,n2)
	fmt.Println("\nN° 1 =",n1,"\nN° 2 =",n2)
	fmt.Println("\nsoma =",soma,"\nsubtração =",sub)
}