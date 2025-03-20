package main

import "fmt"

//gera uma sequência de números
func seqNumeros(init, final int) <- chan int{
	res := make(chan int)

	go func(){
		for i := init; i <= final; i++{
			res <- i
		}
		close(res)
	}()

	return res
}

func main(){
	valores := seqNumeros(1, 50)
	for valor := range valores{
		fmt.Printf("Valor: %v\n", valor)
	}
}