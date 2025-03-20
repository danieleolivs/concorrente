package main
import (
	"fmt"
	"sync"
)

func fib(total int) int{
	if total == 1 || total == 0{
		return 1
	}

	return fib(total - 1) + fib(total - 2)
}

func main(){
	var wg sync.WaitGroup

	wg.Add(1)

	go func(){
		fmt.Println(fib(5))
		defer wg.Done()
		
	}()
	
	fmt.Println("mensagem de status")
	wg.Wait()
}