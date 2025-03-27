package main

import(
	"fmt"
	"time"
	"math/rand"
)

func exec(valor int) int {
	rand.Seed(42)
	v := rand.Intn(valor)

	time.Sleep(time.Duration(v) * time.Millisecond)
	return v
}

func aux(max_sleep_ms int) <- chan int{
	result := make(chan int)

	go func(){
		defer close(result)

		for i:= 0; i < 1000; i++{
			result <- exec(max_sleep_ms)
		}
	}()

	return result

}

func main(){

	sum := 0

	for i := 0; i < 500; i++{
		select{
			case result_menor := <- aux(20):
				sum += result_menor
			case result_maior := <- aux(200):
				sum += result_maior
		}
	}

	fmt.Println(sum)
}

